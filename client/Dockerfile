# Build environment

## To build use next command:
## $ docker build -t trans_client . && docker run --rm --privileged -v /dev/video0:/dev/video0 trans_client:latest

FROM denismakogon/gocv-alpine:3.4.2-buildstage

MAINTAINER Roman Kaporin <rmkaporin@gmail.com>

WORKDIR /go/src/trans/client/
ADD . /go/src/trans/client/

EXPOSE 8080
EXPOSE 8081

CMD ["go", "run", "main.go"]