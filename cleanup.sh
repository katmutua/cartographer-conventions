set -euo pipefail

# make sure to regenerate manifests and run tests 
make all

# delete any existing convention and deployment
kubectl delete -f samples/ytt-based-conventions/ytt-example.yaml
kapp delete -a conventions -n cartographer-system 

# create a new deployment 
kapp deploy -n cartographer-system -a conventions -f  <( \
  ko resolve -f <( \
    ytt \
      -f dist/cartographer-conventions.yaml
    ) \
)

# create the ytt based convention 
kubectl apply -f samples/ytt-based-conventions/ytt-example.yaml
#kubectl delete -f samples/ytt-based-conventions/ytt-example.yaml

# inspect the ytt based convention
kubectl get clusterpodconvention.conventions.carto.run/spring-ytt-sample -o yaml

# create a workload 
kubectl apply -f samples/ytt-based-conventions/workload.yaml
#kubectl delete -f samples/ytt-based-conventions/workload.yaml
