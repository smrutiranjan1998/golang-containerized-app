
FROM golang:1.23.3


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN GOOS=linux GOARCH=amd64 go build -o main .
RUN ls -la /app  # Debug: Confirm the binary 'main' exists


EXPOSE 8080


CMD ["./main"]