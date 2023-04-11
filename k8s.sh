curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install bitnami/mysql --generate-name