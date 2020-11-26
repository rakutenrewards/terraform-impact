resource "datadog_monitor" "standard_monitor" {
  name                = "default_name"
  message             = "default_msg"

  lifecycle {
    ignore_changes = [silenced]
  }
}