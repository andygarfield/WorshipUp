FROM node:latest
ENV GOPATH /go

RUN mkdir /go /go/src /go/pkg /go/bin
RUN apt-get update
RUN apt-get -y install golang
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
# Go dependencies
RUN go get github.com/andygarfield/worshipup/...
RUN go get github.com/NYTimes/gziphandler
RUN go get github.com/boltdb/bolt/...

WORKDIR /go/src/github.com/andygarfield/worshipup

RUN yarn
CMD ["yarn", "run", "test"]