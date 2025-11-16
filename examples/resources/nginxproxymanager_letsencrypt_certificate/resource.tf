resource "nginxproxymanager_letsencrypt_certificate" "test" {
  domain_names             = ["www.example.com"]
  dns_challenge            = true
  dns_provider             = "route53"
  dns_provider_credentials = <<EOF
    [default]
    aws_access_key_id=exmapleaccesskey
    aws_secret_access_key=examplesecretkey
  EOF
  letsencrypt_email        = "test@test.com"
  propagation_seconds      = 5
}