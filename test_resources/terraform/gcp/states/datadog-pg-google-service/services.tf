module "service-datadog" {
    source     = "../../modules/datadog/instance_group_monitor_set"

    name       = "full-blown-service"
}

module "service-pg" {
    source     = "../../modules/db/pg"

    project    = "full-blow-service"
}