variable "tags" {
    type = map(string)
}

variable "vms" {
    type = list(object({
        name = string
        instance_type = string
        availability_zone = string
        subnet_id = string
        private_ips = list(string)
        security_groups = list(string)
    }))
}