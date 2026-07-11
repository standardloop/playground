resource "helm_release" "flux2" {
  name             = "flux2"
  repository       = "https://fluxcd-community.github.io/helm-charts"
  chart            = "flux2"
  namespace        = "flux-system"
  create_namespace = true
  set = [
    {
      name  = "kustomizeController.create"
      value = "false"
    },
    {
      name  = "notificationController.create"
      value = "false"
    },
    {
      name  = "imageAutomationController.create"
      value = "false"
    },
    {
      name  = "imageReflectionController.create"
      value = "false"
    }
  ]
}

data "kustomization_build" "flux_sources" {
  path = "deploy/infrastructure/sources"
}

resource "kustomization_resource" "flux_sources" {
  for_each   = data.kustomization_build.flux_sources.ids
  manifest   = data.kustomization_build.flux_sources.manifests[each.value]
  depends_on = [helm_release.flux2]
}
