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
    'GetDatabaseAccountResult',
    'AwaitableGetDatabaseAccountResult',
    'get_database_account',
]

@pulumi.output_type
class GetDatabaseAccountResult:
    """
    An Azure Cosmos DB database account.
    """
    def __init__(__self__, capabilities=None, connector_offer=None, consistency_policy=None, database_account_offer_type=None, document_endpoint=None, enable_automatic_failover=None, enable_cassandra_connector=None, enable_multiple_write_locations=None, failover_policies=None, ip_range_filter=None, is_virtual_network_filter_enabled=None, kind=None, location=None, name=None, provisioning_state=None, read_locations=None, tags=None, type=None, virtual_network_rules=None, write_locations=None):
        if capabilities and not isinstance(capabilities, list):
            raise TypeError("Expected argument 'capabilities' to be a list")
        pulumi.set(__self__, "capabilities", capabilities)
        if connector_offer and not isinstance(connector_offer, str):
            raise TypeError("Expected argument 'connector_offer' to be a str")
        pulumi.set(__self__, "connector_offer", connector_offer)
        if consistency_policy and not isinstance(consistency_policy, dict):
            raise TypeError("Expected argument 'consistency_policy' to be a dict")
        pulumi.set(__self__, "consistency_policy", consistency_policy)
        if database_account_offer_type and not isinstance(database_account_offer_type, str):
            raise TypeError("Expected argument 'database_account_offer_type' to be a str")
        pulumi.set(__self__, "database_account_offer_type", database_account_offer_type)
        if document_endpoint and not isinstance(document_endpoint, str):
            raise TypeError("Expected argument 'document_endpoint' to be a str")
        pulumi.set(__self__, "document_endpoint", document_endpoint)
        if enable_automatic_failover and not isinstance(enable_automatic_failover, bool):
            raise TypeError("Expected argument 'enable_automatic_failover' to be a bool")
        pulumi.set(__self__, "enable_automatic_failover", enable_automatic_failover)
        if enable_cassandra_connector and not isinstance(enable_cassandra_connector, bool):
            raise TypeError("Expected argument 'enable_cassandra_connector' to be a bool")
        pulumi.set(__self__, "enable_cassandra_connector", enable_cassandra_connector)
        if enable_multiple_write_locations and not isinstance(enable_multiple_write_locations, bool):
            raise TypeError("Expected argument 'enable_multiple_write_locations' to be a bool")
        pulumi.set(__self__, "enable_multiple_write_locations", enable_multiple_write_locations)
        if failover_policies and not isinstance(failover_policies, list):
            raise TypeError("Expected argument 'failover_policies' to be a list")
        pulumi.set(__self__, "failover_policies", failover_policies)
        if ip_range_filter and not isinstance(ip_range_filter, str):
            raise TypeError("Expected argument 'ip_range_filter' to be a str")
        pulumi.set(__self__, "ip_range_filter", ip_range_filter)
        if is_virtual_network_filter_enabled and not isinstance(is_virtual_network_filter_enabled, bool):
            raise TypeError("Expected argument 'is_virtual_network_filter_enabled' to be a bool")
        pulumi.set(__self__, "is_virtual_network_filter_enabled", is_virtual_network_filter_enabled)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if read_locations and not isinstance(read_locations, list):
            raise TypeError("Expected argument 'read_locations' to be a list")
        pulumi.set(__self__, "read_locations", read_locations)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_network_rules and not isinstance(virtual_network_rules, list):
            raise TypeError("Expected argument 'virtual_network_rules' to be a list")
        pulumi.set(__self__, "virtual_network_rules", virtual_network_rules)
        if write_locations and not isinstance(write_locations, list):
            raise TypeError("Expected argument 'write_locations' to be a list")
        pulumi.set(__self__, "write_locations", write_locations)

    @property
    @pulumi.getter
    def capabilities(self) -> Optional[List['outputs.CapabilityResponse']]:
        """
        List of Cosmos DB capabilities for the account
        """
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter(name="connectorOffer")
    def connector_offer(self) -> Optional[str]:
        """
        The cassandra connector offer type for the Cosmos DB database C* account.
        """
        return pulumi.get(self, "connector_offer")

    @property
    @pulumi.getter(name="consistencyPolicy")
    def consistency_policy(self) -> Optional['outputs.ConsistencyPolicyResponse']:
        """
        The consistency policy for the Cosmos DB database account.
        """
        return pulumi.get(self, "consistency_policy")

    @property
    @pulumi.getter(name="databaseAccountOfferType")
    def database_account_offer_type(self) -> str:
        """
        The offer type for the Cosmos DB database account. Default value: Standard.
        """
        return pulumi.get(self, "database_account_offer_type")

    @property
    @pulumi.getter(name="documentEndpoint")
    def document_endpoint(self) -> str:
        """
        The connection endpoint for the Cosmos DB database account.
        """
        return pulumi.get(self, "document_endpoint")

    @property
    @pulumi.getter(name="enableAutomaticFailover")
    def enable_automatic_failover(self) -> Optional[bool]:
        """
        Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
        """
        return pulumi.get(self, "enable_automatic_failover")

    @property
    @pulumi.getter(name="enableCassandraConnector")
    def enable_cassandra_connector(self) -> Optional[bool]:
        """
        Enables the cassandra connector on the Cosmos DB C* account
        """
        return pulumi.get(self, "enable_cassandra_connector")

    @property
    @pulumi.getter(name="enableMultipleWriteLocations")
    def enable_multiple_write_locations(self) -> Optional[bool]:
        """
        Enables the account to write in multiple locations
        """
        return pulumi.get(self, "enable_multiple_write_locations")

    @property
    @pulumi.getter(name="failoverPolicies")
    def failover_policies(self) -> List['outputs.FailoverPolicyResponse']:
        """
        An array that contains the regions ordered by their failover priorities.
        """
        return pulumi.get(self, "failover_policies")

    @property
    @pulumi.getter(name="ipRangeFilter")
    def ip_range_filter(self) -> Optional[str]:
        """
        Cosmos DB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IPs for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        """
        return pulumi.get(self, "ip_range_filter")

    @property
    @pulumi.getter(name="isVirtualNetworkFilterEnabled")
    def is_virtual_network_filter_enabled(self) -> Optional[bool]:
        """
        Flag to indicate whether to enable/disable Virtual Network ACL rules.
        """
        return pulumi.get(self, "is_virtual_network_filter_enabled")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Indicates the type of database account. This can only be set at database account creation.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The location of the resource group to which the resource belongs.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the database account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'Offline' - the Cosmos DB account is not active. 'DeletionFailed' – the Cosmos DB account deletion failed.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="readLocations")
    def read_locations(self) -> List['outputs.LocationResponse']:
        """
        An array that contains of the read locations enabled for the Cosmos DB account.
        """
        return pulumi.get(self, "read_locations")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of Azure resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworkRules")
    def virtual_network_rules(self) -> Optional[List['outputs.VirtualNetworkRuleResponse']]:
        """
        List of Virtual Network ACL rules configured for the Cosmos DB account.
        """
        return pulumi.get(self, "virtual_network_rules")

    @property
    @pulumi.getter(name="writeLocations")
    def write_locations(self) -> List['outputs.LocationResponse']:
        """
        An array that contains the write location for the Cosmos DB account.
        """
        return pulumi.get(self, "write_locations")


