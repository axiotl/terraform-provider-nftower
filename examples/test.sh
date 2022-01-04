#!/bin/bash

trash .terraform 
trash .terraform.lock.hcl
trash terraform.tfstate
terraform init
TF_LOG=debug terraform plan