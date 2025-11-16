resource "nginxproxymanager_user" "user" {
  name     = "My User"
  nickname = "Cool Nickname"
  email    = "user@domain.com"
  roles    = ["admin"]
}

resource "nginxproxymanager_user_password" "user" {
  user_id  = nginxproxymanager_user.user.id
  password = "SuperS3cr3tP4ssw0rd!"
}