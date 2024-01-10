
variable "name" {
  type        = string
  description = "Name suffix of the secret"
}

variable "region" {
  type        = string
  description = "The AWS region the terraform state belongs to."
  default     = "eu-west-1"
}

variable "team" {
  type        = string
  description = "Official name of the team/stream."
}

variable "environment" {
  type        = string
  description = "Name of the continues integration stage this infrastructure belongs to (e.g. production, staging, sandbox)."
}

variable "stack" {
  type        = string
  description = "Union of components forming a customer feature."
}

variable "repository" {
  type        = string
  description = "Reference to the git repository the infrastructure is defined in."
}
