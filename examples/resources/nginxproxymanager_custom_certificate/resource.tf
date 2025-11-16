resource "tls_private_key" "example" {
  algorithm = "ECDSA"
}

resource "tls_self_signed_cert" "example" {
  key_algorithm         = tls_private_key.example.algorithm
  private_key_pem       = tls_private_key.example.private_key_pem
  validity_period_hours = 24 * 365
  allowed_uses = [
    "server_auth",
  ]

  dns_names = ["test.domain.com"]

  subject {
    common_name  = "test.domain.com"
    organization = "ACME Examples, Inc"
  }
}

resource "nginxproxymanager_custom_certificate" "example" {
  nice_name       = "test.domain.com"
  certificate     = tls_self_signed_cert.example.cert_pem
  certificate_key = tls_private_key.example.private_key_pem
}