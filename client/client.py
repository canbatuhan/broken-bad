from datetime import datetime
import socket
import time
from constants import *

SAMPLE_FILE_PATH = "C:/Users/Batuhan/Desktop/slime_.txt"
SAMPLE_FILE_CONTENT = """Well, I am the slime from your video\nOozin' along on your livin' room floor"""
SAMPLE_BASE = 2
SAMPLE_POWER = 32


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
            self.__generate_task("FileWriterService", SAMPLE_FILE_PATH, SAMPLE_FILE_CONTENT),
            self.__generate_task("PowerCalculatorService", SAMPLE_BASE, SAMPLE_POWER))
        self.__socket.sendto(request.encode(ENCODING), (host, port))
        self.__client_log("Request including {} tasks sent to broker".format(self.__n_task))

    def __receive_response(self) -> None:
        for _ in range(self.__n_task):
            ack, addr = self.__socket.recvfrom(BUFFER_SIZE)
        self.__client_log("Response received from broker")

    def __stop(self) -> None:
        self.__socket.close()

    def run(self) -> None:
        HOST, PORT = "127.0.0.1", 8000
        self.__send_request_to(HOST, PORT)
        self.__receive_response()
        self.__stop()

if __name__ == "__main__":
    Client().run()