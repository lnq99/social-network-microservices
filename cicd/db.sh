#!/bin/bash

# deploy db 
helm upgrade --install \
    --set postgresqlPassword=P@ssw0rd \
    --set resources.requests.cpu="100m" \
    --set resources.requests.memory="256Mi" \
    --set resources.limits.cpu="1000m" \
    --set persistence.size=5Gi \
    --set resources.limits.memory="1Gi" \
    --set postgresqlExtendedConf.max_connections="1000" \
    --namespace=zdhamasha \
    zdhamasha-postgresql \
    bitnami/postgresql \
    --version 10.14.3