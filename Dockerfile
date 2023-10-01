FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN apk --no-cache add curl

RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh

CMD go run main.go

EXPOSE 8080

# RUN go build -o /money-tracking-app

# CMD [ "/money-tracking-app" ]