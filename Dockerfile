# messaging Dockerfile
# Start from the 1.18 golang base image 1
FROM golang:1.18.0-alpine as builder

# Add Maintainer Info
LABEL maintainer="Arif W <arif@majoo.id>"

# Update alpine repository and install git and build-base musl-dev and openssh
RUN apk update && apk add -U --no-cache openssh git build-base musl-dev

# go get uses git internally. The following one liners will make git and consequently go get clone your package via SSH.
RUN git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"

# add credentials on build
ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
RUN echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN chmod 400 /root/.ssh/id_rsa

# make sure your domain is accepted
RUN touch /root/.ssh/known_hosts
# RUN apk update && apk add openssh 
RUN ssh-keyscan bitbucket.org >> /root/.ssh/known_hosts

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN cd app/api/ \
    && go build -o api

######## Start a new stage from scratch #######

FROM alpine:latest

# update alpine repository and install tzdata
RUN apk update && apk add tzdata

#install certificat and git
RUN apk add -U --no-cache ca-certificates git

# set timezone
ENV TZ Asia/Jakarta

# Workdir
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder app/app/api/ /root/app/api/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app/api/main"]