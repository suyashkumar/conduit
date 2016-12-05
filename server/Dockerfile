FROM golang

ADD . /go/src/github.com/suyashkumar/home-automation/server

# Get dependencies and build
RUN go get github.com/suyashkumar/home-automation/server
RUN go install github.com/suyashkumar/home-automation/server

ENV PORT 8080 
ENTRYPOINT /go/bin/server

EXPOSE 8080 
EXPOSE 27017
