##################################################
# /push-token
##################################################
resource "aws_api_gateway_resource" "push_token_resource" {
  rest_api_id = aws_api_gateway_rest_api.charalarm_rest_api.id
  parent_id   = aws_api_gateway_rest_api.charalarm_rest_api.root_resource_id
  path_part   = "push-token"
}
