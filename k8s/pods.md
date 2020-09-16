# Pods

### Commands

```bash
# Create
kubectl apply -f yaml/pods_simple.yaml

# List
kubectl get pods

# Describe
kubectl describe pod

# Delete
kubectl delete pod hello-world
```

### Challenge

Create a yaml file and create a new pod with mysql

### Reference

* [Pod livefecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)
* [Pod readiness gate](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-readiness-gate)
