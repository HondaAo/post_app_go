FROM golang:1.18

RUN mkdir -p /app/server
WORKDIR /app/server

COPY ./server/go.mod .
COPY ./server/go.sum .

RUN go mod download
RUN go mod tidy

COPY ./server .

EXPOSE 4000
# binary will be $(go env GOPATH)/bin/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]