# Project Prj Test prj-test
locals {

  version = "0.0.0"

  prj_vars    = yamldecode(file(find_in_parent_folders("prj.yml")))
  env_vars    = yamldecode(file(find_in_parent_folders("env.yml")))
  region_vars = yamldecode(file(find_in_parent_folders("region.yml")))

  name          = local.prj_vars.name  
  environment   = local.env_vars.environment
  region        = local.region_vars.region  
}

# Generate an azure provider block
generate "provider" {
  path      = "provider.tf"
  if_exists = "overwrite_terragrunt"
  
  contents  = <<EOF
provider "aws" {
  region  = "${local.region}"
}
EOF
}

