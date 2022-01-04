package db

import (
	"context"
	"testing"
	"time"

	"github.com/secretnote/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	args := CreateUserParams{
		Username:    util.RandomEmail(),
		Email:       util.RandomEmail(),
		TypeOfLogin: "GOOGLE",
	}

	user, err := testQueries.CreateUser(context.Background(), args)

	require.NoError(t, err)

	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.TypeOfLogin, user.TypeOfLogin)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	account1 := createRandomUser(t)
	account2, err := testQueries.GetUser(context.Background(), account1.Email)

	require.NoError(t, err)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.TypeOfLogin, account2.TypeOfLogin)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}
