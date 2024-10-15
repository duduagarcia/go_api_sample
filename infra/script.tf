provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "teste_t1" {
    ami = "ami-0866a3c8686eaeeba"
    instance_type = "t2.micro"

    vpc_security_group_ids = [aws_security_group.api_access.id]

    tags = {
      Name = "ec2-t1-example"
    }
}

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

  # Regra de saída que permite todo o tráfego
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

