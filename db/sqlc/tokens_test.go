package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateAndGetToken(t *testing.T) {
	// Создаем тестового пользователя
	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Email:        "tokenuser@test.com",
		PasswordHash: "testhash",
		Role:         "client",
	})
	require.NoError(t, err)

	// Создаем токен
	tokenArg := CreateTokenParams{
		UserID:    user.ID,
		Token:     "test_token_value",
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	token, err := testQueries.CreateToken(context.Background(), tokenArg)
	require.NoError(t, err)

	// Получаем токен
	fetchedToken, err := testQueries.GetToken(context.Background(), tokenArg.Token)
	require.NoError(t, err)

	require.Equal(t, token.Token, fetchedToken.Token)
	require.Equal(t, token.UserID, fetchedToken.UserID)
}

func TestDeleteExpiredTokens(t *testing.T) {
	// Создаем просроченный токен
	_, _ = testQueries.CreateToken(context.Background(), CreateTokenParams{
		Token:     "expired_token",
		ExpiresAt: time.Now().Add(-1 * time.Hour),
	})

	// Удаляем просроченные
	err := testQueries.DeleteExpiredTokens(context.Background())
	require.NoError(t, err)

	// Проверяем что токен удален
	_, err = testQueries.GetToken(context.Background(), "expired_token")
	require.Error(t, err)
}
