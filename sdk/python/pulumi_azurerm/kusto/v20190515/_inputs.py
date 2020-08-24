# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'AzureSkuArgs',
    'DatabaseStatisticsArgs',
    'OptimizedAutoscaleArgs',
    'TrustedExternalTenantArgs',
    'VirtualNetworkConfigurationArgs',
]

@pulumi.input_type
class AzureSkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 tier: pulumi.Input[str],
                 capacity: Optional[pulumi.Input[float]] = None):
        """
        Azure SKU definition.
        :param pulumi.Input[str] name: SKU name.
        :param pulumi.Input[str] tier: SKU tier.
        :param pulumi.Input[float] capacity: The number of instances of the cluster.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "tier", tier)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        SKU name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Input[str]:
        """
        SKU tier.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: pulumi.Input[str]):
        pulumi.set(self, "tier", value)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[float]]:
        """
        The number of instances of the cluster.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "capacity", value)


@pulumi.input_type
class DatabaseStatisticsArgs:
    def __init__(__self__, *,
                 size: Optional[pulumi.Input[float]] = None):
        """
        A class that contains database statistics information.
        :param pulumi.Input[float] size: The database size - the total size of compressed data and index in bytes.
        """
        if size is not None:
            pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[float]]:
        """
        The database size - the total size of compressed data and index in bytes.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "size", value)


@pulumi.input_type
class OptimizedAutoscaleArgs:
    def __init__(__self__, *,
                 is_enabled: pulumi.Input[bool],
                 maximum: pulumi.Input[float],
                 minimum: pulumi.Input[float],
                 version: pulumi.Input[float]):
        """
        A class that contains the optimized auto scale definition.
        :param pulumi.Input[bool] is_enabled: A boolean value that indicate if the optimized autoscale feature is enabled or not.
        :param pulumi.Input[float] maximum: Maximum allowed instances count.
        :param pulumi.Input[float] minimum: Minimum allowed instances count.
        :param pulumi.Input[float] version: The version of the template defined, for instance 1.
        """
        pulumi.set(__self__, "is_enabled", is_enabled)
        pulumi.set(__self__, "maximum", maximum)
        pulumi.set(__self__, "minimum", minimum)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Input[bool]:
        """
        A boolean value that indicate if the optimized autoscale feature is enabled or not.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter
    def maximum(self) -> pulumi.Input[float]:
        """
        Maximum allowed instances count.
        """
        return pulumi.get(self, "maximum")

    @maximum.setter
    def maximum(self, value: pulumi.Input[float]):
        pulumi.set(self, "maximum", value)

    @property
    @pulumi.getter
    def minimum(self) -> pulumi.Input[float]:
        """
        Minimum allowed instances count.
        """
        return pulumi.get(self, "minimum")

    @minimum.setter
    def minimum(self, value: pulumi.Input[float]):
        pulumi.set(self, "minimum", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[float]:
        """
        The version of the template defined, for instance 1.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[float]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class TrustedExternalTenantArgs:
    def __init__(__self__, *,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Represents a tenant ID that is trusted by the cluster.
        :param pulumi.Input[str] value: GUID representing an external tenant.
        """
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        GUID representing an external tenant.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class VirtualNetworkConfigurationArgs:
    def __init__(__self__, *,
                 data_management_public_ip_id: pulumi.Input[str],
                 engine_public_ip_id: pulumi.Input[str],
                 subnet_id: pulumi.Input[str]):
        """
        A class that contains virtual network definition.
        :param pulumi.Input[str] data_management_public_ip_id: Data management's service public IP address resource id.
        :param pulumi.Input[str] engine_public_ip_id: Engine service's public IP address resource id.
        :param pulumi.Input[str] subnet_id: The subnet resource id.
        """
        pulumi.set(__self__, "data_management_public_ip_id", data_management_public_ip_id)
        pulumi.set(__self__, "engine_public_ip_id", engine_public_ip_id)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="dataManagementPublicIpId")
    def data_management_public_ip_id(self) -> pulumi.Input[str]:
        """
        Data management's service public IP address resource id.
        """
        return pulumi.get(self, "data_management_public_ip_id")

    @data_management_public_ip_id.setter
    def data_management_public_ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_management_public_ip_id", value)

    @property
    @pulumi.getter(name="enginePublicIpId")
    def engine_public_ip_id(self) -> pulumi.Input[str]:
        """
        Engine service's public IP address resource id.
        """
        return pulumi.get(self, "engine_public_ip_id")

    @engine_public_ip_id.setter
    def engine_public_ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine_public_ip_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The subnet resource id.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)


