FROM navikt/java:8-appdynamics

ENV APPD_ENABLED=true

COPY --chown=apprunner:root java-debug.sh /init-scripts/08-java-debug.sh
COPY --chown=apprunner:root appd-init.sh /init-scripts/09-appd-init.sh

USER root

RUN mkdir /secure-logs && chown apprunner:root /secure-logs

USER apprunner
