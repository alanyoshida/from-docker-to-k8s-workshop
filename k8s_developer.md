# Kubernetes for developers

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
* [configmap](k8s/configmap.md)
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

## Development Life Cycle

* Helm/Charts
* Helm templates
* Kustomize
* Deploy using argo
* Flagger
* Vault Secrets
* Skaffold
* Tilt
* Telepresence

## Troubleshooting

* how to deal with a problem with my deploy
* how to deal with a problem with my application
