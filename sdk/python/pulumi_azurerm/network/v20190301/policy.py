# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Policy(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Gets a unique read-only string that changes whenever the resource is updated.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the web application firewall policy.
      * `custom_rules` (`dict`) - Describes custom rules inside the policy.
        * `rules` (`list`) - List of rules
          * `action` (`str`) - Describes what action to be applied when rule matches.
          * `enabled_state` (`str`) - Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
          * `match_conditions` (`list`) - List of match conditions.
            * `match_value` (`list`) - List of possible match values.
            * `match_variable` (`str`) - Request variable to compare with.
            * `negate_condition` (`bool`) - Describes if the result of this condition should be negated.
            * `operator` (`str`) - Comparison type to use for matching with the variable value.
            * `selector` (`str`) - Match against a specific key from the QueryString, PostArgs, RequestHeader or Cookies variables. Default is null.
            * `transforms` (`list`) - List of transforms.

          * `name` (`str`) - Describes the name of the rule.
          * `priority` (`float`) - Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
          * `rate_limit_duration_in_minutes` (`float`) - Time window for resetting the rate limit count. Default is 1 minute.
          * `rate_limit_threshold` (`float`) - Number of allowed requests per client within the time window.
          * `rule_type` (`str`) - Describes type of rule.

      * `frontend_endpoint_links` (`list`) - Describes Frontend Endpoints associated with this Web Application Firewall policy.
        * `id` (`str`) - Resource ID.

      * `managed_rules` (`dict`) - Describes managed rules inside the policy.
        * `managed_rule_sets` (`list`) - List of rule sets.
          * `rule_group_overrides` (`list`) - Defines the rule group overrides to apply to the rule set.
            * `rule_group_name` (`str`) - Describes the managed rule group to override.
            * `rules` (`list`) - List of rules that will be disabled. If none specified, all rules in the group will be disabled.
              * `action` (`str`) - Describes the override action to be applied when rule matches.
              * `enabled_state` (`str`) - Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
              * `rule_id` (`str`) - Identifier for the managed rule.

          * `rule_set_type` (`str`) - Defines the rule set type to use.
          * `rule_set_version` (`str`) - Defines the version of the rule set to use.

      * `policy_settings` (`dict`) - Describes settings for the policy.
        * `custom_block_response_body` (`str`) - If the action type is block, customer can override the response body. The body must be specified in base64 encoding.
        * `custom_block_response_status_code` (`float`) - If the action type is block, customer can override the response status code.
        * `enabled_state` (`str`) - Describes if the policy is in enabled or disabled state. Defaults to Enabled if not specified.
        * `mode` (`str`) - Describes if it is in detection mode or prevention mode at policy level.
        * `redirect_url` (`str`) - If action type is redirect, this field represents redirect URL for the client.

      * `provisioning_state` (`str`) - Provisioning state of the policy.
      * `resource_state` (`str`)
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, custom_rules=None, etag=None, location=None, managed_rules=None, name=None, policy_settings=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Defines web application firewall policy.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] custom_rules: Describes custom rules inside the policy.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[dict] managed_rules: Describes managed rules inside the policy.
        :param pulumi.Input[str] name: The name of the Web Application Firewall Policy.
        :param pulumi.Input[dict] policy_settings: Describes settings for the policy.
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[dict] tags: Resource tags.

        The **custom_rules** object supports the following:

          * `rules` (`pulumi.Input[list]`) - List of rules
            * `action` (`pulumi.Input[str]`) - Describes what action to be applied when rule matches.
            * `enabled_state` (`pulumi.Input[str]`) - Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
            * `match_conditions` (`pulumi.Input[list]`) - List of match conditions.
              * `match_value` (`pulumi.Input[list]`) - List of possible match values.
              * `match_variable` (`pulumi.Input[str]`) - Request variable to compare with.
              * `negate_condition` (`pulumi.Input[bool]`) - Describes if the result of this condition should be negated.
              * `operator` (`pulumi.Input[str]`) - Comparison type to use for matching with the variable value.
              * `selector` (`pulumi.Input[str]`) - Match against a specific key from the QueryString, PostArgs, RequestHeader or Cookies variables. Default is null.
              * `transforms` (`pulumi.Input[list]`) - List of transforms.

            * `name` (`pulumi.Input[str]`) - Describes the name of the rule.
            * `priority` (`pulumi.Input[float]`) - Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
            * `rate_limit_duration_in_minutes` (`pulumi.Input[float]`) - Time window for resetting the rate limit count. Default is 1 minute.
            * `rate_limit_threshold` (`pulumi.Input[float]`) - Number of allowed requests per client within the time window.
            * `rule_type` (`pulumi.Input[str]`) - Describes type of rule.

        The **managed_rules** object supports the following:

          * `managed_rule_sets` (`pulumi.Input[list]`) - List of rule sets.
            * `rule_group_overrides` (`pulumi.Input[list]`) - Defines the rule group overrides to apply to the rule set.
              * `rule_group_name` (`pulumi.Input[str]`) - Describes the managed rule group to override.
              * `rules` (`pulumi.Input[list]`) - List of rules that will be disabled. If none specified, all rules in the group will be disabled.
                * `action` (`pulumi.Input[str]`) - Describes the override action to be applied when rule matches.
                * `enabled_state` (`pulumi.Input[str]`) - Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
                * `rule_id` (`pulumi.Input[str]`) - Identifier for the managed rule.

            * `rule_set_type` (`pulumi.Input[str]`) - Defines the rule set type to use.
            * `rule_set_version` (`pulumi.Input[str]`) - Defines the version of the rule set to use.

        The **policy_settings** object supports the following:

          * `custom_block_response_body` (`pulumi.Input[str]`) - If the action type is block, customer can override the response body. The body must be specified in base64 encoding.
          * `custom_block_response_status_code` (`pulumi.Input[float]`) - If the action type is block, customer can override the response status code.
          * `enabled_state` (`pulumi.Input[str]`) - Describes if the policy is in enabled or disabled state. Defaults to Enabled if not specified.
          * `mode` (`pulumi.Input[str]`) - Describes if it is in detection mode or prevention mode at policy level.
          * `redirect_url` (`pulumi.Input[str]`) - If action type is redirect, this field represents redirect URL for the client.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['custom_rules'] = custom_rules
            __props__['etag'] = etag
            __props__['location'] = location
            __props__['managed_rules'] = managed_rules
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['policy_settings'] = policy_settings
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(Policy, __self__).__init__(
            'azurerm:network/v20190301:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Policy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
