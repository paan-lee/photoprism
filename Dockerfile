FROM ubuntu:18.04

LABEL maintainer="Michael Mayer <michael@liquidbytes.net>"

RUN apt-get update && apt-get install -y --no-install-recommends \
        build-essential \
        curl \
        libfreetype6-dev \
        libhdf5-serial-dev \
        libpng-dev \
        libzmq3-dev \
        pkg-config \
        rsync \
        software-properties-common \
        unzip \
        g++ \
        gcc \
        libc6-dev \
        gpg-agent \
        apt-utils \
        make \
        nano \
        wget \
        darktable \
        git \
        mysql-client \
        && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN add-apt-repository ppa:pmjdebruijn/darktable-release

RUN apt-get update && apt-get install -y --no-install-recommends \
        darktable \
        && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN apt-get upgrade -y

# Install TensorFlow C library
RUN curl -L \
   "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-1.11.0.tar.gz" | \
   tar -C "/usr/local" -xz
RUN ldconfig

# Hide some warnings
ENV TF_CPP_MIN_LOG_LEVEL 2

# Install NPM (NodeJS)
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN apt-get install -y nodejs

# Install YARN (Package Manager)
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt-get update && apt-get install yarn

ENV GOLANG_VERSION 1.11.2
RUN set -eux; \
	\
	url="https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz"; \
	wget -O go.tgz "$url"; \
	echo "1dfe664fa3d8ad714bbd15a36627992effd150ddabd7523931f077b3926d736d *go.tgz" | sha256sum -c -; \
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz; \
	export PATH="/usr/local/go/bin:$PATH"; \
	go version

ENV GOPATH /go
ENV GOBIN $GOPATH/bin
ENV PATH $GOBIN:/usr/local/go/bin:$PATH
ENV GO111MODULE on
ENV NODE_ENV production

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN env GO111MODULE=off /usr/local/go/bin/go get golang.org/x/tools/cmd/goimports

# Set up project directory
WORKDIR "/go/src/github.com/photoprism/photoprism"
COPY . .

# Build PhotoPrism
RUN make all install

# Expose HTTP port
EXPOSE 80

# Start PhotoPrism server
CMD photoprism start