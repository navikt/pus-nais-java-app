#!/bin/sh
set -e
set -x

# TODO ikke prodklart enda, men det kommer!

# [ -XX:+UnlockExperimentalVMOptions -XX:+UseCGroupMemoryLimitForHeap ]
# https://blogs.oracle.com/java-platform-group/java-se-support-for-docker-cpu-and-memory-limits


java \
-XX:+UnlockExperimentalVMOptions \
-XX:+UseCGroupMemoryLimitForHeap \
-server \
-classpath "${APP_DIR}/WEB-INF/lib/*:${APP_DIR}/WEB-INF/classes" \
Main