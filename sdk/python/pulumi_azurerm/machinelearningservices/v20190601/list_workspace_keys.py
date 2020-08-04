# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListWorkspaceKeysResult:
    def __init__(__self__, app_insights_instrumentation_key=None, container_registry_credentials=None, user_storage_key=None, user_storage_resource_id=None):
        if app_insights_instrumentation_key and not isinstance(app_insights_instrumentation_key, str):
            raise TypeError("Expected argument 'app_insights_instrumentation_key' to be a str")
        __self__.app_insights_instrumentation_key = app_insights_instrumentation_key
        if container_registry_credentials and not isinstance(container_registry_credentials, dict):
            raise TypeError("Expected argument 'container_registry_credentials' to be a dict")
        __self__.container_registry_credentials = container_registry_credentials
        if user_storage_key and not isinstance(user_storage_key, str):
            raise TypeError("Expected argument 'user_storage_key' to be a str")
        __self__.user_storage_key = user_storage_key
        if user_storage_resource_id and not isinstance(user_storage_resource_id, str):
            raise TypeError("Expected argument 'user_storage_resource_id' to be a str")
        __self__.user_storage_resource_id = user_storage_resource_id


class AwaitableListWorkspaceKeysResult(ListWorkspaceKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListWorkspaceKeysResult(
            app_insights_instrumentation_key=self.app_insights_instrumentation_key,
            container_registry_credentials=self.container_registry_credentials,
            user_storage_key=self.user_storage_key,
            user_storage_resource_id=self.user_storage_resource_id)


def list_workspace_keys(resource_group_name=None, workspace_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: Name of the resource group in which workspace is located.
    :param str workspace_name: Name of Azure Machine Learning workspace.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:machinelearningservices/v20190601:listWorkspaceKeys', __args__, opts=opts).value

    return AwaitableListWorkspaceKeysResult(
        app_insights_instrumentation_key=__ret__.get('appInsightsInstrumentationKey'),
        container_registry_credentials=__ret__.get('containerRegistryCredentials'),
        user_storage_key=__ret__.get('userStorageKey'),
        user_storage_resource_id=__ret__.get('userStorageResourceId'))
