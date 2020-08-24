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

__all__ = ['Link']


class Link(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 hub_name: Optional[pulumi.Input[str]] = None,
                 mappings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TypePropertiesMappingArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operation_type: Optional[pulumi.Input[str]] = None,
                 participant_property_references: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ParticipantPropertyReferenceArgs']]]]] = None,
                 reference_only: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_interaction_type: Optional[pulumi.Input[str]] = None,
                 target_profile_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The link resource format.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] description: Localized descriptions for the Link.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] display_name: Localized display name for the Link.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TypePropertiesMappingArgs']]]] mappings: The set of properties mappings between the source and target Types.
        :param pulumi.Input[str] name: The name of the link.
        :param pulumi.Input[str] operation_type: Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ParticipantPropertyReferenceArgs']]]] participant_property_references: The properties that represent the participating profile.
        :param pulumi.Input[bool] reference_only: Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] source_interaction_type: Name of the source Interaction Type.
        :param pulumi.Input[str] target_profile_type: Name of the target Profile Type.
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
            if hub_name is None:
                raise TypeError("Missing required property 'hub_name'")
            __props__['hub_name'] = hub_name
            __props__['mappings'] = mappings
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['operation_type'] = operation_type
            if participant_property_references is None:
                raise TypeError("Missing required property 'participant_property_references'")
            __props__['participant_property_references'] = participant_property_references
            __props__['reference_only'] = reference_only
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source_interaction_type is None:
                raise TypeError("Missing required property 'source_interaction_type'")
            __props__['source_interaction_type'] = source_interaction_type
            if target_profile_type is None:
                raise TypeError("Missing required property 'target_profile_type'")
            __props__['target_profile_type'] = target_profile_type
            __props__['link_name'] = None
            __props__['provisioning_state'] = None
            __props__['tenant_id'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:customerinsights/v20170426:Link")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Link, __self__).__init__(
            'azurerm:customerinsights/v20170101:Link',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Link':
        """
        Get an existing Link resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Link(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> Optional[Mapping[str, str]]:
        """
        Localized descriptions for the Link.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[Mapping[str, str]]:
        """
        Localized display name for the Link.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="linkName")
    def link_name(self) -> str:
        """
        The link name.
        """
        return pulumi.get(self, "link_name")

    @property
    @pulumi.getter
    def mappings(self) -> Optional[List['outputs.TypePropertiesMappingResponse']]:
        """
        The set of properties mappings between the source and target Types.
        """
        return pulumi.get(self, "mappings")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operationType")
    def operation_type(self) -> Optional[str]:
        """
        Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
        """
        return pulumi.get(self, "operation_type")

    @property
    @pulumi.getter(name="participantPropertyReferences")
    def participant_property_references(self) -> List['outputs.ParticipantPropertyReferenceResponse']:
        """
        The properties that represent the participating profile.
        """
        return pulumi.get(self, "participant_property_references")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="referenceOnly")
    def reference_only(self) -> Optional[bool]:
        """
        Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
        """
        return pulumi.get(self, "reference_only")

    @property
    @pulumi.getter(name="sourceInteractionType")
    def source_interaction_type(self) -> str:
        """
        Name of the source Interaction Type.
        """
        return pulumi.get(self, "source_interaction_type")

    @property
    @pulumi.getter(name="targetProfileType")
    def target_profile_type(self) -> str:
        """
        Name of the target Profile Type.
        """
        return pulumi.get(self, "target_profile_type")

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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

