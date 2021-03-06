FROM golang:1.15.5-alpine3.12

RUN mkdir -p /opt/juno-hosting/server

COPY . /opt/juno-hosting/server
WORKDIR /opt/juno-hosting/server

# Download dependencies
RUN go mod download

# Build the binary
RUN go build --ldflags "-s -w" -o bin/juno-hosting ./

# Expose port 8080
EXPOSE 8080

# Run the generated binary
CMD ["/opt/juno-hosting/server/bin/juno-hosting"]
