##################################################
# /user/signup
##################################################
resource "aws_api_gateway_resource" "user_signup_resource" {
  rest_api_id = aws_api_gateway_rest_api.charalarm_rest_api.id
  parent_id   = aws_api_gateway_resource.user_resource.id
  path_part   = "signup"
}


