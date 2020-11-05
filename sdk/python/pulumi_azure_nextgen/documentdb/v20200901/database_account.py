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

__all__ = ['DatabaseAccount']


class DatabaseAccount(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 api_properties: Optional[pulumi.Input[pulumi.InputType['ApiPropertiesArgs']]] = None,
                 backup_policy: Optional[pulumi.Input[Union[pulumi.InputType['ContinuousModeBackupPolicyArgs'], pulumi.InputType['PeriodicModeBackupPolicyArgs']]]] = None,
                 capabilities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CapabilityArgs']]]]] = None,
                 connector_offer: Optional[pulumi.Input[str]] = None,
                 consistency_policy: Optional[pulumi.Input[pulumi.InputType['ConsistencyPolicyArgs']]] = None,
                 cors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CorsPolicyArgs']]]]] = None,
                 database_account_offer_type: Optional[pulumi.Input[str]] = None,
                 disable_key_based_metadata_write_access: Optional[pulumi.Input[bool]] = None,
                 enable_analytical_storage: Optional[pulumi.Input[bool]] = None,
                 enable_automatic_failover: Optional[pulumi.Input[bool]] = None,
                 enable_cassandra_connector: Optional[pulumi.Input[bool]] = None,
                 enable_free_tier: Optional[pulumi.Input[bool]] = None,
                 enable_multiple_write_locations: Optional[pulumi.Input[bool]] = None,
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpAddressOrRangeArgs']]]]] = None,
                 is_virtual_network_filter_enabled: Optional[pulumi.Input[bool]] = None,
                 key_vault_key_uri: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LocationArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_network_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNetworkRuleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An Azure Cosmos DB database account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Cosmos DB database account name.
        :param pulumi.Input[pulumi.InputType['ApiPropertiesArgs']] api_properties: API specific properties. Currently, supported only for MongoDB API.
        :param pulumi.Input[Union[pulumi.InputType['ContinuousModeBackupPolicyArgs'], pulumi.InputType['PeriodicModeBackupPolicyArgs']]] backup_policy: The object representing the policy for taking backups on an account.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CapabilityArgs']]]] capabilities: List of Cosmos DB capabilities for the account
        :param pulumi.Input[str] connector_offer: The cassandra connector offer type for the Cosmos DB database C* account.
        :param pulumi.Input[pulumi.InputType['ConsistencyPolicyArgs']] consistency_policy: The consistency policy for the Cosmos DB account.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CorsPolicyArgs']]]] cors: The CORS policy for the Cosmos DB database account.
        :param pulumi.Input[str] database_account_offer_type: The offer type for the database
        :param pulumi.Input[bool] disable_key_based_metadata_write_access: Disable write operations on metadata resources (databases, containers, throughput) via account keys
        :param pulumi.Input[bool] enable_analytical_storage: Flag to indicate whether to enable storage analytics.
        :param pulumi.Input[bool] enable_automatic_failover: Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
        :param pulumi.Input[bool] enable_cassandra_connector: Enables the cassandra connector on the Cosmos DB C* account
        :param pulumi.Input[bool] enable_free_tier: Flag to indicate whether Free Tier is enabled.
        :param pulumi.Input[bool] enable_multiple_write_locations: Enables the account to write in multiple locations
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpAddressOrRangeArgs']]]] ip_rules: List of IpRules.
        :param pulumi.Input[bool] is_virtual_network_filter_enabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.
        :param pulumi.Input[str] key_vault_key_uri: The URI of the key vault
        :param pulumi.Input[str] kind: Indicates the type of database account. This can only be set at database account creation.
        :param pulumi.Input[str] location: The location of the resource group to which the resource belongs.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LocationArgs']]]] locations: An array that contains the georeplication locations enabled for the Cosmos DB account.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNetworkRuleArgs']]]] virtual_network_rules: List of Virtual Network ACL rules configured for the Cosmos DB account.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['api_properties'] = api_properties
            __props__['backup_policy'] = backup_policy
            __props__['capabilities'] = capabilities
            __props__['connector_offer'] = connector_offer
            __props__['consistency_policy'] = consistency_policy
            __props__['cors'] = cors
            if database_account_offer_type is None:
                raise TypeError("Missing required property 'database_account_offer_type'")
            __props__['database_account_offer_type'] = database_account_offer_type
            __props__['disable_key_based_metadata_write_access'] = disable_key_based_metadata_write_access
            __props__['enable_analytical_storage'] = enable_analytical_storage
            __props__['enable_automatic_failover'] = enable_automatic_failover
            __props__['enable_cassandra_connector'] = enable_cassandra_connector
            __props__['enable_free_tier'] = enable_free_tier
            __props__['enable_multiple_write_locations'] = enable_multiple_write_locations
            __props__['ip_rules'] = ip_rules
            __props__['is_virtual_network_filter_enabled'] = is_virtual_network_filter_enabled
            __props__['key_vault_key_uri'] = key_vault_key_uri
            __props__['kind'] = kind
            __props__['location'] = location
            if locations is None:
                raise TypeError("Missing required property 'locations'")
            __props__['locations'] = locations
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['virtual_network_rules'] = virtual_network_rules
            __props__['document_endpoint'] = None
            __props__['failover_policies'] = None
            __props__['name'] = None
            __props__['private_endpoint_connections'] = None
            __props__['provisioning_state'] = None
            __props__['public_network_access'] = None
            __props__['read_locations'] = None
            __props__['type'] = None
            __props__['write_locations'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:documentdb/latest:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20150401:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20150408:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20151106:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20160319:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20160331:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20190801:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20191212:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20200301:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20200401:DatabaseAccount"), pulumi.Alias(type_="azure-nextgen:documentdb/v20200601preview:DatabaseAccount")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DatabaseAccount, __self__).__init__(
            'azure-nextgen:documentdb/v20200901:DatabaseAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabaseAccount':
        """
        Get an existing DatabaseAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DatabaseAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiProperties")
    def api_properties(self) -> pulumi.Output[Optional['outputs.ApiPropertiesResponse']]:
        """
        API specific properties.
        """
        return pulumi.get(self, "api_properties")

    @property
    @pulumi.getter(name="backupPolicy")
    def backup_policy(self) -> pulumi.Output[Optional[Any]]:
        """
        The object representing the policy for taking backups on an account.
        """
        return pulumi.get(self, "backup_policy")

    @property
    @pulumi.getter
    def capabilities(self) -> pulumi.Output[Optional[Sequence['outputs.CapabilityResponse']]]:
        """
        List of Cosmos DB capabilities for the account
        """
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter(name="connectorOffer")
    def connector_offer(self) -> pulumi.Output[Optional[str]]:
        """
        The cassandra connector offer type for the Cosmos DB database C* account.
        """
        return pulumi.get(self, "connector_offer")

    @property
    @pulumi.getter(name="consistencyPolicy")
    def consistency_policy(self) -> pulumi.Output[Optional['outputs.ConsistencyPolicyResponse']]:
        """
        The consistency policy for the Cosmos DB database account.
        """
        return pulumi.get(self, "consistency_policy")

    @property
    @pulumi.getter
    def cors(self) -> pulumi.Output[Optional[Sequence['outputs.CorsPolicyResponse']]]:
        """
        The CORS policy for the Cosmos DB database account.
        """
        return pulumi.get(self, "cors")

    @property
    @pulumi.getter(name="databaseAccountOfferType")
    def database_account_offer_type(self) -> pulumi.Output[str]:
        """
        The offer type for the Cosmos DB database account. Default value: Standard.
        """
        return pulumi.get(self, "database_account_offer_type")

    @property
    @pulumi.getter(name="disableKeyBasedMetadataWriteAccess")
    def disable_key_based_metadata_write_access(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable write operations on metadata resources (databases, containers, throughput) via account keys
        """
        return pulumi.get(self, "disable_key_based_metadata_write_access")

    @property
    @pulumi.getter(name="documentEndpoint")
    def document_endpoint(self) -> pulumi.Output[str]:
        """
        The connection endpoint for the Cosmos DB database account.
        """
        return pulumi.get(self, "document_endpoint")

    @property
    @pulumi.getter(name="enableAnalyticalStorage")
    def enable_analytical_storage(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to indicate whether to enable storage analytics.
        """
        return pulumi.get(self, "enable_analytical_storage")

    @property
    @pulumi.getter(name="enableAutomaticFailover")
    def enable_automatic_failover(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
        """
        return pulumi.get(self, "enable_automatic_failover")

    @property
    @pulumi.getter(name="enableCassandraConnector")
    def enable_cassandra_connector(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables the cassandra connector on the Cosmos DB C* account
        """
        return pulumi.get(self, "enable_cassandra_connector")

    @property
    @pulumi.getter(name="enableFreeTier")
    def enable_free_tier(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to indicate whether Free Tier is enabled.
        """
        return pulumi.get(self, "enable_free_tier")

    @property
    @pulumi.getter(name="enableMultipleWriteLocations")
    def enable_multiple_write_locations(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables the account to write in multiple locations
        """
        return pulumi.get(self, "enable_multiple_write_locations")

    @property
    @pulumi.getter(name="failoverPolicies")
    def failover_policies(self) -> pulumi.Output[Sequence['outputs.FailoverPolicyResponse']]:
        """
        An array that contains the regions ordered by their failover priorities.
        """
        return pulumi.get(self, "failover_policies")

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> pulumi.Output[Optional[Sequence['outputs.IpAddressOrRangeResponse']]]:
        """
        List of IpRules.
        """
        return pulumi.get(self, "ip_rules")

    @property
    @pulumi.getter(name="isVirtualNetworkFilterEnabled")
    def is_virtual_network_filter_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to indicate whether to enable/disable Virtual Network ACL rules.
        """
        return pulumi.get(self, "is_virtual_network_filter_enabled")

    @property
    @pulumi.getter(name="keyVaultKeyUri")
    def key_vault_key_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The URI of the key vault
        """
        return pulumi.get(self, "key_vault_key_uri")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates the type of database account. This can only be set at database account creation.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The location of the resource group to which the resource belongs.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Sequence['outputs.LocationResponse']]:
        """
        An array that contains all of the locations enabled for the Cosmos DB account.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the ARM resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> pulumi.Output[Sequence['outputs.PrivateEndpointConnectionResponse']]:
        """
        List of Private Endpoint Connections configured for the Cosmos DB account.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> pulumi.Output[str]:
        """
        Whether requests from Public Network are allowed
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter(name="readLocations")
    def read_locations(self) -> pulumi.Output[Sequence['outputs.LocationResponse']]:
        """
        An array that contains of the read locations enabled for the Cosmos DB account.
        """
        return pulumi.get(self, "read_locations")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of Azure resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworkRules")
    def virtual_network_rules(self) -> pulumi.Output[Optional[Sequence['outputs.VirtualNetworkRuleResponse']]]:
        """
        List of Virtual Network ACL rules configured for the Cosmos DB account.
        """
        return pulumi.get(self, "virtual_network_rules")

    @property
    @pulumi.getter(name="writeLocations")
    def write_locations(self) -> pulumi.Output[Sequence['outputs.LocationResponse']]:
        """
        An array that contains the write location for the Cosmos DB account.
        """
        return pulumi.get(self, "write_locations")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

