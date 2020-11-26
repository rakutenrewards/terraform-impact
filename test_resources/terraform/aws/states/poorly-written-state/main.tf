terraform {
  backend "s3" {
    region         = "us-west-2"
    bucket         = "curbside-terraform-us-west-2"
    key            = "production.poorly_written.tfstate"
    dynamodb_table = "terraform-lock"
  }
}

provider "aws" {
  region = var.global_aws_region
}

module "pg" {
  source = "../gateway/modules/db/pg"

  name   = "aws-pg"
}

// this makes purposely no sense
module "not_existing" {
  source = "modules/bob"
}