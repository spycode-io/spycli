include {
  path = find_in_parent_folders()
}

locals {
  
  prj_vars    = read_terragrunt_config(find_in_parent_folders("prj.hcl"))
  env_vars    = read_terragrunt_config(find_in_parent_folders("env.hcl"))
  region_vars = read_terragrunt_config(find_in_parent_folders("region.hcl"))
  
  prj_slug_name = local.prj_vars.locals.slug_name
  environment = local.env_vars.locals.environment
  region = local.region_vars.locals.region

  name = "test-vpc-${local.environment}"
  module = "vpc"

  library = local.env_vars.locals.library
  library_version = local.env_vars.locals.library_version
  tf_module_source = "${local.env_vars.locals.library}//${local.module}${local.library_version}"

  tags = {
    Name = local.name
    Project = local.prj_slug_name
    Environment = local.environment
    Region = local.region
    Module = local.module
    Source = local.tf_module_source
  }
}

terraform {
  source = local.tf_module_source
}

inputs = {
  tags = local.tags
}
