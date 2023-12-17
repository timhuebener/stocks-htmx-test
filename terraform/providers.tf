terraform {
  required_version = ">= 1.5"

  required_providers {
    helm = {
      source  = "hashicorp/helm"
      version = "2.12.1"
    }
    kubectl = {
      source  = "gavinbunney/kubectl"
      version = "1.14.0"
    }
  }
}

provider "helm" {
  kubernetes {
    config_path    = "~/.kube/config"
    config_context = "docker-desktop"
  }
}

provider "kubectl" {
  config_path    = "~/.kube/config"
  config_context = "docker-desktop"
}
