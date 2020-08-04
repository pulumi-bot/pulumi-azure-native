# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetPrivateEndpointConnectionResult:
    """
    The Private Endpoint Connection resource.
    """
    def __init__(__self__, identity=None, location=None, name=None, properties=None, sku=None, tags=None, type=None):
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        __self__.identity = identity
        """
        The identity of the resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Specifies the location of the resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Specifies the name of the resource.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Resource properties.
        """
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        __self__.sku = sku
        """
        The sku of the workspace.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Contains resource tags defined as key/value pairs.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Specifies the type of the resource.
        """


class AwaitableGetPrivateEndpointConnectionResult(GetPrivateEndpointConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateEndpointConnectionResult(
            identity=self.identity,
            location=self.location,
            name=self.name,
            properties=self.properties,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_private_endpoint_connection(name=None, resource_group_name=None, workspace_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the private endpoint connection associated with the workspace
    :param str resource_group_name: Name of the resource group in which workspace is located.
    :param str workspace_name: Name of Azure Machine Learning workspace.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:machinelearningservices/v20200101:getPrivateEndpointConnection', __args__, opts=opts).value

    return AwaitableGetPrivateEndpointConnectionResult(
        identity=__ret__.get('identity'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        sku=__ret__.get('sku'),
        tags=__ret__.get('tags'),
        type=__ret__.get('type'))
