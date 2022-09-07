# DynamoDB
resource "aws_dynamodb_table" "user_table" {
  name           = "user-table"
  hash_key       = "userId"
  billing_mode   = "PROVISIONED"
  read_capacity  = 1
  write_capacity = 1
  stream_enabled = false

  attribute {
    name = "userId"
    type = "S"
  }
}
