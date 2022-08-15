import socket
from constants import *

class Client:
    def __init__(self) -> None:
        self.__socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        self.__n_task = 0

    def __generate_task(self, service, *args):
        cmd = service + HEAD_BODY_SEPERATOR
        for idx, each in enumerate(args):
            cmd += str(each)
            if idx == len(args)-1:
                continue
            cmd += CONTENT_SEPERATOR
        return cmd

    def __generate_req(self, *args):
        return TASK_SEPERATOR.join(args), len(args)

    def __send_request_to(self, host, port):
        request, self.__n_task = self.__generate_req(
            self.__generate_task("A", 42),
            self.__generate_task("B", 42))
        self.__socket.sendto(request.encode('utf-8'), (host, port))

    def __receive_response(self):
        for _ in range(self.__n_task):
            response, addr = self.__socket.recvfrom(2048)
            print(response.decode())

    def run(self):
        HOST, PORT = "127.0.0.1", 8000
        self.__send_request_to(HOST, PORT)
        self.__receive_response()
        
        
if __name__ == "__main__":
    client = Client()
    client.run()