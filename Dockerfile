FROM docker.adeo.no:5000/pus/toolbox as downloader
RUN wget https://repo.adeo.no/repository/raw/appdynamics/appdynamics.zip -O temp.zip
RUN unzip temp.zip

FROM openjdk:8-jre-alpine

COPY --from=downloader /appdynamics /appdynamics

ENV LC_ALL="no_NB.UTF-8"
ENV LANG="no_NB.UTF-8"
ENV TZ="Europe/Oslo"

EXPOSE 8080

WORKDIR /work
ARG APP_DIR=/app
ENV APP_DIR=${APP_DIR}

ADD run.sh /run.sh
ADD proxy.sh /proxy.sh
RUN chmod +x /run.sh
CMD sh /run.sh
