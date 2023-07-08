# Use the official Go base image
FROM golang:latest

# Install IPFS and IPFS Cluster
RUN go get -u github.com/ipfs/go-ipfs
RUN go get -u github.com/ipfs/ipfs-cluster

# Expose the default IPFS ports
EXPOSE 4001 5001 8080

# Set the entrypoint to run IPFS Cluster
ENTRYPOINT ["ipfs-cluster-service"]
