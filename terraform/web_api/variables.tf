variable "domain" {
  type = string
}

variable "route53_zone_id" {
  type = string
}

variable "acm_certificate_arn" {
  type = string
}

variable "application_version" {
  type = string
}

variable "application_bucket_name" {
  type = string
}


locals {
  application_bucket_s3_url = "s3://${aws_s3_bucket.application_bucket.bucket}"
  variables = {
    "BASE_URL" = "https://${var.domain}/"
  }
}