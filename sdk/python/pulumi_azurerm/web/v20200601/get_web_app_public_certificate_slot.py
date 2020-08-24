# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetWebAppPublicCertificateSlotResult',
    'AwaitableGetWebAppPublicCertificateSlotResult',
    'get_web_app_public_certificate_slot',
]

@pulumi.output_type
class GetWebAppPublicCertificateSlotResult:
    """
    Public certificate object
    """
    def __init__(__self__, blob=None, kind=None, name=None, public_certificate_location=None, thumbprint=None, type=None):
        if blob and not isinstance(blob, str):
            raise TypeError("Expected argument 'blob' to be a str")
        pulumi.set(__self__, "blob", blob)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if public_certificate_location and not isinstance(public_certificate_location, str):
            raise TypeError("Expected argument 'public_certificate_location' to be a str")
        pulumi.set(__self__, "public_certificate_location", public_certificate_location)
        if thumbprint and not isinstance(thumbprint, str):
            raise TypeError("Expected argument 'thumbprint' to be a str")
        pulumi.set(__self__, "thumbprint", thumbprint)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def blob(self) -> Optional[str]:
        """
        Public Certificate byte array
        """
        return pulumi.get(self, "blob")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicCertificateLocation")
    def public_certificate_location(self) -> Optional[str]:
        """
        Public Certificate Location
        """
        return pulumi.get(self, "public_certificate_location")

    @property
    @pulumi.getter
    def thumbprint(self) -> str:
        """
        Certificate Thumbprint
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetWebAppPublicCertificateSlotResult(GetWebAppPublicCertificateSlotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebAppPublicCertificateSlotResult(
            blob=self.blob,
            kind=self.kind,
            name=self.name,
            public_certificate_location=self.public_certificate_location,
            thumbprint=self.thumbprint,
            type=self.type)


def get_web_app_public_certificate_slot(name: Optional[str] = None,
                                        resource_group_name: Optional[str] = None,
                                        slot: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebAppPublicCertificateSlotResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Public certificate name.
    :param str resource_group_name: Name of the resource group to which the resource belongs.
    :param str slot: Name of the deployment slot. If a slot is not specified, the API the named binding for the production slot.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['slot'] = slot
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:web/v20200601:getWebAppPublicCertificateSlot', __args__, opts=opts, typ=GetWebAppPublicCertificateSlotResult).value

    return AwaitableGetWebAppPublicCertificateSlotResult(
        blob=__ret__.blob,
        kind=__ret__.kind,
        name=__ret__.name,
        public_certificate_location=__ret__.public_certificate_location,
        thumbprint=__ret__.thumbprint,
        type=__ret__.type)
