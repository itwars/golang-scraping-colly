FROM golang
WORKDIR /go/src/github.com/user/app
COPY . .

ADD https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz /usr/local
RUN set -x && \
    apt update && \
    apt install -y xz-utils && \
    xz -d -c /usr/local/upx-3.95-amd64_linux.tar.xz | \
    tar -xOf - upx-3.95-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app . && \
    strip --strip-unneeded app && \
    upx app

FROM scratch
WORKDIR /root/
COPY --from=0 /go/src/github.com/user/app .
EXPOSE 3000
CMD ["./app"]