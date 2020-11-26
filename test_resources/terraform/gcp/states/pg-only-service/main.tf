terraform {
  backend "gcs" {
    bucket = "bucket-pg-only"
    prefix = "terraform-impact"
  }
}
