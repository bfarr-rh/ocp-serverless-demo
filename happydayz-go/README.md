# Description

A simple K-Native go application with GET and POST functions

## Build

func build --registry quay.io/mancubus77

## Deployment

### Remote deployment with Tekton pipeline

kn func deploy --remote

### Create KN Service

kn service create happydayz \
--image quay.io/mancubus77/kn-happydayz:latest \
--annotation serving.knative.openshift.io/disableRoute=true

### Update service with new ENV variable

kn service update happydayz \
--env MESSAGE=ThisIsRevision2

### Update Image

kn service update happydayz \
--image quay.io/mancubus77/kn-happydayz:v2

### Measure Connection Time 

curl -o /dev/null -s -w 'Establish Connection: %{time_connect}s\nTTFB: %{time_starttransfer}s\nTotal: %{time_total}s\n'  http://127.0.0.1:20001 -kv

