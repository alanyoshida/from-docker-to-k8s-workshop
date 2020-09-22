# Ingress

We are going to make a nginx ingress controller pointing to a service inside a pod.

One ingress controller can have multiple ingresses, each ingress can point to a different service

Create your domain in /etc/hosts

Add `127.0.0.1 $USER.local` to your `/etc/hosts`

Create a self signed certificate, replace with your domain

This create the cert.pem and key.pem

Copy the content to the file in yaml/ingress/tls-secret.yaml

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -subj "/CN=alan.oliveira.local"
```

Change the content of **yaml/ingress/ingress.yaml**:

```yaml
  tls:
  - hosts:
    - <PUT_YOUR_DOMAIN_HERE>
    secretName: tls-secret
  rules:
  - host: <PUT_YOUR_DOMAIN_HERE>
```

```bash
# Create the app namespace
kubectl create ns example-app

# Create the nginx ingress controller
kubectl apply -f yaml/ingress-controller/namespace.yaml
kubectl apply -f yaml/ingress-controller/

# Create the app ingress
kubectl apply -f yaml/ingress/

# Deploy the application
kubectl apply -f app-example/deploy

# Test your application
curl --insecure https://alan.oliveira.test

```

#### This is how roles work

Role is attached to a service account with the Role binding:

Role -> Role binding -> Service Account

#### How ingress redirects

Ingress controller, control the creation of ingresses

Ingress -> Service -> Pod -> Application
