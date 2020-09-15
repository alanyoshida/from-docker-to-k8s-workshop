# Services

Services exposes your service

### Services use namespaces

<service-name>.<namespace>.svc.cluster.local

Example:
`db-service.dev.svc.cluster.local`

### NodePort

Create a port in the node forwarding to the pod

```
+----------------------------------------------------------------+
|                                                                |
|                                                                |
|nodePort         +-------------+----+                           |
+------+          |             |    |port                       |
|      +----------> SERVICE     | 80 +------------+              |
| 3008 |          |             |    |            |              |
|      |          +-------------+----+            |              |
+------+                                          |              |
|                                                 | targetPort   |
|                                     +--------+--v---+-------+  |
|                                     |        |  80  |       |  |
|                                     |        +------+       |  |
|                                     |  +-----------------+  |  |
|                                     |  |     10.244.0.2  |  |  |
|                NODE                 |  +-----------------+  |  |
|                                     |                       |  |
|                                     |          POD          |  |
|                                     |                       |  |
|                                     +-----------------------+  |
|                                                                |
+----------------------------------------------------------------+
```

```bash
kubectl apply -f yaml/services_node_port.yaml
kubectl get services
```

#### Obs

* Services serves also as load balancer when you have more than one pod with the selector
* When you have pods in multiple nodes it distributes the requests as well

### ClusterIp

Create a virtual ip inside the cluster, to internal communication between services

```bash
kubectl create -f yaml/services_cluster_ip.yaml
kubectl get services
kubectl get svc
```

### LoadBalancer

Create a load balancer in supported cloud provider
