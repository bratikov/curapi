FROM golang:1.24.2-alpine3.21
WORKDIR /go/src
RUN apk add git openssh bash
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
#CMD ["go", "run", "cmd/pastor/main.go", "run", "-c", "config/pastor.json"]
CMD ["sleep", "infinite"]