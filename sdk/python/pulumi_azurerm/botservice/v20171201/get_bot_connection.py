# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetBotConnectionResult',
    'AwaitableGetBotConnectionResult',
    'get_bot_connection',
]

@pulumi.output_type
class GetBotConnectionResult:
    """
    Bot channel resource definition
    """
    def __init__(__self__, etag=None, kind=None, location=None, name=None, properties=None, sku=None, tags=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        Entity Tag
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Required. Gets or sets the Kind of the resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Specifies the location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.ConnectionSettingPropertiesResponse':
        """
        The set of properties specific to bot channel resource
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        Gets or sets the SKU of the resource.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Contains resource tags defined as key/value pairs.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetBotConnectionResult(GetBotConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBotConnectionResult(
            etag=self.etag,
            kind=self.kind,
            location=self.location,
            name=self.name,
            properties=self.properties,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_bot_connection(connection_name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       resource_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBotConnectionResult:
    """
    Use this data source to access information about an existing resource.

    :param str connection_name: The name of the Bot Service Connection Setting resource
    :param str resource_group_name: The name of the Bot resource group in the user subscription.
    :param str resource_name: The name of the Bot resource.
    """
    __args__ = dict()
    __args__['connectionName'] = connection_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:botservice/v20171201:getBotConnection', __args__, opts=opts, typ=GetBotConnectionResult).value

    return AwaitableGetBotConnectionResult(
        etag=__ret__.etag,
        kind=__ret__.kind,
        location=__ret__.location,
        name=__ret__.name,
        properties=__ret__.properties,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type)
