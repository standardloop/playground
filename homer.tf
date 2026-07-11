data "kustomization_build" "homer" {
  path = "deploy/infrastructure/homer"
}

resource "kustomization_resource" "homer" {
  for_each = data.kustomization_build.homer.ids
  manifest = data.kustomization_build.homer.manifests[each.value]
}
