data "nginxproxymanager_users" "users" {}

output "usernames" {
  value = [for u in data.nginxproxymanager_users.users.users : u.name]
}

output "admin_users" {
  value = [for u in data.nginxproxymanager_users.users.users : u.name if u.roles[0] == "admin"]
}