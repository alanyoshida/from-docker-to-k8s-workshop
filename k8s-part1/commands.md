# Commands

## Docker

```bash
docker run nginx --image nginx
```

## Kubernetes

```bash
# Show the configured contexts
kubectl config get-contexts

# Use kind context
kubectl config use-context kind-kind

# Run nginx in k8s
kubectl run nginx --image nginx

# List resources
kubectl get pods

# List all pods from all namespaces
kubectl get pods -A
kubectl get pods --all-namespaces

# List all resources
kubectl get all
```

## Reference

- [kubectl cheatsheet](https://kubernetes.io/pt/docs/reference/kubectl/cheatsheet/)
