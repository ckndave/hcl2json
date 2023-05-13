FROM golang:1.18.0
#RUN apt-get update -y
RUN apt-get install ca-certificates libgnutls30 -y
RUN chmod 755 /etc /etc/ssl /etc/ssl/certs
RUN chmod 644 /etc/ssl/certs/ca-certificates.crt
# Setup
WORKDIR /build
COPY go.* ./
COPY *.go ./
COPY ./convert/*.go ./convert/
#RUN ls -R
#ENV GOPATH=$HOME/go
#RUN PATH=$PATH:$GOPATH/bin
##ENV GOROOT=$GOPATH
#ENV GOPHERJS_GOROOT=$GOPATH/bin
# # Install older version of GopherJS
#RUN ls ./ -l
RUN go mod download
RUN go install github.com/gopherjs/gopherjs@v1.18.0-beta2
#RUN ls ./ -l
#RUN go get github.com/ckndave/hclparser
#RUN ls /go/bin -l
# RUN echo $GOROOT
# # Special sauce part 1

RUN go mod vendor



# # Special sauce part 2 - copy to GOPATH
#RUN mkdir -p /go/src/hcl2-parser
#RUN cp -r * /go/src/hcl2-parser
#WORKDIR /go/src/hcl2-parser
#RUN ls ./ -R
#RUN ls ./vendor/ -l
#RUN ls ./vendor/github.com -l
#RUN ls ./vendor/github.com/gopherjs -l
#RUN ls ./vendor/github.com/gopherjs/gopherjs -l
#RUN ls ./vendor/github.com/gopherjs/gopherjs/js/ -l
#RUN ls /go/ -R
#RUN ls /go/bin -l
#RUN ls /go/src/ -l
# # Finally we can build
RUN gopherjs version
RUN gopherjs build parser.go -o index.js -m

COPY build/copy-out.sh .
RUN chmod +x copy-out.sh
ENTRYPOINT ["/build/copy-out.sh"]

