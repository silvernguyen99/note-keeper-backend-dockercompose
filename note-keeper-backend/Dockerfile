FROM golang:1.13.4-alpine3.10 as builder
WORKDIR /app
RUN apk update && apk upgrade && \
    apk add bash git openssh gcc libc-dev
#COPY ./go.mod ./go.sum ./
COPY ./ ./
RUN go mod download && go build /app/cmd/server

FROM alpine:3.9

RUN apk add --update ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
    echo "Asia/Ho_Chi_Minh" > /etc/timezone && \
    rm -rf /var/cache/apk/*

COPY --from=builder /app/server /app/server
COPY wait-for-postgres.sh /app
RUN chmod +x /app/wait-for-postgres.sh && \
    apk add --no-cache postgresql-client

WORKDIR /app
CMD ["/app/server"]
