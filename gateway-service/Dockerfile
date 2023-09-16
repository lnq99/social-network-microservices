FROM golang:alpine as build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api ./main.go

FROM alpine as production-stage
WORKDIR /app
COPY --from=build-stage /app/api /app/config.yaml ./
EXPOSE 8080
CMD ./api
