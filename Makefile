SERVER_FILE=cmd/server/main.go
SERVER_DEST=bin/blog-web

all: build

build:
	go build -o ${SERVER_DEST} ${SERVER_FILE}

run:
	go build -o ${SERVER_DEST} ${SERVER_FILE}
	${SERVER_DEST}

# test:
# 	go test -v ./...

clean:
	rm -f ${SERVER_DEST}