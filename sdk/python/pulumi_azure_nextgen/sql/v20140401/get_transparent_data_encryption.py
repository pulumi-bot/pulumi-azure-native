# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'GetTransparentDataEncryptionResult',
    'AwaitableGetTransparentDataEncryptionResult',
    'get_transparent_data_encryption',
]

@pulumi.output_type
class GetTransparentDataEncryptionResult:
    """
    Represents a database transparent data encryption configuration.
    """
    def __init__(__self__, location=None, name=None, status=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the database transparent data encryption.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetTransparentDataEncryptionResult(GetTransparentDataEncryptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransparentDataEncryptionResult(
            location=self.location,
            name=self.name,
            status=self.status,
            type=self.type)


def get_transparent_data_encryption(database_name: Optional[str] = None,
                                    resource_group_name: Optional[str] = None,
                                    server_name: Optional[str] = None,
                                    transparent_data_encryption_name: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransparentDataEncryptionResult:
    """
    Use this data source to access information about an existing resource.

    :param str database_name: The name of the database for which the transparent data encryption applies.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str server_name: The name of the server.
    :param str transparent_data_encryption_name: The name of the transparent data encryption configuration.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    __args__['transparentDataEncryptionName'] = transparent_data_encryption_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:sql/v20140401:getTransparentDataEncryption', __args__, opts=opts, typ=GetTransparentDataEncryptionResult).value

    return AwaitableGetTransparentDataEncryptionResult(
        location=__ret__.location,
        name=__ret__.name,
        status=__ret__.status,
        type=__ret__.type)
