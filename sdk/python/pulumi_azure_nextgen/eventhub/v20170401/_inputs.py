# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'CaptureDescriptionArgs',
    'DestinationArgs',
    'NWRuleSetIpRulesArgs',
    'NWRuleSetVirtualNetworkRulesArgs',
    'SkuArgs',
    'SubnetArgs',
]

@pulumi.input_type
class CaptureDescriptionArgs:
    def __init__(__self__, *,
                 destination: Optional[pulumi.Input['DestinationArgs']] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 encoding: Optional[pulumi.Input[str]] = None,
                 interval_in_seconds: Optional[pulumi.Input[int]] = None,
                 size_limit_in_bytes: Optional[pulumi.Input[int]] = None,
                 skip_empty_archives: Optional[pulumi.Input[bool]] = None):
        """
        Properties to configure capture description for eventhub
        :param pulumi.Input['DestinationArgs'] destination: Properties of Destination where capture will be stored. (Storage Account, Blob Names)
        :param pulumi.Input[bool] enabled: A value that indicates whether capture description is enabled. 
        :param pulumi.Input[str] encoding: Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be deprecated in New API Version
        :param pulumi.Input[int] interval_in_seconds: The time window allows you to set the frequency with which the capture to Azure Blobs will happen, value should between 60 to 900 seconds
        :param pulumi.Input[int] size_limit_in_bytes: The size window defines the amount of data built up in your Event Hub before an capture operation, value should be between 10485760 to 524288000 bytes
        :param pulumi.Input[bool] skip_empty_archives: A value that indicates whether to Skip Empty Archives
        """
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if encoding is not None:
            pulumi.set(__self__, "encoding", encoding)
        if interval_in_seconds is not None:
            pulumi.set(__self__, "interval_in_seconds", interval_in_seconds)
        if size_limit_in_bytes is not None:
            pulumi.set(__self__, "size_limit_in_bytes", size_limit_in_bytes)
        if skip_empty_archives is not None:
            pulumi.set(__self__, "skip_empty_archives", skip_empty_archives)

    @property
    @pulumi.getter
    def destination(self) -> Optional[pulumi.Input['DestinationArgs']]:
        """
        Properties of Destination where capture will be stored. (Storage Account, Blob Names)
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input['DestinationArgs']]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        A value that indicates whether capture description is enabled. 
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be deprecated in New API Version
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="intervalInSeconds")
    def interval_in_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The time window allows you to set the frequency with which the capture to Azure Blobs will happen, value should between 60 to 900 seconds
        """
        return pulumi.get(self, "interval_in_seconds")

    @interval_in_seconds.setter
    def interval_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval_in_seconds", value)

    @property
    @pulumi.getter(name="sizeLimitInBytes")
    def size_limit_in_bytes(self) -> Optional[pulumi.Input[int]]:
        """
        The size window defines the amount of data built up in your Event Hub before an capture operation, value should be between 10485760 to 524288000 bytes
        """
        return pulumi.get(self, "size_limit_in_bytes")

    @size_limit_in_bytes.setter
    def size_limit_in_bytes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_limit_in_bytes", value)

    @property
    @pulumi.getter(name="skipEmptyArchives")
    def skip_empty_archives(self) -> Optional[pulumi.Input[bool]]:
        """
        A value that indicates whether to Skip Empty Archives
        """
        return pulumi.get(self, "skip_empty_archives")

    @skip_empty_archives.setter
    def skip_empty_archives(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_empty_archives", value)


@pulumi.input_type
class DestinationArgs:
    def __init__(__self__, *,
                 archive_name_format: Optional[pulumi.Input[str]] = None,
                 blob_container: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 storage_account_resource_id: Optional[pulumi.Input[str]] = None):
        """
        Capture storage details for capture description
        :param pulumi.Input[str] archive_name_format: Blob naming convention for archive, e.g. {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order
        :param pulumi.Input[str] blob_container: Blob container Name
        :param pulumi.Input[str] name: Name for capture destination
        :param pulumi.Input[str] storage_account_resource_id: Resource id of the storage account to be used to create the blobs
        """
        if archive_name_format is not None:
            pulumi.set(__self__, "archive_name_format", archive_name_format)
        if blob_container is not None:
            pulumi.set(__self__, "blob_container", blob_container)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if storage_account_resource_id is not None:
            pulumi.set(__self__, "storage_account_resource_id", storage_account_resource_id)

    @property
    @pulumi.getter(name="archiveNameFormat")
    def archive_name_format(self) -> Optional[pulumi.Input[str]]:
        """
        Blob naming convention for archive, e.g. {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order
        """
        return pulumi.get(self, "archive_name_format")

    @archive_name_format.setter
    def archive_name_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "archive_name_format", value)

    @property
    @pulumi.getter(name="blobContainer")
    def blob_container(self) -> Optional[pulumi.Input[str]]:
        """
        Blob container Name
        """
        return pulumi.get(self, "blob_container")

    @blob_container.setter
    def blob_container(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "blob_container", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for capture destination
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="storageAccountResourceId")
    def storage_account_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource id of the storage account to be used to create the blobs
        """
        return pulumi.get(self, "storage_account_resource_id")

    @storage_account_resource_id.setter
    def storage_account_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_resource_id", value)


@pulumi.input_type
class NWRuleSetIpRulesArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 ip_mask: Optional[pulumi.Input[str]] = None):
        """
        Description of NetWorkRuleSet - IpRules resource.
        :param pulumi.Input[str] action: The IP Filter Action
        :param pulumi.Input[str] ip_mask: IP Mask
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if ip_mask is not None:
            pulumi.set(__self__, "ip_mask", ip_mask)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The IP Filter Action
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="ipMask")
    def ip_mask(self) -> Optional[pulumi.Input[str]]:
        """
        IP Mask
        """
        return pulumi.get(self, "ip_mask")

    @ip_mask.setter
    def ip_mask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_mask", value)


