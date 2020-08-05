# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PolicySetDefinitionAtManagementGroup(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the policy set definition.
    """
    properties: pulumi.Output[dict]
    """
    The policy definition properties.
      * `description` (`str`) - The policy set definition description.
      * `display_name` (`str`) - The display name of the policy set definition.
      * `metadata` (`dict`) - The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
      * `parameters` (`dict`) - The policy set definition parameters that can be used in policy definition references.
      * `policy_definition_groups` (`list`) - The metadata describing groups of policy definition references within the policy set definition.
        * `additional_metadata_id` (`str`) - A resource ID of a resource that contains additional metadata about the group.
        * `category` (`str`) - The group's category.
        * `description` (`str`) - The group's description.
        * `display_name` (`str`) - The group's display name.
        * `name` (`str`) - The name of the group.

      * `policy_definitions` (`list`) - An array of policy definition references.
        * `group_names` (`list`) - The name of the groups that this policy definition reference belongs to.
        * `parameters` (`dict`) - The parameter values for the referenced policy rule. The keys are the parameter names.
        * `policy_definition_id` (`str`) - The ID of the policy definition or policy set definition.
        * `policy_definition_reference_id` (`str`) - A unique id (within the policy set definition) for this policy definition reference.

      * `policy_type` (`str`) - The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
    """
    type: pulumi.Output[str]
    """
    The type of the resource (Microsoft.Authorization/policySetDefinitions).
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, management_group_id=None, metadata=None, name=None, parameters=None, policy_definition_groups=None, policy_definitions=None, policy_type=None, __props__=None, __name__=None, __opts__=None):
        """
        The policy set definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The policy set definition description.
        :param pulumi.Input[str] display_name: The display name of the policy set definition.
        :param pulumi.Input[str] management_group_id: The ID of the management group.
        :param pulumi.Input[dict] metadata: The policy set definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.
        :param pulumi.Input[str] name: The name of the policy set definition to create.
        :param pulumi.Input[dict] parameters: The policy set definition parameters that can be used in policy definition references.
        :param pulumi.Input[list] policy_definition_groups: The metadata describing groups of policy definition references within the policy set definition.
        :param pulumi.Input[list] policy_definitions: An array of policy definition references.
        :param pulumi.Input[str] policy_type: The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.

        The **policy_definition_groups** object supports the following:

          * `additional_metadata_id` (`pulumi.Input[str]`) - A resource ID of a resource that contains additional metadata about the group.
          * `category` (`pulumi.Input[str]`) - The group's category.
          * `description` (`pulumi.Input[str]`) - The group's description.
          * `display_name` (`pulumi.Input[str]`) - The group's display name.
          * `name` (`pulumi.Input[str]`) - The name of the group.

        The **policy_definitions** object supports the following:

          * `group_names` (`pulumi.Input[list]`) - The name of the groups that this policy definition reference belongs to.
          * `parameters` (`pulumi.Input[dict]`) - The parameter values for the referenced policy rule. The keys are the parameter names.
          * `policy_definition_id` (`pulumi.Input[str]`) - The ID of the policy definition or policy set definition.
          * `policy_definition_reference_id` (`pulumi.Input[str]`) - A unique id (within the policy set definition) for this policy definition reference.
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
            if management_group_id is None:
                raise TypeError("Missing required property 'management_group_id'")
            __props__['management_group_id'] = management_group_id
            __props__['metadata'] = metadata
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['parameters'] = parameters
            __props__['policy_definition_groups'] = policy_definition_groups
            if policy_definitions is None:
                raise TypeError("Missing required property 'policy_definitions'")
            __props__['policy_definitions'] = policy_definitions
            __props__['policy_type'] = policy_type
            __props__['properties'] = None
            __props__['type'] = None
        super(PolicySetDefinitionAtManagementGroup, __self__).__init__(
            'azurerm:management/v20190901:PolicySetDefinitionAtManagementGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing PolicySetDefinitionAtManagementGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PolicySetDefinitionAtManagementGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
