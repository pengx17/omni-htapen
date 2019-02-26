FROM golang

WORKDIR /

# Avoid reunning `go mod`
COPY go.sum go.mod ./

# Cache go dependencies for go1.11 and npm deps
RUN go mod download

COPY . .

# Build
RUN GOOS=linux CGO_ENABLED=0 go build -o ./server .
