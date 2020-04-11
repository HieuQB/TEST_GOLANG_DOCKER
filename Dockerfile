FROM golang:alpine as build-env

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# All these steps will be cached
RUN mkdir /fly_jx
WORKDIR /fly_jx
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/fly_jx ./main.go

# <- Second step to build minimal image
FROM scratch

# Import from builder.
COPY --from=build-env /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /etc/passwd /etc/passwd

COPY --from=build-env /go/bin/fly_jx /go/bin/fly_jx
COPY ./data/config.yml ./data/config.yml

EXPOSE 8080
ENTRYPOINT ["/go/bin/fly_jx"]


