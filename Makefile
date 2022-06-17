.PHONY: all

all:
	go build -v ./exe/buntrepair.go
	go build -v ./exe/buntdump.go
