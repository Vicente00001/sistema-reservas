FROM golang:alpine AS build
RUN apk add --no-cache git ca-certificates
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/server ./cmd/server

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /app/server ./server
COPY static ./static
COPY migrations ./migrations
COPY sqlc.yaml ./sqlc.yaml
EXPOSE 8080
ENV PORT=8080
CMD ["./server"]
