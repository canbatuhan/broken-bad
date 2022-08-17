import base64
import time
from typing import List
from service import Service

SERVICE_HOST = "127.0.0.1"
SERVICE_PORT = 8020

class PowerCalculatorService(Service):
    def __init__(self, host, port):
        super().__init__("PowerCalculatorService", host, port)
        
    def serve(self, data_arr:List[bytes]) -> int:
        if len(data_arr) != 2:
            return -1
        else:
            base = int(data_arr[0].decode())
            power = int(data_arr[1].decode())
            result = 1
            for _ in range(power):
                result *= base
            return result

if __name__ == "__main__":
    PowerCalculatorService(SERVICE_HOST, SERVICE_PORT).run()