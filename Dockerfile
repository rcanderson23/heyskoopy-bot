FROM golang:1.16 as builder

ENV UID=10001
ENV GOOS=linux
ENV CGO_ENABLED=0

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "$UID" \
    "bot"

RUN apt update && apt install git mercurial gcc -y
WORKDIR /tmp/app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN go build -ldflags='-w -extldflags "-static"' -o /bot ./cmd/heyskoopy/

# Application
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /bot /bot
USER bot:bot
ENTRYPOINT ["/bot"]