@pulumi.input_type
class NWRuleSetVirtualNetworkRulesArgs:
    def __init__(__self__, *,
                 ignore_missing_vnet_service_endpoint: Optional[pulumi.Input[bool]] = None,
                 subnet: Optional[pulumi.Input['SubnetArgs']] = None):
        """
        Description of VirtualNetworkRules - NetworkRules resource.
        :param pulumi.Input[bool] ignore_missing_vnet_service_endpoint: Value that indicates whether to ignore missing VNet Service Endpoint
        :param pulumi.Input['SubnetArgs'] subnet: Subnet properties
        """
        if ignore_missing_vnet_service_endpoint is not None:
            pulumi.set(__self__, "ignore_missing_vnet_service_endpoint", ignore_missing_vnet_service_endpoint)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)

    @property
    @pulumi.getter(name="ignoreMissingVnetServiceEndpoint")
    def ignore_missing_vnet_service_endpoint(self) -> Optional[pulumi.Input[bool]]:
        """
        Value that indicates whether to ignore missing VNet Service Endpoint
        """
        return pulumi.get(self, "ignore_missing_vnet_service_endpoint")

    @ignore_missing_vnet_service_endpoint.setter
    def ignore_missing_vnet_service_endpoint(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_missing_vnet_service_endpoint", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input['SubnetArgs']]:
        """
        Subnet properties
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input['SubnetArgs']]):
        pulumi.set(self, "subnet", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 capacity: Optional[pulumi.Input[int]] = None,
                 tier: Optional[pulumi.Input[str]] = None):
        """
        SKU parameters supplied to the create namespace operation
        :param pulumi.Input[str] name: Name of this SKU.
        :param pulumi.Input[int] capacity: The Event Hubs throughput units, value should be 0 to 20 throughput units.
        :param pulumi.Input[str] tier: The billing tier of this particular SKU.
        """
        pulumi.set(__self__, "name", name)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of this SKU.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The Event Hubs throughput units, value should be 0 to 20 throughput units.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The billing tier of this particular SKU.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)


@pulumi.input_type
class SubnetArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str]):
        """
        Properties supplied for Subnet
        :param pulumi.Input[str] id: Resource ID of Virtual Network Subnet
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        Resource ID of Virtual Network Subnet
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)


