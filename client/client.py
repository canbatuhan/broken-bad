import socket
import time

def run():
    client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    message = "A:42$B:42"
    client.sendto(message.encode('utf-8'), ("127.0.0.1", 8000))
    for _ in message.split("$"):
        response = client.recvfrom(2048)
        print(response)
    client.close()

if __name__ == "__main__":
    run()