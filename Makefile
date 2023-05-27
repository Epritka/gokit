.PHONY: http
http:
	go run ./example/cmd/http/main.go
	
.PHONY: sandbox
sandbox:
	go run ./example/cmd/sandbox/main.go