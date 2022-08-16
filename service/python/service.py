import abc
from datetime import datetime
import socket
from typing import Any, Tuple
from constants import *

class Service:
    def __init__(self, name, host, port) -> None:
        self.__name = name
        self.__host = host
        self.__port = port
        self.__socket = socket.socket(
            socket.AF_INET, socket.SOCK_DGRAM)

    @abc.abstractmethod
    def serve(self, *args):
        pass

    def __service_log(self, message:str) -> None:
        print("[SERVICE] {} - {}".format(
            datetime.now().strftime("%Y/%m/%d %H:%M:%S.%f"), message))

    def __startup(self) -> None:
        self.__socket.bind((self.__host, self.__port))
        self.__service_log("Service {} started on {}:{}".format(
            self.__name, self.__host, self.__port))

    def __receive_data(self) -> Tuple[bytes, Any]:
        data, address = self.__socket.recvfrom(BUFFER_SIZE)
        data = data.decode().split(CONTENT_SEPERATOR)
        self.__service_log("Data received from broker")
        return data, address

    def __send_acknowledgement(self, *service_results, address) -> None:
        service_results = CONTENT_SEPERATOR.join(str(each) for each in service_results)
        self.__socket.sendto(service_results.encode('utf-8'), address)
        self.__service_log("Acknowledgement sent to broker")

    def run(self):
        self.__startup()
        while True:
            data_arr, remote_address = self.__receive_data()
            results = self.serve(data_arr)
            self.__send_acknowledgement(results, address=remote_address)