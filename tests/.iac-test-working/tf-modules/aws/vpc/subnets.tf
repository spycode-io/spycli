#Private subnets
resource "aws_subnet" "private" {
  count = length(var.private_subnets)
  vpc_id            = aws_vpc.this.id
  cidr_block        = var.private_subnets[count.index].cidr_block
  availability_zone = var.private_subnets[count.index].availability_zone

  tags = merge(var.tags, { 
    Name = var.private_subnets[count.index].name
  })
}

#Public subnets
resource "aws_subnet" "public" {
  count = length(var.public_subnets)
  vpc_id            = aws_vpc.this.id
  cidr_block        = var.public_subnets[count.index].cidr_block
  availability_zone = var.public_subnets[count.index].availability_zone
  map_public_ip_on_launch = true

  tags = merge(var.tags, { 
    Name = var.public_subnets[count.index].name
  })
}