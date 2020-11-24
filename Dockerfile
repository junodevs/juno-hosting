FROM golang:1.15.5-alpine3.12

RUN mkdir /opt/hosting-server

COPY . /opt/hosting-server
WORKDIR /opt/hosting-server

# Download dependencies
RUN go mod download

# Build the binary
RUN go build --ldflags "-s -w" -o bin/hosting-server ./

# Expose port 8080
EXPOSE 8080

# Run the generated binary
CMD ["/opt/hosting-server/bin/hosting-server"]
