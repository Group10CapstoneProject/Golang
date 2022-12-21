FROM golang:1.18-alpine3.16 as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -o main.app ./app/rest
RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -o cron.app ./app/cron

FROM alpine:latest

# set image timezoe
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

COPY --from=builder /app/* .

CMD ["./main.app"]