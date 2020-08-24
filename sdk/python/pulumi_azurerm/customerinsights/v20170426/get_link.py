# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetLinkResult',
    'AwaitableGetLinkResult',
    'get_link',
]

@pulumi.output_type
class GetLinkResult:
    """
    The link resource format.
    """
    def __init__(__self__, description=None, display_name=None, link_name=None, mappings=None, name=None, operation_type=None, participant_property_references=None, provisioning_state=None, reference_only=None, source_entity_type=None, source_entity_type_name=None, target_entity_type=None, target_entity_type_name=None, tenant_id=None, type=None):
        if description and not isinstance(description, dict):
            raise TypeError("Expected argument 'description' to be a dict")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, dict):
            raise TypeError("Expected argument 'display_name' to be a dict")
        pulumi.set(__self__, "display_name", display_name)
        if link_name and not isinstance(link_name, str):
            raise TypeError("Expected argument 'link_name' to be a str")
        pulumi.set(__self__, "link_name", link_name)
        if mappings and not isinstance(mappings, list):
            raise TypeError("Expected argument 'mappings' to be a list")
        pulumi.set(__self__, "mappings", mappings)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if operation_type and not isinstance(operation_type, str):
            raise TypeError("Expected argument 'operation_type' to be a str")
        pulumi.set(__self__, "operation_type", operation_type)
        if participant_property_references and not isinstance(participant_property_references, list):
            raise TypeError("Expected argument 'participant_property_references' to be a list")
        pulumi.set(__self__, "participant_property_references", participant_property_references)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if reference_only and not isinstance(reference_only, bool):
            raise TypeError("Expected argument 'reference_only' to be a bool")
        pulumi.set(__self__, "reference_only", reference_only)
        if source_entity_type and not isinstance(source_entity_type, str):
            raise TypeError("Expected argument 'source_entity_type' to be a str")
        pulumi.set(__self__, "source_entity_type", source_entity_type)
        if source_entity_type_name and not isinstance(source_entity_type_name, str):
            raise TypeError("Expected argument 'source_entity_type_name' to be a str")
        pulumi.set(__self__, "source_entity_type_name", source_entity_type_name)
        if target_entity_type and not isinstance(target_entity_type, str):
            raise TypeError("Expected argument 'target_entity_type' to be a str")
        pulumi.set(__self__, "target_entity_type", target_entity_type)
        if target_entity_type_name and not isinstance(target_entity_type_name, str):
            raise TypeError("Expected argument 'target_entity_type_name' to be a str")
        pulumi.set(__self__, "target_entity_type_name", target_entity_type_name)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

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
    @pulumi.getter(name="sourceEntityType")
    def source_entity_type(self) -> str:
        """
        Type of source entity.
        """
        return pulumi.get(self, "source_entity_type")

    @property
    @pulumi.getter(name="sourceEntityTypeName")
    def source_entity_type_name(self) -> str:
        """
        Name of the source Entity Type.
        """
        return pulumi.get(self, "source_entity_type_name")

    @property
    @pulumi.getter(name="targetEntityType")
    def target_entity_type(self) -> str:
        """
        Type of target entity.
        """
        return pulumi.get(self, "target_entity_type")

    @property
    @pulumi.getter(name="targetEntityTypeName")
    def target_entity_type_name(self) -> str:
        """
        Name of the target Entity Type.
        """
        return pulumi.get(self, "target_entity_type_name")

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


class AwaitableGetLinkResult(GetLinkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLinkResult(
            description=self.description,
            display_name=self.display_name,
            link_name=self.link_name,
            mappings=self.mappings,
            name=self.name,
            operation_type=self.operation_type,
            participant_property_references=self.participant_property_references,
            provisioning_state=self.provisioning_state,
            reference_only=self.reference_only,
            source_entity_type=self.source_entity_type,
            source_entity_type_name=self.source_entity_type_name,
            target_entity_type=self.target_entity_type,
            target_entity_type_name=self.target_entity_type_name,
            tenant_id=self.tenant_id,
            type=self.type)


def get_link(hub_name: Optional[str] = None,
             name: Optional[str] = None,
             resource_group_name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLinkResult:
    """
    Use this data source to access information about an existing resource.

    :param str hub_name: The name of the hub.
    :param str name: The name of the link.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['hubName'] = hub_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:customerinsights/v20170426:getLink', __args__, opts=opts, typ=GetLinkResult).value

    return AwaitableGetLinkResult(
        description=__ret__.description,
        display_name=__ret__.display_name,
        link_name=__ret__.link_name,
        mappings=__ret__.mappings,
        name=__ret__.name,
        operation_type=__ret__.operation_type,
        participant_property_references=__ret__.participant_property_references,
        provisioning_state=__ret__.provisioning_state,
        reference_only=__ret__.reference_only,
        source_entity_type=__ret__.source_entity_type,
        source_entity_type_name=__ret__.source_entity_type_name,
        target_entity_type=__ret__.target_entity_type,
        target_entity_type_name=__ret__.target_entity_type_name,
        tenant_id=__ret__.tenant_id,
        type=__ret__.type)
