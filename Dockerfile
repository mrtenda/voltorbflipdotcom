FROM golang
ENV VFLIP_STATIC_CONTENT_PATH /go/src/github.com/mrtenda/voltorbflipdotcom/site
COPY ./server /go/src/github.com/mrtenda/voltorbflipdotcom/server
COPY ./jekyll-site/_site /go/src/github.com/mrtenda/voltorbflipdotcom/site
WORKDIR /go/src/github.com/mrtenda/voltorbflipdotcom/server
RUN go install github.com/mrtenda/voltorbflipdotcom/server
ENTRYPOINT /go/bin/server
EXPOSE 8080
