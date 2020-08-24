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
    'GetNatGatewayResult',
    'AwaitableGetNatGatewayResult',
    'get_nat_gateway',
]

@pulumi.output_type
class GetNatGatewayResult:
    """
    Nat Gateway resource.
    """
    def __init__(__self__, etag=None, idle_timeout_in_minutes=None, location=None, name=None, provisioning_state=None, public_ip_addresses=None, public_ip_prefixes=None, resource_guid=None, sku=None, subnets=None, tags=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if idle_timeout_in_minutes and not isinstance(idle_timeout_in_minutes, float):
            raise TypeError("Expected argument 'idle_timeout_in_minutes' to be a float")
        pulumi.set(__self__, "idle_timeout_in_minutes", idle_timeout_in_minutes)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if public_ip_addresses and not isinstance(public_ip_addresses, list):
            raise TypeError("Expected argument 'public_ip_addresses' to be a list")
        pulumi.set(__self__, "public_ip_addresses", public_ip_addresses)
        if public_ip_prefixes and not isinstance(public_ip_prefixes, list):
            raise TypeError("Expected argument 'public_ip_prefixes' to be a list")
        pulumi.set(__self__, "public_ip_prefixes", public_ip_prefixes)
        if resource_guid and not isinstance(resource_guid, str):
            raise TypeError("Expected argument 'resource_guid' to be a str")
        pulumi.set(__self__, "resource_guid", resource_guid)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if subnets and not isinstance(subnets, list):
            raise TypeError("Expected argument 'subnets' to be a list")
        pulumi.set(__self__, "subnets", subnets)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="idleTimeoutInMinutes")
    def idle_timeout_in_minutes(self) -> Optional[float]:
        """
        The idle timeout of the nat gateway.
        """
        return pulumi.get(self, "idle_timeout_in_minutes")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        The provisioning state of the NatGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicIpAddresses")
    def public_ip_addresses(self) -> Optional[List['outputs.SubResourceResponse']]:
        """
        An array of public ip addresses associated with the nat gateway resource.
        """
        return pulumi.get(self, "public_ip_addresses")

    @property
    @pulumi.getter(name="publicIpPrefixes")
    def public_ip_prefixes(self) -> Optional[List['outputs.SubResourceResponse']]:
        """
        An array of public ip prefixes associated with the nat gateway resource.
        """
        return pulumi.get(self, "public_ip_prefixes")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> Optional[str]:
        """
        The resource GUID property of the nat gateway resource.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.NatGatewaySkuResponse']:
        """
        The nat gateway SKU.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def subnets(self) -> List['outputs.SubResourceResponse']:
        """
        An array of references to the subnets using this nat gateway resource.
        """
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetNatGatewayResult(GetNatGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNatGatewayResult(
            etag=self.etag,
            idle_timeout_in_minutes=self.idle_timeout_in_minutes,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            public_ip_addresses=self.public_ip_addresses,
            public_ip_prefixes=self.public_ip_prefixes,
            resource_guid=self.resource_guid,
            sku=self.sku,
            subnets=self.subnets,
            tags=self.tags,
            type=self.type)


def get_nat_gateway(expand: Optional[str] = None,
                    name: Optional[str] = None,
                    resource_group_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNatGatewayResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: Expands referenced resources.
    :param str name: The name of the nat gateway.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20190201:getNatGateway', __args__, opts=opts, typ=GetNatGatewayResult).value

    return AwaitableGetNatGatewayResult(
        etag=__ret__.etag,
        idle_timeout_in_minutes=__ret__.idle_timeout_in_minutes,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        public_ip_addresses=__ret__.public_ip_addresses,
        public_ip_prefixes=__ret__.public_ip_prefixes,
        resource_guid=__ret__.resource_guid,
        sku=__ret__.sku,
        subnets=__ret__.subnets,
        tags=__ret__.tags,
        type=__ret__.type)
