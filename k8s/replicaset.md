# Replicasets

```bash
# Create
kubectl create -f yaml/replication_controller.yaml

# List
kubectl get replicaSet
kubectl get pods

# Deleting pods, they will come back
kubectl delete pod <pod_name>

# Edit replicaset
kubectl edit rs
kubectl replace -f <.yaml>

# Change scale
kubectl scale --replicas=5 -f <.yaml>
kubectl scale --replicas=5 rs myapp-rs

# Delete
kubectl delete rs myapp-rc
```

## Labels & Selectors

You can use for managing creation of pods
