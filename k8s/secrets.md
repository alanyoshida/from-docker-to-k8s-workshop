# Secrets

Secrets can be files that can be mounted inside the pods, or environment variables

Let's do an example:

```bash
cd secrets-example

# Runs kind with a registry running in docker in port 5000
./kind-registry.sh

kubectl apply -f yaml/secrets.yaml

kubectl get secrets
kubectl get secrets example-secret -oyaml
kubectl describe secrets

# Build the image
docker build . -t localhost:5000/secrets-example
docker push localhost:5000/secrets-example

# Create the deployment
kubectl apply -f yaml/deployment.yaml

kubectl get pods

kubectl logs deploy/secrets-example

```

## Challenge

Create a configmap that uses the config.json file instead of the config inside yaml file.
