# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetP2sVpnServerConfigurationResult:
    """
    P2SVpnServerConfiguration Resource.
    """
    def __init__(__self__, etag=None, name=None, properties=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        __self__.etag = etag
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Properties of the P2SVpnServer configuration.
        """


class AwaitableGetP2sVpnServerConfigurationResult(GetP2sVpnServerConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetP2sVpnServerConfigurationResult(
            etag=self.etag,
            name=self.name,
            properties=self.properties)


def get_p2s_vpn_server_configuration(name=None, resource_group_name=None, virtual_wan_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the P2SVpnServerConfiguration.
    :param str resource_group_name: The resource group name of the P2SVpnServerConfiguration.
    :param str virtual_wan_name: The name of the VirtualWan.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['virtualWanName'] = virtual_wan_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20190701:getP2sVpnServerConfiguration', __args__, opts=opts).value

    return AwaitableGetP2sVpnServerConfigurationResult(
        etag=__ret__.get('etag'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'))
