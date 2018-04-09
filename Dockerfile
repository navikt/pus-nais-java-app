FROM openjdk:8-jre-alpine

ENV LC_ALL="no_NB.UTF-8"
ENV LANG="no_NB.UTF-8"
ENV TZ="Europe/Oslo"

EXPOSE 8080

WORKDIR /work
ARG APP_DIR=/app
ENV APP_DIR=${APP_DIR}

ADD run.sh /run.sh
RUN chmod +x /run.sh
CMD sh /run.sh