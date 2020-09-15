# Namespaces

```bash
# list namespaces
kubectl get namespaces
kubectl get ns

# Create
kubectl create ns dev
kubectl create -f yaml/namespace.yaml

# get pods from namespace
kubectl --namespace=kube-system get pods
kubectl -n kube-system get pods

# Create pods
kubectl create -f yaml/pod_simple.yaml --namespace dev
kubectl create -f yaml/pod_simple_with_namespace.yaml
```

### Services use namespaces

<service-name>.<namespace>.svc.cluster.local

Example:
`db-service.dev.svc.cluster.local`
