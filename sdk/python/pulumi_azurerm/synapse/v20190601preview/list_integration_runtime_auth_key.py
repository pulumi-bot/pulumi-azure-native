# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ListIntegrationRuntimeAuthKeyResult',
    'AwaitableListIntegrationRuntimeAuthKeyResult',
    'list_integration_runtime_auth_key',
]

@pulumi.output_type
class ListIntegrationRuntimeAuthKeyResult:
    """
    The integration runtime authentication keys.
    """
    def __init__(__self__, auth_key1=None, auth_key2=None):
        if auth_key1 and not isinstance(auth_key1, str):
            raise TypeError("Expected argument 'auth_key1' to be a str")
        pulumi.set(__self__, "auth_key1", auth_key1)
        if auth_key2 and not isinstance(auth_key2, str):
            raise TypeError("Expected argument 'auth_key2' to be a str")
        pulumi.set(__self__, "auth_key2", auth_key2)

    @property
    @pulumi.getter(name="authKey1")
    def auth_key1(self) -> Optional[str]:
        """
        The primary integration runtime authentication key.
        """
        return pulumi.get(self, "auth_key1")

    @property
    @pulumi.getter(name="authKey2")
    def auth_key2(self) -> Optional[str]:
        """
        The secondary integration runtime authentication key.
        """
        return pulumi.get(self, "auth_key2")


class AwaitableListIntegrationRuntimeAuthKeyResult(ListIntegrationRuntimeAuthKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListIntegrationRuntimeAuthKeyResult(
            auth_key1=self.auth_key1,
            auth_key2=self.auth_key2)


def list_integration_runtime_auth_key(integration_runtime_name: Optional[str] = None,
                                      resource_group_name: Optional[str] = None,
                                      workspace_name: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListIntegrationRuntimeAuthKeyResult:
    """
    Use this data source to access information about an existing resource.

    :param str integration_runtime_name: Integration runtime name
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str workspace_name: The name of the workspace
    """
    __args__ = dict()
    __args__['integrationRuntimeName'] = integration_runtime_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:synapse/v20190601preview:listIntegrationRuntimeAuthKey', __args__, opts=opts, typ=ListIntegrationRuntimeAuthKeyResult).value

    return AwaitableListIntegrationRuntimeAuthKeyResult(
        auth_key1=__ret__.auth_key1,
        auth_key2=__ret__.auth_key2)
