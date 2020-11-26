provider "random" {
  version = "~> 2.1.2"
}

resource "google_project_service" "network" {
  disable_on_destroy = false
  project            = var.project
  service            = "servicenetworking.googleapis.com"
}

resource "google_sql_database_instance" "master" {
  provider = google-beta
}

module "pg-monitors" {
  source     = "../../datadog/standard_monitor"

  name       = "postgres"
  message    = "postgres_message"
}