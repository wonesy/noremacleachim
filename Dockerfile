FROM golang:1.15.5-alpine AS builder

WORKDIR /src

COPY . .

RUN go build -o /out/noremac .

FROM alpine:latest

COPY --from=builder /out/noremac /

CMD ["/noremac"]