package main

import broker "broker/src"

func Setup() broker.Server {
	server := broker.InitServer("udp", "127.0.0.1", 8000)
	broker.AddService(&server, broker.InitService("A", "127.0.0.1", 8010))
	broker.AddService(&server, broker.InitService("B", "127.0.0.1", 8020))
	return server
}

func Run(mb broker.MessageBroker) {
	for {
		request, clientConnection := mb.ReceiveRequest()
		taskArr, ackArr := mb.ProcessRequest(request)

		for idx, task := range taskArr {
			go func(task broker.Task, ackArr [][]byte, idx int) {
				mb.SendTask(task)
				ackArr[idx] = mb.ReceiveAck(task)
			}(task, ackArr, idx)
		}

		mb.SendResponse(clientConnection, ackArr)
	}
}

func main() {
	server := Setup()
	Run(server)
}
