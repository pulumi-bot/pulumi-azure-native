# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ListNotebookWorkspaceConnectionInfoResult',
    'AwaitableListNotebookWorkspaceConnectionInfoResult',
    'list_notebook_workspace_connection_info',
]

@pulumi.output_type
class ListNotebookWorkspaceConnectionInfoResult:
    """
    The connection info for the given notebook workspace
    """
    def __init__(__self__, auth_token=None, notebook_server_endpoint=None):
        if auth_token and not isinstance(auth_token, str):
            raise TypeError("Expected argument 'auth_token' to be a str")
        pulumi.set(__self__, "auth_token", auth_token)
        if notebook_server_endpoint and not isinstance(notebook_server_endpoint, str):
            raise TypeError("Expected argument 'notebook_server_endpoint' to be a str")
        pulumi.set(__self__, "notebook_server_endpoint", notebook_server_endpoint)

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> str:
        """
        Specifies auth token used for connecting to Notebook server (uses token-based auth).
        """
        return pulumi.get(self, "auth_token")

    @property
    @pulumi.getter(name="notebookServerEndpoint")
    def notebook_server_endpoint(self) -> str:
        """
        Specifies the endpoint of Notebook server.
        """
        return pulumi.get(self, "notebook_server_endpoint")


class AwaitableListNotebookWorkspaceConnectionInfoResult(ListNotebookWorkspaceConnectionInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListNotebookWorkspaceConnectionInfoResult(
            auth_token=self.auth_token,
            notebook_server_endpoint=self.notebook_server_endpoint)


def list_notebook_workspace_connection_info(account_name: Optional[str] = None,
                                            notebook_workspace_name: Optional[str] = None,
                                            resource_group_name: Optional[str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListNotebookWorkspaceConnectionInfoResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: Cosmos DB database account name.
    :param str notebook_workspace_name: The name of the notebook workspace resource.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['notebookWorkspaceName'] = notebook_workspace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:documentdb/v20191212:listNotebookWorkspaceConnectionInfo', __args__, opts=opts, typ=ListNotebookWorkspaceConnectionInfoResult).value

    return AwaitableListNotebookWorkspaceConnectionInfoResult(
        auth_token=__ret__.auth_token,
        notebook_server_endpoint=__ret__.notebook_server_endpoint)
