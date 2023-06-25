locals {
  config = {
    development = {
      aws_profile                  = "sandbox"
      route53_zone_id              = "Z06272247TSQ89OL8QZN"
      api_domain                   = "api.charalarm.sandbox.swiswiswift.com"
      api_acm_certificate_arn      = "arn:aws:acm:ap-northeast-1:397693451628:certificate/766e3ddf-1e97-406f-a3e8-32aedb8c5ce6"
      application_bucket_name      = "application.charalarm.sandbox.swiswiswift.com"
      resource_domain              = "resource.charalarm.sandbox.swiswiswift.com"
      resource_bucket_name         = "resource.charalarm.sandbox.swiswiswift.com"
      resource_bucket_url          = "https://s3.ap-northeast-1.amazonaws.com/resource.charalarm.sandbox.swiswiswift.com"
      resource_acm_certificate_arn = "arn:aws:acm:us-east-1:397693451628:certificate/6f024ec6-82c4-4412-b43e-e7095dc4195e"
      datadog_log_forwarder_arn    = "arn:aws:lambda:ap-northeast-1:397693451628:function:datadog-forwarder"
    }

    production = {
      aws_profile     = "charalarm"
      route53_zone_id = "Z04405773OSCRT4AMPBDO"
      # front_domain            = "review.swiswiswift.com"
      # front_certificate_arn   = "arn:aws:acm:us-east-1:772281501799:certificate/041caa0c-a884-4ef9-a746-2a1db6b8a28c"
      api_domain                = "api.sandbox.swiswiswift.com"
      api_acm_certificate_arn   = "arn:aws:acm:us-east-1:397693451628:certificate/cb4062b6-32b4-48c4-9d46-58c7a906846e"
      application_bucket_name   = "application.charalarm.com"
      resource_bucket_name      = "resource.charalarm.com"
      resource_bucket_url       = "https://s3.ap-northeast-1.amazonaws.com/resource.charalarm.com"
      datadog_log_forwarder_arn = ""
    }
  }

  aws_profile                  = local.config[terraform.workspace].aws_profile
  route53_zone_id              = local.config[terraform.workspace].route53_zone_id
  api_domain                   = local.config[terraform.workspace].api_domain
  api_acm_certificate_arn      = local.config[terraform.workspace].api_acm_certificate_arn
  application_version          = "0.0.1"
  application_bucket_name      = local.config[terraform.workspace].application_bucket_name
  resource_domain              = local.config[terraform.workspace].resource_domain
  resource_bucket_name         = local.config[terraform.workspace].resource_bucket_name
  resource_bucket_url          = local.config[terraform.workspace].resource_bucket_url
  resource_acm_certificate_arn = local.config[terraform.workspace].resource_acm_certificate_arn
  datadog_log_forwarder_arn    = local.config[terraform.workspace].datadog_log_forwarder_arn
}
