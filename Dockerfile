FROM golang:1.20.4 as build-stage

RUN mkdir /build
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go 

FROM alpine:3.18

WORKDIR /root/

COPY --from=build-stage /build/main .
COPY --from=build-stage /build/.env .

EXPOSE 8080

CMD ["./main"]