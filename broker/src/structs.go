package src

import (
	"fmt"
	"net"
	"os"
	"strconv"
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
	port int
}

func InitClient(host string, port int) Client {
	var client Client
	client.host = host
	client.port = port
	return client
}

type Service struct {
	name          string
	host          string
	port          int
	udpConnection net.Conn
}

func InitService(name string, host string, port int) Service {
	var service Service
	service.name = name
	service.host = host
	service.port = port
	service.udpConnection = nil
	return service
}

type Server struct {
	host           string
	port           int
	connectionType string
	udpConnection  *net.UDPConn
	services       []Service
}

func InitServer(connectionType string, host string, port int) Server {
	var server Server
	server.host = host
	server.port = port
	server.connectionType = connectionType

	address := net.UDPAddr{IP: net.ParseIP(server.host), Port: server.port}
	connection, err := net.ListenUDP(server.connectionType, &address)

	if err != nil {
		ServerLog("Error occured during starting: " + err.Error())
		os.Exit(1)
	}

	server.udpConnection = connection
	server.services = make([]Service, 0)
	return server
}

func ServerLog(message string) {
	timestamp := time.Now().Local().String()
	fmt.Println("[SERVER] " + timestamp + " - " + message)
}

func AddService(server *Server, service Service) {
	server.services = append(server.services, service)
	addr := service.host + ":" + strconv.Itoa(service.port)
	fmt.Println(addr)
	connection, err := net.Dial(server.connectionType, addr)

	if err != nil {
		ServerLog("Error occured during connecting to a service: " + err.Error())
		os.Exit(1)
	}

	service.udpConnection = connection
}
