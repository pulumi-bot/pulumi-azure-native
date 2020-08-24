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
    'GetPeeringServiceResult',
    'AwaitableGetPeeringServiceResult',
    'get_peering_service',
]

@pulumi.output_type
class GetPeeringServiceResult:
    """
    Peering Service
    """
    def __init__(__self__, location=None, name=None, peering_service_location=None, peering_service_provider=None, provisioning_state=None, sku=None, tags=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if peering_service_location and not isinstance(peering_service_location, str):
            raise TypeError("Expected argument 'peering_service_location' to be a str")
        pulumi.set(__self__, "peering_service_location", peering_service_location)
        if peering_service_provider and not isinstance(peering_service_provider, str):
            raise TypeError("Expected argument 'peering_service_provider' to be a str")
        pulumi.set(__self__, "peering_service_provider", peering_service_provider)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peeringServiceLocation")
    def peering_service_location(self) -> Optional[str]:
        """
        The PeeringServiceLocation of the Customer.
        """
        return pulumi.get(self, "peering_service_location")

    @property
    @pulumi.getter(name="peeringServiceProvider")
    def peering_service_provider(self) -> Optional[str]:
        """
        The MAPS Provider Name.
        """
        return pulumi.get(self, "peering_service_provider")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.PeeringServiceSkuResponse']:
        """
        The SKU that defines the type of the peering service.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetPeeringServiceResult(GetPeeringServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPeeringServiceResult(
            location=self.location,
            name=self.name,
            peering_service_location=self.peering_service_location,
            peering_service_provider=self.peering_service_provider,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_peering_service(name: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPeeringServiceResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the peering.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:peering/v20200401:getPeeringService', __args__, opts=opts, typ=GetPeeringServiceResult).value

    return AwaitableGetPeeringServiceResult(
        location=__ret__.location,
        name=__ret__.name,
        peering_service_location=__ret__.peering_service_location,
        peering_service_provider=__ret__.peering_service_provider,
        provisioning_state=__ret__.provisioning_state,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type)
