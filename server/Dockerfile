FROM golang:1.16-alpine

RUN apk add --no-cache git

WORKDIR /server_app

COPY . .

RUN go mod download && \
    cd logServer/ && \
    go build
    
EXPOSE 50051

CMD [ "/server_app/logServer/logServer" ]
