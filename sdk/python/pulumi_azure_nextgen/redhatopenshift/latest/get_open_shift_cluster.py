# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetOpenShiftClusterResult',
    'AwaitableGetOpenShiftClusterResult',
    'get_open_shift_cluster',
]

@pulumi.output_type
class GetOpenShiftClusterResult:
    """
    OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
    """
    def __init__(__self__, apiserver_profile=None, cluster_profile=None, console_profile=None, ingress_profiles=None, location=None, master_profile=None, name=None, network_profile=None, provisioning_state=None, service_principal_profile=None, tags=None, type=None, worker_profiles=None):
        if apiserver_profile and not isinstance(apiserver_profile, dict):
            raise TypeError("Expected argument 'apiserver_profile' to be a dict")
        pulumi.set(__self__, "apiserver_profile", apiserver_profile)
        if cluster_profile and not isinstance(cluster_profile, dict):
            raise TypeError("Expected argument 'cluster_profile' to be a dict")
        pulumi.set(__self__, "cluster_profile", cluster_profile)
        if console_profile and not isinstance(console_profile, dict):
            raise TypeError("Expected argument 'console_profile' to be a dict")
        pulumi.set(__self__, "console_profile", console_profile)
        if ingress_profiles and not isinstance(ingress_profiles, list):
            raise TypeError("Expected argument 'ingress_profiles' to be a list")
        pulumi.set(__self__, "ingress_profiles", ingress_profiles)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if master_profile and not isinstance(master_profile, dict):
            raise TypeError("Expected argument 'master_profile' to be a dict")
        pulumi.set(__self__, "master_profile", master_profile)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_profile and not isinstance(network_profile, dict):
            raise TypeError("Expected argument 'network_profile' to be a dict")
        pulumi.set(__self__, "network_profile", network_profile)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if service_principal_profile and not isinstance(service_principal_profile, dict):
            raise TypeError("Expected argument 'service_principal_profile' to be a dict")
        pulumi.set(__self__, "service_principal_profile", service_principal_profile)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if worker_profiles and not isinstance(worker_profiles, list):
            raise TypeError("Expected argument 'worker_profiles' to be a list")
        pulumi.set(__self__, "worker_profiles", worker_profiles)

    @property
    @pulumi.getter(name="apiserverProfile")
    def apiserver_profile(self) -> Optional['outputs.APIServerProfileResponse']:
        """
        The cluster API server profile.
        """
        return pulumi.get(self, "apiserver_profile")

    @property
    @pulumi.getter(name="clusterProfile")
    def cluster_profile(self) -> Optional['outputs.ClusterProfileResponse']:
        """
        The cluster profile.
        """
        return pulumi.get(self, "cluster_profile")

    @property
    @pulumi.getter(name="consoleProfile")
    def console_profile(self) -> Optional['outputs.ConsoleProfileResponse']:
        """
        The console profile.
        """
        return pulumi.get(self, "console_profile")

    @property
    @pulumi.getter(name="ingressProfiles")
    def ingress_profiles(self) -> Optional[Sequence['outputs.IngressProfileResponse']]:
        """
        The cluster ingress profiles.
        """
        return pulumi.get(self, "ingress_profiles")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="masterProfile")
    def master_profile(self) -> Optional['outputs.MasterProfileResponse']:
        """
        The cluster master profile.
        """
        return pulumi.get(self, "master_profile")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkProfile")
    def network_profile(self) -> Optional['outputs.NetworkProfileResponse']:
        """
        The cluster network profile.
        """
        return pulumi.get(self, "network_profile")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        The cluster provisioning state (immutable).
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="servicePrincipalProfile")
    def service_principal_profile(self) -> Optional['outputs.ServicePrincipalProfileResponse']:
        """
        The cluster service principal profile.
        """
        return pulumi.get(self, "service_principal_profile")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="workerProfiles")
    def worker_profiles(self) -> Optional[Sequence['outputs.WorkerProfileResponse']]:
        """
        The cluster worker profiles.
        """
        return pulumi.get(self, "worker_profiles")


class AwaitableGetOpenShiftClusterResult(GetOpenShiftClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOpenShiftClusterResult(
            apiserver_profile=self.apiserver_profile,
            cluster_profile=self.cluster_profile,
            console_profile=self.console_profile,
            ingress_profiles=self.ingress_profiles,
            location=self.location,
            master_profile=self.master_profile,
            name=self.name,
            network_profile=self.network_profile,
            provisioning_state=self.provisioning_state,
            service_principal_profile=self.service_principal_profile,
            tags=self.tags,
            type=self.type,
            worker_profiles=self.worker_profiles)


def get_open_shift_cluster(resource_group_name: Optional[str] = None,
                           resource_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOpenShiftClusterResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str resource_name: The name of the OpenShift cluster resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:redhatopenshift/latest:getOpenShiftCluster', __args__, opts=opts, typ=GetOpenShiftClusterResult).value

    return AwaitableGetOpenShiftClusterResult(
        apiserver_profile=__ret__.apiserver_profile,
        cluster_profile=__ret__.cluster_profile,
        console_profile=__ret__.console_profile,
        ingress_profiles=__ret__.ingress_profiles,
        location=__ret__.location,
        master_profile=__ret__.master_profile,
        name=__ret__.name,
        network_profile=__ret__.network_profile,
        provisioning_state=__ret__.provisioning_state,
        service_principal_profile=__ret__.service_principal_profile,
        tags=__ret__.tags,
        type=__ret__.type,
        worker_profiles=__ret__.worker_profiles)
