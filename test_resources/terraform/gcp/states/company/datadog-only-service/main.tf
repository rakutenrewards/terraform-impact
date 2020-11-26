terraform {
  backend "gcs" {
    bucket = "bucket-dg"
    prefix = "terraform-impact"
  }

  required_providers {
    datadog = {
      source  = "terraform-providers/datadog"
      version = "~> 2.7.0"
    }
  }
}

data "terraform_remote_state" "admin" {
  backend = "gcs"
  config = {
    bucket = "bucket-dg"
    prefix = "terraform-impact"
  }
}

data "terraform_remote_state" "xpn_host" {
  backend = "gcs"
  config = {
    bucket = "bucket-dg"
    prefix = "terraform-impact"
  }
}

provider "datadog" {
  api_key = "some_api_key"
  app_key = "some_app_key"
}

module "datadog" {
  source  = "../../../modules/datadog/standard_monitor"

  name    = "datadog-only"
  message = "datadog-only-message"
}