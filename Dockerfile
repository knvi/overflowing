FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy all source code (stored in src/ directory)
COPY /src/ .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

EXPOSE 4321

CMD ["/server"]