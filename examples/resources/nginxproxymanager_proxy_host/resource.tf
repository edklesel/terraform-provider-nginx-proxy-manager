resource "nginxproxymanager_proxy_host" "test" {
  forward_host   = "server.local"
  forward_port   = 8080
  forward_scheme = "https"
  domain_names   = ["domain123.com"]

  http2_support  = true
  certificate_id = 42
}