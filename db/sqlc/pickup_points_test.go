package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePickupPoint(t *testing.T) {
	arg := CreatePickupPointParams{
		Name:    "ПВЗ Центральный",
		City:    "Москва",
		Address: "ул. Тверская, д. 1",
	}

	pvz, err := testQueries.CreatePickupPoint(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pvz)

	require.Equal(t, arg.Name, pvz.Name)
	require.Equal(t, arg.City, pvz.City)
	require.Equal(t, arg.Address, pvz.Address)

	require.NotZero(t, pvz.ID)
	require.NotZero(t, pvz.RegisteredAt)
}

func TestGetPickupPoint(t *testing.T) {
	// Сначала создаем тестовые данные
	createdPVZ, err := testQueries.CreatePickupPoint(context.Background(), CreatePickupPointParams{
		Name:    "Тестовый ПВЗ",
		City:    "Казань",
		Address: "ул. Баумана, 1",
	})
	require.NoError(t, err)

	// Получаем созданный ПВЗ
	fetchedPVZ, err := testQueries.GetPickupPoint(context.Background(), createdPVZ.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedPVZ)

	require.Equal(t, createdPVZ.ID, fetchedPVZ.ID)
	require.Equal(t, createdPVZ.Name, fetchedPVZ.Name)
}
