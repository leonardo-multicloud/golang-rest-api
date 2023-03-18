#!/bin/bash

#
# Delete in Kubernetes
#

echo "*** Delete to Kubernetes ***"

kubectl delete -f .ci/apply.yaml
rm .ci/apply.yaml