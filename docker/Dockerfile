FROM golang:latest

# Environments
ENV GO111MODULE on
ENV GOPATH /go
ENV TORCH_API_PORT 2800
ENV TORCH_MYSQL_DB torch
ENV TORCH_MYSQL_HOST db
ENV TORCH_MYSQL_PORT 3306
ENV TORCH_MYSQL_USER torch
ENV TORCH_MYSQL_PASS torch-password

# Initialize VM Working Directory
WORKDIR /go/src/github.com/swkoubou/torch/server
ADD ./server /go/src/github.com/swkoubou/torch/server

# External Port
EXPOSE ${TORCH_API_PORT}

# Run
CMD ["make", "run"]
