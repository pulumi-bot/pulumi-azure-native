# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListWebAppHybridConnectionKeysSlotResult:
    """
    Hybrid Connection key contract. This has the send key name and value for a Hybrid Connection.
    """
    def __init__(__self__, kind=None, name=None, properties=None, type=None):
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        __self__.kind = kind
        """
        Kind of resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource Name.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        HybridConnectionKey resource specific properties
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type.
        """


class AwaitableListWebAppHybridConnectionKeysSlotResult(ListWebAppHybridConnectionKeysSlotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListWebAppHybridConnectionKeysSlotResult(
            kind=self.kind,
            name=self.name,
            properties=self.properties,
            type=self.type)


def list_web_app_hybrid_connection_keys_slot(name=None, namespace_name=None, resource_group_name=None, slot=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The relay name for this hybrid connection.
    :param str namespace_name: The namespace for this hybrid connection.
    :param str resource_group_name: Name of the resource group to which the resource belongs.
    :param str slot: The name of the slot for the web app.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['slot'] = slot
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:web/v20181101:listWebAppHybridConnectionKeysSlot', __args__, opts=opts).value

    return AwaitableListWebAppHybridConnectionKeysSlotResult(
        kind=__ret__.get('kind'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
