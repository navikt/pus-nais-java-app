FROM docker.adeo.no:5000/alpine-java:jre8
EXPOSE 8080
WORKDIR /work
ARG APP_DIR=/app
ENV APP_DIR=${APP_DIR}
ADD run.sh /run.sh
RUN chmod +x /run.sh
CMD bash /run.sh