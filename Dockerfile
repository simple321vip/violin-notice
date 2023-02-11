# STAGE 1
FROM golang AS build

ENV location /go/src/github.com/grpc-up-and-running/samples/ch07/grpc-docker/go
WORKDIR ${location}/server

ADD ./server ${location}/server
ADD ./proto-gen ${location}/proto-gen

RUN go get -d ./...
RUN go install ./...

RUN CGO_ENABLED=0 go build -o /bin/grpc-productinfo-server

# STAGE 2
FROM scratch
COPY --from=build /bin/grpc-productinfo-server /bin/grpc-productinfo-server
ENTRYPOINT ["/bin/grpc-productinfo-server"]
