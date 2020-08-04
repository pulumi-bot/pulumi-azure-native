# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetRecordSetResult:
    """
    Describes a DNS record set (a collection of DNS records with the same name and type) in a Private DNS zone.
    """
    def __init__(__self__, etag=None, name=None, properties=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        __self__.etag = etag
        """
        The ETag of the record set.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the resource
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The properties of the record set.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
        """


class AwaitableGetRecordSetResult(GetRecordSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRecordSetResult(
            etag=self.etag,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_record_set(name=None, private_zone_name=None, record_type=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the record set, relative to the name of the zone.
    :param str private_zone_name: The name of the Private DNS zone (without a terminating dot).
    :param str record_type: The type of DNS record in this record set.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['privateZoneName'] = private_zone_name
    __args__['recordType'] = record_type
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20180901:getRecordSet', __args__, opts=opts).value

    return AwaitableGetRecordSetResult(
        etag=__ret__.get('etag'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))