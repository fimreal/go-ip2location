FROM golang:1.17 as builder
ADD . /src/ip2location
WORKDIR /src/ip2location 
#RUN go mod download 
RUN make docker

# FROM scratch
FROM alpine
COPY --from=builder /src/ip2location/bin/ip2location-docker /opt/ip2location
EXPOSE 5000
ENTRYPOINT ["/opt/ip2location"]

ADD DB11LITEBINIPv6.zip /opt/
ENV DB_TYPE="IPv6"
ENV DB_LEVEL="DB11"
ENV TOKEN=""
ENV WORKDIR "/opt/"
WORKDIR /opt
