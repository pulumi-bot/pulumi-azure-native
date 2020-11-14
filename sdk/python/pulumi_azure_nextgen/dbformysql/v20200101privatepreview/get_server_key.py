# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'GetServerKeyResult',
    'AwaitableGetServerKeyResult',
    'get_server_key',
]

@pulumi.output_type
class GetServerKeyResult:
    """
    A MySQL Server key.
    """
    def __init__(__self__, creation_date=None, kind=None, name=None, server_key_type=None, type=None, uri=None):
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if server_key_type and not isinstance(server_key_type, str):
            raise TypeError("Expected argument 'server_key_type' to be a str")
        pulumi.set(__self__, "server_key_type", server_key_type)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uri and not isinstance(uri, str):
            raise TypeError("Expected argument 'uri' to be a str")
        pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        The key creation date.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Kind of encryption protector. This is metadata used for the Azure portal experience.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serverKeyType")
    def server_key_type(self) -> str:
        """
        The key type like 'AzureKeyVault'.
        """
        return pulumi.get(self, "server_key_type")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uri(self) -> Optional[str]:
        """
        The URI of the key.
        """
        return pulumi.get(self, "uri")


class AwaitableGetServerKeyResult(GetServerKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerKeyResult(
            creation_date=self.creation_date,
            kind=self.kind,
            name=self.name,
            server_key_type=self.server_key_type,
            type=self.type,
            uri=self.uri)


def get_server_key(key_name: Optional[str] = None,
                   resource_group_name: Optional[str] = None,
                   server_name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerKeyResult:
    """
    Use this data source to access information about an existing resource.

    :param str key_name: The name of the MySQL Server key to be retrieved.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['keyName'] = key_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:dbformysql/v20200101privatepreview:getServerKey', __args__, opts=opts, typ=GetServerKeyResult).value

    return AwaitableGetServerKeyResult(
        creation_date=__ret__.creation_date,
        kind=__ret__.kind,
        name=__ret__.name,
        server_key_type=__ret__.server_key_type,
        type=__ret__.type,
        uri=__ret__.uri)
