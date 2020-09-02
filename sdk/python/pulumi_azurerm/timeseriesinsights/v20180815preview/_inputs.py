# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'LocalTimestampArgs',
    'LocalTimestampTimeZoneOffsetArgs',
    'ReferenceDataSetKeyPropertyArgs',
    'SkuArgs',
]

@pulumi.input_type
class LocalTimestampArgs:
    def __init__(__self__, *,
                 format: Optional[pulumi.Input[str]] = None,
                 time_zone_offset: Optional[pulumi.Input['LocalTimestampTimeZoneOffsetArgs']] = None):
        """
        An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
        :param pulumi.Input[str] format: An enum that represents the format of the local timestamp property that needs to be set.
        :param pulumi.Input['LocalTimestampTimeZoneOffsetArgs'] time_zone_offset: An object that represents the offset information for the local timestamp format specified. Should not be specified for LocalTimestampFormat - Embedded.
        """
        if format is not None:
            pulumi.set(__self__, "format", format)
        if time_zone_offset is not None:
            pulumi.set(__self__, "time_zone_offset", time_zone_offset)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        An enum that represents the format of the local timestamp property that needs to be set.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="timeZoneOffset")
    def time_zone_offset(self) -> Optional[pulumi.Input['LocalTimestampTimeZoneOffsetArgs']]:
        """
        An object that represents the offset information for the local timestamp format specified. Should not be specified for LocalTimestampFormat - Embedded.
        """
        return pulumi.get(self, "time_zone_offset")

    @time_zone_offset.setter
    def time_zone_offset(self, value: Optional[pulumi.Input['LocalTimestampTimeZoneOffsetArgs']]):
        pulumi.set(self, "time_zone_offset", value)


@pulumi.input_type
class LocalTimestampTimeZoneOffsetArgs:
    def __init__(__self__, *,
                 property_name: Optional[pulumi.Input[str]] = None):
        """
        An object that represents the offset information for the local timestamp format specified. Should not be specified for LocalTimestampFormat - Embedded.
        :param pulumi.Input[str] property_name: The event property that will be contain the offset information to calculate the local timestamp. When the LocalTimestampFormat is Iana, the property name will contain the name of the column which contains IANA Timezone Name (eg: Americas/Los Angeles). When LocalTimestampFormat is Timespan, it contains the name of property which contains values representing the offset (eg: P1D or 1.00:00:00)
        """
        if property_name is not None:
            pulumi.set(__self__, "property_name", property_name)

    @property
    @pulumi.getter(name="propertyName")
    def property_name(self) -> Optional[pulumi.Input[str]]:
        """
        The event property that will be contain the offset information to calculate the local timestamp. When the LocalTimestampFormat is Iana, the property name will contain the name of the column which contains IANA Timezone Name (eg: Americas/Los Angeles). When LocalTimestampFormat is Timespan, it contains the name of property which contains values representing the offset (eg: P1D or 1.00:00:00)
        """
        return pulumi.get(self, "property_name")

    @property_name.setter
    def property_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "property_name", value)


@pulumi.input_type
class ReferenceDataSetKeyPropertyArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        A key property for the reference data set. A reference data set can have multiple key properties.
        :param pulumi.Input[str] name: The name of the key property.
        :param pulumi.Input[str] type: The type of the key property.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the key property.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the key property.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 capacity: pulumi.Input[float],
                 name: pulumi.Input[str]):
        """
        The sku determines the type of environment, either standard (S1 or S2) or long-term (L1). For standard environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
        :param pulumi.Input[float] capacity: The capacity of the sku. For standard environments, this value can be changed to support scale out of environments after they have been created.
        :param pulumi.Input[str] name: The name of this SKU.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Input[float]:
        """
        The capacity of the sku. For standard environments, this value can be changed to support scale out of environments after they have been created.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: pulumi.Input[float]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of this SKU.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


