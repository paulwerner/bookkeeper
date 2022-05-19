FROM golang:1.18.1-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GARCH=amd64

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/bookkeeper-api/main.go

WORKDIR /dist

RUN cp /build/main .
RUN cp -r /build/ops.migrations/ .
RUN cp /build/app.env .

FROM scratch

COPY --from=builder /dist/ /
EXPOSE 8080
ENTRYPOINT [ "./main" ]