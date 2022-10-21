##################################################
# Lambdaにアタッチするためのロール
##################################################
resource "aws_iam_role" "batch_lambda_role" {
  name               = "batch-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.batch_lambda_role_assume_policy_document.json
}

data "aws_iam_policy_document" "batch_lambda_role_assume_policy_document" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}
