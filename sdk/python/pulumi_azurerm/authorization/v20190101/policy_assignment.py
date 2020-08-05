# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PolicyAssignment(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The managed identity associated with the policy assignment.
      * `principal_id` (`str`) - The principal ID of the resource identity.
      * `tenant_id` (`str`) - The tenant ID of the resource identity.
      * `type` (`str`) - The identity type.
    """
    location: pulumi.Output[str]
    """
    The location of the policy assignment. Only required when utilizing managed identity.
    """
    name: pulumi.Output[str]
    """
    The name of the policy assignment.
    """
    properties: pulumi.Output[dict]
    """
    Properties for the policy assignment.
      * `description` (`str`) - This message will be part of response in case of policy violation.
      * `display_name` (`str`) - The display name of the policy assignment.
      * `metadata` (`dict`) - The policy assignment metadata.
      * `not_scopes` (`list`) - The policy's excluded scopes.
      * `parameters` (`dict`) - Required if a parameter is used in policy rule.
      * `policy_definition_id` (`str`) - The ID of the policy definition or policy set definition being assigned.
      * `scope` (`str`) - The scope for the policy assignment.
    """
    sku: pulumi.Output[dict]
    """
    The policy sku. This property is optional, obsolete, and will be ignored.
      * `name` (`str`) - The name of the policy sku. Possible values are A0 and A1.
      * `tier` (`str`) - The policy sku tier. Possible values are Free and Standard.
    """
    type: pulumi.Output[str]
    """
    The type of the policy assignment.
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, identity=None, location=None, metadata=None, name=None, not_scopes=None, parameters=None, policy_definition_id=None, scope=None, sku=None, __props__=None, __name__=None, __opts__=None):
        """
        The policy assignment.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: This message will be part of response in case of policy violation.
        :param pulumi.Input[str] display_name: The display name of the policy assignment.
        :param pulumi.Input[dict] identity: The managed identity associated with the policy assignment.
        :param pulumi.Input[str] location: The location of the policy assignment. Only required when utilizing managed identity.
        :param pulumi.Input[dict] metadata: The policy assignment metadata.
        :param pulumi.Input[str] name: The name of the policy assignment.
        :param pulumi.Input[list] not_scopes: The policy's excluded scopes.
        :param pulumi.Input[dict] parameters: Required if a parameter is used in policy rule.
        :param pulumi.Input[str] policy_definition_id: The ID of the policy definition or policy set definition being assigned.
        :param pulumi.Input[str] scope: The scope for the policy assignment.
        :param pulumi.Input[dict] sku: The policy sku. This property is optional, obsolete, and will be ignored.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the policy sku. Possible values are A0 and A1.
          * `tier` (`pulumi.Input[str]`) - The policy sku tier. Possible values are Free and Standard.
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
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['metadata'] = metadata
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['not_scopes'] = not_scopes
            __props__['parameters'] = parameters
            __props__['policy_definition_id'] = policy_definition_id
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['sku'] = sku
            __props__['properties'] = None
            __props__['type'] = None
        super(PolicyAssignment, __self__).__init__(
            'azurerm:authorization/v20190101:PolicyAssignment',
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
