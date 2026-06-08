// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tests

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccUsersDataSource(t *testing.T) {

	randInt := rand.Int64N(1e5)
	expectedName := fmt.Sprintf("test_%d", randInt)
	expectedNickname := fmt.Sprintf("nicktest_%d", randInt)
	expectedEmail := fmt.Sprintf("test_%d@example.com", randInt)
	expectedRole := "user"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + fmt.Sprintf(`
resource "nginxproxymanager_user" "test" {
	name     = "%s"
	nickname = "%s"
	email    = "%s"
	roles    = ["%s"]
}
	
data "nginxproxymanager_users" "test" {
	depends_on = [nginxproxymanager_user.test]
}
`, expectedName, expectedNickname, expectedEmail, expectedRole),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Then check that the first user is there, which is always
					// the admin user
					resource.TestCheckResourceAttr("data.nginxproxymanager_users.test", "users.0.id", "1"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_users.test", "users.0.name", "Administrator"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_users.test", "users.0.nickname", "Admin"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_users.test", "users.0.email", "admin@example.com"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_users.test", "users.0.roles.0", "admin"),

					// Then check that the number of users is greater than 1
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["data.nginxproxymanager_users.test"]
						if !ok {
							return fmt.Errorf("Unable to list users.")
						}

						numUsers, err := strconv.Atoi(rs.Primary.Attributes["users.#"])
						if err != nil {
							return fmt.Errorf("Error parsing array length: %s", err)
						}

						if numUsers <= 1 {
							return fmt.Errorf("Only found %d user(s), expected more.", numUsers)
						}

						// name := rs.Primary.Attributes[fmt.Sprintf("users.%d.name", lastIndex)]
						// if name != expectedName {
						// 	return fmt.Errorf("Name mistmatch at index %d, got %s, want %s", lastIndex, name, expectedName)
						// }
						// nickname := rs.Primary.Attributes[fmt.Sprintf("users.%d.nickname", lastIndex)]
						// if nickname != expectedNickname {
						// 	return fmt.Errorf("Nickname mistmatch, got %s, want %s", nickname, expectedNickname)
						// }
						// email := rs.Primary.Attributes[fmt.Sprintf("users.%d.email", lastIndex)]
						// if email != expectedEmail {
						// 	return fmt.Errorf("Email mistmatch, got %s, want %s", email, expectedEmail)
						// }
						// role := rs.Primary.Attributes[fmt.Sprintf("users.%d.roles.0", lastIndex)]
						// if role != expectedRole {
						// 	return fmt.Errorf("Role mistmatch, got %s, want %s", role, expectedRole)
						// }

						return nil
					},
				),
			},
		},
	})
}
