provider "ncl" {}

resource "ncl_instance" "sample" {
  image_id = "68"
  key_name = "OjtKoike"
  instance_type = "mini"
  avail_zone = "west-12"
  accounting_type = "2"
  instance_id = "OjtTerra"
  security_groups = [
    { name = "OjtKoike" }
  ]
}
