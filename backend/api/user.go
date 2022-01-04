package api

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/secretnote/backend/db/sqlc"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
// var googleOauthConfig = &oauth2.Config{
// 	RedirectURL:  "http://localhost:8080/google",
// 	ClientID:     "109570213407-a20vnkjpsq7kqmeinbjcef2gs8mfsv15.apps.googleusercontent.com",
// 	ClientSecret: "pLgJtteHz3FHWJs_26drJbVC",
// 	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
// 	Endpoint:     google.Endpoint,
// }

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type urlResponse struct {
	URL string `json:"redirectURL"`
}

func (server *Server) oauthGoogleLogin(ctx *gin.Context) {

	oauthState := generateStateForOauth()
	u := server.googleOauthConfig.AuthCodeURL(oauthState)
	urlRes := urlResponse{
		URL: u,
	}
	ctx.JSON(http.StatusOK, urlRes)
}

type getLoginRequest struct {
	State string `form:"state"`
	Code  string `form:"code"`
	Scope string `form:"scope"`
}

type loginUserResponse struct {
	AccessToken string  `json:"access_token"`
	User        db.User `json:"user"`
}

func (server *Server) oauthGoogleCallback(ctx *gin.Context) {
	// Read oauthState from Cookie
	var req getLoginRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userEmail, err := server.getUserDataFromGoogle(req.Code, ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg, err := server.store.GetUser(ctx, userEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			userCreated := db.CreateUserParams{
				Username:    userEmail,
				Email:       userEmail,
				TypeOfLogin: "GOOGLE",
			}

			userDetail, err := server.store.CreateUser(ctx, userCreated)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
			accessToken, err := server.tokenMaker.CreateToken(userDetail.Email, server.config.AccessTokenDuration)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			response := loginUserResponse{
				AccessToken: accessToken,
				User:        userDetail,
			}
			ctx.JSON(http.StatusOK, response)
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(arg.Email, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := loginUserResponse{
		AccessToken: accessToken,
		User:        arg,
	}
	ctx.JSON(http.StatusOK, response)
}

type userInfor struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail string `json:"verified_email"`
	Picture       string `json:"picture"`
}

func (server *Server) getUserDataFromGoogle(code string, ctx *gin.Context) (string, error) {
	// Use code to get token and get user info from Google.
	token, err := server.googleOauthConfig.Exchange(ctx, code)
	fmt.Println(token.AccessToken)
	if err != nil {
		return "", errors.New("code exchange wrong")
	}
	urlInfor := oauthGoogleUrlAPI + token.AccessToken
	response, err := http.Get(urlInfor)
	if err != nil {
		return "", errors.New("failed to read response")
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("failed to read body")
	}

	data := userInfor{}
	json.Unmarshal(contents, &data)

	if data.Email == "" {
		return "", errors.New("failed to get email")
	}
	return data.Email, nil
}

func generateStateForOauth() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}
