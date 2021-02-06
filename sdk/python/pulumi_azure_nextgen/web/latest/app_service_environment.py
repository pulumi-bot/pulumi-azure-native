# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['AppServiceEnvironment']


class AppServiceEnvironment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_management_account_id: Optional[pulumi.Input[str]] = None,
                 cluster_settings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NameValuePairArgs']]]]] = None,
                 dns_suffix: Optional[pulumi.Input[str]] = None,
                 dynamic_cache_enabled: Optional[pulumi.Input[bool]] = None,
                 front_end_scale_factor: Optional[pulumi.Input[int]] = None,
                 has_linux_workers: Optional[pulumi.Input[bool]] = None,
                 internal_load_balancing_mode: Optional[pulumi.Input[Union[str, 'LoadBalancingMode']]] = None,
                 ipssl_address_count: Optional[pulumi.Input[int]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 multi_role_count: Optional[pulumi.Input[int]] = None,
                 multi_size: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_access_control_list: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAccessControlEntryArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 ssl_cert_key_vault_id: Optional[pulumi.Input[str]] = None,
                 ssl_cert_key_vault_secret_name: Optional[pulumi.Input[str]] = None,
                 suspended: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_whitelisted_ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_network: Optional[pulumi.Input[pulumi.InputType['VirtualNetworkProfileArgs']]] = None,
                 vnet_name: Optional[pulumi.Input[str]] = None,
                 vnet_resource_group_name: Optional[pulumi.Input[str]] = None,
                 vnet_subnet_name: Optional[pulumi.Input[str]] = None,
                 worker_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkerPoolArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        App Service Environment ARM resource.
        Latest API Version: 2020-09-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_account_id: API Management Account associated with the App Service Environment.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NameValuePairArgs']]]] cluster_settings: Custom settings for changing the behavior of the App Service Environment.
        :param pulumi.Input[str] dns_suffix: DNS suffix of the App Service Environment.
        :param pulumi.Input[bool] dynamic_cache_enabled: True/false indicating whether the App Service Environment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
               (most likely because NSG blocked the incoming traffic).
        :param pulumi.Input[int] front_end_scale_factor: Scale factor for front-ends.
        :param pulumi.Input[bool] has_linux_workers: Flag that displays whether an ASE has linux workers or not
        :param pulumi.Input[Union[str, 'LoadBalancingMode']] internal_load_balancing_mode: Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
        :param pulumi.Input[int] ipssl_address_count: Number of IP SSL addresses reserved for the App Service Environment.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] location: Resource Location.
        :param pulumi.Input[int] multi_role_count: Number of front-end instances.
        :param pulumi.Input[str] multi_size: Front-end VM size, e.g. "Medium", "Large".
        :param pulumi.Input[str] name: Name of the App Service Environment.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkAccessControlEntryArgs']]]] network_access_control_list: Access control list for controlling traffic to the App Service Environment.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[str] ssl_cert_key_vault_id: Key Vault ID for ILB App Service Environment default SSL certificate
        :param pulumi.Input[str] ssl_cert_key_vault_secret_name: Key Vault Secret Name for ILB App Service Environment default SSL certificate
        :param pulumi.Input[bool] suspended: <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
                (most likely because NSG blocked the incoming traffic).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_whitelisted_ip_ranges: User added ip ranges to whitelist on ASE db
        :param pulumi.Input[pulumi.InputType['VirtualNetworkProfileArgs']] virtual_network: Description of the Virtual Network.
        :param pulumi.Input[str] vnet_name: Name of the Virtual Network for the App Service Environment.
        :param pulumi.Input[str] vnet_resource_group_name: Resource group of the Virtual Network.
        :param pulumi.Input[str] vnet_subnet_name: Subnet of the Virtual Network.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkerPoolArgs']]]] worker_pools: Description of worker pools with worker size IDs, VM sizes, and number of workers in each pool.
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

            __props__['api_management_account_id'] = api_management_account_id
            __props__['cluster_settings'] = cluster_settings
            __props__['dns_suffix'] = dns_suffix
            __props__['dynamic_cache_enabled'] = dynamic_cache_enabled
            __props__['front_end_scale_factor'] = front_end_scale_factor
            __props__['has_linux_workers'] = has_linux_workers
            __props__['internal_load_balancing_mode'] = internal_load_balancing_mode
            __props__['ipssl_address_count'] = ipssl_address_count
            __props__['kind'] = kind
            __props__['location'] = location
            __props__['multi_role_count'] = multi_role_count
            __props__['multi_size'] = multi_size
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_access_control_list'] = network_access_control_list
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['ssl_cert_key_vault_id'] = ssl_cert_key_vault_id
            __props__['ssl_cert_key_vault_secret_name'] = ssl_cert_key_vault_secret_name
            __props__['suspended'] = suspended
            __props__['tags'] = tags
            __props__['user_whitelisted_ip_ranges'] = user_whitelisted_ip_ranges
            if virtual_network is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_network'")
            __props__['virtual_network'] = virtual_network
            __props__['vnet_name'] = vnet_name
            __props__['vnet_resource_group_name'] = vnet_resource_group_name
            __props__['vnet_subnet_name'] = vnet_subnet_name
            if worker_pools is None and not opts.urn:
                raise TypeError("Missing required property 'worker_pools'")
            __props__['worker_pools'] = worker_pools
            __props__['allowed_multi_sizes'] = None
            __props__['allowed_worker_sizes'] = None
            __props__['database_edition'] = None
            __props__['database_service_objective'] = None
            __props__['default_front_end_scale_factor'] = None
            __props__['environment_capacities'] = None
            __props__['environment_is_healthy'] = None
            __props__['environment_status'] = None
            __props__['last_action'] = None
            __props__['last_action_result'] = None
            __props__['maximum_number_of_machines'] = None
            __props__['provisioning_state'] = None
            __props__['resource_group'] = None
            __props__['status'] = None
            __props__['subscription_id'] = None
            __props__['system_data'] = None
            __props__['type'] = None
            __props__['upgrade_domains'] = None
            __props__['vip_mappings'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:web/v20150801:AppServiceEnvironment"), pulumi.Alias(type_="azure-nextgen:web/v20160901:AppServiceEnvironment"), pulumi.Alias(type_="azure-nextgen:web/v20180201:AppServiceEnvironment"), pulumi.Alias(type_="azure-nextgen:web/v20190801:AppServiceEnvironment"), pulumi.Alias(type_="azure-nextgen:web/v20200601:AppServiceEnvironment"), pulumi.Alias(type_="azure-nextgen:web/v20200901:AppServiceEnvironment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AppServiceEnvironment, __self__).__init__(
            'azure-nextgen:web/latest:AppServiceEnvironment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AppServiceEnvironment':
        """
        Get an existing AppServiceEnvironment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AppServiceEnvironment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedMultiSizes")
    def allowed_multi_sizes(self) -> pulumi.Output[str]:
        """
        List of comma separated strings describing which VM sizes are allowed for front-ends.
        """
        return pulumi.get(self, "allowed_multi_sizes")

    @property
    @pulumi.getter(name="allowedWorkerSizes")
    def allowed_worker_sizes(self) -> pulumi.Output[str]:
        """
        List of comma separated strings describing which VM sizes are allowed for workers.
        """
        return pulumi.get(self, "allowed_worker_sizes")

    @property
    @pulumi.getter(name="apiManagementAccountId")
    def api_management_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        API Management Account associated with the App Service Environment.
        """
        return pulumi.get(self, "api_management_account_id")

    @property
    @pulumi.getter(name="clusterSettings")
    def cluster_settings(self) -> pulumi.Output[Optional[Sequence['outputs.NameValuePairResponse']]]:
        """
        Custom settings for changing the behavior of the App Service Environment.
        """
        return pulumi.get(self, "cluster_settings")

    @property
    @pulumi.getter(name="databaseEdition")
    def database_edition(self) -> pulumi.Output[str]:
        """
        Edition of the metadata database for the App Service Environment, e.g. "Standard".
        """
        return pulumi.get(self, "database_edition")

    @property
    @pulumi.getter(name="databaseServiceObjective")
    def database_service_objective(self) -> pulumi.Output[str]:
        """
        Service objective of the metadata database for the App Service Environment, e.g. "S0".
        """
        return pulumi.get(self, "database_service_objective")

    @property
    @pulumi.getter(name="defaultFrontEndScaleFactor")
    def default_front_end_scale_factor(self) -> pulumi.Output[int]:
        """
        Default Scale Factor for FrontEnds.
        """
        return pulumi.get(self, "default_front_end_scale_factor")

    @property
    @pulumi.getter(name="dnsSuffix")
    def dns_suffix(self) -> pulumi.Output[Optional[str]]:
        """
        DNS suffix of the App Service Environment.
        """
        return pulumi.get(self, "dns_suffix")

    @property
    @pulumi.getter(name="dynamicCacheEnabled")
    def dynamic_cache_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        True/false indicating whether the App Service Environment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
        (most likely because NSG blocked the incoming traffic).
        """
        return pulumi.get(self, "dynamic_cache_enabled")

    @property
    @pulumi.getter(name="environmentCapacities")
    def environment_capacities(self) -> pulumi.Output[Sequence['outputs.StampCapacityResponse']]:
        """
        Current total, used, and available worker capacities.
        """
        return pulumi.get(self, "environment_capacities")

    @property
    @pulumi.getter(name="environmentIsHealthy")
    def environment_is_healthy(self) -> pulumi.Output[bool]:
        """
        True/false indicating whether the App Service Environment is healthy.
        """
        return pulumi.get(self, "environment_is_healthy")

    @property
    @pulumi.getter(name="environmentStatus")
    def environment_status(self) -> pulumi.Output[str]:
        """
        Detailed message about with results of the last check of the App Service Environment.
        """
        return pulumi.get(self, "environment_status")

    @property
    @pulumi.getter(name="frontEndScaleFactor")
    def front_end_scale_factor(self) -> pulumi.Output[Optional[int]]:
        """
        Scale factor for front-ends.
        """
        return pulumi.get(self, "front_end_scale_factor")

    @property
    @pulumi.getter(name="hasLinuxWorkers")
    def has_linux_workers(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag that displays whether an ASE has linux workers or not
        """
        return pulumi.get(self, "has_linux_workers")

    @property
    @pulumi.getter(name="internalLoadBalancingMode")
    def internal_load_balancing_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
        """
        return pulumi.get(self, "internal_load_balancing_mode")

    @property
    @pulumi.getter(name="ipsslAddressCount")
    def ipssl_address_count(self) -> pulumi.Output[Optional[int]]:
        """
        Number of IP SSL addresses reserved for the App Service Environment.
        """
        return pulumi.get(self, "ipssl_address_count")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastAction")
    def last_action(self) -> pulumi.Output[str]:
        """
        Last deployment action on the App Service Environment.
        """
        return pulumi.get(self, "last_action")

    @property
    @pulumi.getter(name="lastActionResult")
    def last_action_result(self) -> pulumi.Output[str]:
        """
        Result of the last deployment action on the App Service Environment.
        """
        return pulumi.get(self, "last_action_result")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource Location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maximumNumberOfMachines")
    def maximum_number_of_machines(self) -> pulumi.Output[int]:
        """
        Maximum number of VMs in the App Service Environment.
        """
        return pulumi.get(self, "maximum_number_of_machines")

    @property
    @pulumi.getter(name="multiRoleCount")
    def multi_role_count(self) -> pulumi.Output[Optional[int]]:
        """
        Number of front-end instances.
        """
        return pulumi.get(self, "multi_role_count")

    @property
    @pulumi.getter(name="multiSize")
    def multi_size(self) -> pulumi.Output[Optional[str]]:
        """
        Front-end VM size, e.g. "Medium", "Large".
        """
        return pulumi.get(self, "multi_size")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkAccessControlList")
    def network_access_control_list(self) -> pulumi.Output[Optional[Sequence['outputs.NetworkAccessControlEntryResponse']]]:
        """
        Access control list for controlling traffic to the App Service Environment.
        """
        return pulumi.get(self, "network_access_control_list")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Provisioning state of the App Service Environment.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> pulumi.Output[str]:
        """
        Resource group of the App Service Environment.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="sslCertKeyVaultId")
    def ssl_cert_key_vault_id(self) -> pulumi.Output[Optional[str]]:
        """
        Key Vault ID for ILB App Service Environment default SSL certificate
        """
        return pulumi.get(self, "ssl_cert_key_vault_id")

    @property
    @pulumi.getter(name="sslCertKeyVaultSecretName")
    def ssl_cert_key_vault_secret_name(self) -> pulumi.Output[Optional[str]]:
        """
        Key Vault Secret Name for ILB App Service Environment default SSL certificate
        """
        return pulumi.get(self, "ssl_cert_key_vault_secret_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Current status of the App Service Environment.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> pulumi.Output[str]:
        """
        Subscription of the App Service Environment.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter
    def suspended(self) -> pulumi.Output[Optional[bool]]:
        """
        <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
         (most likely because NSG blocked the incoming traffic).
        """
        return pulumi.get(self, "suspended")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        The system metadata relating to this resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="upgradeDomains")
    def upgrade_domains(self) -> pulumi.Output[int]:
        """
        Number of upgrade domains of the App Service Environment.
        """
        return pulumi.get(self, "upgrade_domains")

    @property
    @pulumi.getter(name="userWhitelistedIpRanges")
    def user_whitelisted_ip_ranges(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        User added ip ranges to whitelist on ASE db
        """
        return pulumi.get(self, "user_whitelisted_ip_ranges")

    @property
    @pulumi.getter(name="vipMappings")
    def vip_mappings(self) -> pulumi.Output[Sequence['outputs.VirtualIPMappingResponse']]:
        """
        Description of IP SSL mapping for the App Service Environment.
        """
        return pulumi.get(self, "vip_mappings")

    @property
    @pulumi.getter(name="virtualNetwork")
    def virtual_network(self) -> pulumi.Output['outputs.VirtualNetworkProfileResponse']:
        """
        Description of the Virtual Network.
        """
        return pulumi.get(self, "virtual_network")

    @property
    @pulumi.getter(name="vnetName")
    def vnet_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the Virtual Network for the App Service Environment.
        """
        return pulumi.get(self, "vnet_name")

    @property
    @pulumi.getter(name="vnetResourceGroupName")
    def vnet_resource_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        Resource group of the Virtual Network.
        """
        return pulumi.get(self, "vnet_resource_group_name")

    @property
    @pulumi.getter(name="vnetSubnetName")
    def vnet_subnet_name(self) -> pulumi.Output[Optional[str]]:
        """
        Subnet of the Virtual Network.
        """
        return pulumi.get(self, "vnet_subnet_name")

    @property
    @pulumi.getter(name="workerPools")
    def worker_pools(self) -> pulumi.Output[Sequence['outputs.WorkerPoolResponse']]:
        """
        Description of worker pools with worker size IDs, VM sizes, and number of workers in each pool.
        """
        return pulumi.get(self, "worker_pools")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

