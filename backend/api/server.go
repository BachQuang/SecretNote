package api

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/secretnote/backend/db/sqlc"
	"github.com/secretnote/backend/token"
	"github.com/secretnote/backend/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Server struct {
	config            util.Config
	store             db.Store
	tokenMaker        token.Maker
	router            *gin.Engine
	googleOauthConfig *oauth2.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("typeOfLogin", validTypeOfLogin)
	}

	server.SetupRouter()
	server.SetupGoogleOauth(config)
	return server, nil

}

func (server *Server) SetupGoogleOauth(config util.Config) {
	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  config.REDIRECT_URL,
		ClientID:     config.CLIENT_ID,
		ClientSecret: config.CLIENT_SECRET,
		Scopes:       []string{config.SCOPES},
		Endpoint:     google.Endpoint,
	}
	server.googleOauthConfig = googleOauthConfig
}
func (server *Server) SetupRouter() {

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	authapi := router.Group("/auth")
	authapi.GET("/google/login", server.oauthGoogleLogin)
	authapi.POST("/google/callback", server.oauthGoogleCallback)

	authRoutes := router.Group("/api").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/posts", server.createPost)
	authRoutes.GET("/posts/:id", server.getPost)
	authRoutes.GET("/posts", server.listPost)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
