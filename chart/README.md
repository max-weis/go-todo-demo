## Prerequisites

Install postgres helm chart:
```shell script
helm repo add bitnami https://charts.bitnami.com/bitnami
``` 

## Values

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