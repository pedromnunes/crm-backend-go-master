FROM golang:latest

# Working directory
WORKDIR /go/src/app

# Copy files
COPY . .

# Run 
RUN go get github.com/gorilla/mux
RUN go get github.com/google/uuid

# Expose Ports
EXPOSE 8080

# Commands
CMD [ "go","run","main.go" ]
