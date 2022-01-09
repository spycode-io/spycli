include {
  path = find_in_parent_folders()
}

locals {
  
  prj_vars    = yamldecode(file(find_in_parent_folders("prj.yml")))
  env_vars    = yamldecode(file(find_in_parent_folders("env.yml")))
  region_vars = yamldecode(file(find_in_parent_folders("region.yml")))
  
  prj_slug_name = local.prj_vars.slug_name
  environment = local.env_vars.environment
  region = local.region_vars.region

  name = "vpc-${local.environment}"
  module = "aws/vpc"

  library = local.env_vars.library
  library_version = local.env_vars.library_version
  tf_module_source = "${local.env_vars.library}//${local.module}${local.library_version}"

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
