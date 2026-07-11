provider "kustomization" {
  kubeconfig_path = "~/.kube/config"
  context         = "kind-slke-1"
}

provider "helm" {
  kubernetes = {
    config_path    = "~/.kube/config"
    config_context = "kind-slke-1"
  }
}
