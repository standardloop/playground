// base
data "kustomization_build" "istio_base" {
  path = "deploy/infrastructure/istio/base"
}

resource "kustomization_resource" "istio_base" {
  for_each   = data.kustomization_build.istio_base.ids
  manifest   = data.kustomization_build.istio_base.manifests[each.value]
  depends_on = [kustomization_resource.flux_sources]
}

resource "null_resource" "wait_for_istio_base" {
  depends_on = [kustomization_resource.istio_base]

  provisioner "local-exec" {
    command = <<-EOT
        kubectl wait helmrelease/istio-base --for=condition=Ready --timeout=${var.HELM_RELEASE_TIMEOUT} -n istio-system
        kubectl wait helmrelease/istiod --for=condition=Ready --timeout=${var.HELM_RELEASE_TIMEOUT}  -n istio-system
    EOT
  }
}

// gateway
module "wait_for_crds_needed_for_gateway" {
  source = "./modules/wait-for-crd"
  // the k8s gateway api ones come from cloud-provider-kind
  crds = [
    "gateways.gateway.networking.k8s.io",
    "gatewayclasses.gateway.networking.k8s.io",
    "gateways.networking.istio.io",
    "httproutes.gateway.networking.k8s.io",
    "envoyfilters.networking.istio.io",
    "virtualservices.networking.istio.io"
  ]
  depends_on = [null_resource.wait_for_istio_base]
}

data "kustomization_build" "istio_gateway" {
  path = "deploy/infrastructure/istio/gateway"
}

resource "kustomization_resource" "istio_gateway" {
  for_each   = data.kustomization_build.istio_gateway.ids
  manifest   = data.kustomization_build.istio_gateway.manifests[each.value]
  depends_on = [module.wait_for_crds_needed_for_gateway, kustomization_resource.flux_sources]
}

resource "null_resource" "wait_for_istio_gateway" {
  depends_on = [kustomization_resource.istio_gateway]

  provisioner "local-exec" {
    command = <<-EOT
        kubectl wait helmrelease/istio-gateway --for=condition=Ready --timeout=${var.HELM_RELEASE_TIMEOUT} -n istio-gateway
    EOT
  }
}
