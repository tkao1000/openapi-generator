package sumologic

import (
	"fmt"
	"testing"
	"os"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
        "github.com/hashicorp/terraform/helper/acctest"
)

func TestAccUserCreate(t *testing.T) {
        var user User
        testfirstName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
        testlastName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
        testemail := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
        testroleIds := [1]string{acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)}
        testisActive := true
        resource.Test(t, resource.TestCase{
                PreCheck: func() { TestAccPreCheck(t) },
                Providers:    testAccProviders,
                CheckDestroy: testAccCheckUserDestroy(user),
                Steps: []resource.TestStep{
                        {
                                Config: testAccSumologicUser(testfirstName, testlastName, testemail, testroleIds, testisActive),
                                Check: resource.ComposeTestCheckFunc(
                                        testAccCheckUserExists("sumologic_user.test", &user, t),
                                        testAccCheckUserAttributes("sumologic_user.test"),
                                        resource.TestCheckResourceAttr("sumologic_user.test", "first_name", testfirstName),
                                        resource.TestCheckResourceAttr("sumologic_user.test", "last_name", testlastName),
                                        resource.TestCheckResourceAttr("sumologic_user.test", "email", testemail),
                                        resource.TestCheckResourceAttr("sumologic_user.test", "role_ids", testroleIds),
                                        resource.TestCheckResourceAttr("sumologic_user.test", "is_active", testisActive),
                                ),
                        },
                },
        })
}







func testAccCheckUserDestroy(user User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Client)
		_, err := client.GetUser(user.ID)
		if err == nil {
			return fmt.Errorf("User still exists")
		}
		return nil
	}
}



func testAccCheckUserExists(name string, user *User, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("User not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("User ID is not set")
		}

		id := rs.Primary.ID
		c := testAccProvider.Meta().(*Client)
		newUser, err := c.GetUser(id)
		if err != nil {
			return fmt.Errorf("User %s not found", id)
		}
		user = newUser
		return nil
	}
}







func TestAccUserUpdate(t *testing.T) {
  var user User
          testfirstName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
          testlastName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
          testemail := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
          testroleIds := [1]string{acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)}
          testisActive := true

          testUpdatedfirstName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
          testUpdatedlastName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
          testUpdatedemail := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
          testUpdatedroleIds := [1]string{acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)}
          testUpdatedisActive := false

	resource.Test(t, resource.TestCase{
		PreCheck: func() { TestAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy(user),
		Steps: []resource.TestStep{
			{
				Config: testAccSumologic(testfirstName, testlastName, testemail, testroleIds, testisActive),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("sumologic_user.test", &user, t),
					testAccCheckUserAttributes("sumologic_user.test"),
          resource.TestCheckResourceAttr("sumologic_user.test", "first_name", testfirstName),
          resource.TestCheckResourceAttr("sumologic_user.test", "last_name", testlastName),
          resource.TestCheckResourceAttr("sumologic_user.test", "email", testemail),
          resource.TestCheckResourceAttr("sumologic_user.test", "role_ids", testroleIds),
          resource.TestCheckResourceAttr("sumologic_user.test", "is_active", testisActive),
				),
			},
			{
				Config: testAccSumologicUserUpdate(testUpdatedfirstName, testUpdatedlastName, testUpdatedemail, testUpdatedroleIds, testUpdatedisActive),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserExists("sumologic_user.test", &user, t),
					testAccCheckUserAttributes("sumologic_user.test"),
          resource.TestCheckResourceAttr("sumologic_user.test", "first_name", testUpdatedfirstName),
          resource.TestCheckResourceAttr("sumologic_user.test", "last_name", testUpdatedlastName),
          resource.TestCheckResourceAttr("sumologic_user.test", "email", testUpdatedemail),
          resource.TestCheckResourceAttr("sumologic_user.test", "role_ids", testUpdatedroleIds),
          resource.TestCheckResourceAttr("sumologic_user.test", "is_active", testUpdatedisActive),
				),
			},
		},
	})
}










func testAccSumologicUser(firstName string, lastName string, email string, roleIds []string, isActive bool) string {
	return fmt.Sprintf(`
resource "sumologic_user" "test" {
  firstName = %s
  lastName = %s
  email = %s
  roleIds = %v
  isActive = %t
}
`, firstName, lastName, email, roleIds, isActive)
}

func testAccSumologicUserUpdate(firstName string, lastName string, email string, roleIds []string, isActive bool) string {
	return fmt.Sprintf(`
resource "sumologic_user" "test" {
      firstName = %s
      lastName = %s
      email = %s
      roleIds = %v
      isActive = %t
}
`, firstName, lastName, email, roleIds, isActive)
}

func testAccCheckUserAttributes(name string) resource.TestCheckFunc {
  return func(s *terraform.State) error {
      f := resource.ComposeTestCheckFunc(
        resource.TestCheckResourceAttrSet(name, "firstName"),
        resource.TestCheckResourceAttrSet(name, "lastName"),
        resource.TestCheckResourceAttrSet(name, "email"),
        resource.TestCheckResourceAttrSet(name, "roleIds"),
        resource.TestCheckResourceAttrSet(name, "isActive"),
      )
      return f(s)
   }
}
