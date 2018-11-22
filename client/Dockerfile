# Build environment

## To build use next command:
## $ docker build -t trans_client . && docker run -t -i trans_client:latest

FROM denismakogon/gocv-alpine:3.4.2-buildstage

MAINTAINER Roman Kaporin <rmkaporin@gmail.com>

WORKDIR /go/src/trans/client/
ADD . /go/src/trans/client/

RUN go build  -a -o _tclient main.go