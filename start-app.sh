# kill port 8888 if it is already in use
lsof -i:8887 | grep LISTEN | awk '{print $2}' | xargs kill -9
go run cmd/apiserver.go apiserver --port=8887