# STAGE 1
FROM golang AS build

WORKDIR /violin-notice

COPY . /violin-notice

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /violin-notice/common

RUN go mod tidy

WORKDIR /violin-notice/violin-notice

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o grpc-notice-server

# STAGE 2
FROM alpine as work

## wo can change it from to debug, wo should redefine it in kuberntes env.
ENV GIN_MODE=release

WORKDIR /violin-notice/violin-notice

COPY --from=build /violin-notice/violin-notice/grpc-notice-server ./grpc-notice-server

## we should comment this record, and mount config file with kubernetes pv or configMap
COPY --from=build /violin-notice/violin-notice/config ./config

# notice grpc service port
EXPOSE 8081

EXPOSE 80

ENTRYPOINT ["./grpc-notice-server"]
