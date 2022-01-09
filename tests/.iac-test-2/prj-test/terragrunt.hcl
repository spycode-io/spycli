# Project Prj Test prj-test
locals {

  version = "0.0.0"

  prj_vars      = read_terragrunt_config(find_in_parent_folders("prj.hcl"))
  env_vars      = read_terragrunt_config(find_in_parent_folders("env.hcl"))
  region_vars   = read_terragrunt_config(find_in_parent_folders("region.hcl"))

  name          = local.prj_vars.locals.name  
  environment   = local.env_vars.locals.environment
  region        = local.region_vars.locals.region  
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

