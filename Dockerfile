FROM golang AS build
ADD . /go/src/github.com/special-tactical-service/slotlist
WORKDIR /go/src/github.com/special-tactical-service/slotlist
RUN apt update && \
	apt upgrade -y && \
	apt install curl -y
RUN curl -sL https://deb.nodesource.com/setup_8.x -o nodesource_setup.sh && bash nodesource_setup.sh
RUN apt-get install -y nodejs

# build backend
ENV GOPATH=/go
RUN go build -ldflags "-s -w" main.go

# build frontend
RUN cd /go/src/github.com/special-tactical-service/slotlist/public && npm i && npm rebuild node-sass && npm run build

FROM ubuntu
COPY --from=build /go/src/github.com/special-tactical-service/slotlist /app
WORKDIR /app

# default config
ENV STS_SLOTLIST_LOGLEVEL=info
ENV STS_SLOTLIST_WATCH_BUILD_JS=false
ENV STS_SLOTLIST_ALLOWED_ORIGINS=*
ENV STS_SLOTLIST_HOST=0.0.0.0:8080

CMD ["/app/main"]
