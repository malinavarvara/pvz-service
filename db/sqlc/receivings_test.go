package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateReceiving(t *testing.T) {
	// Сначала создаем ПВЗ для связи
	pvz, err := testQueries.CreatePickupPoint(context.Background(), CreatePickupPointParams{
		Name:    "Для приёмки",
		City:    "Санкт-Петербург",
		Address: "Невский пр., 10",
	})
	require.NoError(t, err)

	arg := CreateReceivingParams{
		PickupPointID: pvz.ID,
		Status:        "in_progress",
	}

	receiving, err := testQueries.CreateReceiving(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, receiving)

	require.Equal(t, arg.PickupPointID, receiving.PickupPointID)
	require.Equal(t, arg.Status, receiving.Status)
	require.False(t, receiving.ClosedAt.Valid)

	require.NotZero(t, receiving.ID)
	require.NotZero(t, receiving.StartedAt)
}

func TestCloseReceiving(t *testing.T) {
	// 1. Создаем тестовые данные
	pvz, err := testQueries.CreatePickupPoint(context.Background(), CreatePickupPointParams{
		Name:    "Test PVZ",
		City:    "Москва",
		Address: "ул. Тестовая, 1",
	})
	require.NoError(t, err)

	receiving, err := testQueries.CreateReceiving(context.Background(), CreateReceivingParams{
		PickupPointID: pvz.ID,
		Status:        "in_progress",
	})
	require.NoError(t, err)

	// 2. Закрываем приёмку (с явным приведением типа ID)
	updated, err := testQueries.UpdateReceivingStatus(context.Background(), UpdateReceivingStatusParams{
		Status: "closed",     // Параметр $1
		ID:     receiving.ID, // Параметр $2 (int64)
	})
	require.NoError(t, err)

	// 3. Проверяем результат
	require.Equal(t, "closed", updated.Status)
	require.True(t, updated.ClosedAt.Valid) // Проверяем что closed_at установлен
	require.WithinDuration(t, time.Now(), updated.ClosedAt.Time, 2*time.Second)
}
