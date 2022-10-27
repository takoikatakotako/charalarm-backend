provider "aws" {
  profile                  = local.aws_profile
  region                   = "ap-northeast-1"
  shared_credentials_files = ["~/.aws/credentials"]
}

module "dynamodb" {
  source = "./dynamodb"
}

# module "application" {
#   source = "./application"
#   bucket_name = local.application_bucket_name
# }

# module "web_front" {
#   source              = "./web_front"
#   domain              = local.front_domain
#   route53_zone_id     = local.route53_zone_id
#   acm_certificate_arn = local.front_acm_certificate_arn
# }

module "web_api" {
  source                  = "./web_api"
  domain                  = local.api_domain
  route53_zone_id         = local.route53_zone_id
  acm_certificate_arn     = local.api_acm_certificate_arn
  application_version     = local.application_version
  application_bucket_name = local.application_bucket_name
}


module "batch" {
  source = "./batch"

}



module "datadog" {
  source = "./datadog"
  dd_api_key = "DD_API_KEY"
}
