FROM docker.adeo.no:5000/pus/toolbox as downloader
RUN wget https://repo.adeo.no/repository/raw/appdynamics/appdynamics.zip -O temp.zip
RUN unzip temp.zip

FROM golang:1.10 as go
ADD proxyopts.go /proxyopts.go
ADD proxyopts_test.go /proxyopts_test.go
RUN go test /proxyopts.go /proxyopts_test.go
RUN GOOS=linux GOARCH=amd64 go build -o /proxyopts /proxyopts.go

FROM openjdk:8-jre-alpine

COPY --from=downloader /appdynamics /appdynamics
COPY --from=go /proxyopts /proxyopts

ENV LC_ALL="no_NB.UTF-8"
ENV LANG="no_NB.UTF-8"
ENV TZ="Europe/Oslo"

EXPOSE 8080

WORKDIR /work

ADD run.sh /run.sh
RUN chmod +x /run.sh
CMD sh /run.sh
