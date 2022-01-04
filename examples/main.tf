terraform {
  required_providers {
    nftower = {
      version = "0.2"
      source  = "hashicorp.com/edu/nftower"
    }
  }
}

provider "nftower" {
  token = "eyJ0aWQiOiA0MjIwfS4yYzNlMmQ5NTI5MTVjODllNjIxMDU1ZGJjYWY5Y2IyZTNmOGU0Y2Rk"
}

resource "nftower_compute_env" "e" {
  name           = "test1"
  workspace_id   = "197562422694202"
  credentials_id = "7Y9dwU2JKHwuASqOdDLJ77"
  config = {

    region   = "us-east-1"
    work_dir = "s3://convergence-default-data"
    forge = {
      fusion_enabled = true
      min_cpus       = 0
      max_cpus       = 512
    }
  }
}


output "test" {
  value = nftower_compute_env.e
}
