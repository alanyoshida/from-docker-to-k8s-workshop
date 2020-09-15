# Kubernetes for administrators

### Summary

* Core concepts
  * etcd
  * kube api server
  * kube controller manager
  * kube scheduler
  * kubelet
  * kube proxy

[Diagrams](https://github.com/cloudogu/k8s-diagrams)

### Resources

Kind:

* [pods](k8s/pods.md)
* [replicasets](k8s/replicasets.md)
* [deployment](k8s/deployment.md)
* [namespaces](k8s/namespaces.md)
* [services](k8s/services.md)
* configmap
* statefulset

## Scheduling

* Scheduling
* Labels & Selectors
* Taints & Tolerations
* Node selector
* Node Affinity
* [deamonset](k8s/deamonset.md)

## Logging & Monitoring

* Loki / Graylog / Kibana ELK
* Prometheus operator
* Grafana

## Security

* Secrets
* TLS
* Certificates
* [serviceAccount](k8s/serviceAccount.md)
* [policy](k8s/policy.md)

## Storage

* Volumes
* Persistent Volumes
* Volume Claims

## Networking

* CoreDNS
* Ingress Controller (nginx, traefik)
* Service Mesh (Istio, linkerd)

## Development Life Cycle

* Helm
* Helm templates
* Deploy using argo
* Flagger
* Vault

## Troubleshooting

* how to deal with a problem with my deploy
* how to deal with a problem with my application
