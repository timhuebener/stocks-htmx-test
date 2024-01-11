resource "helm_release" "otel-operator" {
  name       = "otel-operator"
  repository = "https://open-telemetry.github.io/opentelemetry-helm-charts"
  chart      = "opentelemetry-operator"
}

resource "helm_release" "otel-collector" {
  name       = "otel-collector"
  repository = "https://open-telemetry.github.io/opentelemetry-helm-charts"
  chart      = "opentelemetry-collector"
  values     = [file("values/otel-collector.yml")]

  depends_on = [
    helm_release.otel-operator
  ]
}

