# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Cluster']


class Cluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_on_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 application_type_versions_cleanup_policy: Optional[pulumi.Input[pulumi.InputType['ApplicationTypeVersionsCleanupPolicyArgs']]] = None,
                 azure_active_directory: Optional[pulumi.Input[pulumi.InputType['AzureActiveDirectoryArgs']]] = None,
                 certificate: Optional[pulumi.Input[pulumi.InputType['CertificateDescriptionArgs']]] = None,
                 certificate_common_names: Optional[pulumi.Input[pulumi.InputType['ServerCertificateCommonNamesArgs']]] = None,
                 client_certificate_common_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClientCertificateCommonNameArgs']]]]] = None,
                 client_certificate_thumbprints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClientCertificateThumbprintArgs']]]]] = None,
                 cluster_code_version: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 diagnostics_storage_account_config: Optional[pulumi.Input[pulumi.InputType['DiagnosticsStorageAccountConfigArgs']]] = None,
                 event_store_service_enabled: Optional[pulumi.Input[bool]] = None,
                 fabric_settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SettingsSectionDescriptionArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_endpoint: Optional[pulumi.Input[str]] = None,
                 node_types: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NodeTypeDescriptionArgs']]]]] = None,
                 reliability_level: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 reverse_proxy_certificate: Optional[pulumi.Input[pulumi.InputType['CertificateDescriptionArgs']]] = None,
                 reverse_proxy_certificate_common_names: Optional[pulumi.Input[pulumi.InputType['ServerCertificateCommonNamesArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upgrade_description: Optional[pulumi.Input[pulumi.InputType['ClusterUpgradePolicyArgs']]] = None,
                 upgrade_mode: Optional[pulumi.Input[str]] = None,
                 vm_image: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The cluster resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] add_on_features: The list of add-on features to enable in the cluster.
        :param pulumi.Input[pulumi.InputType['ApplicationTypeVersionsCleanupPolicyArgs']] application_type_versions_cleanup_policy: The policy used to clean up unused versions.
        :param pulumi.Input[pulumi.InputType['AzureActiveDirectoryArgs']] azure_active_directory: The AAD authentication settings of the cluster.
        :param pulumi.Input[pulumi.InputType['CertificateDescriptionArgs']] certificate: The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
        :param pulumi.Input[pulumi.InputType['ServerCertificateCommonNamesArgs']] certificate_common_names: Describes a list of server certificates referenced by common name that are used to secure the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClientCertificateCommonNameArgs']]]] client_certificate_common_names: The list of client certificates referenced by common name that are allowed to manage the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClientCertificateThumbprintArgs']]]] client_certificate_thumbprints: The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
        :param pulumi.Input[str] cluster_code_version: The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        :param pulumi.Input[str] cluster_name: The name of the cluster resource.
        :param pulumi.Input[pulumi.InputType['DiagnosticsStorageAccountConfigArgs']] diagnostics_storage_account_config: The storage account information for storing Service Fabric diagnostic logs.
        :param pulumi.Input[bool] event_store_service_enabled: Indicates if the event store service is enabled.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SettingsSectionDescriptionArgs']]]] fabric_settings: The list of custom fabric settings to configure the cluster.
        :param pulumi.Input[str] location: Azure resource location.
        :param pulumi.Input[str] management_endpoint: The http management endpoint of the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NodeTypeDescriptionArgs']]]] node_types: The list of node types in the cluster.
        :param pulumi.Input[str] reliability_level: The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
               
                 - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
                 - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
                 - Silver - Run the System services with a target replica set count of 5.
                 - Gold - Run the System services with a target replica set count of 7.
                 - Platinum - Run the System services with a target replica set count of 9.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['CertificateDescriptionArgs']] reverse_proxy_certificate: The server certificate used by reverse proxy.
        :param pulumi.Input[pulumi.InputType['ServerCertificateCommonNamesArgs']] reverse_proxy_certificate_common_names: Describes a list of server certificates referenced by common name that are used to secure the cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Azure resource tags.
        :param pulumi.Input[pulumi.InputType['ClusterUpgradePolicyArgs']] upgrade_description: The policy to use when upgrading the cluster.
        :param pulumi.Input[str] upgrade_mode: The upgrade mode of the cluster when new Service Fabric runtime version is available.
               
                 - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
                 - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        :param pulumi.Input[str] vm_image: The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['add_on_features'] = add_on_features
            __props__['application_type_versions_cleanup_policy'] = application_type_versions_cleanup_policy
            __props__['azure_active_directory'] = azure_active_directory
            __props__['certificate'] = certificate
            __props__['certificate_common_names'] = certificate_common_names
            __props__['client_certificate_common_names'] = client_certificate_common_names
            __props__['client_certificate_thumbprints'] = client_certificate_thumbprints
            __props__['cluster_code_version'] = cluster_code_version
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['diagnostics_storage_account_config'] = diagnostics_storage_account_config
            __props__['event_store_service_enabled'] = event_store_service_enabled
            __props__['fabric_settings'] = fabric_settings
            __props__['location'] = location
            if management_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'management_endpoint'")
            __props__['management_endpoint'] = management_endpoint
            if node_types is None and not opts.urn:
                raise TypeError("Missing required property 'node_types'")
            __props__['node_types'] = node_types
            __props__['reliability_level'] = reliability_level
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['reverse_proxy_certificate'] = reverse_proxy_certificate
            __props__['reverse_proxy_certificate_common_names'] = reverse_proxy_certificate_common_names
            __props__['tags'] = tags
            __props__['upgrade_description'] = upgrade_description
            __props__['upgrade_mode'] = upgrade_mode
            __props__['vm_image'] = vm_image
            __props__['available_cluster_versions'] = None
            __props__['cluster_endpoint'] = None
            __props__['cluster_id'] = None
            __props__['cluster_state'] = None
            __props__['etag'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:servicefabric/latest:Cluster"), pulumi.Alias(type_="azure-nextgen:servicefabric/v20160901:Cluster"), pulumi.Alias(type_="azure-nextgen:servicefabric/v20170701preview:Cluster"), pulumi.Alias(type_="azure-nextgen:servicefabric/v20180201:Cluster"), pulumi.Alias(type_="azure-nextgen:servicefabric/v20190301:Cluster"), pulumi.Alias(type_="azure-nextgen:servicefabric/v20190301preview:Cluster"), pulumi.Alias(type_="azure-nextgen:servicefabric/v20190601preview:Cluster"), pulumi.Alias(type_="azure-nextgen:servicefabric/v20200301:Cluster")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Cluster, __self__).__init__(
            'azure-nextgen:servicefabric/v20191101preview:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addOnFeatures")
    def add_on_features(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of add-on features to enable in the cluster.
        """
        return pulumi.get(self, "add_on_features")

    @property
    @pulumi.getter(name="applicationTypeVersionsCleanupPolicy")
    def application_type_versions_cleanup_policy(self) -> pulumi.Output[Optional['outputs.ApplicationTypeVersionsCleanupPolicyResponse']]:
        """
        The policy used to clean up unused versions.
        """
        return pulumi.get(self, "application_type_versions_cleanup_policy")

    @property
    @pulumi.getter(name="availableClusterVersions")
    def available_cluster_versions(self) -> pulumi.Output[Sequence['outputs.ClusterVersionDetailsResponse']]:
        """
        The Service Fabric runtime versions available for this cluster.
        """
        return pulumi.get(self, "available_cluster_versions")

    @property
    @pulumi.getter(name="azureActiveDirectory")
    def azure_active_directory(self) -> pulumi.Output[Optional['outputs.AzureActiveDirectoryResponse']]:
        """
        The AAD authentication settings of the cluster.
        """
        return pulumi.get(self, "azure_active_directory")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional['outputs.CertificateDescriptionResponse']]:
        """
        The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="certificateCommonNames")
    def certificate_common_names(self) -> pulumi.Output[Optional['outputs.ServerCertificateCommonNamesResponse']]:
        """
        Describes a list of server certificates referenced by common name that are used to secure the cluster.
        """
        return pulumi.get(self, "certificate_common_names")

    @property
    @pulumi.getter(name="clientCertificateCommonNames")
    def client_certificate_common_names(self) -> pulumi.Output[Optional[Sequence['outputs.ClientCertificateCommonNameResponse']]]:
        """
        The list of client certificates referenced by common name that are allowed to manage the cluster.
        """
        return pulumi.get(self, "client_certificate_common_names")

    @property
    @pulumi.getter(name="clientCertificateThumbprints")
    def client_certificate_thumbprints(self) -> pulumi.Output[Optional[Sequence['outputs.ClientCertificateThumbprintResponse']]]:
        """
        The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
        """
        return pulumi.get(self, "client_certificate_thumbprints")

    @property
    @pulumi.getter(name="clusterCodeVersion")
    def cluster_code_version(self) -> pulumi.Output[Optional[str]]:
        """
        The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        """
        return pulumi.get(self, "cluster_code_version")

    @property
    @pulumi.getter(name="clusterEndpoint")
    def cluster_endpoint(self) -> pulumi.Output[str]:
        """
        The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
        """
        return pulumi.get(self, "cluster_endpoint")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        A service generated unique identifier for the cluster resource.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterState")
    def cluster_state(self) -> pulumi.Output[str]:
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
    def diagnostics_storage_account_config(self) -> pulumi.Output[Optional['outputs.DiagnosticsStorageAccountConfigResponse']]:
        """
        The storage account information for storing Service Fabric diagnostic logs.
        """
        return pulumi.get(self, "diagnostics_storage_account_config")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Azure resource etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="eventStoreServiceEnabled")
    def event_store_service_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the event store service is enabled.
        """
        return pulumi.get(self, "event_store_service_enabled")

    @property
    @pulumi.getter(name="fabricSettings")
    def fabric_settings(self) -> pulumi.Output[Optional[Sequence['outputs.SettingsSectionDescriptionResponse']]]:
        """
        The list of custom fabric settings to configure the cluster.
        """
        return pulumi.get(self, "fabric_settings")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Azure resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementEndpoint")
    def management_endpoint(self) -> pulumi.Output[str]:
        """
        The http management endpoint of the cluster.
        """
        return pulumi.get(self, "management_endpoint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeTypes")
    def node_types(self) -> pulumi.Output[Sequence['outputs.NodeTypeDescriptionResponse']]:
        """
        The list of node types in the cluster.
        """
        return pulumi.get(self, "node_types")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the cluster resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="reliabilityLevel")
    def reliability_level(self) -> pulumi.Output[Optional[str]]:
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
    def reverse_proxy_certificate(self) -> pulumi.Output[Optional['outputs.CertificateDescriptionResponse']]:
        """
        The server certificate used by reverse proxy.
        """
        return pulumi.get(self, "reverse_proxy_certificate")

    @property
    @pulumi.getter(name="reverseProxyCertificateCommonNames")
    def reverse_proxy_certificate_common_names(self) -> pulumi.Output[Optional['outputs.ServerCertificateCommonNamesResponse']]:
        """
        Describes a list of server certificates referenced by common name that are used to secure the cluster.
        """
        return pulumi.get(self, "reverse_proxy_certificate_common_names")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Azure resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="upgradeDescription")
    def upgrade_description(self) -> pulumi.Output[Optional['outputs.ClusterUpgradePolicyResponse']]:
        """
        The policy to use when upgrading the cluster.
        """
        return pulumi.get(self, "upgrade_description")

    @property
    @pulumi.getter(name="upgradeMode")
    def upgrade_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The upgrade mode of the cluster when new Service Fabric runtime version is available.

          - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
          - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        """
        return pulumi.get(self, "upgrade_mode")

    @property
    @pulumi.getter(name="vmImage")
    def vm_image(self) -> pulumi.Output[Optional[str]]:
        """
        The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
        """
        return pulumi.get(self, "vm_image")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

