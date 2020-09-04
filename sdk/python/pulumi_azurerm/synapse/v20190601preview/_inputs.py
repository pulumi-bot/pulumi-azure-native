# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'AutoPausePropertiesArgs',
    'AutoScalePropertiesArgs',
    'DataLakeStorageAccountDetailsArgs',
    'IntegrationRuntimeArgs',
    'LibraryRequirementsArgs',
    'ManagedIdentityArgs',
    'PrivateEndpointConnectionArgs',
    'PrivateLinkServiceConnectionStateArgs',
    'SkuArgs',
    'VirtualNetworkProfileArgs',
    'VulnerabilityAssessmentRecurringScansPropertiesArgs',
]

@pulumi.input_type
class AutoPausePropertiesArgs:
    def __init__(__self__, *,
                 delay_in_minutes: Optional[pulumi.Input[float]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        Auto-pausing properties of a Big Data pool powered by Apache Spark
        :param pulumi.Input[float] delay_in_minutes: Number of minutes of idle time before the Big Data pool is automatically paused.
        :param pulumi.Input[bool] enabled: Whether auto-pausing is enabled for the Big Data pool.
        """
        if delay_in_minutes is not None:
            pulumi.set(__self__, "delay_in_minutes", delay_in_minutes)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="delayInMinutes")
    def delay_in_minutes(self) -> Optional[pulumi.Input[float]]:
        """
        Number of minutes of idle time before the Big Data pool is automatically paused.
        """
        return pulumi.get(self, "delay_in_minutes")

    @delay_in_minutes.setter
    def delay_in_minutes(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "delay_in_minutes", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether auto-pausing is enabled for the Big Data pool.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class AutoScalePropertiesArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 max_node_count: Optional[pulumi.Input[float]] = None,
                 min_node_count: Optional[pulumi.Input[float]] = None):
        """
        Auto-scaling properties of a Big Data pool powered by Apache Spark
        :param pulumi.Input[bool] enabled: Whether automatic scaling is enabled for the Big Data pool.
        :param pulumi.Input[float] max_node_count: The maximum number of nodes the Big Data pool can support.
        :param pulumi.Input[float] min_node_count: The minimum number of nodes the Big Data pool can support.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if max_node_count is not None:
            pulumi.set(__self__, "max_node_count", max_node_count)
        if min_node_count is not None:
            pulumi.set(__self__, "min_node_count", min_node_count)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether automatic scaling is enabled for the Big Data pool.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="maxNodeCount")
    def max_node_count(self) -> Optional[pulumi.Input[float]]:
        """
        The maximum number of nodes the Big Data pool can support.
        """
        return pulumi.get(self, "max_node_count")

    @max_node_count.setter
    def max_node_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_node_count", value)

    @property
    @pulumi.getter(name="minNodeCount")
    def min_node_count(self) -> Optional[pulumi.Input[float]]:
        """
        The minimum number of nodes the Big Data pool can support.
        """
        return pulumi.get(self, "min_node_count")

    @min_node_count.setter
    def min_node_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "min_node_count", value)


@pulumi.input_type
class DataLakeStorageAccountDetailsArgs:
    def __init__(__self__, *,
                 account_url: Optional[pulumi.Input[str]] = None,
                 filesystem: Optional[pulumi.Input[str]] = None):
        """
        Details of the data lake storage account associated with the workspace
        :param pulumi.Input[str] account_url: Account URL
        :param pulumi.Input[str] filesystem: Filesystem name
        """
        if account_url is not None:
            pulumi.set(__self__, "account_url", account_url)
        if filesystem is not None:
            pulumi.set(__self__, "filesystem", filesystem)

    @property
    @pulumi.getter(name="accountUrl")
    def account_url(self) -> Optional[pulumi.Input[str]]:
        """
        Account URL
        """
        return pulumi.get(self, "account_url")

    @account_url.setter
    def account_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_url", value)

    @property
    @pulumi.getter
    def filesystem(self) -> Optional[pulumi.Input[str]]:
        """
        Filesystem name
        """
        return pulumi.get(self, "filesystem")

    @filesystem.setter
    def filesystem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filesystem", value)


