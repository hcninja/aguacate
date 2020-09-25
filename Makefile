# Config
BINARY=aguacate
TARGET=all
LDFLAGS=-ldflags="\
	-s \
	-w"


.DEFAULT_GOAL: $(BINARY)

.PHONY: all
$(TARGET):
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}_darwin
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}_linux
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}.exe
	tar czvf ${BINARY}_darwin.tgz ${BINARY}_darwin
	tar czvf ${BINARY}_linux.tgz ${BINARY}_linux
	zip -9 ${BINARY}_windows.zip ${BINARY}.exe

.PHONY: macos
macos:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}

.PHONY: windows
windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}.exe

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}

.PHONY: install
install:
	mv ${BINARY} ${GOBIN}

.PHONY: clean
clean:
	rm -rf ${BINARY}.exe ${BINARY}_darwin ${BINARY}_linux 
	rm -rf ${BINARY}_darwin.tgz ${BINARY}_linux.tgz ${BINARY}_windows.zip
