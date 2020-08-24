# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetTagByOperationResult',
    'AwaitableGetTagByOperationResult',
    'get_tag_by_operation',
]

@pulumi.output_type
class GetTagByOperationResult:
    """
    Tag Contract details.
    """
    def __init__(__self__, display_name=None, name=None, type=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Tag name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetTagByOperationResult(GetTagByOperationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagByOperationResult(
            display_name=self.display_name,
            name=self.name,
            type=self.type)


def get_tag_by_operation(api_id: Optional[str] = None,
                         name: Optional[str] = None,
                         operation_id: Optional[str] = None,
                         resource_group_name: Optional[str] = None,
                         service_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTagByOperationResult:
    """
    Use this data source to access information about an existing resource.

    :param str api_id: API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
    :param str name: Tag identifier. Must be unique in the current API Management service instance.
    :param str operation_id: Operation identifier within an API. Must be unique in the current API Management service instance.
    :param str resource_group_name: The name of the resource group.
    :param str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['name'] = name
    __args__['operationId'] = operation_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:apimanagement/v20191201:getTagByOperation', __args__, opts=opts, typ=GetTagByOperationResult).value

    return AwaitableGetTagByOperationResult(
        display_name=__ret__.display_name,
        name=__ret__.name,
        type=__ret__.type)
