FROM golang:1.12

#
# NOTE: The RPC server listens on localhost by default.
#       If you require access to the RPC server,
#       rpclisten should be set to an empty value.
#
# NOTE: When running simnet, you may not want to preserve
#       the data and logs.  This can be achieved by specifying
#       a location outside the default ~/.ucd.  For example:
#          rpclisten=
#          simnet=1
#          datadir=~/simnet-data
#          logdir=~/simnet-logs
#
# Example testnet instance with RPC server access:
# $ mkdir -p /local/path/ucd
#
# Place a ucd.conf into a local directory, i.e. /var/ucd
# $ mv ucd.conf /var/ucd
#
# Verify basic configuration
# $ cat /var/ucd/ucd.conf
# rpclisten=
# testnet=1
#
# Build the docker image
# $ docker build -t user/ucd .
#
# Run the docker image, mapping the testnet ucd RPC port.
# $ docker run -d --rm -p 127.0.0.1:11509:11509 -v /var/ucd:/root/.ucd user/ucd
#

WORKDIR /go/src/github.com/UtopiaCoinOrg/ucd
COPY . .

RUN env GO111MODULE=on go install . ./cmd/...

# mainnet
EXPOSE 10508 10509

# testnet
EXPOSE 11508 11509

# simnet
EXPOSE 12508 12509

CMD [ "ucd" ]
