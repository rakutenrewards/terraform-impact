# Terraform module: gcp/datadog/instance_group_monitor_set
# Creates a set of basic DataDog monitors for an instance_group.

# Datadog ---------------------------------------------------------------------

module "unhealthy_instances_monitor" {
  source  = "../standard_monitor"

  name    = "${var.name}_unhealthy"
  message = "${var.name}_unhealthy_msg"
}

module "high_memory_usage_monitor" {
  source  = "../standard_monitor"

  name    = "${var.name}_high_memory"
  message = "${var.name}_high_memory_msg"
}

module "high_cpu_usage_monitor" {
  source  = "../standard_monitor"

  name    = "${var.name}_high_cpu"
  message = "${var.name}_high_cpu_msg"
}

module "high_error_logs" {
  source  = "../standard_monitor"

  name    = "${var.name}_high_error_logs"
  message = "${var.name}_high_error_logs_msg"
}
