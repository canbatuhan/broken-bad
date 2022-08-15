import socket

def run():
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect(("localhost", 8000))
    message = input("Your message: ")
    client.send(message.encode())
    client.close()

if __name__ == "__main__":
    run()