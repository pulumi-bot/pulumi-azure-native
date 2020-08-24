# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetVirtualRouterPeeringResult',
    'AwaitableGetVirtualRouterPeeringResult',
    'get_virtual_router_peering',
]

@pulumi.output_type
class GetVirtualRouterPeeringResult:
    """
    Virtual Router Peering resource.
    """
    def __init__(__self__, etag=None, name=None, peer_asn=None, peer_ip=None, provisioning_state=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if peer_asn and not isinstance(peer_asn, float):
            raise TypeError("Expected argument 'peer_asn' to be a float")
        pulumi.set(__self__, "peer_asn", peer_asn)
        if peer_ip and not isinstance(peer_ip, str):
            raise TypeError("Expected argument 'peer_ip' to be a str")
        pulumi.set(__self__, "peer_ip", peer_ip)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the virtual router peering that is unique within a virtual router.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> Optional[float]:
        """
        Peer ASN.
        """
        return pulumi.get(self, "peer_asn")

    @property
    @pulumi.getter(name="peerIp")
    def peer_ip(self) -> Optional[str]:
        """
        Peer IP.
        """
        return pulumi.get(self, "peer_ip")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Peering type.
        """
        return pulumi.get(self, "type")


class AwaitableGetVirtualRouterPeeringResult(GetVirtualRouterPeeringResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualRouterPeeringResult(
            etag=self.etag,
            name=self.name,
            peer_asn=self.peer_asn,
            peer_ip=self.peer_ip,
            provisioning_state=self.provisioning_state,
            type=self.type)


def get_virtual_router_peering(name: Optional[str] = None,
                               resource_group_name: Optional[str] = None,
                               virtual_router_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualRouterPeeringResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the Virtual Router Peering.
    :param str resource_group_name: The name of the resource group.
    :param str virtual_router_name: The name of the Virtual Router.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['virtualRouterName'] = virtual_router_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20200401:getVirtualRouterPeering', __args__, opts=opts, typ=GetVirtualRouterPeeringResult).value

    return AwaitableGetVirtualRouterPeeringResult(
        etag=__ret__.etag,
        name=__ret__.name,
        peer_asn=__ret__.peer_asn,
        peer_ip=__ret__.peer_ip,
        provisioning_state=__ret__.provisioning_state,
        type=__ret__.type)
