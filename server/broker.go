package server

import (
	"net"
	"os"
	"strings"
)

type MessageBroker interface {
	ReceiveRequest() (Request, net.Conn) // Receives request from a client
	ProcessRequest(Request) []Task       // Processes raw request
	SendTask(Task)                       // Sends task to a service
	ReceiveAck(Task) []byte              // Receives acknowledgment from a service
	SendResponse(net.Conn, []byte)       // Sends response to a client
}

func (server Server) ReceiveRequest() (Request, net.Conn) {
	connection, err := server.listener.Accept()

	if err != nil {
		ServerLog("Error occured during connecting: " + err.Error())
		os.Exit(1)
	}

	messageBuffer := make([]byte, BUFFER_SIZE)
	messageLength, err := connection.Read(messageBuffer)

	if err != nil {
		ServerLog("Error occured during reading: " + err.Error())
	}

	clientHost := strings.Split(connection.LocalAddr().String(), ":")[0]
	clientPort := strings.Split(connection.LocalAddr().String(), ":")[1]
	message := messageBuffer[:messageLength]
	request := Request{Client{clientHost, clientPort}, message}
	return request, connection
}

func (server Server) ProcessRequest(request Request) []Task {
	decodedRequest := string(request.message)
	taskStringArr := strings.Split(decodedRequest, "$")
	taskArr := make([]Task, 0)

	for _, taskString := range taskStringArr {
		taskType := strings.Split(taskString, ":")[0]
		taskData := []byte(strings.Split(taskString, ":")[1])
		taskArr = append(taskArr, Task{taskType, taskData})
	}

	return taskArr
}

func (server Server) SendTask(task Task) {
	for _, service := range server.services {
		if service.name == task.taskType {
			service.connection.Write(task.data)
			break
		}
	}
}

func (server Server) ReceiveAck(task Task) []byte {
	var acknowledgment []byte
	for _, service := range server.services {
		if service.name == task.taskType {
			ackBuffer := make([]byte, BUFFER_SIZE)
			ackLength, err := service.connection.Read(ackBuffer)
			if err != nil {
				ServerLog("Error while receiving acknowledgement: " + err.Error())
				os.Exit(1)
			}
			acknowledgment = ackBuffer[:ackLength]
			break
		}
	}
	return acknowledgment
}

func (server Server) SendResponse(connection net.Conn, acknowledgment []byte) {
	connection.Write(acknowledgment)
	connection.Close()
}
