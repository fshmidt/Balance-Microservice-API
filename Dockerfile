
FROM golang:1.19-buster AS build

RUN go version
ENV GOPATH=/

WORKDIR /app

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client


# build go app
RUN go mod download
RUN go build -o /balance-app ./cmd/main.go


## Deploy

FROM gcr.io/distroless/base-debian11

ENV GO111MODULE=on

WORKDIR /

COPY --from=build /balance-app /balance-app

EXPOSE 8000

USER root:root

ENTRYPOINT ["./balance-app"]
