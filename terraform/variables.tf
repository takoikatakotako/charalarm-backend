locals {
  config = {
    development = {
      application_version          = "0.0.1"
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

    staging = {
      application_version          = "0.0.1"
      aws_profile                  = "charalarm-staging"
      route53_zone_id              = "Z1001429NNWJ0CTVGUIG"
      api_domain                   = "api.charalarm.swiswiswift.com"
      api_acm_certificate_arn      = "arn:aws:acm:ap-northeast-1:334832660826:certificate/05220010-3029-4b61-827a-fac783808a8c"
      application_bucket_name      = "application.charalarm.swiswiswift.com"
      resource_domain              = "resource.charalarm.swiswiswift.com"
      resource_bucket_name         = "resource.charalarm.swiswiswift.com"
      resource_bucket_url          = "https://s3.ap-northeast-1.amazonaws.com/resource.charalarm.swiswiswift.com"
      resource_acm_certificate_arn = "arn:aws:acm:us-east-1:334832660826:certificate/cbd20721-8637-4079-9843-37169da6daa9"
      datadog_log_forwarder_arn    = "arn:aws:lambda:ap-northeast-1:334832660826:function:datadog-forwarder"
    }

    production = {
      application_version          = "0.0.1"
      aws_profile                  = "charalarm-production"
      route53_zone_id              = "Z04405773OSCRT4AMPBDO"
      api_domain                   = "api.charalarm.com"
      api_acm_certificate_arn      = "arn:aws:acm:us-east-1:397693451628:certificate/cb4062b6-32b4-48c4-9d46-58c7a906846e"
      application_bucket_name      = "application.charalarm.com"
      resource_domain              = "resource.charalarm.com"
      resource_bucket_name         = "resource.charalarm.com"
      resource_bucket_url          = "https://s3.ap-northeast-1.amazonaws.com/resource.charalarm.com"
      resource_acn_certificate_arn = "arn:aws:acm:us-east-1:334832660826:certificate/cbd20721-8637-4079-9843-37169da6daa9"
      datadog_log_forwarder_arn    = ""
    }
  }

  aws_profile                  = local.config[terraform.workspace].aws_profile
  route53_zone_id              = local.config[terraform.workspace].route53_zone_id
  api_domain                   = local.config[terraform.workspace].api_domain
  api_acm_certificate_arn      = local.config[terraform.workspace].api_acm_certificate_arn
  application_version          = local.config[terraform.workspace].application_version
  application_bucket_name      = local.config[terraform.workspace].application_bucket_name
  resource_domain              = local.config[terraform.workspace].resource_domain
  resource_bucket_name         = local.config[terraform.workspace].resource_bucket_name
  resource_acm_certificate_arn = local.config[terraform.workspace].resource_acm_certificate_arn
  datadog_log_forwarder_arn    = local.config[terraform.workspace].datadog_log_forwarder_arn
}
