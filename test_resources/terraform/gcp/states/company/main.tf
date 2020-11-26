terraform {
  backend "gcs" {
    bucket = "bucket-company"
    prefix = "terraform-impact"
  }
}