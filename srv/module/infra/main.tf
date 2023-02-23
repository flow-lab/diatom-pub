locals {
  config_dir = "${path.module}/sql"
}

# read file from local disk
variable "name" {
  default = "diatom-pub"
}

resource "kubernetes_config_map" "sql_files" {
  metadata {
    name = "${var.name}-sql-files"
  }

  data = {
    for f in fileset(local.config_dir, "**/*.sql") :
    f => file("${local.config_dir}/${f}")
  }
}