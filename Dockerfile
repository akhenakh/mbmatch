FROM golang:1.13-alpine as builder
RUN apk add --no-cache git gcc musl-dev
COPY . /workdir
RUN go get github.com/gobuffalo/packr/packr
RUN cd /workdir/cmd/mbmatch && $GOPATH/bin/packr build

FROM gcr.io/distroless/base
WORKDIR /root/
COPY --from=builder /workdir/cmd/mbmatch/mbmatch .
COPY --from=builder /workdir/cmd/mbmatch/hawaii.mbtiles .
EXPOSE 8000
ENTRYPOINT ["./mbmatch"]

