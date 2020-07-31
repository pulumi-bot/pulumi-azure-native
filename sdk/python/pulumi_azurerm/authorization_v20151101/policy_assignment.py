# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class PolicyAssignment(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Gets or sets the policy assignment name.
    """
    properties: pulumi.Output[dict]
    """
    Gets or sets the policy assignment properties.
      * `display_name` (`str`) - Gets or sets the policy assignment display name.
      * `policy_definition_id` (`str`) - Gets or sets the policy definition Id.
      * `scope` (`str`) - Gets or sets the policy assignment scope.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, scope=None, __props__=None, __name__=None, __opts__=None):
        """
        Policy assignment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Policy assignment name.
        :param pulumi.Input[dict] properties: Gets or sets the policy assignment properties.
        :param pulumi.Input[str] scope: Scope.

        The **properties** object supports the following:

          * `display_name` (`pulumi.Input[str]`) - Gets or sets the policy assignment display name.
          * `policy_definition_id` (`pulumi.Input[str]`) - Gets or sets the policy definition Id.
          * `scope` (`pulumi.Input[str]`) - Gets or sets the policy assignment scope.
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
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
        super(PolicyAssignment, __self__).__init__(
            'azurerm:authorization/v20151101:PolicyAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing PolicyAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PolicyAssignment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop