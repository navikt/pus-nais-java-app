#!/bin/sh
set -e
set -x

APP_DIR=/app
HOSTNAME=`hostname`

# ikke aktiver appdynamic med mindre man har
if [ "${APPDYNAMICS_AGENT_ACCOUNT_ACCESS_KEY}" != "" ]; then

APPDYNAMICS_OPTS="
-javaagent:/appdynamics/javaagent.jar
-Dappdynamics.agent.applicationName=${APP_NAME}
-Dappdynamics.agent.nodeName=${HOSTNAME}
-Dappdynamics.agent.tierName=${FASIT_ENVIRONMENT_NAME}_${APP_NAME}
"

fi

# Convert proxy settings to Java form
PROXY_OPTS=$(/proxyopts)

# [ -XX:+UnlockExperimentalVMOptions -XX:+UseCGroupMemoryLimitForHeap ]
# https://blogs.oracle.com/java-platform-group/java-se-support-for-docker-cpu-and-memory-limits

# [ -XX:+MaxRAMFraction=2 ]
# https://stackoverflow.com/questions/49854237/is-xxmaxramfraction-1-safe-for-production-in-a-containered-environment

exec java \
-XX:+UnlockExperimentalVMOptions \
-XX:+UseCGroupMemoryLimitForHeap \
-XX:MaxRAMFraction=2 \
${JAVA_OPTS} \
${APPDYNAMICS_OPTS} \
${PROXY_OPTS} \
-server \
-classpath "${APP_DIR}/WEB-INF/classes:${APP_DIR}/WEB-INF/lib/*" \
Main
