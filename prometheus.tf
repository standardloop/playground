// base
data "kustomization_build" "prometheus" {
  path = "deploy/infrastructure/kube-prom-stack"
}

resource "kustomization_resource" "prometheus" {
  for_each   = data.kustomization_build.prometheus.ids
  manifest   = data.kustomization_build.prometheus.manifests[each.value]
  depends_on = [helm_release.flux2, kustomization_resource.flux_sources]
}

resource "null_resource" "wait_for_prometheus" {
  depends_on = [kustomization_resource.prometheus]

  provisioner "local-exec" {
    command = "kubectl wait helmrelease/kube-prometheus-stack --for=condition=Ready --timeout=180s -n kube-prometheus-stack"
  }
}

// scrapers

module "wait_for_crds_needed_for_scrapers" {
  source = "./modules/wait-for-crd"
  // the k8s gateway api ones come from cloud-provider-kind
  crds = [
    "podmonitors.monitoring.coreos.com",
    "servicemonitors.monitoring.coreos.com",
  ]
  depends_on = [null_resource.wait_for_prometheus]
}

data "kustomization_build" "prometheus_scrapers" {
  path = "deploy/infrastructure/kube-prom-stack/scrapers"
}

resource "kustomization_resource" "prometheus_scrapers" {
  for_each   = data.kustomization_build.prometheus_scrapers.ids
  manifest   = data.kustomization_build.prometheus_scrapers.manifests[each.value]
  depends_on = [module.wait_for_crds_needed_for_scrapers]
}
