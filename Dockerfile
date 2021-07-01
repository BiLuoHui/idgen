ARG BUILDER_IMAGE=golang:alpine
############################
# step 1: build executeable binary
############################
FROM ${BUILDER_IMAGE} AS builder

LABEL version="1.0"
LABEL description="全局唯一ID生成器"
MAINTAINER BiLuoHui(junlingorg@gmail.com)

# install git
RUN apk update && \
    apk add --no-cache git ca-certificates tzdata && \
    update-ca-certificates

# create user
ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go mod download && \
    go mod verify && \
    cd cmd/server && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/idgen .

############################
# STEP 2 build a small image
############################
FROM scratch

# timezone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# cert
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# import user and group files from builder
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# copy static executable
COPY --from=builder /go/bin/idgen /go/bin/idgen

# use unprivileged user
USER appuser:appuser

EXPOSE 8081

ENTRYPOINT ["/go/bin/idgen", "1", "8081"]