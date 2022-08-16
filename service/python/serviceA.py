import time
from service import Service

class ServiceA(Service):
    def __init__(self, host, port):
        super().__init__("A", host, port)
        
    def serve(self, data_arr) -> int:
        if len(data_arr) != 1:
            return -1
        else:
            input_data = int(data_arr[0])
            time.sleep(3)
            return input_data*2

if __name__ == "__main__":
    ServiceA("127.0.0.1", 8010).run()