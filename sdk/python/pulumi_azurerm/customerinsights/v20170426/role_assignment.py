# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['RoleAssignment']


class RoleAssignment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conflation_policies: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 connectors: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 description: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 hub_name: Optional[pulumi.Input[str]] = None,
                 interactions: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 kpis: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 links: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 principals: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AssignmentPrincipalArgs']]]]] = None,
                 profiles: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 relationship_links: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 relationships: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 role_assignments: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 sas_policies: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 segments: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 views: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 widget_types: Optional[pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The Role Assignment resource format.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] conflation_policies: Widget types set for the assignment.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] connectors: Connectors set for the assignment.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] description: Localized description for the metadata.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] display_name: Localized display names for the metadata.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] interactions: Interactions set for the assignment.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] kpis: Kpis set for the assignment.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] links: Links set for the assignment.
        :param pulumi.Input[str] name: The assignment name
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AssignmentPrincipalArgs']]]] principals: The principals being assigned to.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] profiles: Profiles set for the assignment.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] relationship_links: The Role assignments set for the relationship links.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] relationships: The Role assignments set for the relationships.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] role: Type of roles.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] role_assignments: The Role assignments set for the assignment.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] sas_policies: Sas Policies set for the assignment.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] segments: The Role assignments set for the assignment.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] views: Views set for the assignment.
        :param pulumi.Input[pulumi.InputType['ResourceSetDescriptionArgs']] widget_types: Widget types set for the assignment.
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
            __props__['assignment_name'] = None
            __props__['provisioning_state'] = None
            __props__['tenant_id'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:customerinsights/v20170101:RoleAssignment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RoleAssignment, __self__).__init__(
            'azurerm:customerinsights/v20170426:RoleAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RoleAssignment':
        """
        Get an existing RoleAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RoleAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="assignmentName")
    def assignment_name(self) -> str:
        """
        The name of the metadata object.
        """
        return pulumi.get(self, "assignment_name")

    @property
    @pulumi.getter(name="conflationPolicies")
    def conflation_policies(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        Widget types set for the assignment.
        """
        return pulumi.get(self, "conflation_policies")

    @property
    @pulumi.getter
    def connectors(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        Connectors set for the assignment.
        """
        return pulumi.get(self, "connectors")

    @property
    @pulumi.getter
    def description(self) -> Optional[Mapping[str, str]]:
        """
        Localized description for the metadata.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[Mapping[str, str]]:
        """
        Localized display names for the metadata.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def interactions(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        Interactions set for the assignment.
        """
        return pulumi.get(self, "interactions")

    @property
    @pulumi.getter
    def kpis(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        Kpis set for the assignment.
        """
        return pulumi.get(self, "kpis")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        Links set for the assignment.
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def principals(self) -> List['outputs.AssignmentPrincipalResponse']:
        """
        The principals being assigned to.
        """
        return pulumi.get(self, "principals")

    @property
    @pulumi.getter
    def profiles(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        Profiles set for the assignment.
        """
        return pulumi.get(self, "profiles")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="relationshipLinks")
    def relationship_links(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        The Role assignments set for the relationship links.
        """
        return pulumi.get(self, "relationship_links")

    @property
    @pulumi.getter
    def relationships(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        The Role assignments set for the relationships.
        """
        return pulumi.get(self, "relationships")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        Type of roles.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="roleAssignments")
    def role_assignments(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        The Role assignments set for the assignment.
        """
        return pulumi.get(self, "role_assignments")

    @property
    @pulumi.getter(name="sasPolicies")
    def sas_policies(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        Sas Policies set for the assignment.
        """
        return pulumi.get(self, "sas_policies")

    @property
    @pulumi.getter
    def segments(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        The Role assignments set for the assignment.
        """
        return pulumi.get(self, "segments")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The hub name.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def views(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        Views set for the assignment.
        """
        return pulumi.get(self, "views")

    @property
    @pulumi.getter(name="widgetTypes")
    def widget_types(self) -> Optional['outputs.ResourceSetDescriptionResponse']:
        """
        Widget types set for the assignment.
        """
        return pulumi.get(self, "widget_types")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

