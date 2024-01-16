d:
	@go run server
r:
	@go build server
	./server
	rm -f ./server
