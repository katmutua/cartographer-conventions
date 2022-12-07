set -euo pipefail

# make sure to regenerate manifests and run tests 
make all

# delete existing workload
if [[ -n $( kubectl get podintent.conventions.carto.run/spring-sample -n my-apps) ]]; then
    kubectl delete -f samples/ytt-based-conventions/workload.yaml 
fi 

# delete any existing convention and deployment
if [[ -n $(kubectl get clusterpodconvention.conventions.carto.run/spring-ytt-sample) ]]; then
  kubectl delete -f samples/ytt-based-conventions/ytt-example.yaml
fi 

# delete convention deployent if it exists 
if [[ -n $( kapp list -n cartographer-system  | grep 'conventions') ]]; then
  kapp delete -a conventions -n cartographer-system -y 
fi 


if [[ -n $(kubectl get secret image-pull-secret -n my-apps) ]]; then 
# delete existing secret 
  kubectl delete secret image-pull-secret -n my-apps 
else 
 # create secret  with credentials
  kubectl create secret docker-registry image-pull-secret \
          --docker-server="${REGISTRY_SERVER}"\
          --docker-username="${REGISTRY_USER}" \
          --docker-password="${REGISTRY_PASSWORD}" \
          -n my-apps

  fi 

# download ytt to ko data folder 


# make a new deployment 
kapp deploy -n cartographer-system -a conventions -f  <( \
ko resolve -f <( \
      ytt \
        -f dist/cartographer-conventions.yaml
      ) \
  )

# # create the ytt based convention 
# kubectl apply -f samples/ytt-based-conventions/ytt-example.yaml
# #kubectl delete -f samples/ytt-based-conventions/ytt-example.yaml

# # inspect the ytt based convention
# kubectl get clusterpodconvention.conventions.carto.run/spring-ytt-sample -o yaml

# # create a workload 
# kubectl apply -f samples/ytt-based-conventions/workload.yaml
# #kubectl delete -f samples/ytt-based-conventions/workload.yaml
