resource "aws_apprunner_service" "app" {
  service_name = "${var.app_name}-${var.env}"

  source_configuration {
    image_repository {
      image_configuration {
        port = "8080"
        runtime_environment_variables = {
          "BASE_URL"             = "https://${var.app_base_url}",
          "LOG_JSON"             = "true",
          "LINKS_DYNAMODB_TABLE" = aws_dynamodb_table.links.name,
          "IP_STACK_TOKEN"       = var.ip_stack_token
        }
      }
      image_identifier      = "${aws_ecr_repository.gotidy.repository_url}:latest"
      image_repository_type = "ECR"
    }

    authentication_configuration {
      access_role_arn = aws_iam_role.apprunner_service_role.arn
    }

    auto_deployments_enabled = true
  }


  auto_scaling_configuration_arn = aws_apprunner_auto_scaling_configuration_version.app.arn

  health_check_configuration {
    healthy_threshold   = 1
    interval            = 10
    path                = "/api/health"
    protocol            = "HTTP"
    timeout             = 5
    unhealthy_threshold = 2

  }

  instance_configuration {
    cpu               = 256
    memory            = 512
    instance_role_arn = aws_iam_role.apprunner_instance_role.arn
  }

  tags = {
    App = var.app_name
    Env = var.env
  }
}

resource "aws_apprunner_auto_scaling_configuration_version" "app" {
  auto_scaling_configuration_name = "${var.app_name}-${var.env}-scaling-config"

  max_size = 10
  min_size = 1

  tags = {
    App = var.app_name
    Env = var.env
  }
}

resource "aws_apprunner_custom_domain_association" "app" {
  domain_name          = var.app_base_url
  service_arn          = aws_apprunner_service.app.arn
  enable_www_subdomain = false
}

output "app_certificate_validation_records" {
  value     = aws_apprunner_custom_domain_association.app.certificate_validation_records
  sensitive = true
}

output "app_dns_target" {
  value     = aws_apprunner_custom_domain_association.app.dns_target
  sensitive = true
}
