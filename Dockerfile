FROM ubuntu:18.04
MAINTAINER  Aaron Janes
WORKDIR /app

RUN apt-get -y update && apt-get -y upgrade && apt-get install -y \
    build-essential \
    curl \
    git \
    wget \
    unzip \
    vim \
    && rm -rf /var/lib/apt/lists/*
# Install go
RUN curl -s https://dl.google.com/go/go1.20.5.linux-amd64.tar.gz | tar -v -C /usr/local -xz

ENV PATH $PATH:/usr/local/go/bin

#Install nodejs
RUN apt-get install -y curl gnupg
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -
RUN apt-get install -y nodejs

# Install sass compiler
RUN npm install -g sass




COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN sass styles:static/css
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-server

EXPOSE 8080

CMD ["/docker-server"]
