---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "GKEHub"
description: |-
  RBACRoleBinding represents a rbacrolebinding across the Fleet.
---

# google\_gke\_hub\_membership\_rbac\_role\_binding

RBACRoleBinding represents a rbacrolebinding across the Fleet.


To get more information about MembershipRBACRoleBinding, see:

* [API documentation](https://cloud.google.com/anthos/fleet-management/docs/reference/rest/v1/projects.locations.memberships)
* How-to Guides
    * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)

## Example Usage - Gkehub Membership Rbac Role Binding Basic


```hcl
resource "google_container_cluster" "primary" {
  name               = "basiccluster"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membershiprbacrolebinding" {
  membership_id = "tf-test-membership%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }
  
  depends_on = [google_container_cluster.primary]
}

resource "google_gke_hub_membership_rbac_role_binding" "membershiprbacrolebinding" {
  membership_rbac_role_binding_id = "tf-test-membership-rbac-role-binding%{random_suffix}"
  membership_id = "tf-test-membership%{random_suffix}"
  user = "service-${data.google_project.project.number}@gcp-sa-anthossupport.iam.gserviceaccount.com"
  role {
    predefined_role = "ANTHOS_SUPPORT"
  }
  location = "global"
  depends_on = [google_gke_hub_membership.membershiprbacrolebinding]
}

data "google_project" "project" {}
```

## Argument Reference

The following arguments are supported:


* `membership_rbac_role_binding_id` -
  (Required)
  The client-provided identifier of the RBAC Role Binding.

* `user` -
  (Required)
  Principal that is be authorized in the cluster (at least of one the oneof
  is required). Updating one will unset the other automatically.
  user is the name of the user as seen by the kubernetes cluster, example
  "alice" or "alice@domain.tld"

* `role` -
  (Required)
  Role to bind to the principal.
  Structure is [documented below](#nested_role).

* `membership_id` -
  (Required)
  Id of the membership

* `location` -
  (Required)
  Location of the Membership


<a name="nested_role"></a>The `role` block supports:

* `predefined_role` -
  (Required)
  PredefinedRole is an ENUM representation of the default Kubernetes Roles
  Possible values are: `UNKNOWN`, `ADMIN`, `EDIT`, `VIEW`, `ANTHOS_SUPPORT`.

- - -


* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}/rbacrolebindings/{{membership_rbac_role_binding_id}}`

* `name` -
  The resource name for the RBAC Role Binding

* `uid` -
  Google-generated UUID for this resource.

* `create_time` -
  Time the RBAC Role Binding was created in UTC.

* `update_time` -
  Time the RBAC Role Binding was updated in UTC.

* `delete_time` -
  Time the RBAC Role Binding was deleted in UTC.

* `state` -
  State of the RBAC Role Binding resource.
  Structure is [documented below](#nested_state).


<a name="nested_state"></a>The `state` block contains:

* `code` -
  (Output)
  Code describes the state of a RBAC Role Binding resource.

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


MembershipRBACRoleBinding can be imported using any of these accepted formats:

```
$ terraform import google_gke_hub_membership_rbac_role_binding.default projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}/rbacrolebindings/{{membership_rbac_role_binding_id}}
$ terraform import google_gke_hub_membership_rbac_role_binding.default {{project}}/{{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}
$ terraform import google_gke_hub_membership_rbac_role_binding.default {{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
