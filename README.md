# Countme


## Deploy Countme on local Kubernetes
```
cd helm/
```

```
helm repo update
```

```
helm dependency update
```

```
helm install -n countme-local --dry-run --debug .
```

```
$ kubectl get all --selector release=countme-local
NAME                           DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/countme-local-countme   3         3         3            3           4s
deploy/countme-local-redis     1         1         1            1           4s

NAME                                  DESIRED   CURRENT   READY     AGE
rs/countme-local-countme-6f5b699764   3         3         3         4s

NAME                           DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/countme-local-countme   3         3         3            3           4s
deploy/countme-local-redis     1         1         1            1           4s

NAME                                  DESIRED   CURRENT   READY     AGE
rs/countme-local-countme-6f5b699764   3         3         3         4s

NAME                                        READY     STATUS    RESTARTS   AGE
po/countme-local-countme-6f5b699764-bjl5p   1/1       Running   0          4s
po/countme-local-countme-6f5b699764-g8nbf   1/1       Running   0          4s
po/countme-local-countme-6f5b699764-mgq6c   1/1       Running   0          4s

NAME                        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
svc/countme-local-countme   ClusterIP   10.109.108.216   <none>        8000/TCP   4s
svc/countme-local-redis     ClusterIP   10.109.58.109    <none>        6379/TCP   4s
```

## Test countme
```
$ kubectl port-forward countme-local-countme-6f5b699764-bjl5p 8000:8000
Forwarding from 127.0.0.1:8000 -> 8000
```

```
$ curl http://:8000/
PUT /incr
GET /count
GET /version
```

```
$ curl http://:8000/incr
1
[~]
$ curl http://:8000/incr
2
[~]
$ curl http://:8000/incr
3
```

## Upgrade countme

```
cd  helm/
```

```
helm upgrade countme-local --dry-run --debug .
```

```
helm upgrade countme-local --debug .
```