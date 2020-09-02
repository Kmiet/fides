FROM golang:alpine as builder

ARG service
RUN mkdir /build 
COPY go.mod go.sum /build/
COPY cmd /build/cmd 
COPY internal /build/internal
COPY services /build/services
WORKDIR /build 

RUN go build -o main /build/cmd/$service/main.go

FROM alpine
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]