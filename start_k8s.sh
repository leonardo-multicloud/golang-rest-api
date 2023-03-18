#!/bin/bash

#
# Deploy in Kubernetes
#

# Carrega as variáveis de ambiente do arquivo ENV_STG.env
source .ci/ENV_STG.env

echo "*** Criando manifestos do : $APP_NAME ***"

echo "*** Exporando variáveis da aplicação ***"
export APP_NAME=$APP_NAME
export APP_NAMESPACE=$APP_NAMESPACE
export IMAGE=$IMAGE
export REQUEST_NODE_MEMORIA=$REQUEST_NODE_MEMORIA
export REQUEST_NODE_CPU=$REQUEST_NODE_CPU
export LIMITE_NODE_MEMORIA=$LIMITE_NODE_MEMORIA
export LIMITE_NODE_CPU=$LIMITE_NODE_CPU
export APP_URL=$APP_URL

echo "*** Lendo arquivos de manifesto yaml ***"
# Lê o arquivo deployment.yml linha por linha e substitui as variáveis de ambiente
while IFS= read -r line; do
  echo "$line" | envsubst
done < .ci/deployment.yml > .ci/k8s-deployment.yml


while IFS= read -r line; do
  echo "$line" | envsubst
done < .ci/service.yml > .ci/k8s-service.yml

while IFS= read -r line; do
  echo "$line" | envsubst
done < .ci/ingress.yml > .ci/k8s-ingress.yml

awk 'FNR==1 && NR!=1 {printf "---\n"}{print}' .ci/k8s-*.yml > .ci/apply.yaml

echo "*** Removendo arquivos temporário ***"
rm .ci/k8s-*.yml

echo "*** Deploy to Kubernetes ***"
kubectl config use-context docker-desktop

kubectl apply -f .ci/apply.yaml
