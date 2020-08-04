# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListRegistryCredentialsResult:
    """
    The response from the ListCredentials operation.
    """
    def __init__(__self__, passwords=None, username=None):
        if passwords and not isinstance(passwords, list):
            raise TypeError("Expected argument 'passwords' to be a list")
        __self__.passwords = passwords
        """
        The list of passwords for a container registry.
        """
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        __self__.username = username
        """
        The username for a container registry.
        """


class AwaitableListRegistryCredentialsResult(ListRegistryCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListRegistryCredentialsResult(
            passwords=self.passwords,
            username=self.username)


def list_registry_credentials(registry_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str registry_name: The name of the container registry.
    :param str resource_group_name: The name of the resource group to which the container registry belongs.
    """
    __args__ = dict()
    __args__['registryName'] = registry_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:containerregistry/v20171001:listRegistryCredentials', __args__, opts=opts).value

    return AwaitableListRegistryCredentialsResult(
        passwords=__ret__.get('passwords'),
        username=__ret__.get('username'))
