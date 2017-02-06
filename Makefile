export GOPATH
TARGET := $(CURDIR)/bin

all:
	@echo "start building..."
	GOOS=darwin GOARCH=amd64 go build -o "$(TARGET)/darwin"
	GOOS=linux GOARCH=amd64 go build -o "$(TARGET)/linux"
	GOOS=windows GOARCH=amd64 go build -o "$(TARGET)/windows.exe"
	@echo "finish"


