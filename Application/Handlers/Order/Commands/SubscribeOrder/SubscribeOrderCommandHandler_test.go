package SubscribeOrder

//
//import (
//	"GolangCodeBase/Application/Common/Interfaces"
//	"GolangCodeBase/Application/Mocks"
//	"GolangCodeBase/Infrastructure/Config"
//	"context"
//	"github.com/golang/mock/gomock"
//	"os"
//	"testing"
//)
//
//func TestMain(m *testing.M) {
//	os.Exit(m.Run())
//}
//
//func setup(t *testing.T) (Interfaces.IOrderHandlerCommands, *Mocks.MockIUnitOfWork, *Mocks.MockIRedis, Config.SConfig, func()) {
//	mockCtrl := gomock.NewController(t)
//	mockILogger := Mocks.NewMockILogger(mockCtrl)
//	mockIUnitOfWork := Mocks.NewMockIUnitOfWork(mockCtrl)
//	mockIRedis := Mocks.NewMockIRedis(mockCtrl)
//	mockIConfiguration, _ := Config.NewConfig()
//
//	service := NewSubscribeOrderCommandHandler(mockIConfiguration, mockILogger, mockIUnitOfWork, mockIRedis)
//
//	return service, mockIUnitOfWork, mockIRedis, mockIConfiguration, func() {
//		mockCtrl.Finish()
//	}
//}
//
//func TestSubscribeOrderCommand(t *testing.T) {
//	t.Run("fail on store creation error", func(t *testing.T) {
//		orderHandlerCommands, _, redis, conf, teardown := setup(t)
//		defer teardown()
//		ctx := context.Background()
//		//unitOfWork.OrderRepository().Add(gomock.Any())
//		redis.EXPECT().Subscribe(ctx, conf.Redis.Queues[string], make(chan string)).Do(func() {})
//		go orderHandlerCommands.SubscribeOrderCommand(ctx)
//	})
//}
