package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testInterfaceBondingAddress = "routeros_interface_bonding.test"

func TestAccInterfaceBondingTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/bonding", "routeros_interface_bonding"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfaceBondingConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfaceBondingExists(testInterfaceBondingAddress),
							resource.TestCheckResourceAttr(testInterfaceBondingAddress, "name", "bonding-test"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckInterfaceBondingExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccInterfaceBondingConfig() string {
	return providerConfig + `

resource "routeros_interface_bonding" "test" {
  name   = "bonding-test"
  slaves =  ["ether3", "ether4"]
}
`
}
