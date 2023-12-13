FROM golang:1.21-alpine as builder
WORKDIR /
COPY . ./
RUN go mod download


RUN go build -o /hello-world


FROM alpine
COPY --from=builder /hello-world .


EXPOSE 80
CMD [ "/hello-world" ]