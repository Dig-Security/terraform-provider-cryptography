terraform {
  required_providers {
    sha = {
      source  = "terraform.local/local/sha"
      version = "1.0.0"
    }
  }
}

data "sha" "example" {
  input = local.input
}


output "res" {
  value = data.sha.example.sha
}

output "in" {
  value = local.input
}

locals {
  input  = textdecodebase64 (base64sha512("d9241543-c5ca-4d08-8291-ff41775f5af7"), "ISO-8859-1")
}