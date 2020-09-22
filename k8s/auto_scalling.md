# Auto scalling

```bash
kubectl top pods

kubectl autoscale deploy/application-cpu --cpu-percent=95 --min=1 --max=10

kubectl get hpa/application-cpu -owide

```
