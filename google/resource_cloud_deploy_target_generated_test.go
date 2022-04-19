// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudDeployTarget_cloudDeployTargetExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudDeployTargetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudDeployTarget_cloudDeployTargetExample(context),
			},
			{
				ResourceName:            "google_cloud_deploy_target.pipeline",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "region"},
			},
		},
	})
}

func testAccCloudDeployTarget_cloudDeployTargetExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_deploy_target" "pipeline" {
  name          = "tf-test-tf-test%{random_suffix}"
  description   = "Target Cluster"
  gke {
    cluster = "projects/%{project}/locations/us-central1/clusters/dev"
  }
}
`, context)
}

func TestAccCloudDeployTarget_cloudDeployTargetFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudDeployTargetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudDeployTarget_cloudDeployTargetFullExample(context),
			},
			{
				ResourceName:            "google_cloud_deploy_target.pipeline",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "region"},
			},
		},
	})
}

func testAccCloudDeployTarget_cloudDeployTargetFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_deploy_target" "pipeline" {
  name          = "tf-test-tf-test%{random_suffix}"
  description   = "Target Cluster"
  annotations = {
    generated-by = "magic-modules"
  }

  require_approval = true
  labels = {
    env = "dev"
  }
  gke {
    cluster = "projects/%{project}/locations/us-central1/clusters/dev"
  }
  execution_configs {
    usages = ["RENDER", "DEPLOY"]
    service_account = "%{project}@appspot.gserviceaccount.com"
  }
}
`, context)
}

func testAccCheckCloudDeployTargetDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_deploy_target" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudDeployBasePath}}projects/{{project}}/locations/{{region}}/targets/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("CloudDeployTarget still exists at %s", url)
			}
		}

		return nil
	}
}