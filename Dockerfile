FROM golang:1.11-alpine3.8 AS builder
RUN apk add --no-cache make git
WORKDIR /go/src/github.com/kakakakakku/env-viewer/
COPY . .
RUN go build

FROM alpine:3.8
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/kakakakakku/env-viewer/env-viewer .
ENTRYPOINT ["./env-viewer"]
