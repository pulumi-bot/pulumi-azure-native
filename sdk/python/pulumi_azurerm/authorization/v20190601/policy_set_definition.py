# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PolicySetDefinition(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the policy set definition.
    """
    properties: pulumi.Output[dict]
    """
    The policy definition properties.
      * `description` (`str`) - The policy set definition description.
      * `display_name` (`str`) - The display name of the policy set definition.
      * `metadata` (`dict`) - The policy set definition metadata.
      * `parameters` (`dict`) - The policy set definition parameters that can be used in policy definition references.
      * `policy_definitions` (`list`) - An array of policy definition references.
        * `parameters` (`dict`) - Required if a parameter is used in policy rule.
        * `policy_definition_id` (`str`) - The ID of the policy definition or policy set definition.

      * `policy_type` (`str`) - The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
    """
    type: pulumi.Output[str]
    """
    The type of the resource (Microsoft.Authorization/policySetDefinitions).
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, __props__=None, __name__=None, __opts__=None):
        """
        The policy set definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the policy set definition to create.
        :param pulumi.Input[dict] properties: The policy definition properties.

        The **properties** object supports the following:

          * `description` (`pulumi.Input[str]`) - The policy set definition description.
          * `display_name` (`pulumi.Input[str]`) - The display name of the policy set definition.
          * `metadata` (`pulumi.Input[dict]`) - The policy set definition metadata.
          * `parameters` (`pulumi.Input[dict]`) - The policy set definition parameters that can be used in policy definition references.
          * `policy_definitions` (`pulumi.Input[list]`) - An array of policy definition references.
            * `parameters` (`pulumi.Input[dict]`) - Required if a parameter is used in policy rule.
            * `policy_definition_id` (`pulumi.Input[str]`) - The ID of the policy definition or policy set definition.

          * `policy_type` (`pulumi.Input[str]`) - The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            __props__['type'] = None
        super(PolicySetDefinition, __self__).__init__(
            'azurerm:authorization/v20190601:PolicySetDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing PolicySetDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PolicySetDefinition(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
