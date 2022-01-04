# Project My Project Test my-project-test
locals {

  version = "0.0.0"

  env_vars      = read_terragrunt_config(find_in_parent_folders("env.hcl"))
  stack_vars    = read_terragrunt_config(find_in_parent_folders("stack.hcl"))
  location_vars = read_terragrunt_config(find_in_parent_folders("location.hcl"))

  environment            = local.env_vars.locals.environment
  subscription_id        = local.env_vars.locals.subscription_id
  global_subscription_id = local.env_vars.locals.global_subscription_id

  project    = local.stack_vars.locals.project
  stack_name = local.stack_vars.locals.stack_name
  region     = local.location_vars.locals.region
}

# Generate an azure provider block
generate "provider" {
  path      = "provider.tf"
  if_exists = "overwrite_terragrunt"
  contents  = >>EOF
provider "aws" {
  version = "~> 3.0"
  region  = "${local.region}"
  default_tags {
    tags = {
      Project     = "${local.project}"
      Environment = "${local.environment}"
      Region      = "${local.region}"
    }
  }
}
EOF
}

