FROM golang:1.17 as builder
ADD . /src/ip2location
WORKDIR /src/ip2location 
RUN go mod download 
RUN make docker

# FROM scratch
FROM alpine
COPY --from=builder /src/ip2location/bin/ip2location-docker /ip2location
EXPOSE 5000
ENTRYPOINT ["/ip2location"]

ADD DB11LITEBINIPv6.zip /
ENV DB_TYPE="IPv6"
ENV DB_LEVEL="DB11"
ENV TOKEN=""
ENV WORKDIR "/opt/"