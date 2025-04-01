FROM golang:1.19-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app .
ENV PORT=8080
CMD ["./app"]