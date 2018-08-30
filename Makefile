all: get-rice embed

embed:
	${GOPATH}/bin/rice embed-go -i .

get-rice:
	go get -u github.com/GeertJohan/go.rice
	go get -u github.com/GeertJohan/go.rice/rice
