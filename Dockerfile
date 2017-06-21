FROM golang:1.8
EXPOSE 8080
WORKDIR /go/src/app
COPY . .
RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"]
