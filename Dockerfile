FROM golang:1.24.2-bookworm as builder

ARG GOPROXY=direct
ENV GOPROXY=${GOPROXY}

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o go-api ./adapter/http/main.go

FROM scratch
COPY --from=builder /app/go-api /
COPY --from=builder /app/database/migrations /database/migrations
EXPOSE 3000
CMD ["/go-api"]
