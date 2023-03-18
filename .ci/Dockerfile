FROM golang:1.20.2-alpine3.16

WORKDIR /usr/src/app

ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

COPY . .

RUN go mod init github.com/leonardo-multicloud/golang-rest-api && go mod tidy
RUN go test -v ./api/

RUN go build -o /api

CMD [ "/api" ]