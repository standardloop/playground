data "kustomization_build" "argo_rollouts" {
  path = "deploy/infrastructure/argo-rollouts"
}

resource "kustomization_resource" "argo_rollouts" {
  depends_on = [helm_release.flux2, kustomization_resource.flux_sources]

  for_each = data.kustomization_build.argo_rollouts.ids
  manifest = data.kustomization_build.argo_rollouts.manifests[each.value]

}

resource "null_resource" "wait_for_argo_rollouts" {
  depends_on = [kustomization_resource.argo_rollouts]

  provisioner "local-exec" {
    command = "kubectl wait helmrelease/argo-rollouts --for=condition=Ready --timeout=${var.HELM_RELEASE_TIMEOUT} -n argo-rollouts"
  }
}
