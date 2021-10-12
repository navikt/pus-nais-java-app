FROM navikt/java:16-appdynamics

ENV APPD_ENABLED=true

COPY --chown=apprunner:root java-debug.sh /init-scripts/08-java-debug.sh
COPY --chown=apprunner:root appd-init.sh /init-scripts/09-appd-init.sh