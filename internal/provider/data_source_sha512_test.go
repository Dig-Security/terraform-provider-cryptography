package provider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestDataSource_Sha(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `data "sha" "sha_test" {
							input = "some text" 
						 }`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sha.sha_test",
						"sha",
						"e2732baedca3eac1407828637de1dbca702c3fc9ece16cf536ddb8d6139cd85dfe7464b8235b29826f608ccf4ac643e29b19c637858a3d8710a59111df42ddb5",
					),
				),
			},
		},
	})
}
