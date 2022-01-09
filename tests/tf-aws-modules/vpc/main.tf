locals {
  ig_tags = merge(var.tags, { Name = "igw-${var.tags.Name}" })
}

#VPC
resource "aws_vpc" "this" {
  cidr_block           = var.cidr_block
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = var.tags
}

#Internet gateway
resource "aws_internet_gateway" "this" {
  vpc_id = aws_vpc.this.id
  tags =  merge(var.tags, { 
    Name = "igw-${var.tags.Name}" 
  })
}

#Route table
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.this.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.this.id    
  }

  tags =  merge(var.tags, { 
    Name = "rt-${var.tags.Name}" 
  })
}

#Route table association
resource "aws_route_table_association" "public" {
  count = length(var.public_subnets)
  subnet_id      = "${aws_subnet.public[count.index].id}"
  route_table_id = "${aws_route_table.public.id}"
}
