terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.62.0"
    }
  }
}

provider "aws" {
  access_key          = "mockaccesskey"
  secret_key          = "mocksecretkey"
  region              = "ap-northeast-1"
  skip_credentials_validation = true
  skip_requesting_account_id  = true
  skip_metadata_api_check = true
}