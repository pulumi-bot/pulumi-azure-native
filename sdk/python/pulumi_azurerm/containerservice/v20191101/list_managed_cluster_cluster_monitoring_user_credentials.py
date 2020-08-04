# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListManagedClusterClusterMonitoringUserCredentialsResult:
    """
    The list of credential result response.
    """
    def __init__(__self__, kubeconfigs=None):
        if kubeconfigs and not isinstance(kubeconfigs, list):
            raise TypeError("Expected argument 'kubeconfigs' to be a list")
        __self__.kubeconfigs = kubeconfigs
        """
        Base64-encoded Kubernetes configuration file.
        """


class AwaitableListManagedClusterClusterMonitoringUserCredentialsResult(ListManagedClusterClusterMonitoringUserCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListManagedClusterClusterMonitoringUserCredentialsResult(
            kubeconfigs=self.kubeconfigs)


def list_managed_cluster_cluster_monitoring_user_credentials(resource_group_name=None, resource_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group.
    :param str resource_name: The name of the managed cluster resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:containerservice/v20191101:listManagedClusterClusterMonitoringUserCredentials', __args__, opts=opts).value

    return AwaitableListManagedClusterClusterMonitoringUserCredentialsResult(
        kubeconfigs=__ret__.get('kubeconfigs'))
