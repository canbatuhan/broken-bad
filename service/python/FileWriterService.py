import time
from typing import List
from service import Service

SERVICE_HOST = "127.0.0.1"
SERVICE_PORT = 8010

class FileWriterService(Service):
    def __init__(self, host, port):
        super().__init__("FileWriterService", host, port)
        
    def serve(self, data_arr:List[bytes]) -> None:
        if len(data_arr) != 2:
            return -1
        else:
            name, content = data_arr[0].decode(), data_arr[1].decode()
            file = open(name, 'w')
            file.write(content)

if __name__ == "__main__":
    FileWriterService(SERVICE_HOST, SERVICE_PORT).run()