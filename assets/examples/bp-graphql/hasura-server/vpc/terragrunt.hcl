include {
  path = find_in_parent_folders()
}

locals {
  prj_cfg    = read_terragrunt_config(find_in_parent_folders("prj.hcl"))
  env_cfg    = read_terragrunt_config(find_in_parent_folders("env.hcl"))
  region_cfg = read_terragrunt_config(find_in_parent_folders("region.hcl"))
  
  prj_name = local.prj_cfg.locals.name
  prj_slug_name = local.prj_cfg.locals.slug_name
  stack = local.prj_cfg.locals.stack
  environment = local.env_cfg.environment  
  tf_lib_version = local.env_cfg.tf_lib_version

  tf_module = "vpc"
  tf_lib = local.prj_cfg.locals.tf_lib  
  tf_module_source = "${local.tf_lib}//${local.tf_module}?ref=${local.tf_lib_version}"
}

terraform {
  source = local.tf_module_source
}

inputs = {
  prj_name = local.prj_name
  prj_slug_name = local.prj_slug_name
  stack = local.stack
  environment = local.environment
  tf_lib = local.tf_lib
  tf_module = local.tf_module
  tf_lib_version = local.tf_lib_version
}
