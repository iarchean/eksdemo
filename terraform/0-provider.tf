provider "aws" {
  region = "ap-northeast-1"
  profile = "baidu"
}

terraform {
  required_providers {
    aws = {
        source = "hashicorp/aws"
        version = "~> 3.0"
    }
  }
  backend "s3" {
    bucket         = "terraform-backend.baidujp"
    key            = "eks"
    region         = "ap-northeast-1"
    # encrypt        = true
    # dynamodb_table = "terraform-lock"
  }
}