package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kungfucoding23/testeable_code/services"
	"github.com/gin-gonic/gin"
)

type pingServiceMock struct {
	handlePingFn func() (string, error)
}

//implement our method in our mock
func (mock pingServiceMock) HandlePing() (string, error) {
	fmt.Println("mocking complex things...")
	return mock.handlePingFn()
}

func TestPingNoError(t *testing.T) {
	serviceMock := pingServiceMock{}
	serviceMock.handlePingFn = func() (string, error) {
		return "pong", nil
	}
	services.PingService = serviceMock
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	Ping(context)
	if response.Code != http.StatusOK {
		t.Error("response code should be 200")
	}

	if response.Body.String() != "pong" {
		t.Error("response body should say 'pong'")
	}
}

func TestPingWithError(t *testing.T) {
	serviceMock := pingServiceMock{}
	serviceMock.handlePingFn = func() (string, error) {
		return "", errors.New("error executing ping")
	}
	services.PingService = serviceMock
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	Ping(context)
	if response.Code != http.StatusInternalServerError {
		t.Error("response code should be 500")
	}

	if response.Body.String() != "error executing ping" {
		t.Error("response body should say 'error'")
	}
}
