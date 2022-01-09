include {
  path = find_in_parent_folders()
}

dependency "vpc" {
  config_path = "../test-vpc"
}

locals {
  
  prj_vars    = yamldecode(file(find_in_parent_folders("prj.yml")))
  env_vars    = yamldecode(file(find_in_parent_folders("env.yml")))
  region_vars = yamldecode(file(find_in_parent_folders("region.yml")))
  
  prj_slug_name = local.prj_vars.slug_name
  environment = local.env_vars.environment
  region = local.region_vars.region

  name = "vm-${local.environment}"
  module = "aws/vm"

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
  vms = [
    {
      name = "app"
      instance_type = "t3.micro"
      availability_zone = "us-east-1a"
      subnet_id = dependency.vpc.outputs.private_subnet_ids[0]
      private_ips = ["10.0.1.100"]
      security_groups = []
    }
  ]
}
