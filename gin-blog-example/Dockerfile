#FROM golang:latest
FROM scratch

# ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/gin-blog-example
COPY . $GOPATH/src/gin-blog-example
# RUN go build .

EXPOSE 8000
CMD ["./gin-blog-example"]