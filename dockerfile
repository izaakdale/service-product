# build a tiny docker image
FROM alpine:latest 

RUN mkdir /app
# requires a built binary
COPY service-product /app

CMD ["/app/service-product"]