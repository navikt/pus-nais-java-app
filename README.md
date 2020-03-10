# pus-nais-java-app
Baseimage for java apps running on nais

## usage
Create a dockerfile with the contents:
```docker
FROM navikt/java:8
COPY /target/<app-name> .
```
