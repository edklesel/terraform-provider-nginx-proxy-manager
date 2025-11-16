resource "nginxproxymanager_user" "user" {
  name     = "My User"
  nickname = "Cool Nickname"
  email    = "user@domain.com"
  roles    = ["admin"]
}