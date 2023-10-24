FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY ./source .
RUN go mod download
RUN go build -o /app/bin/program

FROM alpine:latest
ENV TZ="Asia/Bangkok"
RUN apk --no-cache add tzdata
RUN addgroup -S golang && adduser -S golang -G golang
RUN mkdir -p /app && chown golang:golang /app
WORKDIR /app
USER golang
COPY --from=builder --chown=golang:golang /app/bin/program /app/program
COPY --from=builder --chown=golang:golang /app/settings /app/settings
EXPOSE 8000
CMD [ "/app/program"]
