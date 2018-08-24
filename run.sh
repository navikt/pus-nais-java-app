#!/bin/sh
set -e
set -x


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

exec java \
-XX:+UnlockExperimentalVMOptions \
-XX:+UseCGroupMemoryLimitForHeap \
${JAVA_OPTS} \
${APPDYNAMICS_OPTS} \
${PROXY_OPTS} \
-server \
-classpath "${APP_DIR}/WEB-INF/classes:${APP_DIR}/WEB-INF/lib/*" \
Main
