import requests as req
import random
import json

ROOT_URL = "http://127.0.0.1:8000"

def generate_random_floats(n, start, end):
    return [random.uniform(start, end) for _ in range(n)]

class ServerClient:
    def get_hello(self):
        url = "http://127.0.0.1:8000/hello"
        rcv = req.get(url=url)
        print(rcv.content.decode())

    def get_counter(self):
        url = "http://127.0.0.1:800/counter"
        rcv = req.get(url=url)
        print(rcv.content)

    def post_data(self):
        list = generate_random_floats(24, 1.0, 100)
        url = ROOT_URL + "/senddata"
        json_data = json.dumps(list)
        rcv = req.post(url=url, data=json_data, headers={'Content-Type': 'application/json'})
        print(rcv.content)


if __name__ == "__main__":
    server = ServerClient()
    # server.get_hello()
    for i in range(100):
        server.post_data()
