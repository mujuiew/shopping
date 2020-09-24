FROM golang:1.14 as builder

WORKDIR /src

COPY go.* ./

#BUILD APP
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./cmd/ .

ENV PORT 8080

CMD ["./main"]