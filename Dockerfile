########### BUILD STAGE ###########
FROM golang:1.11-alpine as builder

ARG SUBDIR

LABEL maintainer="David Kleiber <kleiberd93@gmail.com>"

# install dependencies
RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

# create backenduser
RUN adduser -D -g '' backenduser

# copy project files
COPY $SUBDIR/. $GOPATH/src/budapest-car-sharing-backend/$SUBDIR
WORKDIR $GOPATH/src/budapest-car-sharing-backend/$SUBDIR

# install project dependencies
RUN go get -u github.com/golang/dep/cmd/dep github.com/gobuffalo/packr/...
RUN dep ensure

# build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $GOPATH/bin/packr build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/$SUBDIR

########### FINAL STAGE ###########
FROM scratch

# copy user credentials, ca certs and static executable
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/$SUBDIR /go/bin/$SUBDIR

# set user
USER backenduser