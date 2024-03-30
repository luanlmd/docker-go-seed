FROM golang:alpine3.19

RUN go install github.com/HenriBeck/gowatch@latest

WORKDIR /src

CMD ["gowatch"]