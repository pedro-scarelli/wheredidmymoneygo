FROM golang:1.24.2-bookworm as builder
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api /adapter/http/main.go

FROM scratch
WORKDIR /
COPY --from=builder /go-api .
COPY --from=builder /database/migrations /database/migrations
EXPOSE 3000
CMD ["./go-api"]
