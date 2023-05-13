package api

import (
	"github.com/ehsandavari/go-clean-architecture/application/common/mocks"
	"github.com/ehsandavari/go-clean-architecture/infrastructure/config"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_NewSApi_Start_Stop(t *testing.T) {
	sConfig := &config.SConfig{
		Service: &config.SService{
			Api: &config.SApi{
				IsEnabled: true,
				Host:      "localhost",
				Port:      "8080",
				Mode:      "debug",
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := mocks.NewMockILogger(ctrl)
	mockLogger.EXPECT().Info(gomock.Any()).Times(1)

	sApi := NewSApi(sConfig, mockLogger)

	// Start the server
	sApi.Start()

	// Act
	resp, err := http.Get("http://localhost:8080/-/health")
	defer resp.Body.Close()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Stop the server
	sApi.Stop()

	_, err = http.Get("http://localhost:8080/-/health")
	assert.Error(t, err)
}
