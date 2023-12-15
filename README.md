# stocks-htmx-test

A repro to test out different software:
- htmx
- shoelace
- copilot
- go templating


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