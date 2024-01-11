# stocks-htmx-test

A repro to test out different software:

- htmx
- shoelace
- copilot
- go templating
- grafana/tempo
- elastic-stack
- terraform
- otel

## TODO

- add OTEL

## Docker

### Build

```bash
docker build -f ./cmd/stocks-htmx/dockerfile -t "timhuebener/go-htmx:latest" .
```

### Run

```bash
docker run --publish 8080:8080 go-htmx
```

## Local Kubernetes

```bash
kubectl config use-context docker-desktop
```

```bash
kubectl apply -f ./k8s/stocks-htmx.yml
```

## Terraform

Init Terraform

```bash
terraform init
```

Apply Terraform

```bash
terraform apply
```

### Run Grafana

```bash
kubectl port-forward svc/tempo 4318 -n monitoring
```

```bash
kubectl port-forward svc/grafana 3000:80 -n monitoring
```

Add datasource `http://tempo.monitoring:3100`