@pulumi.input_type
class IntegrationRuntimeArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        Azure Synapse nested object which serves as a compute resource for activities.
        :param pulumi.Input[str] type: Type of integration runtime.
        :param pulumi.Input[str] description: Integration runtime description.
        """
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of integration runtime.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Integration runtime description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class LibraryRequirementsArgs:
    def __init__(__self__, *,
                 content: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[str]] = None):
        """
        Library requirements for a Big Data pool powered by Apache Spark
        :param pulumi.Input[str] content: The library requirements.
        :param pulumi.Input[str] filename: The filename of the library requirements file.
        """
        if content is not None:
            pulumi.set(__self__, "content", content)
        if filename is not None:
            pulumi.set(__self__, "filename", filename)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The library requirements.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def filename(self) -> Optional[pulumi.Input[str]]:
        """
        The filename of the library requirements file.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filename", value)


@pulumi.input_type
class ManagedIdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The workspace managed identity
        :param pulumi.Input[str] type: The type of managed identity for the workspace
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of managed identity for the workspace
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class PrivateEndpointConnectionArgs:
    def __init__(__self__, *,
                 private_link_service_connection_state: Optional[pulumi.Input['PrivateLinkServiceConnectionStateArgs']] = None):
        """
        A private endpoint connection
        :param pulumi.Input['PrivateLinkServiceConnectionStateArgs'] private_link_service_connection_state: Connection state of the private endpoint connection.
        """
        if private_link_service_connection_state is not None:
            pulumi.set(__self__, "private_link_service_connection_state", private_link_service_connection_state)

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> Optional[pulumi.Input['PrivateLinkServiceConnectionStateArgs']]:
        """
        Connection state of the private endpoint connection.
        """
        return pulumi.get(self, "private_link_service_connection_state")

    @private_link_service_connection_state.setter
    def private_link_service_connection_state(self, value: Optional[pulumi.Input['PrivateLinkServiceConnectionStateArgs']]):
        pulumi.set(self, "private_link_service_connection_state", value)


@pulumi.input_type
class PrivateLinkServiceConnectionStateArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Connection state details of the private endpoint
        :param pulumi.Input[str] description: The private link service connection description.
        :param pulumi.Input[str] status: The private link service connection status.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The private link service connection description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The private link service connection status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 capacity: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tier: Optional[pulumi.Input[str]] = None):
        """
        SQL pool SKU
        :param pulumi.Input[float] capacity: If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
        :param pulumi.Input[str] name: The SKU name
        :param pulumi.Input[str] tier: The service tier
        """
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[float]]:
        """
        If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The SKU name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The service tier
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)


@pulumi.input_type
class VirtualNetworkProfileArgs:
    def __init__(__self__, *,
                 compute_subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Virtual Network Profile
        :param pulumi.Input[str] compute_subnet_id: Subnet ID used for computes in workspace
        """
        if compute_subnet_id is not None:
            pulumi.set(__self__, "compute_subnet_id", compute_subnet_id)

    @property
    @pulumi.getter(name="computeSubnetId")
    def compute_subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet ID used for computes in workspace
        """
        return pulumi.get(self, "compute_subnet_id")

    @compute_subnet_id.setter
    def compute_subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_subnet_id", value)


@pulumi.input_type
class VulnerabilityAssessmentRecurringScansPropertiesArgs:
    def __init__(__self__, *,
                 email_subscription_admins: Optional[pulumi.Input[bool]] = None,
                 emails: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None):
        """
        Properties of a Vulnerability Assessment recurring scans.
        :param pulumi.Input[bool] email_subscription_admins: Specifies that the schedule scan notification will be is sent to the subscription administrators.
        :param pulumi.Input[List[pulumi.Input[str]]] emails: Specifies an array of e-mail addresses to which the scan notification is sent.
        :param pulumi.Input[bool] is_enabled: Recurring scans state.
        """
        if email_subscription_admins is not None:
            pulumi.set(__self__, "email_subscription_admins", email_subscription_admins)
        if emails is not None:
            pulumi.set(__self__, "emails", emails)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)

    @property
    @pulumi.getter(name="emailSubscriptionAdmins")
    def email_subscription_admins(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies that the schedule scan notification will be is sent to the subscription administrators.
        """
        return pulumi.get(self, "email_subscription_admins")

    @email_subscription_admins.setter
    def email_subscription_admins(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "email_subscription_admins", value)

    @property
    @pulumi.getter
    def emails(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Specifies an array of e-mail addresses to which the scan notification is sent.
        """
        return pulumi.get(self, "emails")

    @emails.setter
    def emails(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "emails", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Recurring scans state.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)