class AwaitableGetDatabaseAccountResult(GetDatabaseAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseAccountResult(
            capabilities=self.capabilities,
            connector_offer=self.connector_offer,
            consistency_policy=self.consistency_policy,
            database_account_offer_type=self.database_account_offer_type,
            document_endpoint=self.document_endpoint,
            enable_automatic_failover=self.enable_automatic_failover,
            enable_cassandra_connector=self.enable_cassandra_connector,
            enable_multiple_write_locations=self.enable_multiple_write_locations,
            failover_policies=self.failover_policies,
            ip_range_filter=self.ip_range_filter,
            is_virtual_network_filter_enabled=self.is_virtual_network_filter_enabled,
            kind=self.kind,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            read_locations=self.read_locations,
            tags=self.tags,
            type=self.type,
            virtual_network_rules=self.virtual_network_rules,
            write_locations=self.write_locations)


def get_database_account(name: Optional[str] = None,
                         resource_group_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseAccountResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Cosmos DB database account name.
    :param str resource_group_name: Name of an Azure resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:documentdb/v20160331:getDatabaseAccount', __args__, opts=opts, typ=GetDatabaseAccountResult).value

    return AwaitableGetDatabaseAccountResult(
        capabilities=__ret__.capabilities,
        connector_offer=__ret__.connector_offer,
        consistency_policy=__ret__.consistency_policy,
        database_account_offer_type=__ret__.database_account_offer_type,
        document_endpoint=__ret__.document_endpoint,
        enable_automatic_failover=__ret__.enable_automatic_failover,
        enable_cassandra_connector=__ret__.enable_cassandra_connector,
        enable_multiple_write_locations=__ret__.enable_multiple_write_locations,
        failover_policies=__ret__.failover_policies,
        ip_range_filter=__ret__.ip_range_filter,
        is_virtual_network_filter_enabled=__ret__.is_virtual_network_filter_enabled,
        kind=__ret__.kind,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        read_locations=__ret__.read_locations,
        tags=__ret__.tags,
        type=__ret__.type,
        virtual_network_rules=__ret__.virtual_network_rules,
        write_locations=__ret__.write_locations)
