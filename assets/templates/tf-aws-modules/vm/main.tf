#Chave rsa
resource "aws_key_pair" "this" {
  key_name   = "${var.tags.Name}-key"
  public_key = file("~/.ssh/id_rsa.pub")
}

#Placa de rede que será usada pela instância
resource "aws_network_interface" "this" {
  count = length(var.vms)
  subnet_id = var.vms[count.index].subnet_id
  private_ips = var.vms[count.index].private_ips
  security_groups = var.vms[count.index].security_groups

  tags =  merge(var.tags, { 
    Name = "neti-${var.tags.Name}" 
  })
}

#IP públic
# resource "aws_eip" "ips" {
#   count = 3
#   vpc                       = true
#   network_interface         = aws_network_interface.interface[count.index].id
#   associate_with_private_ip = "10.0.0.10${count.index}"
# }

#Instância aline-vm-01
resource "aws_instance" "this" {
  count             = length(var.vms)
  ami               = data.aws_ami.ubuntu.id
  instance_type     = var.vms[count.index].instance_type
  availability_zone = var.vms[count.index].availability_zone
  key_name          = aws_key_pair.this.key_name
  
  network_interface {
    network_interface_id = aws_network_interface.this[count.index].id
    device_index         = 0
  }

  tags =  merge(var.tags, { 
    Name = "vm-${var.vms[count.index].name}" 
  })
}

