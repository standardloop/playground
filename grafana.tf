// operator
data "kustomization_build" "grafana_operator" {
  path = "deploy/infrastructure/grafana/operator"
}

resource "kustomization_resource" "grafana_operator" {
  for_each   = data.kustomization_build.grafana_operator.ids
  manifest   = data.kustomization_build.grafana_operator.manifests[each.value]
  depends_on = [helm_release.flux2, kustomization_resource.flux_sources]
}

resource "null_resource" "wait_for_grafana_operator" {
  depends_on = [kustomization_resource.grafana_operator]

  provisioner "local-exec" {
    command = "kubectl wait helmrelease/grafana-operator --for=condition=Ready --timeout=180s -n grafana"
  }
}

// instance
module "wait_for_crds_needed_for_grafana_instance" {
    depends_on = [ null_resource.wait_for_grafana_operator ]
  source = "./modules/wait-for-crd"
  crds = [
    "grafanas.grafana.integreatly.org",
    "grafanadatasources.grafana.integreatly.org"
  ]
}

data "kustomization_build" "grafana_instance" {
  path = "deploy/infrastructure/grafana/instance"
}

resource "kustomization_resource" "grafana_instance" {
  for_each   = data.kustomization_build.grafana_instance.ids
  manifest   = data.kustomization_build.grafana_instance.manifests[each.value]
  depends_on = [helm_release.flux2, kustomization_resource.flux_sources, null_resource.wait_for_grafana_operator, module.wait_for_crds_needed_for_grafana_instance]
}

// dashboards
module "wait_for_crds_needed_for_grafana_dashboards" {
    depends_on = [ null_resource.wait_for_grafana_operator ]
  source = "./modules/wait-for-crd"
  crds = [
    "grafanadashboards.grafana.integreatly.org",
  ]
}

resource "helm_release" "grafana_dashboards" {
  depends_on = [module.wait_for_crds_needed_for_grafana_dashboards, kustomization_resource.grafana_instance]
  name       = "grafana-dashboards"
  chart      = "./deploy/infrastructure/grafana/dashboards"
  namespace  = "grafana"

  values = [
    file("./deploy/infrastructure/grafana/dashboards/values.yaml"),
    file("./deploy/infrastructure/grafana/dashboards/env/playground.values.yaml")
  ]
}
