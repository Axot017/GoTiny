resource "aws_dynamodb_table" "links" {
  name         = "links-${var.app_name}-${var.env}"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "PK"
  range_key    = "SK"

  attribute {
    name = "PK"
    type = "S"
  }

  attribute {
    name = "SK"
    type = "S"
  }

  attribute {
    name = "LSI_1"
    type = "S"
  }

  attribute {
    name = "LSI_2"
    type = "S"
  }

  attribute {
    name = "LSI_3"
    type = "S"
  }

  attribute {
    name = "LSI_4"
    type = "N"
  }

  attribute {
    name = "LSI_5"
    type = "N"
  }

  attribute {
    name = "GSI_1_PK"
    type = "S"
  }

  attribute {
    name = "GSI_1_SK"
    type = "S"
  }

  local_secondary_index {
    name            = "LSI_1"
    range_key       = "LSI_1"
    projection_type = "ALL"
  }

  local_secondary_index {
    name            = "LSI_2"
    range_key       = "LSI_2"
    projection_type = "ALL"
  }

  local_secondary_index {
    name            = "LSI_3"
    range_key       = "LSI_3"
    projection_type = "ALL"
  }

  local_secondary_index {
    name            = "LSI_4"
    range_key       = "LSI_4"
    projection_type = "ALL"
  }

  local_secondary_index {
    name            = "LSI_5"
    range_key       = "LSI_5"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "GSI_1"
    hash_key        = "GSI_1_PK"
    range_key       = "GSI_1_SK"
    projection_type = "ALL"
  }

  ttl {
    attribute_name = "TTL"
    enabled        = true
  }

  point_in_time_recovery {
    enabled = true
  }

  tags = {
    App = var.app_name
    Env = var.env
  }
}

output "dynamodb_table_name" {
  value = aws_dynamodb_table.links.name
}

output "dynamodb_table_arn" {
  value = aws_dynamodb_table.links.arn
}
