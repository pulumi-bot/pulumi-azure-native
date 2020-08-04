# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetVolumeResult:
    """
    The volume.
    """
    def __init__(__self__, kind=None, name=None, properties=None, type=None):
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        __self__.kind = kind
        """
        The Kind of the object. Currently only Series8000 is supported
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the object.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The properties of the volume.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The hierarchical type of the object.
        """


class AwaitableGetVolumeResult(GetVolumeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeResult(
            kind=self.kind,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_volume(device_name=None, manager_name=None, name=None, resource_group_name=None, volume_container_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str device_name: The device name
    :param str manager_name: The manager name
    :param str name: The volume name.
    :param str resource_group_name: The resource group name
    :param str volume_container_name: The volume container name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['managerName'] = manager_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['volumeContainerName'] = volume_container_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:storsimple/v20170601:getVolume', __args__, opts=opts).value

    return AwaitableGetVolumeResult(
        kind=__ret__.get('kind'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
