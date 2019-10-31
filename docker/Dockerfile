FROM golang:latest

# Environments
ENV GO111MODULE on
ENV GOPATH /go
ENV TORCH_API_PORT 2800

# Initialize VM Working Directory
WORKDIR /go/src/github.com/swkoubou/torch/server
ADD ./server /go/src/github.com/usagiga/torch/server

# External Port
EXPOSE ${TORCH_API_PORT}

# Run
CMD ["make", "run"]
