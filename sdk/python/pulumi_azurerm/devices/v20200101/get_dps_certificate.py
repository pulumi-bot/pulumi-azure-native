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
    'GetDpsCertificateResult',
    'AwaitableGetDpsCertificateResult',
    'get_dps_certificate',
]

@pulumi.output_type
class GetDpsCertificateResult:
    """
    The X509 Certificate.
    """
    def __init__(__self__, etag=None, name=None, properties=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
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
    def etag(self) -> str:
        """
        The entity tag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the certificate.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.CertificatePropertiesResponse':
        """
        properties of a certificate
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetDpsCertificateResult(GetDpsCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDpsCertificateResult(
            etag=self.etag,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_dps_certificate(name: Optional[str] = None,
                        provisioning_service_name: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDpsCertificateResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Name of the certificate to retrieve.
    :param str provisioning_service_name: Name of the provisioning service the certificate is associated with.
    :param str resource_group_name: Resource group identifier.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['provisioningServiceName'] = provisioning_service_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:devices/v20200101:getDpsCertificate', __args__, opts=opts, typ=GetDpsCertificateResult).value

    return AwaitableGetDpsCertificateResult(
        etag=__ret__.etag,
        name=__ret__.name,
        properties=__ret__.properties,
        type=__ret__.type)
