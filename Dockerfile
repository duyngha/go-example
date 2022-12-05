FROM ubuntu:latest

# Set variables for docker
# ARG DEBIAN_FRONTEND=noninteractive
ARG GOLANG_VERSION=1.19.3

ENV GOROOT=/usr/local/go
ENV GOPATH=/home/app

WORKDIR /home/app

# Install linux packages
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y wget

# Install golang
RUN wget https://go.dev/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
# RUN sha256sum go$GOLANG_VERSION.linux-amd64.tar.gz
RUN tar -xvf go$GOLANG_VERSION.linux-amd64.tar.gz
RUN mv go /usr/local
# RUN touch /etc/profile
RUN echo "export GOROOT=/usr/local/go" >> /etc/profile
RUN echo "export GOPATH=/home/app" >> /etc/profile
RUN echo 'export PATH="$GOPATH/bin:$GOROOT/bin:$PATH"' >> /etc/profile

RUN cat /etc/profile
RUN . /etc/profile
RUN -- bash -c source /etc/profile
RUN echo $PATH

# RUN cat ~/.profile


# RUN tar -xzf go$GOLANG_VERSION.linux-amd64.tar.gz -C /usr/local
# RUN ls -alh /usr/local
# RUN echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
# SHELL ["/bin/bash", "-c"]
# RUN cat /etc/profile
# RUN source /etc/profile
# RUN export GOROOT_BOOTSTRAP=/usr/local/go
# RUN cd /usr/local/go/src && ./make.bash

# ENV PATH=$PATH:/usr/local/go/bin

# RUN echo "export PATH=$PATH:/usr/local/go/bin" >> /home/.bash_profile
# RUN cat /home/.bash_profile
# RUN go version

# Install and configure Supervisor
# RUN apt-get install -y openssh-server supervisor \
    # && mkdir -p /var/log/supervisor