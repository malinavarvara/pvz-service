package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddProduct(t *testing.T) {
	// Создаем тестовый ПВЗ
	pvz, err := testQueries.CreatePickupPoint(context.Background(), CreatePickupPointParams{
		Name:    "Тестовый ПВЗ",
		City:    "Москва",
		Address: "ул. Тестовая, 1",
	})
	require.NoError(t, err)

	// Создаем тестовую приёмку
	receiving, err := testQueries.CreateReceiving(context.Background(), CreateReceivingParams{
		PickupPointID: pvz.ID,
		Status:        "in_progress",
	})
	require.NoError(t, err)

	arg := AddProductParams{
		ReceivingID:  receiving.ID,
		Type:         "electronics",
		Description:  sql.NullString{String: "Смартфон", Valid: true},
		AvitoOrderID: "AVITO-12345",
	}

	product, err := testQueries.AddProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.ReceivingID, product.ReceivingID)
	require.Equal(t, arg.Type, product.Type)
	require.Equal(t, arg.AvitoOrderID, product.AvitoOrderID)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.AddedAt)
}

func TestDeleteLastProduct(t *testing.T) {
	// 1. Создаем тестовый ПВЗ
	pvz, err := testQueries.CreatePickupPoint(context.Background(), CreatePickupPointParams{
		Name:    "ПВЗ для теста удаления",
		City:    "Санкт-Петербург",
		Address: "Невский пр., 100",
	})
	require.NoError(t, err)
	require.NotZero(t, pvz.ID)

	// 2. Создаем приёмку
	receiving, err := testQueries.CreateReceiving(context.Background(), CreateReceivingParams{
		PickupPointID: pvz.ID,
		Status:        "in_progress",
	})
	require.NoError(t, err)
	require.NotZero(t, receiving.ID)

	// 3. Добавляем два тестовых товара
	product1, err := testQueries.AddProduct(context.Background(), AddProductParams{
		ReceivingID:  receiving.ID,
		Type:         "electronics",
		Description:  sql.NullString{String: "Ноутбук", Valid: true},
		AvitoOrderID: "AVITO-111",
	})
	require.NoError(t, err)
	require.NotZero(t, product1.ID)

	product2, err := testQueries.AddProduct(context.Background(), AddProductParams{
		ReceivingID:  receiving.ID,
		Type:         "clothing",
		Description:  sql.NullString{String: "Футболка", Valid: true},
		AvitoOrderID: "AVITO-222",
	})
	require.NoError(t, err)
	require.NotZero(t, product2.ID)

	// 4. Проверяем что товары добавились
	productsBefore, err := testQueries.ListProductsInReceiving(context.Background(), receiving.ID)
	require.NoError(t, err)
	require.Len(t, productsBefore, 2)

	// 5. Удаляем последний товар (LIFO)
	deleted, err := testQueries.DeleteLastProduct(context.Background(), receiving.ID)
	require.NoError(t, err)

	// 6. Проверяем что удалился именно последний товар
	require.Equal(t, product2.ID, deleted.ID)
	require.Equal(t, product2.Type, deleted.Type)

	// 7. Проверяем что остался только первый товар
	productsAfter, err := testQueries.ListProductsInReceiving(context.Background(), receiving.ID)
	require.NoError(t, err)
	require.Len(t, productsAfter, 1)
	require.Equal(t, product1.ID, productsAfter[0].ID)
}

func TestListPickupPoints(t *testing.T) {
	// Создаем несколько тестовых ПВЗ
	for i := 0; i < 5; i++ {
		_, err := testQueries.CreatePickupPoint(context.Background(), CreatePickupPointParams{
			Name:    fmt.Sprintf("ПВЗ %d", i),
			City:    "Москва",
			Address: fmt.Sprintf("ул. Тестовая, %d", i),
		})
		require.NoError(t, err)
	}

	// Получаем список
	pvzList, err := testQueries.ListPickupPointsByCity(context.Background(), "Москва")
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(pvzList), 5)

	for _, pvz := range pvzList {
		require.Equal(t, "Москва", pvz.City)
		require.NotEmpty(t, pvz.Name)
	}
}
