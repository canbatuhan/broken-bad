from datetime import datetime
import socket
from constants import *

class Client:
    def __init__(self) -> None:
        self.__socket = socket.socket(
            socket.AF_INET, socket.SOCK_DGRAM)
        self.__n_task = 0

    def __client_log(self, message:str) -> None:
        print("[CLIENT] {} - {}".format(
            datetime.now().strftime("%Y/%m/%d %H:%M:%S.%f"), message))

    def __generate_task(self, service_name:str, *args) -> str:
        cmd = service_name + HEAD_BODY_SEPERATOR
        cmd += CONTENT_SEPERATOR.join(str(each) for each in args)
        return cmd

    def __generate_req(self, *args) -> str:
        return TASK_SEPERATOR.join(args), len(args)

    def __send_request_to(self, host:str, port:int) -> None:
        request, self.__n_task = self.__generate_req(
            self.__generate_task("A", 89),
            self.__generate_task("B", 51))
        self.__socket.sendto(request.encode('utf-8'), (host, port))
        self.__client_log("Request including {} tasks sent to broker".format(self.__n_task))

    def __receive_response(self) -> None:
        for _ in range(self.__n_task):
            response, addr = self.__socket.recvfrom(2048)
        self.__client_log("Response received from broker")

    def run(self) -> None:
        HOST, PORT = "127.0.0.1", 8000
        self.__send_request_to(HOST, PORT)
        self.__receive_response()

if __name__ == "__main__":
    client = Client()
    client.run()