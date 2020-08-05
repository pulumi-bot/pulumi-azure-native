# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PolicyDefinition(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the policy definition.
    """
    properties: pulumi.Output[dict]
    """
    The policy definition properties.
      * `description` (`str`) - The policy definition description.
      * `display_name` (`str`) - The display name of the policy definition.
      * `metadata` (`dict`) - The policy definition metadata.
      * `mode` (`str`) - The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
      * `parameters` (`dict`) - Required if a parameter is used in policy rule.
      * `policy_rule` (`dict`) - The policy rule.
      * `policy_type` (`str`) - The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
    """
    type: pulumi.Output[str]
    """
    The type of the resource (Microsoft.Authorization/policyDefinitions).
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, metadata=None, mode=None, name=None, parameters=None, policy_rule=None, policy_type=None, __props__=None, __name__=None, __opts__=None):
        """
        The policy definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The policy definition description.
        :param pulumi.Input[str] display_name: The display name of the policy definition.
        :param pulumi.Input[dict] metadata: The policy definition metadata.
        :param pulumi.Input[str] mode: The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
        :param pulumi.Input[str] name: The name of the policy definition to create.
        :param pulumi.Input[dict] parameters: Required if a parameter is used in policy rule.
        :param pulumi.Input[dict] policy_rule: The policy rule.
        :param pulumi.Input[str] policy_type: The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
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

            __props__['description'] = description
            __props__['display_name'] = display_name
            __props__['metadata'] = metadata
            __props__['mode'] = mode
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['policy_rule'] = policy_rule
            __props__['policy_type'] = policy_type
            __props__['properties'] = None
            __props__['type'] = None
        super(PolicyDefinition, __self__).__init__(
            'azurerm:authorization/v20190601:PolicyDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing PolicyDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PolicyDefinition(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
