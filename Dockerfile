########### BUILD STAGE ###########
FROM golang:1.10.3-alpine as builder

LABEL maintainer="David Kleiber <kleiberd93@gmail.com>"

# install dependencies
RUN apk update && apk add git && apk add ca-certificates

# create backenduser
RUN adduser -D -g '' backenduser

# copy project files
COPY . $GOPATH/src/budapest-car-sharing-backend
WORKDIR $GOPATH/src/budapest-car-sharing-backend

# install project dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/backend

########### FINAL STAGE ###########
# start from scratch
FROM scratch

# copy user credentials, ca certs and static executable
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/backend /go/bin/backend

# set user
USER backenduser

# expose port
EXPOSE 8080

# run executable
ENTRYPOINT ["/go/bin/backend"]