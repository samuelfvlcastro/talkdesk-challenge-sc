FROM golang:1.12 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR /go/src/gitlab.com/samuelfvlcastro/talkdesk-challenge-sc
COPY . ./
RUN dep ensure --vendor-only

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app/phoneval main.go

FROM alpine:3.7
RUN apk add ca-certificates
RUN apk add dumb-init
COPY --from=builder /app /usr/local/bin
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
