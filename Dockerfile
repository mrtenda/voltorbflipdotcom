FROM golang
ENV VFLIP_STATIC_CONTENT_PATH /go/src/github.com/mrtenda/voltorbflipdotcom/site
ADD ./server /go/src/github.com/mrtenda/voltorbflipdotcom/server
ADD ./jekyll-site/_site /go/src/github.com/mrtenda/voltorbflipdotcom/site
RUN go install github.com/mrtenda/voltorbflipdotcom/server
ENTRYPOINT /go/bin/server
EXPOSE 8080