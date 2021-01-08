data "okta_group" "all" {
  name = "Everyone"
}

resource okta_idp_social google {
  type                = "GOOGLE"
  protocol_type       = "OAUTH2"
  name                = "Google"
  provisioning_action = "DISABLED"
  scopes = [
    "profile",
    "email",
    "openid",
  ]

  client_id         = "foo"
  client_secret     = "bar"
  username_template = "idpuser.email"
}

resource "okta_policy_signon" "test" {
  name            = "testAcc_replace_with_uuid"
  status          = "ACTIVE"
  description     = "Terraform Acceptance Test SignOn Policy"
  groups_included = [data.okta_group.all.id]
}

resource "okta_policy_rule_signon" "test" {
  policyid           = okta_policy_signon.test.id
  name               = "testAcc_replace_with_uuid"
  status             = "ACTIVE"
  mfa_required      = true
  identity_provider  = "SPECIFIC_IDP"
  identity_provider_idps = [okta_idp_social.google.id]
}
