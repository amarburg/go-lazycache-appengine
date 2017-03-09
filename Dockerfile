FROM amarburg/golang-ffmpeg:wheezy-1.8

RUN go get github.com/amarburg/go-lazycache-app
RUN go install github.com/amarburg/go-lazycache-app

CMD ["/go/bin/go-lazycache-app", "--port", "8080"]
EXPOSE 8080
