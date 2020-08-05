# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class DeviceSecurityGroup(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Device Security group data
      * `allowlist_rules` (`list`) - The allow-list custom alert rules.
        * `allowlist_values` (`list`) - The values to allow. The format of the values depends on the rule type.
        * `description` (`str`) - The description of the custom alert.
        * `display_name` (`str`) - The display name of the custom alert.
        * `is_enabled` (`bool`) - Status of the custom alert.
        * `rule_type` (`str`) - The type of the custom alert rule.
        * `value_type` (`str`) - The value type of the items in the list.

      * `denylist_rules` (`list`) - The deny-list custom alert rules.
        * `denylist_values` (`list`) - The values to deny. The format of the values depends on the rule type.
        * `description` (`str`) - The description of the custom alert.
        * `display_name` (`str`) - The display name of the custom alert.
        * `is_enabled` (`bool`) - Status of the custom alert.
        * `rule_type` (`str`) - The type of the custom alert rule.
        * `value_type` (`str`) - The value type of the items in the list.

      * `threshold_rules` (`list`) - The list of custom alert threshold rules.
        * `description` (`str`) - The description of the custom alert.
        * `display_name` (`str`) - The display name of the custom alert.
        * `is_enabled` (`bool`) - Status of the custom alert.
        * `max_threshold` (`float`) - The maximum threshold.
        * `min_threshold` (`float`) - The minimum threshold.
        * `rule_type` (`str`) - The type of the custom alert rule.

      * `time_window_rules` (`list`) - The list of custom alert time-window rules.
        * `description` (`str`) - The description of the custom alert.
        * `display_name` (`str`) - The display name of the custom alert.
        * `is_enabled` (`bool`) - Status of the custom alert.
        * `max_threshold` (`float`) - The maximum threshold.
        * `min_threshold` (`float`) - The minimum threshold.
        * `rule_type` (`str`) - The type of the custom alert rule.
        * `time_window_size` (`str`) - The time window size in iso8601 format.
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, allowlist_rules=None, denylist_rules=None, name=None, resource_id=None, threshold_rules=None, time_window_rules=None, __props__=None, __name__=None, __opts__=None):
        """
        The device security group resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowlist_rules: The allow-list custom alert rules.
        :param pulumi.Input[list] denylist_rules: The deny-list custom alert rules.
        :param pulumi.Input[str] name: The name of the device security group. Note that the name of the device security group is case insensitive.
        :param pulumi.Input[str] resource_id: The identifier of the resource.
        :param pulumi.Input[list] threshold_rules: The list of custom alert threshold rules.
        :param pulumi.Input[list] time_window_rules: The list of custom alert time-window rules.

        The **allowlist_rules** object supports the following:

          * `allowlist_values` (`pulumi.Input[list]`) - The values to allow. The format of the values depends on the rule type.
          * `is_enabled` (`pulumi.Input[bool]`) - Status of the custom alert.
          * `rule_type` (`pulumi.Input[str]`) - The type of the custom alert rule.

        The **denylist_rules** object supports the following:

          * `denylist_values` (`pulumi.Input[list]`) - The values to deny. The format of the values depends on the rule type.
          * `is_enabled` (`pulumi.Input[bool]`) - Status of the custom alert.
          * `rule_type` (`pulumi.Input[str]`) - The type of the custom alert rule.

        The **threshold_rules** object supports the following:

          * `is_enabled` (`pulumi.Input[bool]`) - Status of the custom alert.
          * `max_threshold` (`pulumi.Input[float]`) - The maximum threshold.
          * `min_threshold` (`pulumi.Input[float]`) - The minimum threshold.
          * `rule_type` (`pulumi.Input[str]`) - The type of the custom alert rule.

        The **time_window_rules** object supports the following:

          * `is_enabled` (`pulumi.Input[bool]`) - Status of the custom alert.
          * `max_threshold` (`pulumi.Input[float]`) - The maximum threshold.
          * `min_threshold` (`pulumi.Input[float]`) - The minimum threshold.
          * `rule_type` (`pulumi.Input[str]`) - The type of the custom alert rule.
          * `time_window_size` (`pulumi.Input[str]`) - The time window size in iso8601 format.
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

            __props__['allowlist_rules'] = allowlist_rules
            __props__['denylist_rules'] = denylist_rules
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_id is None:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            __props__['threshold_rules'] = threshold_rules
            __props__['time_window_rules'] = time_window_rules
            __props__['properties'] = None
            __props__['type'] = None
        super(DeviceSecurityGroup, __self__).__init__(
            'azurerm:security/v20190801:DeviceSecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing DeviceSecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DeviceSecurityGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
