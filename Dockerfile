FROM golang:1.16-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

RUN mkdir -p /server/conf
WORKDIR /server

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /server/main ./main.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata curl
ENV TZ Asia/Shanghai

RUN mkdir -p /app/conf && mkdir -p /app/logs
WORKDIR /app
COPY --from=builder /server/main /app/main
COPY --from=builder /server/conf /app/conf
COPY --from=builder /server/logs /app/logs
EXPOSE 7890
CMD ["./main"]
