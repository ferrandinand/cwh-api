FROM golang:alpine as builder

RUN mkdir /build 
ADD . /build/
WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o api .

FROM alpine:3
RUN apk --update add git less openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*
RUN ls -la
COPY --from=builder /build/api /app/
WORKDIR /app
CMD ["./api"]