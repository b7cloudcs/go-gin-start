FROM golang as builder

ENV GO111MODULE=on
# ENV GOPROXY=https://goproxy.cn,direct

# build
WORKDIR /workdir
COPY . .
RUN go mod download
RUN go generate ./ent
RUN CGO_ENABLED=0 go build -o main ./main.go

# -------

FROM alpine

# set timezone
ADD zoneinfo.tar.gz /
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# file copy
WORKDIR /workdir
RUN pwd
RUN ls -al
COPY --from=builder /workdir/app.ini .
COPY --from=builder /workdir/main .

# run
EXPOSE 8080
CMD ["/workdir/main"]