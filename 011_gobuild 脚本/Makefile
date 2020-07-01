# Go parameters
GO_CMD = go

GO_BUILD_CROSS_PRE =  CGO_ENABLED=0  GOARCH=amd64
GO_BUILD = $(GO_CMD) build
GO_TEST = $(GO_CMD) test

BINARY_NAME = shengjing
BINARY_NAME_WIN = shengjing_win.exe
BINARY_NAME_LINUX = shengjing_linux
BINARY_NAME_MAC = shengjing_darwin

all:  build win linux darwin
build:
	@echo "---------------------------------->编译为all<--------------------------------------------"
	$(GO_BUILD) -a -v -o $(BINARY_NAME)
win:
	@echo "---------------------------->编译为build_win<------------------------------------------"
	$(GO_BUILD_CROSS_PRE) GOOS=windows $(GO_BUILD) -a -v -o $(BINARY_NAME_WIN)
	@echo "...build success!!! $(BINARY_NAME_WIN)编译成功..."
linux:
	@echo "---------------------------->编译为build_linux<------------------------------------------"
	$(GO_BUILD_CROSS_PRE) GOOS=linux $(GO_BUILD) -a -v -o $(BINARY_NAME_LINUX)
	@echo "...build success!!! $(BINARY_NAME_WIN)编译成功..."
darwin:
	@echo "---------------------------->编译为build_darwin<------------------------------------------"
	$(GO_BUILD_CROSS_PRE) GOOS=darwin $(GO_BUILD) -a -v -o $(BINARY_NAME_MAC)
	@echo "...build success!!! $(BINARY_NAME_MAC)编译成功..."



