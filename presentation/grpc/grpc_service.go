package grpc

import (
	"context"
	bankAccountService "github.com/ehsandavari/go-clean-architecture/presentation/grpc/proto/bankAccount"
	"time"
)

type Service struct {
}

func NewGrpcService() *Service {
	return &Service{}
}

func (g *Service) CreateBankAccount(ctx context.Context, request *bankAccountService.CreateBankAccountRequest) (*bankAccountService.CreateBankAccountResponse, error) {
	time.Sleep(10 * time.Second)
	return &bankAccountService.CreateBankAccountResponse{}, nil
}

func (g *Service) DepositBalance(ctx context.Context, request *bankAccountService.DepositBalanceRequest) (*bankAccountService.DepositBalanceResponse, error) {
	time.Sleep(10 * time.Second)
	return &bankAccountService.DepositBalanceResponse{}, nil
}

func (g *Service) WithdrawBalance(ctx context.Context, request *bankAccountService.WithdrawBalanceRequest) (*bankAccountService.WithdrawBalanceResponse, error) {
	time.Sleep(10 * time.Second)
	return &bankAccountService.WithdrawBalanceResponse{}, nil
}

func (g *Service) ChangeEmail(ctx context.Context, request *bankAccountService.ChangeEmailRequest) (*bankAccountService.ChangeEmailResponse, error) {
	time.Sleep(10 * time.Second)
	return &bankAccountService.ChangeEmailResponse{}, nil
}

func (g *Service) GetById(ctx context.Context, request *bankAccountService.GetByIdRequest) (*bankAccountService.GetByIdResponse, error) {
	time.Sleep(10 * time.Second)
	return &bankAccountService.GetByIdResponse{}, nil
}

func (g *Service) SearchBankAccounts(ctx context.Context, request *bankAccountService.SearchBankAccountsRequest) (*bankAccountService.SearchBankAccountsResponse, error) {
	time.Sleep(10 * time.Second)
	return &bankAccountService.SearchBankAccountsResponse{}, nil
}
