set -euo pipefail

readonly DEV_NAMESPACE="my-apps"
readonly YTT_VERSION=0.41.0
# make sure to regenerate manifests and run tests 
make all

# delete existing workload
if [[ -n $( kubectl get podintent.conventions.carto.run/spring-sample -n my-apps) ]]; then
    kubectl delete -f samples/ytt-based-conventions/workload.yaml 
fi 

# delete any existing convention and deployment
if [[ -n $(kubectl get clusterpodconvention.conventions.carto.run/spring-ytt-sample) ]]; then
  kubectl delete -f samples/ytt-based-conventions/server.yaml
fi 

# delete convention deployent if it exists 
if [[ -n $( kapp list -n cartographer-system  | grep 'conventions') ]]; then
  kapp delete -a conventions -n cartographer-system -y 
fi 


if [[ -n $(kubectl get secret image-pull-secret -n my-apps) ]]; then 
# delete existing secret 
    kubectl delete secret image-pull-secret -n ${DEV_NAMESPACE}
else 
 # create secret  with credentials
    kubectl create secret docker-registry image-pull-secret \
          --docker-server="${REGISTRY_SERVER}"\
          --docker-username="${REGISTRY_USER}" \
          --docker-password="${REGISTRY_PASSWORD}" \
          -n my-apps
  
    kubectl -n my-apps  patch serviceaccount default -p '{"imagePullSecrets": [{"name": "image-pull-secret"}]}'
  fi 

# download ytt to ko data folder 
download_ytt_to_kodata() {
  local url=https://github.com/vmware-tanzu/carvel-ytt/releases/download/v${YTT_VERSION}/ytt-linux-amd64
  local dest=$(realpath ./cmd/kodata/ytt-linux-amd64)

  test -x $dest && {
          echo "ytt already found in kodata."
          return
  }

  curl -sSL $url -o $dest
  chmod +x $dest
}

download_ytt_to_kodata

# make a new deployment 
ytt -f "dist/cartographer-conventions.yaml"
ko resolve -f 
kapp deploy -n cartographer-system -a conventions  -f 

# # create the ytt based convention 
# kubectl apply -f samples/ytt-based-conventions/ytt-example.yaml
# #kubectl delete -f samples/ytt-based-conventions/ytt-example.yaml

# # inspect the ytt based convention
# kubectl get clusterpodconvention.conventions.carto.run/spring-ytt-sample -o yaml

# # create a workload 
# kubectl apply -f samples/ytt-based-conventions/workload.yaml
# #kubectl delete -f samples/ytt-based-conventions/workload.yaml
