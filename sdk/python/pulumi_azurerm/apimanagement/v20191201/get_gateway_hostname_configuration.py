# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetGatewayHostnameConfigurationResult',
    'AwaitableGetGatewayHostnameConfigurationResult',
    'get_gateway_hostname_configuration',
]

@pulumi.output_type
class GetGatewayHostnameConfigurationResult:
    """
    Gateway hostname configuration details.
    """
    def __init__(__self__, certificate_id=None, hostname=None, name=None, negotiate_client_certificate=None, type=None):
        if certificate_id and not isinstance(certificate_id, str):
            raise TypeError("Expected argument 'certificate_id' to be a str")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if negotiate_client_certificate and not isinstance(negotiate_client_certificate, bool):
            raise TypeError("Expected argument 'negotiate_client_certificate' to be a bool")
        pulumi.set(__self__, "negotiate_client_certificate", negotiate_client_certificate)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[str]:
        """
        Identifier of Certificate entity that will be used for TLS connection establishment
        """
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter
    def hostname(self) -> Optional[str]:
        """
        Hostname value. Supports valid domain name, partial or full wildcard
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="negotiateClientCertificate")
    def negotiate_client_certificate(self) -> Optional[bool]:
        """
        Determines whether gateway requests client certificate
        """
        return pulumi.get(self, "negotiate_client_certificate")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetGatewayHostnameConfigurationResult(GetGatewayHostnameConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayHostnameConfigurationResult(
            certificate_id=self.certificate_id,
            hostname=self.hostname,
            name=self.name,
            negotiate_client_certificate=self.negotiate_client_certificate,
            type=self.type)


def get_gateway_hostname_configuration(gateway_id: Optional[str] = None,
                                       name: Optional[str] = None,
                                       resource_group_name: Optional[str] = None,
                                       service_name: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayHostnameConfigurationResult:
    """
    Use this data source to access information about an existing resource.

    :param str gateway_id: Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
    :param str name: Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
    :param str resource_group_name: The name of the resource group.
    :param str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['gatewayId'] = gateway_id
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:apimanagement/v20191201:getGatewayHostnameConfiguration', __args__, opts=opts, typ=GetGatewayHostnameConfigurationResult).value

    return AwaitableGetGatewayHostnameConfigurationResult(
        certificate_id=__ret__.certificate_id,
        hostname=__ret__.hostname,
        name=__ret__.name,
        negotiate_client_certificate=__ret__.negotiate_client_certificate,
        type=__ret__.type)
