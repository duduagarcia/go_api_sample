provider "aws" {
  region = "us-east-1"
}

# Cria uma EC2
resource "aws_instance" "teste_t1" {
    ami = "ami-0866a3c8686eaeeba"
    instance_type = "t2.micro"

    vpc_security_group_ids = [aws_security_group.api_access.id]

    key_name = aws_key_pair.my_key_pair.key_name
    tags = {
      Name = "ec2-t1"
    }
}

# Cria um security group(firewall) para ec2
resource "aws_security_group" "api_access" {
  name        = "API-security-group-T1"
  description = "Security group para permitir SSH, HTTP e HTTPS"

  # Regra de entrada para SSH (porta 22)
  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Regra de entrada para HTTP (porta 80)
  ingress {
    description = "HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Regra de entrada para HTTPS (porta 443)
  ingress {
    description = "HTTPS"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # porta da api 
  ingress {
    description = "api port"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # porta do banco de dados
  ingress {
    description = "pg port"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # porta do pgAdmin
    ingress {
    description = "pgAdmin port"
    from_port   = 5050
    to_port     = 5050
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Regra de saída que permite todo o tráfego
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # pode ser uma boa botar uma regra para liberar um ICMP, para possíveis debugs com ping
}

# Criação da ssh keypair diretamente no terraform para acessar a ec2
resource "tls_private_key" "my_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "my_key_pair" {
  key_name   = "my-ec2-key"
  public_key = tls_private_key.my_key.public_key_openssh
}

output "private_key_pem" {
  value     = tls_private_key.my_key.private_key_pem
  sensitive = true
}