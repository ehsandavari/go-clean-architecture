package SubscribeOrder

import (
	ApplicationInterfaces "GolangCodeBase/Application/Common/Interfaces"
	MockIRedis "GolangCodeBase/Application/Mocks"
	MockIUnitOfWork "GolangCodeBase/Domain/Mocks"
	"GolangCodeBase/Infrastructure/Config"
	"context"
	"github.com/golang/mock/gomock"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setup(t *testing.T) (ApplicationInterfaces.IOrderHandlerCommands, *MockIUnitOfWork.MockIUnitOfWork, *MockIRedis.MockIRedis, Config.SConfig, func()) {
	mockCtrl := gomock.NewController(t)
	mockIUnitOfWork := MockIUnitOfWork.NewMockIUnitOfWork(mockCtrl)
	mockIRedis := MockIRedis.NewMockIRedis(mockCtrl)
	mockIConfiguration := Config.NewConfig()

	service := NewSubscribeOrderCommand(mockIConfiguration, mockIUnitOfWork, mockIRedis)

	return service, mockIUnitOfWork, mockIRedis, mockIConfiguration, func() {
		mockCtrl.Finish()
	}
}

func TestSubscribeOrderCommand(t *testing.T) {
	t.Run("fail on store creation error", func(t *testing.T) {
		orderHandlerCommands, _, redis, conf, teardown := setup(t)
		defer teardown()
		ctx := context.Background()
		//unitOfWork.OrderRepository().Add(gomock.Any())
		redis.EXPECT().Subscribe(ctx, conf.Redis.Queues[string], make(chan string)).Do(func() {})
		go orderHandlerCommands.SubscribeOrderCommand(ctx)
	})
}
