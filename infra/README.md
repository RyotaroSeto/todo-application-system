## Kind

### Create Cluster

```bash
$ kind create cluster --config=infra/multi-node.yaml
```

### Delete Cluster

```bash
$ kind delete cluster
```

### Sample Nginx Deploy

```bash
$ kubectl apply -f infra/app/app-nginx.yaml
```
