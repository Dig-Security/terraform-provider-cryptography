package provider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestDataSource_Sha512(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `data "sha512" "sha_test" {
							input = "some text" 
							encoding = "ISO-8859-1"	
						 }`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sha512.sha_test",
						"sha",
						"e2732baedca3eac1407828637de1dbca702c3fc9ece16cf536ddb8d6139cd85dfe7464b8235b29826f608ccf4ac643e29b19c637858a3d8710a59111df42ddb5",
					),
				),
			},
			{
				Config: `data "sha512" "sha_test" {
							input = "some text"
							encoding = "UTF-16"
						 }`,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sha512.sha_test",
						"sha",
						"6b6138aacb0dc7ca5451669fc11c2f71f2a6228a4de9e3fcc8e214fb62a638eca2562df1bb27fc1068bdbd54c8eb402aae5718065787ac60150b6c4b9387b8fd",
					),
				),
			},
		},
	})
}
