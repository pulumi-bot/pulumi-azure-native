# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetSavedSearchResult:
    """
    Value object for saved search results.
    """
    def __init__(__self__, e_tag=None, name=None, properties=None, type=None):
        if e_tag and not isinstance(e_tag, str):
            raise TypeError("Expected argument 'e_tag' to be a str")
        __self__.e_tag = e_tag
        """
        The ETag of the saved search.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the saved search.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The properties of the saved search.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of the saved search.
        """


class AwaitableGetSavedSearchResult(GetSavedSearchResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSavedSearchResult(
            e_tag=self.e_tag,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_saved_search(name=None, resource_group_name=None, workspace_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The id of the saved search.
    :param str resource_group_name: The Resource Group name.
    :param str workspace_name: The Log Analytics Workspace name.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:operationalinsights/v20150320:getSavedSearch', __args__, opts=opts).value

    return AwaitableGetSavedSearchResult(
        e_tag=__ret__.get('eTag'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))