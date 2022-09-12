##################################################
# Lambda
##################################################
resource "aws_lambda_function" "lambda_function" {
  function_name    = var.function_name
  role             = var.role
  runtime          = "go1.x"
  handler          = var.handler
  timeout          = 15
  filename         = data.archive_file.python_script_archive_file.output_path
  source_code_hash = data.archive_file.python_script_archive_file.output_base64sha256

  architectures = [
    "x86_64"
  ]

  environment {
    variables = {
      DISCORD_WEBHOOK_URL = "xxx"
    }
  }
}

data "archive_file" "python_script_archive_file" {
  type        = "zip"
  source_file = "${path.root}/../application/build/${var.filename}"
  output_path = "${path.root}/../application/build/${var.archive_filename}"

  # source {
  #   content  = file("${path.root}/../application/build/${var.filename}")
  #   filename = "lambda_function.py"
  # }
}

# data "template_file" "python_script_file" {
#   template = file("${path.root}/../application/build/${var.filename}")
# }


resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "api-gateway-${var.function_name}-statement-id"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_function.function_name
  principal     = "apigateway.amazonaws.com"

  # The /*/*/* part allows invocation from any stage, method and resource path
  source_arn = "${var.execution_arn}/*/${var.method}${var.path}"
}
