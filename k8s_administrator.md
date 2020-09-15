# Kubernetes for administrators

### Summary

* Core concepts
  * etcd
  * kube api server
  * kube controller manager
  * kube scheduler
  * kubelet
  * kube proxy

#### Reference Links

* [Diagrams](https://github.com/cloudogu/k8s-diagrams)
* [Kubernetes objects](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/)

### Resources

Kind:

* [Pods](k8s/pods.md)
* [Replicasets](k8s/replicasets.md)
* [Deployment](k8s/deployment.md)
* [Namespaces](k8s/namespaces.md)
* [Services](k8s/services.md)
* Configmap
* Statefulset

## Scheduling

* Scheduling
* Labels & Selectors
* Taints & Tolerations
* Node selector
* Node Affinity
* deamonset

## Logging & Monitoring

* Loki / Graylog / Kibana ELK
* Prometheus operator
* Grafana

## Security

* Secrets
* TLS
* Certificates
* Service Account
* Policy
* Vault

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
* Charts
* Helm templates
* Deploy using argo
* Flagger
* Vault

## Troubleshooting

* how to deal with a problem with my deploy
* how to deal with a problem with my application
