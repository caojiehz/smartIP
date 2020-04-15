FROM golang:1.12.12 as build

ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release
ADD . .

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -a -o smartIP main.go

FROM scratch as prod

COPY --from=build /go/release/smartIP /

CMD ["/smartIP"]