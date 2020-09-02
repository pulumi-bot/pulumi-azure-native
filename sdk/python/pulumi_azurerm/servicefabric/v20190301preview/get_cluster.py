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
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
]

@pulumi.output_type
class GetClusterResult:
    """
    The cluster resource
    """
    def __init__(__self__, add_on_features=None, available_cluster_versions=None, azure_active_directory=None, certificate=None, certificate_common_names=None, client_certificate_common_names=None, client_certificate_thumbprints=None, cluster_code_version=None, cluster_endpoint=None, cluster_id=None, cluster_state=None, diagnostics_storage_account_config=None, etag=None, event_store_service_enabled=None, fabric_settings=None, location=None, management_endpoint=None, name=None, node_types=None, provisioning_state=None, reliability_level=None, reverse_proxy_certificate=None, reverse_proxy_certificate_common_names=None, tags=None, type=None, upgrade_description=None, upgrade_mode=None, vm_image=None):
        if add_on_features and not isinstance(add_on_features, list):
            raise TypeError("Expected argument 'add_on_features' to be a list")
        pulumi.set(__self__, "add_on_features", add_on_features)
        if available_cluster_versions and not isinstance(available_cluster_versions, list):
            raise TypeError("Expected argument 'available_cluster_versions' to be a list")
        pulumi.set(__self__, "available_cluster_versions", available_cluster_versions)
        if azure_active_directory and not isinstance(azure_active_directory, dict):
            raise TypeError("Expected argument 'azure_active_directory' to be a dict")
        pulumi.set(__self__, "azure_active_directory", azure_active_directory)
        if certificate and not isinstance(certificate, dict):
            raise TypeError("Expected argument 'certificate' to be a dict")
        pulumi.set(__self__, "certificate", certificate)
        if certificate_common_names and not isinstance(certificate_common_names, dict):
            raise TypeError("Expected argument 'certificate_common_names' to be a dict")
        pulumi.set(__self__, "certificate_common_names", certificate_common_names)
        if client_certificate_common_names and not isinstance(client_certificate_common_names, list):
            raise TypeError("Expected argument 'client_certificate_common_names' to be a list")
        pulumi.set(__self__, "client_certificate_common_names", client_certificate_common_names)
        if client_certificate_thumbprints and not isinstance(client_certificate_thumbprints, list):
            raise TypeError("Expected argument 'client_certificate_thumbprints' to be a list")
        pulumi.set(__self__, "client_certificate_thumbprints", client_certificate_thumbprints)
        if cluster_code_version and not isinstance(cluster_code_version, str):
            raise TypeError("Expected argument 'cluster_code_version' to be a str")
        pulumi.set(__self__, "cluster_code_version", cluster_code_version)
        if cluster_endpoint and not isinstance(cluster_endpoint, str):
            raise TypeError("Expected argument 'cluster_endpoint' to be a str")
        pulumi.set(__self__, "cluster_endpoint", cluster_endpoint)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_state and not isinstance(cluster_state, str):
            raise TypeError("Expected argument 'cluster_state' to be a str")
        pulumi.set(__self__, "cluster_state", cluster_state)
        if diagnostics_storage_account_config and not isinstance(diagnostics_storage_account_config, dict):
            raise TypeError("Expected argument 'diagnostics_storage_account_config' to be a dict")
        pulumi.set(__self__, "diagnostics_storage_account_config", diagnostics_storage_account_config)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if event_store_service_enabled and not isinstance(event_store_service_enabled, bool):
            raise TypeError("Expected argument 'event_store_service_enabled' to be a bool")
        pulumi.set(__self__, "event_store_service_enabled", event_store_service_enabled)
        if fabric_settings and not isinstance(fabric_settings, list):
            raise TypeError("Expected argument 'fabric_settings' to be a list")
        pulumi.set(__self__, "fabric_settings", fabric_settings)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if management_endpoint and not isinstance(management_endpoint, str):
            raise TypeError("Expected argument 'management_endpoint' to be a str")
        pulumi.set(__self__, "management_endpoint", management_endpoint)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_types and not isinstance(node_types, list):
            raise TypeError("Expected argument 'node_types' to be a list")
        pulumi.set(__self__, "node_types", node_types)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if reliability_level and not isinstance(reliability_level, str):
            raise TypeError("Expected argument 'reliability_level' to be a str")
        pulumi.set(__self__, "reliability_level", reliability_level)
        if reverse_proxy_certificate and not isinstance(reverse_proxy_certificate, dict):
            raise TypeError("Expected argument 'reverse_proxy_certificate' to be a dict")
        pulumi.set(__self__, "reverse_proxy_certificate", reverse_proxy_certificate)
        if reverse_proxy_certificate_common_names and not isinstance(reverse_proxy_certificate_common_names, dict):
            raise TypeError("Expected argument 'reverse_proxy_certificate_common_names' to be a dict")
        pulumi.set(__self__, "reverse_proxy_certificate_common_names", reverse_proxy_certificate_common_names)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if upgrade_description and not isinstance(upgrade_description, dict):
            raise TypeError("Expected argument 'upgrade_description' to be a dict")
        pulumi.set(__self__, "upgrade_description", upgrade_description)
        if upgrade_mode and not isinstance(upgrade_mode, str):
            raise TypeError("Expected argument 'upgrade_mode' to be a str")
        pulumi.set(__self__, "upgrade_mode", upgrade_mode)
        if vm_image and not isinstance(vm_image, str):
            raise TypeError("Expected argument 'vm_image' to be a str")
        pulumi.set(__self__, "vm_image", vm_image)

    @property
    @pulumi.getter(name="addOnFeatures")
    def add_on_features(self) -> Optional[List[str]]:
        """
        The list of add-on features to enable in the cluster.
        """
        return pulumi.get(self, "add_on_features")

    @property
    @pulumi.getter(name="availableClusterVersions")
    def available_cluster_versions(self) -> List['outputs.ClusterVersionDetailsResponse']:
        """
        The Service Fabric runtime versions available for this cluster.
        """
        return pulumi.get(self, "available_cluster_versions")

    @property
    @pulumi.getter(name="azureActiveDirectory")
    def azure_active_directory(self) -> Optional['outputs.AzureActiveDirectoryResponse']:
        """
        The AAD authentication settings of the cluster.
        """
        return pulumi.get(self, "azure_active_directory")

    @property
    @pulumi.getter
    def certificate(self) -> Optional['outputs.CertificateDescriptionResponse']:
        """
        The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="certificateCommonNames")
    def certificate_common_names(self) -> Optional['outputs.ServerCertificateCommonNamesResponse']:
        """
        Describes a list of server certificates referenced by common name that are used to secure the cluster.
        """
        return pulumi.get(self, "certificate_common_names")

    @property
    @pulumi.getter(name="clientCertificateCommonNames")
    def client_certificate_common_names(self) -> Optional[List['outputs.ClientCertificateCommonNameResponse']]:
        """
        The list of client certificates referenced by common name that are allowed to manage the cluster.
        """
        return pulumi.get(self, "client_certificate_common_names")

    @property
    @pulumi.getter(name="clientCertificateThumbprints")
    def client_certificate_thumbprints(self) -> Optional[List['outputs.ClientCertificateThumbprintResponse']]:
        """
        The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
        """
        return pulumi.get(self, "client_certificate_thumbprints")

    @property
    @pulumi.getter(name="clusterCodeVersion")
    def cluster_code_version(self) -> Optional[str]:
        """
        The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        """
        return pulumi.get(self, "cluster_code_version")

    @property
    @pulumi.getter(name="clusterEndpoint")
    def cluster_endpoint(self) -> str:
        """
        The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
        """
        return pulumi.get(self, "cluster_endpoint")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        A service generated unique identifier for the cluster resource.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterState")
    def cluster_state(self) -> str:
        """
        The current state of the cluster.

          - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
          - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
          - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
          - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
          - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
          - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
          - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
          - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
          - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
          - Ready - Indicates that the cluster is in a stable state.
        """
        return pulumi.get(self, "cluster_state")

    @property
    @pulumi.getter(name="diagnosticsStorageAccountConfig")
    def diagnostics_storage_account_config(self) -> Optional['outputs.DiagnosticsStorageAccountConfigResponse']:
        """
        The storage account information for storing Service Fabric diagnostic logs.
        """
        return pulumi.get(self, "diagnostics_storage_account_config")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Azure resource etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="eventStoreServiceEnabled")
    def event_store_service_enabled(self) -> Optional[bool]:
        """
        Indicates if the event store service is enabled.
        """
        return pulumi.get(self, "event_store_service_enabled")

    @property
    @pulumi.getter(name="fabricSettings")
    def fabric_settings(self) -> Optional[List['outputs.SettingsSectionDescriptionResponse']]:
        """
        The list of custom fabric settings to configure the cluster.
        """
        return pulumi.get(self, "fabric_settings")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Azure resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementEndpoint")
    def management_endpoint(self) -> str:
        """
        The http management endpoint of the cluster.
        """
        return pulumi.get(self, "management_endpoint")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Azure resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeTypes")
    def node_types(self) -> List['outputs.NodeTypeDescriptionResponse']:
        """
        The list of node types in the cluster.
        """
        return pulumi.get(self, "node_types")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the cluster resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="reliabilityLevel")
    def reliability_level(self) -> Optional[str]:
        """
        The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).

          - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
          - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
          - Silver - Run the System services with a target replica set count of 5.
          - Gold - Run the System services with a target replica set count of 7.
          - Platinum - Run the System services with a target replica set count of 9.
        """
        return pulumi.get(self, "reliability_level")

    @property
    @pulumi.getter(name="reverseProxyCertificate")
    def reverse_proxy_certificate(self) -> Optional['outputs.CertificateDescriptionResponse']:
        """
        The server certificate used by reverse proxy.
        """
        return pulumi.get(self, "reverse_proxy_certificate")

    @property
    @pulumi.getter(name="reverseProxyCertificateCommonNames")
    def reverse_proxy_certificate_common_names(self) -> Optional['outputs.ServerCertificateCommonNamesResponse']:
        """
        Describes a list of server certificates referenced by common name that are used to secure the cluster.
        """
        return pulumi.get(self, "reverse_proxy_certificate_common_names")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Azure resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Azure resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="upgradeDescription")
    def upgrade_description(self) -> Optional['outputs.ClusterUpgradePolicyResponse']:
        """
        The policy to use when upgrading the cluster.
        """
        return pulumi.get(self, "upgrade_description")

    @property
    @pulumi.getter(name="upgradeMode")
    def upgrade_mode(self) -> Optional[str]:
        """
        The upgrade mode of the cluster when new Service Fabric runtime version is available.

          - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
          - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        """
        return pulumi.get(self, "upgrade_mode")

    @property
    @pulumi.getter(name="vmImage")
    def vm_image(self) -> Optional[str]:
        """
        The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
        """
        return pulumi.get(self, "vm_image")


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            add_on_features=self.add_on_features,
            available_cluster_versions=self.available_cluster_versions,
            azure_active_directory=self.azure_active_directory,
            certificate=self.certificate,
            certificate_common_names=self.certificate_common_names,
            client_certificate_common_names=self.client_certificate_common_names,
            client_certificate_thumbprints=self.client_certificate_thumbprints,
            cluster_code_version=self.cluster_code_version,
            cluster_endpoint=self.cluster_endpoint,
            cluster_id=self.cluster_id,
            cluster_state=self.cluster_state,
            diagnostics_storage_account_config=self.diagnostics_storage_account_config,
            etag=self.etag,
            event_store_service_enabled=self.event_store_service_enabled,
            fabric_settings=self.fabric_settings,
            location=self.location,
            management_endpoint=self.management_endpoint,
            name=self.name,
            node_types=self.node_types,
            provisioning_state=self.provisioning_state,
            reliability_level=self.reliability_level,
            reverse_proxy_certificate=self.reverse_proxy_certificate,
            reverse_proxy_certificate_common_names=self.reverse_proxy_certificate_common_names,
            tags=self.tags,
            type=self.type,
            upgrade_description=self.upgrade_description,
            upgrade_mode=self.upgrade_mode,
            vm_image=self.vm_image)


def get_cluster(cluster_name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_name: The name of the cluster resource.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:servicefabric/v20190301preview:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        add_on_features=__ret__.add_on_features,
        available_cluster_versions=__ret__.available_cluster_versions,
        azure_active_directory=__ret__.azure_active_directory,
        certificate=__ret__.certificate,
        certificate_common_names=__ret__.certificate_common_names,
        client_certificate_common_names=__ret__.client_certificate_common_names,
        client_certificate_thumbprints=__ret__.client_certificate_thumbprints,
        cluster_code_version=__ret__.cluster_code_version,
        cluster_endpoint=__ret__.cluster_endpoint,
        cluster_id=__ret__.cluster_id,
        cluster_state=__ret__.cluster_state,
        diagnostics_storage_account_config=__ret__.diagnostics_storage_account_config,
        etag=__ret__.etag,
        event_store_service_enabled=__ret__.event_store_service_enabled,
        fabric_settings=__ret__.fabric_settings,
        location=__ret__.location,
        management_endpoint=__ret__.management_endpoint,
        name=__ret__.name,
        node_types=__ret__.node_types,
        provisioning_state=__ret__.provisioning_state,
        reliability_level=__ret__.reliability_level,
        reverse_proxy_certificate=__ret__.reverse_proxy_certificate,
        reverse_proxy_certificate_common_names=__ret__.reverse_proxy_certificate_common_names,
        tags=__ret__.tags,
        type=__ret__.type,
        upgrade_description=__ret__.upgrade_description,
        upgrade_mode=__ret__.upgrade_mode,
        vm_image=__ret__.vm_image)
