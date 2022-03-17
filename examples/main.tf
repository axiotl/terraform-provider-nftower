terraform {
  required_providers {
    nftower = {
      source  = "hashicorp.com/edu/nftower"
      version = "0.2"
    }
  }
}

provider "nftower" {
  token = var.nf_tower_token
}

resource "nftower_compute_env" "main" {
  name           = "testing2"
  workspace_id   = var.nf_tower_workpace_id
  credentials_id = var.credentials_id
  config = {
    compute_job_role = "alsdfjalsdkfj"
    head_job_role    = "alsdkfalsdkjfalskdjf"
    region           = "us-east-1"
    work_dir         = "s3://convergence-beta-run"
    forge = {
      type           = "EC2"
      vpc_id         = var.vpc_id
      subnets        = [var.subnet]
      fusion_enabled = true
      min_cpus       = 0
      max_cpus       = 512
    }
  }
}
