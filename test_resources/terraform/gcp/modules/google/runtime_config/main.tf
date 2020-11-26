# Terraform module: gcp/google/service_config

resource "google_runtimeconfig_config" "config" {
  name    = "${var.name}_config"
  project = var.project
}
