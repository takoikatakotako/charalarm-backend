variable "resource_bucket_url" {
  type = string
}

locals {
  environment_variables = {
    "RESOURCE_BASE_URL" = var.resource_bucket_url
  }
}