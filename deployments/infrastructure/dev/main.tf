terraform {
  backend "s3" {
    bucket         = "go-tiny-tf-state"
    key            = "dev/terraform.tfstate"
    region         = "eu-central-1"
    dynamodb_table = "go-tiny-tf-state-lock"
    encrypt        = true
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.5"
    }
  }
}

provider "aws" {
  region = var.region
}

module "app" {
  source = "../base"

  env          = "dev"
  app_name     = var.app_name
  region       = var.region
  app_base_url = "dev.goti.one"
}

output "all_outputs" {
  value     = module.app
  sensitive = true
}
