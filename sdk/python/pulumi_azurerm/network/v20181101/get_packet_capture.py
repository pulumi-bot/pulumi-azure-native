# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetPacketCaptureResult:
    """
    Information about packet capture session.
    """
    def __init__(__self__, etag=None, name=None, properties=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        __self__.etag = etag
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Name of the packet capture session.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Describes the properties of a packet capture session.
        """


class AwaitableGetPacketCaptureResult(GetPacketCaptureResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPacketCaptureResult(
            etag=self.etag,
            name=self.name,
            properties=self.properties)


def get_packet_capture(name=None, network_watcher_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the packet capture session.
    :param str network_watcher_name: The name of the network watcher.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['networkWatcherName'] = network_watcher_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20181101:getPacketCapture', __args__, opts=opts).value

    return AwaitableGetPacketCaptureResult(
        etag=__ret__.get('etag'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'))
