go build:
	go build server.go
	
docker image:
	docker build -t golang/xuedev.grpc .
	
docker run:
	docker run -itd --name go.rpc -p 8888:8888 golang/xuedev.grpc
