build:
	go mod tidy
	go build -o stl_generator -v generator_main.go

clean:
	rm -f stl_generator
	go clean ./..
