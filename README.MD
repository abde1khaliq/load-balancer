## Load Balancer Made in Golang

Made to learn and simulate how a real Load Balancer would behave like.

## Features

- **Health Checks**: Always checks servers' health
- **Round Robin**: Algorithm for Load Balancing

## whilst building this project

![Vibe](https://ik.imagekit.io/cin2tn3bj/screeny.jpeg)


## How to run the project locally

* Make sure you have python and go installed on your machine, then:

1. Clone the repo
```bash
git clone https://github.com/abde1khaliq/load-balancer
cd load-balancer
```

2. Install Go Packages
```bash
go mod tidy
```

3. Install dependencies for test servers:
```bash
cd servers/
pip install -r requirements.txt
```

4. Start the test servers:
* Test server ports are hardcoded in `main.go` but you can change them.
```bash
uvicorn server:app --port 8000 --reload &
uvicorn server:app --port 8001 --reload &
uvicorn server:app --port 8002 --reload &
uvicorn server:app --port 8003 --reload &
```

5. Set the algorithm in `main.go` (Round Robin by default):
```go
patterns["roundrobin"] = true
```

6. Run the load balancer
```bash
go run main.go
```

7. Send requests to the load balancer NOT the test servers:
```bash
curl http://localhost/message
```

- You can also run concurrent requests to test the load balancing using the following command:
```bash
for i in {1..10}; do curl http://localhost:8080/message & done
wait
```

# And You're Done!