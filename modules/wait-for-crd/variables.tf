variable "crds" {
  type    = list(string)
  default = []
}

variable "timeout" {
  type    = string
  default = "180s"
}

variable "attempts" {
  type    = number
  default = 30
}
