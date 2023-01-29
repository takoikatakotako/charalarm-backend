##################################################
# /user/signup/anonymous
##################################################
resource "aws_api_gateway_resource" "user_signup_anonymous_resource" {
  rest_api_id = aws_api_gateway_rest_api.charalarm_rest_api.id
  parent_id   = aws_api_gateway_resource.user_signup_resource.id
  path_part   = "anonymous"
}

##################################################
# Lambda
##################################################
module "user_signup_anonymous_post_lambda_function" {
  source        = "./lambda_function"
  function_name = "user-signup-anonymous-post-function"
  role          = aws_iam_role.api_gateway_lambda_role.arn
  handler       = "user_signup_anonymous_user"
  s3_bucket     = local.application_bucket_s3_url
  s3_key        = "/${var.application_version}/signup_anonymous_user.zip"
  execution_arn = aws_api_gateway_rest_api.charalarm_rest_api.execution_arn
  method        = "POST"
  path          = "/user/signup/anonymous"
}

##################################################
# Method
##################################################
module "user_signup_anonymous_post_method" {
  source      = "./method"
  rest_api_id = aws_api_gateway_rest_api.charalarm_rest_api.id
  resource_id = aws_api_gateway_resource.user_signup_anonymous_resource.id
  http_method = "POST"
  lambda_uri  = module.user_signup_anonymous_post_lambda_function.invoke_arn
}

module "user_signup_anonymous_options_method" {
  source      = "./options_method"
  rest_api_id = aws_api_gateway_rest_api.charalarm_rest_api.id
  resource_id = aws_api_gateway_resource.user_signup_anonymous_resource.id
}

