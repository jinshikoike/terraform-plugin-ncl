provider "ncl" {}

resource "ncl_keypair" "ssh_key" {
  name = "koikeOJT"
  public_key_material = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCztl4uCOS3M+JMcDVJyWn2HjyLTVEOWWS5Fm5573iMFVF9y/XcPiXqNdVnxkWqjaxycnmyLOXYWMKurZnRF8qvLVl+MqzUYxypjcQKGySo5MxYfayUd53TWv2p+ZpykJ6omg+HBD2CEtV+4XRGb+/Q5OC40qD8d9T1XdZu6f/jUSO3RNeqRWARKmaFcVfoKYzA8p0RjLRmdJus2ir9kH3OYfSzglqmtw5m8Cj8ikgfs9C99M2KAQUflBcMeHNbIdHhTvuclA86ESRnZNyi3hUCLCme2EaClgl3wMKUxfmqTAHZvnaRs4BhOvi3BFPQXzM8dk+frtCNa+4Ut9yZZSAuKyddGcJeOGNp7ev0752JZtiG+QLwCMZ30aibImFQYhAInhRxSGq0b6UYMgETUXHwj3uJ4pm/ts8r4EODRs2PLbMQjcy41Gnnf52DgIHppNYC8zmrVfZ9wzJtuNdlp/XgiTJJlvUx1Ng+b86WkbGvVIHXaD+hokKKy5KqYF5YmlQcM0GesErvJ9iA8OKsE2t8Yt3UYooG6BMr5zwu6YntVdI2yxzYkbym5zFEMHFiu12hscp9EKo87D9Q2fdyQgBWJj4mMrNirRwFqq+rOOiZujU777Leu+fxJLnKliCo56bBIGhYh7/LUU7Eopjm9VdYeNX3mn0oP+WzPDxuWbSuxw== koike@ubuntu"
}

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

resource "ncl_instance" "sample2" {
  image_id = "68"
  key_name = "OjtKoike"
  instance_type = "mini"
  avail_zone = "west-12"
  accounting_type = "2"
  instance_id = "OjtTerra3"
  security_groups = [
    { name = "OjtKoike" }
  ]
}

resource "ncl_instance" "sample3" {
  image_id = "68"
  key_name = "OjtKoike"
  instance_type = "mini"
  avail_zone = "west-12"
  accounting_type = "2"
  instance_id = "OjtTerra2"
  security_groups = [
    { name = "OjtKoike" }
  ]
}
