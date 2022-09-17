##################################################
# Lambda
##################################################
resource "aws_lambda_function" "lambda_function" {
  function_name = var.function_name
  role          = var.role
  runtime       = "go1.x"
  handler       = var.handler
  timeout       = 15
  s3_bucket     = var.s3_bucket
  s3_key        = var.s3_key

  architectures = [
    "x86_64"
  ]

  environment {
    variables = {
      DISCORD_WEBHOOK_URL = "xxx"
    }
  }
}


##################################################
# Log Group
##################################################
resource "aws_cloudwatch_log_group" "log_group" {
  name              = "/aws/lambda/${aws_lambda_function.lambda_function.function_name}"
  retention_in_days = 90
}


##################################################
# Permission
##################################################
resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "api-gateway-${var.function_name}-statement-id"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_function.arn
  principal     = "apigateway.amazonaws.com"

  # The /*/*/* part allows invocation from any stage, method and resource path
  source_arn = "${var.execution_arn}/*/${var.method}${var.path}"
}


# aws lambda get-policy --function-name user-signup-anonymous-post-function --profile sandbox | jq


# user-signup-anonymous-post-function/e081b2d8-db12-4c3e-a167-130bb7a56591

# terraform import module.web_api.module.user_signup_anonymous_post_lambda_function.aws_lambda_permission.lambda_permission user-signup-anonymous-post-function/e081b2d8-db12-4c3e-a167-130bb7a56591
