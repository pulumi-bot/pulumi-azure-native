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
    'GetPrivateEndpointConnectionResult',
    'AwaitableGetPrivateEndpointConnectionResult',
    'get_private_endpoint_connection',
]

@pulumi.output_type
class GetPrivateEndpointConnectionResult:
    """
    Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
    """
    def __init__(__self__, name=None, properties=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.PrivateEndpointConnectionPropertiesResponse':
        """
        Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")


class AwaitableGetPrivateEndpointConnectionResult(GetPrivateEndpointConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateEndpointConnectionResult(
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_private_endpoint_connection(private_endpoint_connection_name: Optional[str] = None,
                                    resource_group_name: Optional[str] = None,
                                    search_service_name: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateEndpointConnectionResult:
    """
    Use this data source to access information about an existing resource.

    :param str private_endpoint_connection_name: The name of the private endpoint connection to the Azure Cognitive Search service with the specified resource group.
    :param str resource_group_name: The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str search_service_name: The name of the Azure Cognitive Search service associated with the specified resource group.
    """
    __args__ = dict()
    __args__['privateEndpointConnectionName'] = private_endpoint_connection_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['searchServiceName'] = search_service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:search/v20200801:getPrivateEndpointConnection', __args__, opts=opts, typ=GetPrivateEndpointConnectionResult).value

    return AwaitableGetPrivateEndpointConnectionResult(
        name=__ret__.name,
        properties=__ret__.properties,
        type=__ret__.type)
