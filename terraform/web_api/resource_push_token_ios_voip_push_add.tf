##################################################
# /push-token/ios/voip-push/add
##################################################
resource "aws_api_gateway_resource" "push_token_ios_voip_push_add_resource" {
  rest_api_id = aws_api_gateway_rest_api.charalarm_rest_api.id
  parent_id   = aws_api_gateway_resource.push_token_ios_voip_push_resource.id
  path_part   = "add"
}

##################################################
# Lambda
##################################################
module "push_token_ios_voip_push_add_post_lambda_function" {
  source        = "./lambda_function"
  function_name = "push-token-ios-voip-push-add-post-function"
  role          = aws_iam_role.api_gateway_lambda_role.arn
  handler       = "push_token_ios_voip_push_add"
  s3_bucket     = local.application_bucket_s3_url
  s3_key        = "/${var.application_version}/push_token_ios_voip_push_add.zip"
  execution_arn = aws_api_gateway_rest_api.charalarm_rest_api.execution_arn
  method        = "POST"
  path          = "/push-token/ios/voip-push/add"
}

##################################################
# Method
##################################################
module "push_token_ios_voip_push_add_post_method" {
  source      = "./method"
  rest_api_id = aws_api_gateway_rest_api.charalarm_rest_api.id
  resource_id = aws_api_gateway_resource.push_token_ios_voip_push_add_resource.id
  http_method = "POST"
  lambda_uri  = module.push_token_ios_push_add_post_lambda_function.invoke_arn
}

module "push_token_ios_voip_push_add_options_method" {
  source      = "./options_method"
  rest_api_id = aws_api_gateway_rest_api.charalarm_rest_api.id
  resource_id = aws_api_gateway_resource.push_token_ios_voip_push_add_resource.id
}
