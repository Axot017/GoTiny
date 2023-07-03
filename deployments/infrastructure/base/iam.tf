resource "aws_iam_role" "apprunner_service_role" {
  name = "${var.app_name}-${var.env}-apprunner-service-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Principal = {
          Service = "build.apprunner.amazonaws.com"
        }
        Effect = "Allow"
        Sid    = ""
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "apprunner_service_role_attachment" {
  role       = aws_iam_role.apprunner_service_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSAppRunnerServicePolicyForECRAccess"
}

resource "aws_iam_role" "apprunner_instance_role" {
  name = "${var.app_name}-${var.env}-apprunner-instance-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Principal = {
          Service = "tasks.apprunner.amazonaws.com"
        }
        Effect = "Allow"
        Sid    = ""
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "apprunner_instance_role" {
  role       = aws_iam_role.apprunner_instance_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
}
