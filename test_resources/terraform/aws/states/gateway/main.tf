terraform {
  backend "s3" {
    region         = "us-west-2"
    bucket         = "curbside-terraform-us-west-2"
    key            = "production.gw.tfstate"
    dynamodb_table = "terraform-lock"
  }
}

data "terraform_remote_state" "gw" {
  backend = "s3"
  config {
    region         = "us-west-2"
    bucket         = "aws-gw-us-west-2"
    key            = "production.gw.tfstate"
    dynamodb_table = "terraform-lock"
  }
}

provider "aws" {
  region = "${var.global_aws_region}"
}

module "pg" {
  source = "./modules/db/pg"

  name   = "aws-pg"
}

module "monitor" {
  // purposely doesn't start with ./
  source = "modules/db/pg/monitor"

  name   = "aws-pg"
}