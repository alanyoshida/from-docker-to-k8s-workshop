# k8s Workshop

## Part 1

### Install

Requirements:

* docker
* kind
* kubectl

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

* [pods](k8s-part1/pods.md)
* [replicasets](k8s-part1/replicasets.md)
* [deployment](k8s-part1/deployment.md)
* [namespaces](k8s-part1/namespaces.md)
* [services](k8s-part1/services.md)
* configmap
* statefulset

## Part 2 - Scheduling

* Scheduling
* Labels & Selectors
* Taints & Tolerations
* Node selector
* Node Affinity
* [deamonset](k8s-part1/deamonset.md)

## Part 3 - Logging & Monitoring

* Loki / Graylog / Kibana ELK
* Prometheus operator
* Grafana

## Part 4 - Security

* Secrets
* TLS
* Certificates
* [serviceAccount](k8s-part1/serviceAccount.md)
* [policy](k8s-part1/policy.md)

## Part 5 - Storage

* Volumes
* Persistent Volumes
* Volume Claims

## Part 6 - Networking

* CoreDNS
* Ingress Controller (nginx, traefik)
* Service Mesh (Istio, linkerd)

## Part 7 - Development Life Cycle

* Helm
* Helm templates
* Deploy using argo
* Flagger
* Valt

## Troubleshooting

* how to deal with a problem with my deploy
* how to deal with a problem with my application
