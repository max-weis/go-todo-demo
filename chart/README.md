## Prerequisites

Install postgres helm chart:
```shell script
helm repo add bitnami https://charts.bitnami.com/bitnami
``` 

## Values

Here you can change the container image and tag:
```yaml
image:
  repository: maxweis/go-todo-demo
  version: v1.0
```

Change the env var:
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

```shell script
helm install gotodo gotodo
``` 

### Update Values
```shell script
helm upgrade gotodo gotodo
``` 

### Delete Chart
```shell script
helm delete gotodo
``` 