package services

import "fmt"

const (
	pong = "pong"
)

type pingService interface {
	HandlePing() (string, error)
}

type pingServiceImpl struct{}

var (
	//PingService is a variable of type pingService interface
	PingService pingService = pingServiceImpl{}
)

//HandlePing ..
func (service pingServiceImpl) HandlePing() (string, error) {
	fmt.Println("doing some complex things...")
	return pong, nil
}
