package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Email:        "12345678@gmail.com",
		PasswordHash: "knvkjdfnbkdrsejgbnjktfnbfnmgfk",
		Role:         "employee",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.Role, user.Role)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}

func TestDeleteUser(t *testing.T) {
	// 1. Сначала создаем пользователя (аналогично Create тесту)
	createArg := CreateUserParams{
		Email:        "user.to.delete@example.com",
		PasswordHash: "hashed_password_to_delete",
		Role:         "client",
	}

	createdUser, err := testQueries.CreateUser(context.Background(), createArg)
	require.NoError(t, err)
	require.NotEmpty(t, createdUser)

	// 2. Выполняем удаление
	err = testQueries.DeleteUser(context.Background(), createdUser.ID)
	require.NoError(t, err)

	// 3. Проверяем что пользователь действительно удален
	fetchedUser, err := testQueries.GetUserByEmail(context.Background(), createArg.Email)
	require.Error(t, err)         // Ожидаем ошибку
	require.Empty(t, fetchedUser) // Должен быть пустой объект

	// 4. Дополнительная проверка - пытаемся удалить уже удаленного пользователя
	err = testQueries.DeleteUser(context.Background(), createdUser.ID)
	require.NoError(t, err) // Удаление несуществующей записи не должно возвращать ошибку
}
