FROM golang:1.17-alpine

WORKDIR /app

COPY . ./

ARG VERSION
ENV VERSION=$VERSION

RUN echo "building version $VERSION" &&\
    go build -o skadi

EXPOSE 5000

CMD ["./skadi"]