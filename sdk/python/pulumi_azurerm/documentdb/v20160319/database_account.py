# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class DatabaseAccount(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Indicates the type of database account. This can only be set at database account creation.
    """
    location: pulumi.Output[str]
    """
    The location of the resource group to which the resource belongs.
    """
    name: pulumi.Output[str]
    """
    The name of the database account.
    """
    properties: pulumi.Output[dict]
    """
    Properties for the database account.
      * `capabilities` (`list`) - List of Cosmos DB capabilities for the account
        * `name` (`str`) - Name of the Cosmos DB capability. For example, "name": "EnableCassandra". Current values also include "EnableTable" and "EnableGremlin".

      * `connector_offer` (`str`) - The cassandra connector offer type for the Cosmos DB database C* account.
      * `consistency_policy` (`dict`) - The consistency policy for the Cosmos DB database account.
        * `default_consistency_level` (`str`) - The default consistency level and configuration settings of the Cosmos DB account.
        * `max_interval_in_seconds` (`float`) - When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'.
        * `max_staleness_prefix` (`float`) - When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'.

      * `database_account_offer_type` (`str`) - The offer type for the Cosmos DB database account. Default value: Standard.
      * `document_endpoint` (`str`) - The connection endpoint for the Cosmos DB database account.
      * `enable_automatic_failover` (`bool`) - Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
      * `enable_cassandra_connector` (`bool`) - Enables the cassandra connector on the Cosmos DB C* account
      * `enable_multiple_write_locations` (`bool`) - Enables the account to write in multiple locations
      * `failover_policies` (`list`) - An array that contains the regions ordered by their failover priorities.
        * `failover_priority` (`float`) - The failover priority of the region. A failover priority of 0 indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists.
        * `id` (`str`) - The unique identifier of the region in which the database account replicates to. Example: &lt;accountName&gt;-&lt;locationName&gt;.
        * `location_name` (`str`) - The name of the region in which the database account exists.

      * `ip_range_filter` (`str`) - Cosmos DB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IPs for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
      * `is_virtual_network_filter_enabled` (`bool`) - Flag to indicate whether to enable/disable Virtual Network ACL rules.
      * `provisioning_state` (`str`) - The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'Offline' - the Cosmos DB account is not active. 'DeletionFailed' – the Cosmos DB account deletion failed.
      * `read_locations` (`list`) - An array that contains of the read locations enabled for the Cosmos DB account.
        * `document_endpoint` (`str`) - The connection endpoint for the specific region. Example: https://&lt;accountName&gt;-&lt;locationName&gt;.documents.azure.com:443/
        * `failover_priority` (`float`) - The failover priority of the region. A failover priority of 0 indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists.
        * `id` (`str`) - The unique identifier of the region within the database account. Example: &lt;accountName&gt;-&lt;locationName&gt;.
        * `is_zone_redundant` (`bool`) - Flag to indicate whether or not this region is an AvailabilityZone region
        * `location_name` (`str`) - The name of the region.
        * `provisioning_state` (`str`) - The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'Offline' - the Cosmos DB account is not active. 'DeletionFailed' – the Cosmos DB account deletion failed.

      * `virtual_network_rules` (`list`) - List of Virtual Network ACL rules configured for the Cosmos DB account.
        * `id` (`str`) - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.
        * `ignore_missing_v_net_service_endpoint` (`bool`) - Create firewall rule before the virtual network has vnet service endpoint enabled.

      * `write_locations` (`list`) - An array that contains the write location for the Cosmos DB account.
    """
    tags: pulumi.Output[dict]
    """
    Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
    """
    type: pulumi.Output[str]
    """
    The type of Azure resource.
    """
    def __init__(__self__, resource_name, opts=None, capabilities=None, connector_offer=None, consistency_policy=None, database_account_offer_type=None, enable_automatic_failover=None, enable_cassandra_connector=None, enable_multiple_write_locations=None, ip_range_filter=None, is_virtual_network_filter_enabled=None, kind=None, location=None, locations=None, name=None, resource_group_name=None, tags=None, virtual_network_rules=None, __props__=None, __name__=None, __opts__=None):
        """
        An Azure Cosmos DB database account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] capabilities: List of Cosmos DB capabilities for the account
        :param pulumi.Input[str] connector_offer: The cassandra connector offer type for the Cosmos DB database C* account.
        :param pulumi.Input[dict] consistency_policy: The consistency policy for the Cosmos DB account.
        :param pulumi.Input[str] database_account_offer_type: The offer type for the database
        :param pulumi.Input[bool] enable_automatic_failover: Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
        :param pulumi.Input[bool] enable_cassandra_connector: Enables the cassandra connector on the Cosmos DB C* account
        :param pulumi.Input[bool] enable_multiple_write_locations: Enables the account to write in multiple locations
        :param pulumi.Input[str] ip_range_filter: Cosmos DB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IPs for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        :param pulumi.Input[bool] is_virtual_network_filter_enabled: Flag to indicate whether to enable/disable Virtual Network ACL rules.
        :param pulumi.Input[str] kind: Indicates the type of database account. This can only be set at database account creation.
        :param pulumi.Input[str] location: The location of the resource group to which the resource belongs.
        :param pulumi.Input[list] locations: An array that contains the georeplication locations enabled for the Cosmos DB account.
        :param pulumi.Input[str] name: Cosmos DB database account name.
        :param pulumi.Input[str] resource_group_name: Name of an Azure resource group.
        :param pulumi.Input[dict] tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        :param pulumi.Input[list] virtual_network_rules: List of Virtual Network ACL rules configured for the Cosmos DB account.

        The **capabilities** object supports the following:

          * `name` (`pulumi.Input[str]`) - Name of the Cosmos DB capability. For example, "name": "EnableCassandra". Current values also include "EnableTable" and "EnableGremlin".

        The **consistency_policy** object supports the following:

          * `default_consistency_level` (`pulumi.Input[str]`) - The default consistency level and configuration settings of the Cosmos DB account.
          * `max_interval_in_seconds` (`pulumi.Input[float]`) - When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated. Accepted range for this value is 5 - 86400. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'.
          * `max_staleness_prefix` (`pulumi.Input[float]`) - When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated. Accepted range for this value is 1 – 2,147,483,647. Required when defaultConsistencyPolicy is set to 'BoundedStaleness'.

        The **locations** object supports the following:

          * `failover_priority` (`pulumi.Input[float]`) - The failover priority of the region. A failover priority of 0 indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists.
          * `is_zone_redundant` (`pulumi.Input[bool]`) - Flag to indicate whether or not this region is an AvailabilityZone region
          * `location_name` (`pulumi.Input[str]`) - The name of the region.
          * `provisioning_state` (`pulumi.Input[str]`) - The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'Offline' - the Cosmos DB account is not active. 'DeletionFailed' – the Cosmos DB account deletion failed.

        The **virtual_network_rules** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.
          * `ignore_missing_v_net_service_endpoint` (`pulumi.Input[bool]`) - Create firewall rule before the virtual network has vnet service endpoint enabled.
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

            __props__['capabilities'] = capabilities
            __props__['connector_offer'] = connector_offer
            __props__['consistency_policy'] = consistency_policy
            if database_account_offer_type is None:
                raise TypeError("Missing required property 'database_account_offer_type'")
            __props__['database_account_offer_type'] = database_account_offer_type
            __props__['enable_automatic_failover'] = enable_automatic_failover
            __props__['enable_cassandra_connector'] = enable_cassandra_connector
            __props__['enable_multiple_write_locations'] = enable_multiple_write_locations
            __props__['ip_range_filter'] = ip_range_filter
            __props__['is_virtual_network_filter_enabled'] = is_virtual_network_filter_enabled
            __props__['kind'] = kind
            __props__['location'] = location
            if locations is None:
                raise TypeError("Missing required property 'locations'")
            __props__['locations'] = locations
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['virtual_network_rules'] = virtual_network_rules
            __props__['properties'] = None
            __props__['type'] = None
        super(DatabaseAccount, __self__).__init__(
            'azurerm:documentdb/v20160319:DatabaseAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing DatabaseAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DatabaseAccount(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
