# build stage
FROM golang:alpine AS build-env
ADD . /go/src/github.com/pratikjagrut
RUN apk update
RUN apk add -q git dep 
RUN cd /go/src/github.com/pratikjagrut && dep ensure && go build -o app

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/pratikjagrut/app /app/
ENTRYPOINT ./app