FROM golang:1.16

RUN mkdir -p /app/server
WORKDIR /app/server

COPY go.mod /app/server/
COPY go.sum /app/server/
RUN go mod download

COPY . .

EXPOSE 4000
# binary will be $(go env GOPATH)/bin/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]