## Kind

### Create Cluster

```bash
$ kind create cluster --config=infra/multi-node.yaml
```

### Sample Nginx Deploy

```bash
$ kubectl apply -f infra/app/app-nginx.yaml
```

### Go to the following and the nginx screen will appear

`http:/ホストマシンのIP/:30070/`

### Delete Cluster

```bash
$ kind delete cluster
```
