FROM golang:1.12-alpine3.9 as go

ENV GOBIN /go/bin

RUN mkdir /app

RUN apk add git

COPY ./src /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/main .

#######################################
FROM busybox

RUN mkdir /app
WORKDIR /app

COPY --from=go /app/main /app/main

EXPOSE 8000/tcp

CMD ["/app/main"]
