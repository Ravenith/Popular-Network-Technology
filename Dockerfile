FROM golang:1.23.2-bookworm

WORKDIR /app
COPY .. /app/
RUN go get
RUN go build -o server server.go


ENTRYPOINT [ "./server" ]

FROM golang:1.23.2-bookworm

WORKDIR /app
COPY .. /app/
RUN go get
RUN go build -o client client.go

ENTRYPOINT [ "./client" ]