terraform {
  required_providers {
    sha = {
      source  = "terraform.local/local/terraform-provider-cryptography"
      version = "1.0.0"
    }
  }
}

data "sha" "example" {
  input = "value"
}


output "res" {
  value = data.sha.example.sha
}
