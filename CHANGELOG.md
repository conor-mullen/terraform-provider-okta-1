# Changelog

## 3.7.4 (December 28, 2020)

ENHANCEMENTS:

* Add `dependabot` to automate dependency updates [#259](https://github.com/oktadeveloper/terraform-provider-okta/pull/259). Thanks [@jlosito](https://github.com/jlosito)!
* Add `max_clock_skew` property to IdP SAML resource [#263](https://github.com/oktadeveloper/terraform-provider-okta/pull/263). Thanks [@me](https://github.com/bogdanprodan-okta)!

BUGS:

* Fix panic caused by a null pointer in `okta_policy_password` resource. [#262](https://github.com/oktadeveloper/terraform-provider-okta/pull/262). Thanks [@me](https://github.com/bogdanprodan-okta)!
* Add retries for creating/updating `okta_user_schema` resource. [#262](https://github.com/oktadeveloper/terraform-provider-okta/pull/262). Thanks [@me](https://github.com/bogdanprodan-okta)!

## 3.7.3 (December 24, 2020)

ENHANCEMENTS:

* Add call recovery for Okta password policy [#248](https://github.com/oktadeveloper/terraform-provider-okta/pull/248). Thanks [@me](https://github.com/bogdanprodan-okta)!
* Update data okta_group docs [#251](https://github.com/oktadeveloper/terraform-provider-okta/pull/251). Thanks [@ymylei](https://github.com/ymylei)!
* Adds `pattern` property for `okta_*_schema` resources [#159](https://github.com/oktadeveloper/terraform-provider-okta/pull/159). Thanks [@fitzoh](https://github.com/fitzoh) and [@me](https://github.com/bogdanprodan-okta)!
* Add retries on connection timeouts errors [#246](https://github.com/oktadeveloper/terraform-provider-okta/issues/246). Thanks [@me](https://github.com/bogdanprodan-okta)!

BUGS:

* Fixed handling rule with `INVALID` status [#250](https://github.com/oktadeveloper/terraform-provider-okta/pull/250). Thanks [@ymylei](https://github.com/ymylei)!

## 3.7.2 (December 18, 2020)

ENHANCEMENTS:

* Add logs to group data source for different cases [#150](https://github.com/oktadeveloper/terraform-provider-okta/pull/150). Thanks [@nathanbartlett](https://github.com/nathanbartlett)!
* Added missing documentation [#245](https://github.com/oktadeveloper/terraform-provider-okta/pull/245). Thanks [@me](https://github.com/bogdanprodan-okta)!

BUGS:

* Fix default name for idp_discovery [#244](https://github.com/oktadeveloper/terraform-provider-okta/pull/244). Thanks [@nickerzb](https://github.com/nickerzb)!
* Fix okta auth server policy rule resource causing panic [#245](https://github.com/oktadeveloper/terraform-provider-okta/pull/245). Thanks [@SBerda](https://github.com/SBerda) for submitting the [issue](https://github.com/oktadeveloper/terraform-provider-okta/issues/202) and [@me](https://github.com/bogdanprodan-okta) for fixing it!
* Fix `key_years_valid` defaulting to `2` during resource import [#245](https://github.com/oktadeveloper/terraform-provider-okta/pull/245). Thanks [@btsteve](https://github.com/btsteve) for submitting the [issue](https://github.com/oktadeveloper/terraform-provider-okta/issues/201) and [@me](https://github.com/bogdanprodan-okta) for fixing it!

## 3.7.1 (December 16, 2020)

ENHANCEMENTS:

* Add validation for user type [#242](https://github.com/oktadeveloper/terraform-provider-okta/pull/242).

BUGS:

* Fix state refresh for `okta_user_base_schema` and `okta_user_schema` [#242](https://github.com/oktadeveloper/terraform-provider-okta/pull/242).

## 3.7.0 (December 15, 2020)

ENHANCEMENTS:

* Add user types support [#183](https://github.com/oktadeveloper/terraform-provider-okta/pull/183). Thanks, [@rajnadimpalli](https://github.com/rajnadimpalli) and [@bogdanprodan-okta](https://github.com/bogdanprodan-okta)!
* Add type to data okta group [#217](https://github.com/oktadeveloper/terraform-provider-okta/pull/217). Thanks, [@dangoslen](https://github.com/dangoslen)!
* Add `acs_endpoints` to SAML app (okta_app_saml) definition [#226](https://github.com/oktadeveloper/terraform-provider-okta/pull/226). Thanks, [@pranjalranjan](https://github.com/pranjalranjan)!
* Update terraform-plugin-sdk libraries, added possibility to set provider's log level [#220](https://github.com/oktadeveloper/terraform-provider-okta/pull/220). Thanks, [@bryantbiggs](https://github.com/bryantbiggs) and [@bogdanprodan-okta!](https://github.com/bogdanprodan-okta)
* Overhaul idp_discovery_rule documentation [#228](https://github.com/oktadeveloper/terraform-provider-okta/pull/228). Thanks [@eatplaysleep](https://github.com/eatplaysleep)!
* General documentation updates [#224](https://github.com/oktadeveloper/terraform-provider-okta/pull/224). Thanks, [@bryantbiggs](https://github.com/bryantbiggs)!

BUGS:

* Changed `okta_app_basic_auth` optional fields to required [issue 223](https://github.com/oktadeveloper/terraform-provider-okta/issues/223). Thanks, [@bryantbiggs](https://github.com/bryantbiggs)!
* Add idp discovery to allowed list of default policies [#233](https://github.com/oktadeveloper/terraform-provider-okta/pull/233). Thanks, [@nickerzb](https://github.com/nickerzb)!

## 3.6.1 (November 14, 2020)

ENHANCEMENTS:

* Remove 3rd party Okta SDK [#215](https://github.com/oktadeveloper/terraform-provider-okta/pull/215). Thanks, [@bogdanprodan-okta](https://github.com/bogdanprodan-okta)
* Enhance `okta_app_auto_login` resource [#164](https://github.com/oktadeveloper/terraform-provider-okta/pull/164). Thanks, [@isometry](https://github.com/isometry)!
* Add group name to the error for group data call [#156](https://github.com/oktadeveloper/terraform-provider-okta/pull/156). Thanks, [@ymylei](https://github.com/ymylei)!

BUGS:

* Fix population of the user 'status' attribute [#206](https://github.com/oktadeveloper/terraform-provider-okta/pull/206). Thanks, [@isometry](https://github.com/isometry)!

## 3.6.0 (October 12, 2020)

ENHANCEMENTS:

* Upgrade to Okta SDK 2.0.0 [#203](https://github.com/oktadeveloper/terraform-provider-okta/pull/203). Thanks a ton! [@bogdanprodan-okta](https://github.com/bogdanprodan-okta)
* Fix validation false positive when api_token is set via environment variable. [#147](https://github.com/oktadeveloper/terraform-provider-okta/pull/147). Thanks, [@jgeurts](https://github.com/jgeurts)
* Update required to optional and more [#208](https://github.com/oktadeveloper/terraform-provider-okta/pull/208), Thanks, me! :smile:

BUGS:

* Update config.go [#207](https://github.com/oktadeveloper/terraform-provider-okta/pull/207), Thanks, me! :smile:

## 3.5.1 (October 9, 2020)

ENHANCEMENTS:

* Update config.go [#192](https://github.com/oktadeveloper/terraform-provider-okta/pull/192), Thanks, [@bretterer](https://github.com/bretterer)!

BUGS:

* Documentation: Update okta_idp_metadata_saml correct example [#173](https://github.com/oktadeveloper/terraform-provider-okta/pull/173), Thanks, [@gaurdro](https://github.com/gaurdro) and [@netflash](https://github.com/netflash)!
* Documentation: Update warning in app_group_assignment.html.markdown [#172](https://github.com/oktadeveloper/terraform-provider-okta/pull/172), Thanks, [@ssttgg](https://github.com/ssttgg)!
* Renaming Go module as per the organization move [#195](https://github.com/oktadeveloper/terraform-provider-okta/pull/195), Thanks, [@stack72](https://github.com/stack72)!

## 3.5.0 (August 31, 2020)

ENHANCEMENTS:

* Add password import inline hook type. [#168](https://github.com/oktadeveloper/terraform-provider-okta/pull/168), Thanks, [@noinarisak](https://github.com/noinarisak) aka me! :tada:
* Add external_namespace property for app_user_schema and user_schema. [#102](https://github.com/oktadeveloper/terraform-provider-okta/pull/102), Thanks, [@thehunt33r](https://github.com/thehunt33r)!

BUGS:

* Fix inline hook example code to match version that is supported. [#175](https://github.com/oktadeveloper/terraform-provider-okta/pull/175), Thanks, [@noinarisak](https://github.com/noinarisak) me again! :smiley:
* Update app_group_assignment.html.markdown. [#165](https://github.com/oktadeveloper/terraform-provider-okta/pull/165), Thanks, [snolan-amount](https://github.com/snolan-amount)!

## 3.4.1 (July 31, 2020)

RELEASE:

* First release under oktadeveloper organization with binary published to [registry.hashicorp.com](https://registry.terraform.io/).

## 3.4.0 (July 30, 2020)

ENHANCEMENTS:

* Add resource definition for Okta Event Hooks. [#14](https://github.com/terraform-providers/terraform-provider-okta/pull/14), Thanks, [@mbudnek](https://github.com/mbudnek)!
* Adding support for GROUP_MEMBERSHIP_ADMIN & REPORT_ADMIN. [#138](https://github.com/terraform-providers/terraform-provider-okta/pull/138) Thanks, [ymylei](https://github.com/ymylei)!

BUG FIXES:

* Documentation corrections. Thanks, to these fine individuals!
  * [#126](https://github.com/terraform-providers/terraform-provider-okta/pull/126) [@ChristophShyper](https://github.com/ChristophShyper)
  * [#127](https://github.com/terraform-providers/terraform-provider-okta/pull/127) [@thekbb](https://github.com/thekbb)
  * [#151](https://github.com/terraform-providers/terraform-provider-okta/pull/151) [@varrunramani-okta](https://github.com/varrunramani-okta)

## 3.3.0 (May 29, 2020)

ENHANCEMENTS:

* Add user lockout notification channels. [#15](https://github.com/terraform-providers/terraform-provider-okta/pull/15), Thanks, [@thehunt33r](https://github.com/thehunt33r)!
* Adding support for SMS template changes. [#18](https://github.com/terraform-providers/terraform-provider-okta/pull/18) Thanks, [@gusChan](https://github.com/gusChan)!

## 3.2.0 (April 03, 2020)

BUG FIXES:

* Documentation, `id` is an output of `app_oauth`. [#98]() Thanks, [beyondbill](https://github.com/beyondbill)!

ENHANCEMENTS:

* Improve app filtering and update Terraform SDK. [#97](https://github.com/terraform-providers/terraform-provider-okta/pull/97) Thanks, [quantumew](https://github.com/quantumew)! :tada:

## 3.1.1 (March 18, 2020)

ENHANCEMENTS:

* Add unique property to UserSchema. [#12](https://github.com/terraform-providers/terraform-provider-okta/pull/12) Thanks, [@gusChan](https://github.com/gusChan)!

## 3.1.0 (February 19, 2020)

RELEASE:

* First release under terraform-providers organization with binary published to releases.hashicorp.com

## 3.0.0 (October 16, 2019)

FEATURES:

* Updated provider to support Terraform v0.12.0

## 3.0.1

FEATURES:

* **New Resource:** `okta_inline_hook`

ENHANCEMENTS:

* Add missing okta_idp_saml settings

## 3.0.2

ENHANCEMENTS:

* Use backoff/retries functionality for XML API calls

## 3.0.3

FEATURES:

* **New Data Source:** okta_idp_saml

ENHANCEMENTS:

* Support import user by email

## 3.0.4

FEATURES:

* **New Data Source:** Add okta_app_saml data source
* **New Data Source:** Add okta_app_metadata_saml data source
* **New Data Source:** Add okta_idp_metadata_saml data source

ENHANCEMENTS:

* Change type of custom_profile_attributes from map to JSON string to support all types

BUG FIXES:

* Fix group filter bug, filter_type and filter_value were not being sync'd

## 3.0.5

BUG FIXES:

* Fix bug introduced in v3.0.4. User data source was not updated to the new caustom_profile_attribute type
* Added test to cover this scenario, tests were passiing

## 3.0.6

ENHANCEMENTS:

* Allow client_id to be set on OIDC application, while also maintaining the computed version. With some auth methods, such as basic auth, this is possible.

## 3.0.7

ENHANCEMENTS:

* Add group_assignments for SAML and social IdPs

## 3.0.8

ENHANCEMENTS:

* Add issuer_mode to social IdP. Our test org does not have a custom domain setup, thus it was working there but not in other orgs. Hard to test both scenarios in one org.

## 3.0.9

FEATURES:

* **New Resource:** `okta_template_email`
* **New Resource:** `okta_group_roles`

## 3.0.10

FEATURES:

* **New Resource:** `okta_network_zone`

## 3.0.11

BUG FIXES:

* Fix occasional panic when creating a user schema see [issue 144](https://github.com/terraform-providers/terraform-provider-okta/issues/144)
* Users in LOCKED_OUT state are unlocked when config is ACTIVE [issue 225](https://github.com/terraform-providers/terraform-provider-okta/issues/225)

## 3.0.12

BUG FIXES:

* Ensure schema does not panic after retry

## 3.0.13

FEATURES:

* **New Resource:** `okta_user_base_schema`

ENHANCEMENTS:

* Add missing attribute, match_type and match_attribute, on social idp resource

## 3.0.14

BUG FIXES:

* Fix logic around including/excluding networks on policy rules

## 3.0.15

ENHANCEMENTS:

* Update Okta SDK
* Filter out GROUP based admin roles when processing user `admin_roles` attribute

## 3.0.16

* Fix issues around `okta_policy_rule_idp_discovery`
  * `app_include` and `app_exlcude` were missing required properties
  * `user_identifier_type` was being added even when not defined, causing API errors
* Fix integer array type

## 3.0.17

FEATURES:

* **New Resource:** `okta_app_user_schema`
* **New Resource:** `okta_app_user_base_schema`
* **New Resource:** `okta_app_user` resource
* **New Resource:** `okta_app_group` resource

ENHANCEMENTS:

* Add `required` field to base schema

## 3.0.18

ENHANCEMENTS:

* Support SHA-1 signing algorithm on IdPs

BUG FIXES:

* Fix bug where audience is reset on IdP update because it is omitted from the payload

## 3.0.19

BUG FIXES:

* Fix diff issues around `okta_policy_rule_idp_discovery`
* Allow `provisioning_action` for IdPs to be set to `DISABLED`

## 3.0.20

BUG FIXES:

* Fix `okta_auth_server_claim`, `group_filter_type` could not be set to `STARTS_WITH` due to a typo

## 3.0.21

ENHANCEMENTS:

* Expose scope property on `okta_user_schema`
* Allow setting of OAuth application visibility settings

## 3.0.22

BUG FIXES:

* Send `profileMaster` along with IdP, so the config is recognized by Okta API
* Fix bug in SDK related to retries and the request body being empty on subsequent requests.

## 3.0.23

ENHANCEMENTS:

* Add `external_name` property to the `okta_app_user_schema` and `okta_user_schema`

## 3.0.24

ENHANCEMENTS:

* Support `profile` on `okta_oauth_app` resource

## 3.0.25

ENHANCEMENTS:

* Support setting an auth server scope as the default
* Support `profile` and `priority` on `okta_app_group_assignment`
* Support `profile` on `okta_app_user`

BUG FIXES:

* Fix bug with supporting `profile` on `okta_oauth_app` resource

## 3.0.26

ENHANCEMENTS:

* Support array enums in `okta_user_schema` and `okta_app_user_schema` as `array_enum` and `array_one_of`

## 3.0.27

ENHANCEMENTS:

* Update refresh token window validation to account for new upper limit of 5 years

## 3.0.28

BUG FIXES:

* Remove resource from state on 404. ([#269](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/269))

## 3.0.29

BUG FIXES:

* Ensure we safely sync auth server properties. ([#299](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/299))
* MANUAL rotation mode can only be set on an auth server on update. Ensure we run update after create for that scenario. ([#287](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/287))

## 3.0.30

ENHANCEMENT:

* Update to new separate Terraform SDK ([#307](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/307))

## 3.0.31

BUG FIXES:

* Ensure `okta_app_group_assignment` resource syncs using the right read function. ([#307](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/307))

## 3.0.32

BUG FIXES:

* Ensure `okta_app_group_assignment` and `okta_app_user` resources properly take multiple ids on the import functions. ([#307](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/307))
* Ensure `okta_user` does not error on 404 ([#313](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/313))

## 3.0.33

FEATURES:

* **New Resource:** `okta_profile_mapping` ([#246](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/246))
* **New Resource:** `okta_app_basic_auth` ([#329](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/329))

## 3.0.34

BUG FIXES:

* Policy values could not be set to 0. Doing so resulted in the SDK omitting them, resulting in Okta resetting the values to default.

## 3.0.35

ENHANCEMENT:

* Require target_id on `okta_profile_mapping` to avoid ambiguity

FEATURES:

* **New Data Source:** `okta_user_profile_mapping_source` ([#340](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/340))

## 3.0.36

BUG FIXES

* Schema merging helper function was mutating input schema causing side effects when used in a particular way. Used shallow copying to avoid this side effect. ([#338](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/338))
* Ensure response is not nil when checking status code ([#307](https://github.com/terraform-providers/terraform-provider-cherryservers/issues/307))

## 3.0.37

BUG FIXES

* Ensure `index` is sync'd on import to avoid recreation.

## 3.0.38

ENHANCEMENT:

* Support `password`, `recovery_answer`, and `recovery_question` as attributes on the `okta_user` resource.
