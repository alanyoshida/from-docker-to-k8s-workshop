# Configmaps

Configmaps are files that can be mounted inside the pods

```
               +-------------------+
               |                   |
               |    deployment     |
               |                   |
               |                   |
               |  +-----+ +-----+  |
               |  | pod | | pod |  |
               |  +--^--+ +--^--+  |
               |     |       |     |
               +-------------------+
                     |       |
+-----------+        |       |   mount
| configmap +--------+-------+
+-----------+

```

Let's do an example:

```bash
# With config inside the yaml
kubectl apply -f configmap-example/configmap.yaml

# From file
kubectl create configmap example-config --from-file /path-to-config/config.json


kubectl get configmaps
kubectl get configmaps example-config -oyaml

```

## Challenge

Create a configmap that uses the config.json file instead of the config inside yaml file.
