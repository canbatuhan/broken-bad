import socket

def serve(num):
    return "[SERVICE-A] Served in Python, result={}".format(num*2)

if __name__ == "__main__":
    service = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    service.bind(("localhost", 8010))

    print("[SERVICE] Started on {}".format(service.getsockname()))

    while True:
        pair = service.recvfrom(1024)
        data, addr = pair[0], pair[1]
        print("[SERVICE] Received from: {}".format(addr))
        print("[SERVICE] Data received: {}".format(data))
        if not addr:
            break
        response = serve(int(data))
        service.sendto(response.encode(), addr)
        print("[SERVICE] Response send: {}".format(response))

    service.close()
    print("[SERVICE] Closed.")