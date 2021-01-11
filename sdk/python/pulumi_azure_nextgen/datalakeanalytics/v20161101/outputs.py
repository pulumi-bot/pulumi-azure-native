# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = [
    'ComputePolicyResponse',
    'DataLakeStoreAccountInformationResponse',
    'FirewallRuleResponse',
    'HiveMetastoreResponse',
    'SasTokenInformationResponseResult',
    'StorageAccountInformationResponse',
    'VirtualNetworkRuleResponse',
]

@pulumi.output_type
class ComputePolicyResponse(dict):
    """
    Data Lake Analytics compute policy information.
    """
    def __init__(__self__, *,
                 id: str,
                 max_degree_of_parallelism_per_job: int,
                 min_priority_per_job: int,
                 name: str,
                 object_id: str,
                 object_type: str,
                 type: str):
        """
        Data Lake Analytics compute policy information.
        :param str id: The resource identifier.
        :param int max_degree_of_parallelism_per_job: The maximum degree of parallelism per job this user can use to submit jobs.
        :param int min_priority_per_job: The minimum priority per job this user can use to submit jobs.
        :param str name: The resource name.
        :param str object_id: The AAD object identifier for the entity to create a policy for.
        :param str object_type: The type of AAD object the object identifier refers to.
        :param str type: The resource type.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "max_degree_of_parallelism_per_job", max_degree_of_parallelism_per_job)
        pulumi.set(__self__, "min_priority_per_job", min_priority_per_job)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "object_id", object_id)
        pulumi.set(__self__, "object_type", object_type)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="maxDegreeOfParallelismPerJob")
    def max_degree_of_parallelism_per_job(self) -> int:
        """
        The maximum degree of parallelism per job this user can use to submit jobs.
        """
        return pulumi.get(self, "max_degree_of_parallelism_per_job")

    @property
    @pulumi.getter(name="minPriorityPerJob")
    def min_priority_per_job(self) -> int:
        """
        The minimum priority per job this user can use to submit jobs.
        """
        return pulumi.get(self, "min_priority_per_job")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> str:
        """
        The AAD object identifier for the entity to create a policy for.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="objectType")
    def object_type(self) -> str:
        """
        The type of AAD object the object identifier refers to.
        """
        return pulumi.get(self, "object_type")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DataLakeStoreAccountInformationResponse(dict):
    """
    Data Lake Store account information.
    """
    def __init__(__self__, *,
                 id: str,
                 name: str,
                 suffix: str,
                 type: str):
        """
        Data Lake Store account information.
        :param str id: The resource identifier.
        :param str name: The resource name.
        :param str suffix: The optional suffix for the Data Lake Store account.
        :param str type: The resource type.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "suffix", suffix)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def suffix(self) -> str:
        """
        The optional suffix for the Data Lake Store account.
        """
        return pulumi.get(self, "suffix")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FirewallRuleResponse(dict):
    """
    Data Lake Analytics firewall rule information.
    """
    def __init__(__self__, *,
                 end_ip_address: str,
                 id: str,
                 name: str,
                 start_ip_address: str,
                 type: str):
        """
        Data Lake Analytics firewall rule information.
        :param str end_ip_address: The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
        :param str id: The resource identifier.
        :param str name: The resource name.
        :param str start_ip_address: The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
        :param str type: The resource type.
        """
        pulumi.set(__self__, "end_ip_address", end_ip_address)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "start_ip_address", start_ip_address)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="endIpAddress")
    def end_ip_address(self) -> str:
        """
        The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
        """
        return pulumi.get(self, "end_ip_address")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startIpAddress")
    def start_ip_address(self) -> str:
        """
        The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
        """
        return pulumi.get(self, "start_ip_address")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class HiveMetastoreResponse(dict):
    def __init__(__self__, *,
                 database_name: str,
                 id: str,
                 name: str,
                 nested_resource_provisioning_state: str,
                 password: str,
                 runtime_version: str,
                 server_uri: str,
                 type: str,
                 user_name: str):
        """
        :param str database_name: The databaseName for the Hive MetaStore
        :param str id: The resource identifier.
        :param str name: The resource name.
        :param str nested_resource_provisioning_state: The current state of the NestedResource
        :param str password: The password for the Hive MetaStore
        :param str runtime_version: The runtimeVersion for the Hive MetaStore
        :param str server_uri: The serverUri for the Hive MetaStore
        :param str type: The resource type.
        :param str user_name: The userName for the Hive MetaStore
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "nested_resource_provisioning_state", nested_resource_provisioning_state)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "runtime_version", runtime_version)
        pulumi.set(__self__, "server_uri", server_uri)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        """
        The databaseName for the Hive MetaStore
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nestedResourceProvisioningState")
    def nested_resource_provisioning_state(self) -> str:
        """
        The current state of the NestedResource
        """
        return pulumi.get(self, "nested_resource_provisioning_state")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        The password for the Hive MetaStore
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="runtimeVersion")
    def runtime_version(self) -> str:
        """
        The runtimeVersion for the Hive MetaStore
        """
        return pulumi.get(self, "runtime_version")

    @property
    @pulumi.getter(name="serverUri")
    def server_uri(self) -> str:
        """
        The serverUri for the Hive MetaStore
        """
        return pulumi.get(self, "server_uri")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The userName for the Hive MetaStore
        """
        return pulumi.get(self, "user_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SasTokenInformationResponseResult(dict):
    """
    SAS token information.
    """
    def __init__(__self__, *,
                 access_token: str):
        """
        SAS token information.
        :param str access_token: The access token for the associated Azure Storage Container.
        """
        pulumi.set(__self__, "access_token", access_token)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> str:
        """
        The access token for the associated Azure Storage Container.
        """
        return pulumi.get(self, "access_token")


@pulumi.output_type
class StorageAccountInformationResponse(dict):
    """
    Azure Storage account information.
    """
    def __init__(__self__, *,
                 id: str,
                 name: str,
                 suffix: str,
                 type: str):
        """
        Azure Storage account information.
        :param str id: The resource identifier.
        :param str name: The resource name.
        :param str suffix: The optional suffix for the storage account.
        :param str type: The resource type.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "suffix", suffix)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def suffix(self) -> str:
        """
        The optional suffix for the storage account.
        """
        return pulumi.get(self, "suffix")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VirtualNetworkRuleResponse(dict):
    """
    Data Lake Analytics  VirtualNetwork Rule information.
    """
    def __init__(__self__, *,
                 id: str,
                 name: str,
                 subnet_id: str,
                 type: str,
                 virtual_network_rule_state: str):
        """
        Data Lake Analytics  VirtualNetwork Rule information.
        :param str id: The resource identifier.
        :param str name: The resource name.
        :param str subnet_id: The resource identifier for the subnet
        :param str type: The resource type.
        :param str virtual_network_rule_state: The current state of the VirtualNetwork Rule
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "virtual_network_rule_state", virtual_network_rule_state)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The resource identifier for the subnet
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworkRuleState")
    def virtual_network_rule_state(self) -> str:
        """
        The current state of the VirtualNetwork Rule
        """
        return pulumi.get(self, "virtual_network_rule_state")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


