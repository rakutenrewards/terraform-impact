module "service-pg" {
    source     = "../../modules/db/pg"

    project    = "pg-only"
}