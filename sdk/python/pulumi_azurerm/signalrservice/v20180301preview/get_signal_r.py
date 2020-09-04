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
    'GetSignalRResult',
    'AwaitableGetSignalRResult',
    'get_signal_r',
]

@pulumi.output_type
class GetSignalRResult:
    """
    A class represent a SignalR service resource.
    """
    def __init__(__self__, external_ip=None, host_name=None, host_name_prefix=None, location=None, name=None, provisioning_state=None, public_port=None, server_port=None, sku=None, tags=None, type=None, version=None):
        if external_ip and not isinstance(external_ip, str):
            raise TypeError("Expected argument 'external_ip' to be a str")
        pulumi.set(__self__, "external_ip", external_ip)
        if host_name and not isinstance(host_name, str):
            raise TypeError("Expected argument 'host_name' to be a str")
        pulumi.set(__self__, "host_name", host_name)
        if host_name_prefix and not isinstance(host_name_prefix, str):
            raise TypeError("Expected argument 'host_name_prefix' to be a str")
        pulumi.set(__self__, "host_name_prefix", host_name_prefix)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if public_port and not isinstance(public_port, float):
            raise TypeError("Expected argument 'public_port' to be a float")
        pulumi.set(__self__, "public_port", public_port)
        if server_port and not isinstance(server_port, float):
            raise TypeError("Expected argument 'server_port' to be a float")
        pulumi.set(__self__, "server_port", server_port)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="externalIP")
    def external_ip(self) -> str:
        """
        The publicly accessible IP of the SignalR service.
        """
        return pulumi.get(self, "external_ip")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> str:
        """
        FQDN of the SignalR service instance. Format: xxx.service.signalr.net
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter(name="hostNamePrefix")
    def host_name_prefix(self) -> Optional[str]:
        """
        Prefix for the hostName of the SignalR service. Retained for future use.
        The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
        """
        return pulumi.get(self, "host_name_prefix")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
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
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> float:
        """
        The publicly accessibly port of the SignalR service which is designed for browser/client side usage.
        """
        return pulumi.get(self, "public_port")

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> float:
        """
        The publicly accessibly port of the SignalR service which is designed for customer server side usage.
        """
        return pulumi.get(self, "server_port")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ResourceSkuResponse']:
        """
        SKU of the service.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Tags of the service which is a list of key value pairs that describe the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the service - e.g. "Microsoft.SignalRService/SignalR"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
        """
        return pulumi.get(self, "version")


class AwaitableGetSignalRResult(GetSignalRResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSignalRResult(
            external_ip=self.external_ip,
            host_name=self.host_name,
            host_name_prefix=self.host_name_prefix,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            public_port=self.public_port,
            server_port=self.server_port,
            sku=self.sku,
            tags=self.tags,
            type=self.type,
            version=self.version)


def get_signal_r(resource_group_name: Optional[str] = None,
                 resource_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSignalRResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str resource_name: The name of the SignalR resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:signalrservice/v20180301preview:getSignalR', __args__, opts=opts, typ=GetSignalRResult).value

    return AwaitableGetSignalRResult(
        external_ip=__ret__.external_ip,
        host_name=__ret__.host_name,
        host_name_prefix=__ret__.host_name_prefix,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        public_port=__ret__.public_port,
        server_port=__ret__.server_port,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type,
        version=__ret__.version)
