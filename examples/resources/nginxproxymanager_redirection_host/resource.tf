resource "nginxproxymanager_redirection_host" "test" {
  forward_domain_name = "test123.com"
  forward_http_code   = 301
  forward_scheme      = "https"
  domain_names        = ["domain123.com"]

  http2_support  = true
  certificate_id = 42
}