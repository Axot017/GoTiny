variable "env" {
  type = string

  validation {
    condition     = contains(["dev", "prod"], var.env)
    error_message = "env must be one of [dev, prod]"
  }
}

variable "region" {
  type = string
}

variable "app_name" {
  type = string
}

variable "app_base_url" {
  type = string
}

variable "ip_stack_token" {
  type = string
}
