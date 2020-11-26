resource "datadog_monitor" "standard_monitor" {
  message             = var.message
  name                = var.name

  lifecycle {
    ignore_changes = [silenced]
  }
}