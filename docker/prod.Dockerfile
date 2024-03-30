FROM golang:alpine3.19
WORKDIR /app
ADD ./src /app/
RUN go build

FROM alpine:3.19
COPY --from=0 /app/seed /opt
CMD ["/opt/seed"]
