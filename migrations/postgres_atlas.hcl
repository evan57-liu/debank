data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader/postgres",
  ]
}

variable "postgres_url" {
  type = string
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = var.postgres_url
  migration {
    dir = "file://postgres"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}