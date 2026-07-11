data "kustomization_build" "metrics_server" {
  path = "deploy/infrastructure/metrics-server"
}

resource "kustomization_resource" "metrics_server" {
  for_each   = data.kustomization_build.metrics_server.ids
  manifest   = data.kustomization_build.metrics_server.manifests[each.value]
  depends_on = [helm_release.flux2, kustomization_resource.flux_sources]
}

resource "null_resource" "wait_for_metrics_server" {
  depends_on = [kustomization_resource.metrics_server]

  provisioner "local-exec" {
    command = "kubectl wait helmrelease/metrics-server --for=condition=Ready --timeout=${var.HELM_RELEASE_TIMEOUT} -n kube-system"
  }
}
