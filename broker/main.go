package main

import (
	broker "broker/src"
	"sync"
)

func Setup() broker.Server {
	server := broker.InitServer("127.0.0.1", 8000)
	broker.AddService(&server, broker.InitService("FileWriterService", "127.0.0.1", 8010))
	broker.AddService(&server, broker.InitService("PowerCalculatorService", "127.0.0.1", 8020))
	return server
}

func Run(mb broker.MessageBroker) {
	for {
		request, clientConnection := mb.ReceiveRequest()
		taskArr, ackArr := mb.ProcessRequest(request)

		var taskGroup sync.WaitGroup
		for idx, task := range taskArr {
			taskGroup.Add(1)
			go func(taskGroup *sync.WaitGroup, task broker.Task, ackArr [][]byte, idx int) {
				mb.SendTask(task)
				ackArr[idx] = mb.ReceiveAck(task)
				taskGroup.Done()
			}(&taskGroup, task, ackArr, idx)
		}

		taskGroup.Wait()
		mb.SendResponse(clientConnection, ackArr)
	}
}

func main() {
	server := Setup()
	Run(&server)
}
