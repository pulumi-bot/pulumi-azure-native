# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetConnectionMonitorResult:
    """
    Information about the connection monitor.
    """
    def __init__(__self__, etag=None, location=None, name=None, properties=None, tags=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        __self__.etag = etag
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Connection monitor location.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Name of the connection monitor.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Properties of the connection monitor result.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Connection monitor tags.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Connection monitor type.
        """


class AwaitableGetConnectionMonitorResult(GetConnectionMonitorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectionMonitorResult(
            etag=self.etag,
            location=self.location,
            name=self.name,
            properties=self.properties,
            tags=self.tags,
            type=self.type)


def get_connection_monitor(name=None, network_watcher_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the connection monitor.
    :param str network_watcher_name: The name of the Network Watcher resource.
    :param str resource_group_name: The name of the resource group containing Network Watcher.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['networkWatcherName'] = network_watcher_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20191201:getConnectionMonitor', __args__, opts=opts).value

    return AwaitableGetConnectionMonitorResult(
        etag=__ret__.get('etag'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        tags=__ret__.get('tags'),
        type=__ret__.get('type'))
