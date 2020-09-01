## Prerequisites

Install postgres helm chart:
```shell script
helm repo add bitnami https://charts.bitnami.com/bitnami
``` 

## Values

Here you can change the container image and tag:

**Chart.yaml**
```yaml
image:
  repository: maxweis/go-todo-demo
  version: v1.0
```

Change the env var:

**Chart.yaml**
```yaml
env:
  - name: DB_HOST
    value: gotodo-postgresql
  - name: DB_PORT
    value: 5432
  - name: DB_NAME
    value: postgres
  - name: DB_USER
    value: postgres
  - name: DB_PASS
    value: postgres
```

Change the postgres user credentials:

**Chart.yaml**
```yaml
postgresql:
  postgresqlDatabase: postgres
  global:
    postgresql:
      postgresqlPassword: postgres
```

## Commands

### Lint

To check if the chart is correct run:
```shell script
helm lint ./gotodo
```

### Template

View the final manifests run:

```shell script
helm template ./gotodo
``` 

### Install

Get postgres dependency:
```shell script
helm dependency build gotodo
``` 

Install the chart in your k8s cluster
```shell script
helm install gotodo gotodo
``` 

### Update Values

Change the values in the `Chart.yaml` and run `helm upgrade gotodo gotodo` to apply them 

### Delete Chart
```shell script
helm delete gotodo
``` 