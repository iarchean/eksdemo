[reference](https://antonputra.com/terraform/how-to-create-eks-cluster-using-terraform/#create-iam-oidc-provider-eks-using-terraform)

## 0. prerequisites

- [terraform cli](https://learn.hashicorp.com/tutorials/terraform/install-cli)
- [aws cli](https://aws.amazon.com/cli/) / [gcloud cli](https://cloud.google.com/sdk/gcloud)
- [kubectl](https://kubernetes.io/zh-cn/docs/tasks/tools/install-kubectl-linux/)

## 1. deploy kubernetes cluster (EKS) on AWS

terraform plan & apply:

```
cd terraform
terraform plan
terraform apply
```

apply EKS kubeconfig:

```
aws eks --region ap-northeast-1 update-kubeconfig --name demo
```

---

## 2. check the cluster

Show running nodes:

```
kubectl get nodes
```

Show running pods:

```
kubectl get pods --all-namespaces
```

---

## 3. deploy cluster infrastructure

deploy nginx ingress controller and metrics server:

```
kubectl apply -f kubernetes/
```

---

## 4. deploy example app

dockerlize example app:

```
docker build -t archean/echo:1.0.0 .
```

push to hub.docker.com:

```
docker push archean/echo:1.0.0
```

deploy example app: echo:

```
kubectl apply -f deploy/
```

check web service:

```
ip=$(kubectl get ingress echo -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')
curl -H 'Host:example.com' http://$ip/ping
```

---

## 5. scale example app

manual scale up echo app:

```
kubectl scale deployment echo --replicas=2
```

automatic scale up nodes:

```
kubectl apply -f kubernetes/cluster-autoscaler.yaml
kubectl scale deployment echo --replicas=5

```
see nodes auto scale up:

```
kubectl get nodes -w
```

## 6. destroy cluster

delete example app:

```
kubectl delete -f deploy/
```

delete cluster infrastructure:

```
kubectl delete -f kubernetes/
```

destroy EKS cluster:

```
terraform destory
```
