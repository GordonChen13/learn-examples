FROM golang
WORKDIR /go/src/app
COPY . .
RUN go test