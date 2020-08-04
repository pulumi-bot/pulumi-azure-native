# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetCacheResult:
    """
    A Cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
    """
    def __init__(__self__, identity=None, location=None, name=None, properties=None, sku=None, tags=None, type=None):
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        __self__.identity = identity
        """
        The identity of the cache, if configured.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Region name string.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Name of Cache.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Properties of the Cache.
        """
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        __self__.sku = sku
        """
        SKU for the Cache.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        ARM tags as name/value pairs.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Type of the Cache; Microsoft.StorageCache/Cache
        """


class AwaitableGetCacheResult(GetCacheResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCacheResult(
            identity=self.identity,
            location=self.location,
            name=self.name,
            properties=self.properties,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_cache(name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: Name of Cache. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
    :param str resource_group_name: Target resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:storagecache/v20200301:getCache', __args__, opts=opts).value

    return AwaitableGetCacheResult(
        identity=__ret__.get('identity'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        sku=__ret__.get('sku'),
        tags=__ret__.get('tags'),
        type=__ret__.get('type'))