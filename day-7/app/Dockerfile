FROM golang:1.19 as go_builder
LABEL stage="binary_builder"
WORKDIR /appTemp
COPY . .
WORKDIR /appTemp/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

# Multistage
FROM alpine:3.14
LABEL stage="binary_exec"
RUN apk add --no-cache ca-certificates
COPY --from=go_builder /appTemp/main /app/main
COPY --from=go_builder /appTemp/.env /app/.env
WORKDIR /app/