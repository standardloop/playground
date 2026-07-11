terraform {
  required_providers {
    kustomization = {
      source = "kbst/kustomization"
    }
    helm = {
      source = "hashicorp/helm"
    }
    null = {
      source = "hashicorp/null"
    }
  }
}
