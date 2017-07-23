FROM golang:1.7
EXPOSE 8080
RUN mkdir -p /usr/local/go/src/github.com/nehayward/dos/
ADD . /usr/local/go/src/github.com/nehayward/dos/
WORKDIR /usr/local/go/src/github.com/nehayward/dos/
RUN /usr/local/go/src/github.com/nehayward/dos/start.sh
CMD ["/usr/local/go/src/github.com/nehayward/dos/main"]
