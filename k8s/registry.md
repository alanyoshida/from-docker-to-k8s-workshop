# Local docker registry

## Requirements

* Helm 3

```bash
# With Helm
helm install stable/docker-registry --generate-name
kubectl port-forward deployment/docker-registry 5000

# OR run with docker
docker run -d -p 5000:5000 --restart always --name registry registry:2

# Now the docker part
docker build . -t configmap-app

docker tag configmap-app localhost:5000/configmap-app
docker push localhost:5000/configmap-app
```
