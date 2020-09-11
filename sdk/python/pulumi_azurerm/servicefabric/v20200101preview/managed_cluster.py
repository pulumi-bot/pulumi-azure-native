# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ManagedCluster']


class ManagedCluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_user_name: Optional[pulumi.Input[str]] = None,
                 azure_active_directory: Optional[pulumi.Input[pulumi.InputType['AzureActiveDirectoryArgs']]] = None,
                 client_connection_port: Optional[pulumi.Input[float]] = None,
                 clients: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ClientCertificateArgs']]]]] = None,
                 cluster_code_version: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 cluster_upgrade_description: Optional[pulumi.Input[pulumi.InputType['ClusterUpgradePolicyArgs']]] = None,
                 cluster_upgrade_mode: Optional[pulumi.Input[str]] = None,
                 dns_name: Optional[pulumi.Input[str]] = None,
                 fabric_settings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SettingsSectionDescriptionArgs']]]]] = None,
                 http_gateway_connection_port: Optional[pulumi.Input[float]] = None,
                 load_balancing_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['LoadBalancingRuleArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['SkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The manged cluster resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_password: vm admin user password.
        :param pulumi.Input[str] admin_user_name: vm admin user name.
        :param pulumi.Input[pulumi.InputType['AzureActiveDirectoryArgs']] azure_active_directory: Azure active directory.
        :param pulumi.Input[float] client_connection_port: The port used for client connections to the cluster.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ClientCertificateArgs']]]] clients: client certificates for the cluster.
        :param pulumi.Input[str] cluster_code_version: The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        :param pulumi.Input[str] cluster_name: The name of the cluster resource.
        :param pulumi.Input[pulumi.InputType['ClusterUpgradePolicyArgs']] cluster_upgrade_description: Describes the policy used when upgrading the cluster.
        :param pulumi.Input[str] cluster_upgrade_mode: The upgrade mode of the cluster when new Service Fabric runtime version is available.
               
                 - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
                 - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        :param pulumi.Input[str] dns_name: The cluster dns name.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SettingsSectionDescriptionArgs']]]] fabric_settings: The list of custom fabric settings to configure the cluster.
        :param pulumi.Input[float] http_gateway_connection_port: The port used for http connections to the cluster.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['LoadBalancingRuleArgs']]]] load_balancing_rules: Describes load balancing rules.
        :param pulumi.Input[str] location: Azure resource location.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['SkuArgs']] sku: The sku of the managed cluster
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Azure resource tags.
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

            __props__['admin_password'] = admin_password
            if admin_user_name is None:
                raise TypeError("Missing required property 'admin_user_name'")
            __props__['admin_user_name'] = admin_user_name
            __props__['azure_active_directory'] = azure_active_directory
            __props__['client_connection_port'] = client_connection_port
            __props__['clients'] = clients
            __props__['cluster_code_version'] = cluster_code_version
            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['cluster_upgrade_description'] = cluster_upgrade_description
            __props__['cluster_upgrade_mode'] = cluster_upgrade_mode
            if dns_name is None:
                raise TypeError("Missing required property 'dns_name'")
            __props__['dns_name'] = dns_name
            __props__['fabric_settings'] = fabric_settings
            __props__['http_gateway_connection_port'] = http_gateway_connection_port
            __props__['load_balancing_rules'] = load_balancing_rules
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['cluster_certificate_thumbprint'] = None
            __props__['cluster_id'] = None
            __props__['cluster_state'] = None
            __props__['etag'] = None
            __props__['fqdn'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        super(ManagedCluster, __self__).__init__(
            'azurerm:servicefabric/v20200101preview:ManagedCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ManagedCluster':
        """
        Get an existing ManagedCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ManagedCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> pulumi.Output[Optional[str]]:
        """
        vm admin user password.
        """
        return pulumi.get(self, "admin_password")

    @property
    @pulumi.getter(name="adminUserName")
    def admin_user_name(self) -> pulumi.Output[str]:
        """
        vm admin user name.
        """
        return pulumi.get(self, "admin_user_name")

    @property
    @pulumi.getter(name="azureActiveDirectory")
    def azure_active_directory(self) -> pulumi.Output[Optional['outputs.AzureActiveDirectoryResponse']]:
        """
        Azure active directory.
        """
        return pulumi.get(self, "azure_active_directory")

    @property
    @pulumi.getter(name="clientConnectionPort")
    def client_connection_port(self) -> pulumi.Output[Optional[float]]:
        """
        The port used for client connections to the cluster.
        """
        return pulumi.get(self, "client_connection_port")

    @property
    @pulumi.getter
    def clients(self) -> pulumi.Output[Optional[List['outputs.ClientCertificateResponse']]]:
        """
        client certificates for the cluster.
        """
        return pulumi.get(self, "clients")

    @property
    @pulumi.getter(name="clusterCertificateThumbprint")
    def cluster_certificate_thumbprint(self) -> pulumi.Output[str]:
        """
        The cluster certificate thumbprint used node to node communication.
        """
        return pulumi.get(self, "cluster_certificate_thumbprint")

    @property
    @pulumi.getter(name="clusterCodeVersion")
    def cluster_code_version(self) -> pulumi.Output[Optional[str]]:
        """
        The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        """
        return pulumi.get(self, "cluster_code_version")

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
    @pulumi.getter(name="clusterUpgradeDescription")
    def cluster_upgrade_description(self) -> pulumi.Output[Optional['outputs.ClusterUpgradePolicyResponse']]:
        """
        Describes the policy used when upgrading the cluster.
        """
        return pulumi.get(self, "cluster_upgrade_description")

    @property
    @pulumi.getter(name="clusterUpgradeMode")
    def cluster_upgrade_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The upgrade mode of the cluster when new Service Fabric runtime version is available.

          - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
          - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        """
        return pulumi.get(self, "cluster_upgrade_mode")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> pulumi.Output[str]:
        """
        The cluster dns name.
        """
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Azure resource etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="fabricSettings")
    def fabric_settings(self) -> pulumi.Output[Optional[List['outputs.SettingsSectionDescriptionResponse']]]:
        """
        The list of custom fabric settings to configure the cluster.
        """
        return pulumi.get(self, "fabric_settings")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        the cluster Fully qualified domain name.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="httpGatewayConnectionPort")
    def http_gateway_connection_port(self) -> pulumi.Output[Optional[float]]:
        """
        The port used for http connections to the cluster.
        """
        return pulumi.get(self, "http_gateway_connection_port")

    @property
    @pulumi.getter(name="loadBalancingRules")
    def load_balancing_rules(self) -> pulumi.Output[Optional[List['outputs.LoadBalancingRuleResponse']]]:
        """
        Describes load balancing rules.
        """
        return pulumi.get(self, "load_balancing_rules")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Azure resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the managed cluster resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        The sku of the managed cluster
        """
        return pulumi.get(self, "sku")

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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

