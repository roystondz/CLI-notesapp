build:
	@go build -o terminotes .

run: build
	@./terminotes