variable "tags" {
    type = map(string)
}

variable "cidr_block" {
    default = "10.0.0.0/16"
}

variable "private_subnets" {
    type = list(object({
      name = string
      cidr_block = string
      availability_zone = string            
    }))

    default = [ {
      name = "prv-subnet"
      availability_zone = "us-east-1a"
      cidr_block = "10.0.1.0/24"
    } ]
}

variable "public_subnets" {
    type = list(object({
      name = string
      cidr_block = string
      availability_zone = string            
    }))

    default = [ {
      name = "pub-subnet"
      availability_zone = "us-east-1a"
      cidr_block = "10.0.2.0/24"
    } ]
}