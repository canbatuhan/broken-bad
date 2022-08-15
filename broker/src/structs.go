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
	udpConnection *net.UDPConn
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
	host          string
	port          int
	udpConnection *net.UDPConn
	services      []Service
}

func InitServer(host string, port int) Server {
	var server Server
	server.host = host
	server.port = port

	address := net.UDPAddr{IP: net.ParseIP(server.host), Port: server.port}
	connection, err := net.ListenUDP(CONNECTION_TYPE, &address)

	if err != nil {
		ServerLog("Error occured during starting: " + err.Error())
		os.Exit(1)
	}

	server.udpConnection = connection
	server.services = make([]Service, 0)
	ServerLog("Started")
	return server
}

func ServerLog(message string) {
	timestamp := time.Now().Local().String()
	fmt.Println("[SERVER] " + timestamp + " - " + message)
}

func AddService(server *Server, service Service) {
	listenerAddress := net.UDPAddr{IP: net.ParseIP(service.host), Port: service.port}
	connection, err := net.DialUDP(CONNECTION_TYPE, nil, &listenerAddress)

	if err != nil {
		ServerLog("Error occured during connecting to a service: " + err.Error())
		os.Exit(1)
	}

	ServerLog("UDP Connection is made with Service: " + service.name + " on " + listenerAddress.String())
	service.udpConnection = connection
	server.services = append(server.services, service)
}
