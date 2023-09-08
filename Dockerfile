FROM golang:alpine3.18 
WORKDIR /go/bin
COPY main . 
EXPOSE 3000
CMD ["main"]