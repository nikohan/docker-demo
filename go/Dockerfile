FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.19.3-alpine3.16 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

ARG ADDITIONAL_PACKAGE
ARG CGO_ENABLED=0
ARG GO111MODULE="off"
ARG GOPROXY=""
ARG GOFLAGS=""

ENV CGO_ENABLED=0

# Set the working directory to /home/app
WORKDIR /home/app

# Copy the current directory contents into the container at /home/app
ADD . /home/app

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=${CGO_ENABLED} \
    go build --ldflags "-s -w" -a -installsuffix cgo -o health .

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:3.16 as ship

# Set the working directory to /home/app
WORKDIR /home/app

COPY --from=builder /home/app  .

# Make port 80 available to the world outside this container
EXPOSE 80

# Run app.py when the container launches
CMD ["./health"]
