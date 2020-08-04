# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetConnectorMappingResult:
    """
    The connector mapping resource format.
    """
    def __init__(__self__, name=None, properties=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource name.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The connector mapping definition.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type.
        """


class AwaitableGetConnectorMappingResult(GetConnectorMappingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectorMappingResult(
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_connector_mapping(connector_name=None, hub_name=None, name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str connector_name: The name of the connector.
    :param str hub_name: The name of the hub.
    :param str name: The name of the connector mapping.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['connectorName'] = connector_name
    __args__['hubName'] = hub_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:customerinsights/v20170101:getConnectorMapping', __args__, opts=opts).value

    return AwaitableGetConnectorMappingResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
