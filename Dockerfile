FROM golang:1.18-alpine
ENV PORT=8080
ENV GIN_MODE=release

RUN apk update && apk upgrade && apk add --no-cache bash git
WORKDIR /var/www/rawleydotxyz/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o rawleydotxyz .

EXPOSE $PORT
CMD ["./rawleydotxyz"]
