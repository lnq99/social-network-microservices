FROM golang:alpine as build-stage
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o api ./main.go

FROM alpine as production-stage
WORKDIR /app
COPY --from=build-stage /app/config.yaml ./
COPY --from=build-stage /app/api ./
EXPOSE 80
CMD ./api