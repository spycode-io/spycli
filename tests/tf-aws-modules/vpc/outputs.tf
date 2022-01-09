output "vpc_id" {
  value = aws_vpc.this.id
  description = "VPC id"
}

output "vpc_arn" {
  value = aws_vpc.this.arn
  description = "VPC arn"
}

output "private_subnet_ids" {
  value = aws_subnet.private.*.id
  description = "Private subnets id"
}

output "private_subnet_arn" {
  value = aws_subnet.private.*.arn
  description = "Private subnets arn"
}

output "public_subnet_ids" {
  value = aws_subnet.public.*.id
  description = "Private subnets id"
}

output "public_subnet_arn" {
  value = aws_subnet.public.*.arn
  description = "Private subnets arn"
}