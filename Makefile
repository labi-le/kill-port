build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'

install:build
	sudo mv kill-port /usr/local/bin/

uninstall:
	sudo rm /usr/local/bin/kill-port