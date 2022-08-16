package src

import (
	"net"
	"os"
	"strings"
)

type MessageBroker interface {
	ReceiveRequest() (Request, *net.UDPAddr)   // Receives request from a client
	ProcessRequest(Request) ([]Task, [][]byte) // Processes raw request
	SendTask(Task)                             // Sends task to a service
	ReceiveAck(Task) []byte                    // Receives acknowledgment from a service
	SendResponse(*net.UDPAddr, [][]byte)       // Sends response to a client
}

func (server *Server) ReceiveRequest() (Request, *net.UDPAddr) {
	messageBuffer := make([]byte, BUFFER_SIZE)
	messageLength, clientAddr, err := server.udpConnection.ReadFromUDP(messageBuffer)

	if err != nil {
		ServerLog("Error occured during receiving a request: " + err.Error())
		return Request{}, nil
	}

	clientHost := clientAddr.IP.String()
	clientPort := clientAddr.Port
	message := messageBuffer[:messageLength]
	request := Request{Client{clientHost, clientPort}, message}
	ServerLog("Request received from " + clientAddr.String())
	return request, clientAddr
}

func (server *Server) ProcessRequest(request Request) ([]Task, [][]byte) {
	decodedRequest := string(request.message)
	taskStringArr := strings.Split(decodedRequest, TASK_SEPERATOR)
	taskArr := make([]Task, 0)

	for _, taskString := range taskStringArr {
		taskType := strings.Split(taskString, HEAD_BODY_SEPERATOR)[0]
		taskData := []byte(strings.Split(taskString, HEAD_BODY_SEPERATOR)[1])
		taskArr = append(taskArr, Task{taskType, taskData})
	}

	ackArr := make([][]byte, len(taskStringArr))
	for idx := range ackArr {
		ackArr[idx] = make([]byte, 0)
	}

	return taskArr, ackArr
}

func (server *Server) SendTask(task Task) {
	for _, service := range server.services {
		if service.name == task.taskType {
			service.udpConnection.Write(task.data)
			break
		}
	}
	ServerLog("Task sent to Service:" + task.taskType)
}

func (server *Server) ReceiveAck(task Task) []byte {
	var acknowledgment []byte
	for _, service := range server.services {
		if service.name == task.taskType {
			ackBuffer := make([]byte, BUFFER_SIZE)
			ackLength, err := service.udpConnection.Read(ackBuffer)
			if err != nil {
				ServerLog("Error during receiving acknowledgement: " + err.Error())
				os.Exit(1)
			}
			acknowledgment = ackBuffer[:ackLength]
			break
		}
	}
	ServerLog("Acknowledgment received from a service")
	return acknowledgment
}

func (server *Server) SendResponse(udpAddres *net.UDPAddr, ackArr [][]byte) {
	for _, acknowledgment := range ackArr {
		server.udpConnection.WriteToUDP(acknowledgment, udpAddres)
	}
	ServerLog("Responsed to " + udpAddres.String())
}
