resource "aws_ecr_repository" "gotidy" {
  name = "${var.app_name}-${var.env}"

  tags = {
    Environment = "${var.env}"
    AppName     = "${var.app_name}"
  }
}
