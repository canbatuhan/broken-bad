package server

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	CONNECTION_TYPE = "tcp"
	BUFFER_SIZE     = 1024
)

type Request struct {
	client  Client
	message []byte
}

type Task struct {
	taskType string
	data     []byte
}

type Client struct {
	host string
	port string
}

type Service struct {
	name       string
	host       string
	port       string
	connection net.Conn
}

type Server struct {
	host     string
	port     string
	listener net.Listener
	services []Service
}

func ServerLog(message string) {
	timestamp := time.Now().Local().String()
	fmt.Println("[SERVER] " + timestamp + " - " + message)
}

func InitServer(connectionType string, host string, port string) Server {
	var server Server
	server.host = host
	server.port = port

	address := server.host + ":" + server.port
	listener, err := net.Listen(connectionType, address)

	if err != nil {
		ServerLog("Error occured during starting: " + err.Error())
		os.Exit(1)
	}

	server.listener = listener
	server.services = make([]Service, 0)
	return server
}

func AddService(server *Server, service Service) {
	server.services = append(server.services, service)
	address := service.host + ":" + service.port
	connection, err := net.Dial(CONNECTION_TYPE, address)

	if err != nil {
		ServerLog("Error occured while connecting to service: " + err.Error())
	}

	service.connection = connection
}
