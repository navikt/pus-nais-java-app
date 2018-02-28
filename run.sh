#!/bin/sh
set -e
set -x

# [ -XX:+UnlockExperimentalVMOptions -XX:+UseCGroupMemoryLimitForHeap ]
# https://blogs.oracle.com/java-platform-group/java-se-support-for-docker-cpu-and-memory-limits

exec java \
-XX:+UnlockExperimentalVMOptions \
-XX:+UseCGroupMemoryLimitForHeap \
-server \
-classpath "${APP_DIR}/WEB-INF/lib/*:${APP_DIR}/WEB-INF/classes" \
Main