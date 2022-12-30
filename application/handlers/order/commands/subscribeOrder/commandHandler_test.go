package subscribeOrder

//
//import (
//	"golangCodeBase/application/common/interfaces"
//	"golangCodeBase/application/mocks"
//	"golangCodeBase/infrastructure/config"
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
//func setup(t *testing.T) (interfaces.IOrderHandlerCommands, *mocks.MockIUnitOfWork, *mocks.MockIRedis, config.SConfig, func()) {
//	mockCtrl := gomock.NewController(t)
//	mockILogger := mocks.NewMockILogger(mockCtrl)
//	mockIUnitOfWork := mocks.NewMockIUnitOfWork(mockCtrl)
//	mockIRedis := mocks.NewMockIRedis(mockCtrl)
//	mockIConfiguration, _ := config.NewConfig()
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
//		redis.EXPECT().Subscribe(ctx, conf.redis.Queues[string], make(chan string)).Do(func() {})
//		go orderHandlerCommands.SubscribeOrderCommand(ctx)
//	})
//}
