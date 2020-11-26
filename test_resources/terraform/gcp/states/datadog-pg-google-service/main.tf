terraform {
  backend "gcs" {
    bucket = "bucket"
    prefix = "terraform-impact"
  }

  required_providers {
    datadog = {
      source  = "terraform-providers/datadog"
      version = "~> 2.7.0"
    }
    google = {
      source = "hashicorp/google"
      version = "~> 3.23.0"
    }
    google-beta = {
      source = "hashicorp/google-beta"
      version = "~> 3.23.0"
    }
  }
}

data "terraform_remote_state" "admin" {
  backend = "gcs"
  config = {
    bucket = "bucket"
    prefix = "terraform-impact"
  }
}

data "terraform_remote_state" "xpn_host" {
  backend = "gcs"
  config = {
    bucket = "bucket"
    prefix = "terraform-impact"
  }
}

# Providers ------------------------------------------------------------------
provider "google" {
  project = "terraform-impact-s1234"
}

provider "google-beta" {
  project = "terraform-impact-s1234"
}

provider "datadog" {
  api_key = "some_api_key"
  app_key = "some_app_key"
}