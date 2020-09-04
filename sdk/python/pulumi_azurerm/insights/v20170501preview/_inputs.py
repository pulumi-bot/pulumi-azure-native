# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'LogSettingsArgs',
    'MetricSettingsArgs',
    'RetentionPolicyArgs',
    'SubscriptionLogSettingsArgs',
]

@pulumi.input_type
class LogSettingsArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 category: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input['RetentionPolicyArgs']] = None):
        """
        Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
        :param pulumi.Input[bool] enabled: a value indicating whether this log is enabled.
        :param pulumi.Input[str] category: Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
        :param pulumi.Input['RetentionPolicyArgs'] retention_policy: the retention policy for this log.
        """
        pulumi.set(__self__, "enabled", enabled)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if retention_policy is not None:
            pulumi.set(__self__, "retention_policy", retention_policy)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        a value indicating whether this log is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> Optional[pulumi.Input['RetentionPolicyArgs']]:
        """
        the retention policy for this log.
        """
        return pulumi.get(self, "retention_policy")

    @retention_policy.setter
    def retention_policy(self, value: Optional[pulumi.Input['RetentionPolicyArgs']]):
        pulumi.set(self, "retention_policy", value)


@pulumi.input_type
class MetricSettingsArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 category: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input['RetentionPolicyArgs']] = None,
                 time_grain: Optional[pulumi.Input[str]] = None):
        """
        Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.
        :param pulumi.Input[bool] enabled: a value indicating whether this category is enabled.
        :param pulumi.Input[str] category: Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.
        :param pulumi.Input['RetentionPolicyArgs'] retention_policy: the retention policy for this category.
        :param pulumi.Input[str] time_grain: the timegrain of the metric in ISO8601 format.
        """
        pulumi.set(__self__, "enabled", enabled)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if retention_policy is not None:
            pulumi.set(__self__, "retention_policy", retention_policy)
        if time_grain is not None:
            pulumi.set(__self__, "time_grain", time_grain)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        a value indicating whether this category is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> Optional[pulumi.Input['RetentionPolicyArgs']]:
        """
        the retention policy for this category.
        """
        return pulumi.get(self, "retention_policy")

    @retention_policy.setter
    def retention_policy(self, value: Optional[pulumi.Input['RetentionPolicyArgs']]):
        pulumi.set(self, "retention_policy", value)

    @property
    @pulumi.getter(name="timeGrain")
    def time_grain(self) -> Optional[pulumi.Input[str]]:
        """
        the timegrain of the metric in ISO8601 format.
        """
        return pulumi.get(self, "time_grain")

    @time_grain.setter
    def time_grain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_grain", value)


@pulumi.input_type
class RetentionPolicyArgs:
    def __init__(__self__, *,
                 days: pulumi.Input[float],
                 enabled: pulumi.Input[bool]):
        """
        Specifies the retention policy for the log.
        :param pulumi.Input[float] days: the number of days for the retention in days. A value of 0 will retain the events indefinitely.
        :param pulumi.Input[bool] enabled: a value indicating whether the retention policy is enabled.
        """
        pulumi.set(__self__, "days", days)
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def days(self) -> pulumi.Input[float]:
        """
        the number of days for the retention in days. A value of 0 will retain the events indefinitely.
        """
        return pulumi.get(self, "days")

    @days.setter
    def days(self, value: pulumi.Input[float]):
        pulumi.set(self, "days", value)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        a value indicating whether the retention policy is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class SubscriptionLogSettingsArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 category: Optional[pulumi.Input[str]] = None):
        """
        Part of Subscription diagnostic setting. Specifies the settings for a particular log.
        :param pulumi.Input[bool] enabled: a value indicating whether this log is enabled.
        :param pulumi.Input[str] category: Name of a Subscription Diagnostic Log category for a resource type this setting is applied to.
        """
        pulumi.set(__self__, "enabled", enabled)
        if category is not None:
            pulumi.set(__self__, "category", category)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        a value indicating whether this log is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        Name of a Subscription Diagnostic Log category for a resource type this setting is applied to.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)


