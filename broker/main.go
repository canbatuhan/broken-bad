package main

import broker "broker/src"

func Setup() broker.Server {
	server := broker.InitServer("tcp", "localhost", "8000")
	broker.AddService(&server, broker.InitService("A", "localhost", "8010"))
	broker.AddService(&server, broker.InitService("B", "localhost", "8020"))
	return server
}

func Run(mb broker.MessageBroker) {
	for {
		request, connection := mb.ReceiveRequest()
		taskArr, ackArr := mb.ProcessRequest(request)

		for idx, task := range taskArr {
			go func(task broker.Task, ackArr [][]byte, idx int) {
				mb.SendTask(task)
				ackArr[idx] = mb.ReceiveAck(task)
			}(task, ackArr, idx)
		}

		mb.SendResponse(connection, ackArr)
	}
}

func main() {
	server := Setup()
	Run(server)
}
