FROM golang:1.20

WORKDIR /app

COPY go.* ./
RUN go mod download

# Copy all source code (stored in src/ directory)
COPY . ./

# Build the Go app
RUN go build -o /server

EXPOSE 4321

CMD ["/server"]