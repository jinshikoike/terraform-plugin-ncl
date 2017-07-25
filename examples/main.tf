variable "access_key" {}
variable "secret_key" {}
variable "region" {}

provider "ncl" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region = "${var.region}"
}

resource "ncl_instance" "sample" {
  image_id = "68"
}

