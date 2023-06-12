FROM golang:1.20 as builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o app .

FROM scratch

COPY --from=builder /app/app /app

ENTRYPOINT [ "/app" ]