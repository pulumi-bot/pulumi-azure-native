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
    'GetApplicationResult',
    'AwaitableGetApplicationResult',
    'get_application',
]

@pulumi.output_type
class GetApplicationResult:
    """
    Contains information about an application in a Batch account.
    """
    def __init__(__self__, allow_updates=None, default_version=None, display_name=None, packages=None):
        if allow_updates and not isinstance(allow_updates, bool):
            raise TypeError("Expected argument 'allow_updates' to be a bool")
        pulumi.set(__self__, "allow_updates", allow_updates)
        if default_version and not isinstance(default_version, str):
            raise TypeError("Expected argument 'default_version' to be a str")
        pulumi.set(__self__, "default_version", default_version)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if packages and not isinstance(packages, list):
            raise TypeError("Expected argument 'packages' to be a list")
        pulumi.set(__self__, "packages", packages)

    @property
    @pulumi.getter(name="allowUpdates")
    def allow_updates(self) -> Optional[bool]:
        """
        A value indicating whether packages within the application may be overwritten using the same version string.
        """
        return pulumi.get(self, "allow_updates")

    @property
    @pulumi.getter(name="defaultVersion")
    def default_version(self) -> Optional[str]:
        """
        The package to use if a client requests the application but does not specify a version.
        """
        return pulumi.get(self, "default_version")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The display name for the application.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def packages(self) -> Optional[List['outputs.ApplicationPackageResponse']]:
        """
        The list of packages under this application.
        """
        return pulumi.get(self, "packages")


class AwaitableGetApplicationResult(GetApplicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationResult(
            allow_updates=self.allow_updates,
            default_version=self.default_version,
            display_name=self.display_name,
            packages=self.packages)


def get_application(account_name: Optional[str] = None,
                    application_id: Optional[str] = None,
                    resource_group_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: The name of the Batch account.
    :param str application_id: The ID of the application.
    :param str resource_group_name: The name of the resource group that contains the Batch account.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['applicationId'] = application_id
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:batch/v20170901:getApplication', __args__, opts=opts, typ=GetApplicationResult).value

    return AwaitableGetApplicationResult(
        allow_updates=__ret__.allow_updates,
        default_version=__ret__.default_version,
        display_name=__ret__.display_name,
        packages=__ret__.packages)
