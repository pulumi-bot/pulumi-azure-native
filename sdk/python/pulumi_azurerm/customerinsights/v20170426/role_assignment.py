# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class RoleAssignment(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    The Role Assignment definition.
      * `assignment_name` (`str`) - The name of the metadata object.
      * `conflation_policies` (`dict`) - Widget types set for the assignment.
        * `elements` (`list`) - The elements included in the set.
        * `exceptions` (`list`) - The elements that are not included in the set, in case elements contains '*' indicating 'all'.

      * `connectors` (`dict`) - Connectors set for the assignment.
      * `description` (`dict`) - Localized description for the metadata.
      * `display_name` (`dict`) - Localized display names for the metadata.
      * `interactions` (`dict`) - Interactions set for the assignment.
      * `kpis` (`dict`) - Kpis set for the assignment.
      * `links` (`dict`) - Links set for the assignment.
      * `principals` (`list`) - The principals being assigned to.
        * `principal_id` (`str`) - The principal id being assigned to.
        * `principal_metadata` (`dict`) - Other metadata for the principal.
        * `principal_type` (`str`) - The Type of the principal ID.

      * `profiles` (`dict`) - Profiles set for the assignment.
      * `provisioning_state` (`str`) - Provisioning state.
      * `relationship_links` (`dict`) - The Role assignments set for the relationship links.
      * `relationships` (`dict`) - The Role assignments set for the relationships.
      * `role` (`str`) - Type of roles.
      * `role_assignments` (`dict`) - The Role assignments set for the assignment.
      * `sas_policies` (`dict`) - Sas Policies set for the assignment.
      * `segments` (`dict`) - The Role assignments set for the assignment.
      * `tenant_id` (`str`) - The hub name.
      * `views` (`dict`) - Views set for the assignment.
      * `widget_types` (`dict`) - Widget types set for the assignment.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, conflation_policies=None, connectors=None, description=None, display_name=None, hub_name=None, interactions=None, kpis=None, links=None, name=None, principals=None, profiles=None, relationship_links=None, relationships=None, resource_group_name=None, role=None, role_assignments=None, sas_policies=None, segments=None, views=None, widget_types=None, __props__=None, __name__=None, __opts__=None):
        """
        The Role Assignment resource format.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] conflation_policies: Widget types set for the assignment.
        :param pulumi.Input[dict] connectors: Connectors set for the assignment.
        :param pulumi.Input[dict] description: Localized description for the metadata.
        :param pulumi.Input[dict] display_name: Localized display names for the metadata.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[dict] interactions: Interactions set for the assignment.
        :param pulumi.Input[dict] kpis: Kpis set for the assignment.
        :param pulumi.Input[dict] links: Links set for the assignment.
        :param pulumi.Input[str] name: The assignment name
        :param pulumi.Input[list] principals: The principals being assigned to.
        :param pulumi.Input[dict] profiles: Profiles set for the assignment.
        :param pulumi.Input[dict] relationship_links: The Role assignments set for the relationship links.
        :param pulumi.Input[dict] relationships: The Role assignments set for the relationships.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] role: Type of roles.
        :param pulumi.Input[dict] role_assignments: The Role assignments set for the assignment.
        :param pulumi.Input[dict] sas_policies: Sas Policies set for the assignment.
        :param pulumi.Input[dict] segments: The Role assignments set for the assignment.
        :param pulumi.Input[dict] views: Views set for the assignment.
        :param pulumi.Input[dict] widget_types: Widget types set for the assignment.

        The **conflation_policies** object supports the following:

          * `elements` (`pulumi.Input[list]`) - The elements included in the set.
          * `exceptions` (`pulumi.Input[list]`) - The elements that are not included in the set, in case elements contains '*' indicating 'all'.

        The **principals** object supports the following:

          * `principal_id` (`pulumi.Input[str]`) - The principal id being assigned to.
          * `principal_metadata` (`pulumi.Input[dict]`) - Other metadata for the principal.
          * `principal_type` (`pulumi.Input[str]`) - The Type of the principal ID.
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

            __props__['conflation_policies'] = conflation_policies
            __props__['connectors'] = connectors
            __props__['description'] = description
            __props__['display_name'] = display_name
            if hub_name is None:
                raise TypeError("Missing required property 'hub_name'")
            __props__['hub_name'] = hub_name
            __props__['interactions'] = interactions
            __props__['kpis'] = kpis
            __props__['links'] = links
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if principals is None:
                raise TypeError("Missing required property 'principals'")
            __props__['principals'] = principals
            __props__['profiles'] = profiles
            __props__['relationship_links'] = relationship_links
            __props__['relationships'] = relationships
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['role_assignments'] = role_assignments
            __props__['sas_policies'] = sas_policies
            __props__['segments'] = segments
            __props__['views'] = views
            __props__['widget_types'] = widget_types
            __props__['properties'] = None
            __props__['type'] = None
        super(RoleAssignment, __self__).__init__(
            'azurerm:customerinsights/v20170426:RoleAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing RoleAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RoleAssignment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
