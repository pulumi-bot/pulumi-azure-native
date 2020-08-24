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
    'GetNetworkInterfaceTapConfigurationResult',
    'AwaitableGetNetworkInterfaceTapConfigurationResult',
    'get_network_interface_tap_configuration',
]

@pulumi.output_type
class GetNetworkInterfaceTapConfigurationResult:
    """
    Tap configuration in a Network Interface.
    """
    def __init__(__self__, etag=None, name=None, provisioning_state=None, type=None, virtual_network_tap=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_network_tap and not isinstance(virtual_network_tap, dict):
            raise TypeError("Expected argument 'virtual_network_tap' to be a dict")
        pulumi.set(__self__, "virtual_network_tap", virtual_network_tap)

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
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the network interface tap configuration resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Sub Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworkTap")
    def virtual_network_tap(self) -> Optional['outputs.VirtualNetworkTapResponse']:
        """
        The reference to the Virtual Network Tap resource.
        """
        return pulumi.get(self, "virtual_network_tap")


class AwaitableGetNetworkInterfaceTapConfigurationResult(GetNetworkInterfaceTapConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkInterfaceTapConfigurationResult(
            etag=self.etag,
            name=self.name,
            provisioning_state=self.provisioning_state,
            type=self.type,
            virtual_network_tap=self.virtual_network_tap)


def get_network_interface_tap_configuration(name: Optional[str] = None,
                                            network_interface_name: Optional[str] = None,
                                            resource_group_name: Optional[str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkInterfaceTapConfigurationResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the tap configuration.
    :param str network_interface_name: The name of the network interface.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['networkInterfaceName'] = network_interface_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20200501:getNetworkInterfaceTapConfiguration', __args__, opts=opts, typ=GetNetworkInterfaceTapConfigurationResult).value

    return AwaitableGetNetworkInterfaceTapConfigurationResult(
        etag=__ret__.etag,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        type=__ret__.type,
        virtual_network_tap=__ret__.virtual_network_tap)
