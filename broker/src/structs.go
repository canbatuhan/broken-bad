package src

import (
	"fmt"
	"net"
	"os"
	"time"
)

type Request struct {
	client  Client
	message []byte
}

func InitRequest(client Client, message []byte) Request {
	var request Request
	request.client = client
	request.message = message
	return request
}

type Task struct {
	taskType string
	data     []byte
}

func InitTask(taskType string, data []byte) Task {
	var task Task
	task.taskType = taskType
	task.data = data
	return task
}

type Client struct {
	host string
	port string
}

func InitClient(host string, port string) Client {
	var client Client
	client.host = host
	client.port = port
	return client
}

type Service struct {
	name       string
	host       string
	port       string
	connection net.Conn
}

func InitService(name string, host string, port string) Service {
	var service Service
	service.name = name
	service.host = host
	service.port = port
	service.connection = nil
	return service
}

type Server struct {
	host           string
	port           string
	connectionType string
	listener       net.Listener
	services       []Service
}

func InitServer(connectionType string, host string, port string) Server {
	var server Server
	server.host = host
	server.port = port
	server.connectionType = connectionType

	address := server.host + ":" + server.port
	listener, err := net.Listen(server.connectionType, address)

	if err != nil {
		ServerLog("Error occured during starting: " + err.Error())
		os.Exit(1)
	}

	server.listener = listener
	server.services = make([]Service, 0)
	return server
}

func ServerLog(message string) {
	timestamp := time.Now().Local().String()
	fmt.Println("[SERVER] " + timestamp + " - " + message)
}

func AddService(server *Server, service Service) {
	server.services = append(server.services, service)
	address := service.host + ":" + service.port
	connection, err := net.Dial(server.connectionType, address)

	if err != nil {
		ServerLog("Error occured while connecting to service: " + err.Error())
	}

	service.connection = connection
}
