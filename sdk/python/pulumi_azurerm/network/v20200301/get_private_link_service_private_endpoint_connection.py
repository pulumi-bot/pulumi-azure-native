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
    'GetPrivateLinkServicePrivateEndpointConnectionResult',
    'AwaitableGetPrivateLinkServicePrivateEndpointConnectionResult',
    'get_private_link_service_private_endpoint_connection',
]

@pulumi.output_type
class GetPrivateLinkServicePrivateEndpointConnectionResult:
    """
    PrivateEndpointConnection resource.
    """
    def __init__(__self__, etag=None, link_identifier=None, name=None, private_endpoint=None, private_link_service_connection_state=None, provisioning_state=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if link_identifier and not isinstance(link_identifier, str):
            raise TypeError("Expected argument 'link_identifier' to be a str")
        pulumi.set(__self__, "link_identifier", link_identifier)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_endpoint and not isinstance(private_endpoint, dict):
            raise TypeError("Expected argument 'private_endpoint' to be a dict")
        pulumi.set(__self__, "private_endpoint", private_endpoint)
        if private_link_service_connection_state and not isinstance(private_link_service_connection_state, dict):
            raise TypeError("Expected argument 'private_link_service_connection_state' to be a dict")
        pulumi.set(__self__, "private_link_service_connection_state", private_link_service_connection_state)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="linkIdentifier")
    def link_identifier(self) -> str:
        """
        The consumer link id.
        """
        return pulumi.get(self, "link_identifier")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> 'outputs.PrivateEndpointResponse':
        """
        The resource of private end point.
        """
        return pulumi.get(self, "private_endpoint")

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> Optional['outputs.PrivateLinkServiceConnectionStateResponse']:
        """
        A collection of information about the state of the connection between service consumer and provider.
        """
        return pulumi.get(self, "private_link_service_connection_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the private endpoint connection resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetPrivateLinkServicePrivateEndpointConnectionResult(GetPrivateLinkServicePrivateEndpointConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateLinkServicePrivateEndpointConnectionResult(
            etag=self.etag,
            link_identifier=self.link_identifier,
            name=self.name,
            private_endpoint=self.private_endpoint,
            private_link_service_connection_state=self.private_link_service_connection_state,
            provisioning_state=self.provisioning_state,
            type=self.type)


def get_private_link_service_private_endpoint_connection(expand: Optional[str] = None,
                                                         name: Optional[str] = None,
                                                         resource_group_name: Optional[str] = None,
                                                         service_name: Optional[str] = None,
                                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateLinkServicePrivateEndpointConnectionResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: Expands referenced resources.
    :param str name: The name of the private end point connection.
    :param str resource_group_name: The name of the resource group.
    :param str service_name: The name of the private link service.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20200301:getPrivateLinkServicePrivateEndpointConnection', __args__, opts=opts, typ=GetPrivateLinkServicePrivateEndpointConnectionResult).value

    return AwaitableGetPrivateLinkServicePrivateEndpointConnectionResult(
        etag=__ret__.etag,
        link_identifier=__ret__.link_identifier,
        name=__ret__.name,
        private_endpoint=__ret__.private_endpoint,
        private_link_service_connection_state=__ret__.private_link_service_connection_state,
        provisioning_state=__ret__.provisioning_state,
        type=__ret__.type)
