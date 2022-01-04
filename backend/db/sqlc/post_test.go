package db

import (
	"context"
	"testing"
	"time"

	"github.com/secretnote/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T, user User) Post {
	args := CreatePostParams{
		Email:   user.Email,
		Title:   util.RandomString(100),
		Content: util.RandomString(150),
	}

	post, err := testQueries.CreatePost(context.Background(), args)

	require.NoError(t, err)

	require.NotEmpty(t, user)

	require.Equal(t, args.Email, post.Email)
	require.Equal(t, args.Title, post.Title)
	require.Equal(t, args.Content, post.Content)
	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)

	return post
}

func TestCreatePost(t *testing.T) {
	user := createRandomUser(t)
	createRandomPost(t, user)
}

func TestGetPost(t *testing.T) {
	user := createRandomUser(t)
	post1 := createRandomPost(t, user)
	post2, err := testQueries.GetPost(context.Background(), post1.ID)

	require.NoError(t, err)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.Email, post2.Email)
	require.Equal(t, post1.Title, post2.Title)
	require.Equal(t, post1.Content, post2.Content)
	require.WithinDuration(t, post1.CreatedAt, post2.CreatedAt, time.Second)
}

func TestGetListPost(t *testing.T) {
	user := createRandomUser(t)

	for i := 0; i < 5; i++ {
		createRandomPost(t, user)
	}
	args := ListPostsParams{
		Email:  user.Email,
		Limit:  5,
		Offset: 0,
	}
	listPosts, err := testQueries.ListPosts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, listPosts, 5)
	for _, post := range listPosts {
		require.NotEmpty(t, post)
		require.Equal(t, post.Email, user.Email)
	}
}
