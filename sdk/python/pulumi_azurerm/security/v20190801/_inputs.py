# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'AllowlistCustomAlertRuleArgs',
    'DenylistCustomAlertRuleArgs',
    'RecommendationConfigurationPropertiesArgs',
    'ThresholdCustomAlertRuleArgs',
    'TimeWindowCustomAlertRuleArgs',
    'UserDefinedResourcesPropertiesArgs',
]

@pulumi.input_type
class AllowlistCustomAlertRuleArgs:
    def __init__(__self__, *,
                 allowlist_values: pulumi.Input[List[pulumi.Input[str]]],
                 is_enabled: pulumi.Input[bool],
                 rule_type: pulumi.Input[str]):
        """
        A custom alert rule that checks if a value (depends on the custom alert type) is allowed.
        :param pulumi.Input[List[pulumi.Input[str]]] allowlist_values: The values to allow. The format of the values depends on the rule type.
        :param pulumi.Input[bool] is_enabled: Status of the custom alert.
        :param pulumi.Input[str] rule_type: The type of the custom alert rule.
        """
        pulumi.set(__self__, "allowlist_values", allowlist_values)
        pulumi.set(__self__, "is_enabled", is_enabled)
        pulumi.set(__self__, "rule_type", rule_type)

    @property
    @pulumi.getter(name="allowlistValues")
    def allowlist_values(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The values to allow. The format of the values depends on the rule type.
        """
        return pulumi.get(self, "allowlist_values")

    @allowlist_values.setter
    def allowlist_values(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "allowlist_values", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Input[bool]:
        """
        Status of the custom alert.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> pulumi.Input[str]:
        """
        The type of the custom alert rule.
        """
        return pulumi.get(self, "rule_type")

    @rule_type.setter
    def rule_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_type", value)


@pulumi.input_type
class DenylistCustomAlertRuleArgs:
    def __init__(__self__, *,
                 denylist_values: pulumi.Input[List[pulumi.Input[str]]],
                 is_enabled: pulumi.Input[bool],
                 rule_type: pulumi.Input[str]):
        """
        A custom alert rule that checks if a value (depends on the custom alert type) is denied.
        :param pulumi.Input[List[pulumi.Input[str]]] denylist_values: The values to deny. The format of the values depends on the rule type.
        :param pulumi.Input[bool] is_enabled: Status of the custom alert.
        :param pulumi.Input[str] rule_type: The type of the custom alert rule.
        """
        pulumi.set(__self__, "denylist_values", denylist_values)
        pulumi.set(__self__, "is_enabled", is_enabled)
        pulumi.set(__self__, "rule_type", rule_type)

    @property
    @pulumi.getter(name="denylistValues")
    def denylist_values(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The values to deny. The format of the values depends on the rule type.
        """
        return pulumi.get(self, "denylist_values")

    @denylist_values.setter
    def denylist_values(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "denylist_values", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Input[bool]:
        """
        Status of the custom alert.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> pulumi.Input[str]:
        """
        The type of the custom alert rule.
        """
        return pulumi.get(self, "rule_type")

    @rule_type.setter
    def rule_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_type", value)


@pulumi.input_type
class RecommendationConfigurationPropertiesArgs:
    def __init__(__self__, *,
                 recommendation_type: pulumi.Input[str],
                 status: pulumi.Input[str]):
        """
        The type of IoT Security recommendation.
        :param pulumi.Input[str] recommendation_type: The type of IoT Security recommendation.
        :param pulumi.Input[str] status: Recommendation status. When the recommendation status is disabled recommendations are not generated.
        """
        pulumi.set(__self__, "recommendation_type", recommendation_type)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="recommendationType")
    def recommendation_type(self) -> pulumi.Input[str]:
        """
        The type of IoT Security recommendation.
        """
        return pulumi.get(self, "recommendation_type")

    @recommendation_type.setter
    def recommendation_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "recommendation_type", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        Recommendation status. When the recommendation status is disabled recommendations are not generated.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class ThresholdCustomAlertRuleArgs:
    def __init__(__self__, *,
                 is_enabled: pulumi.Input[bool],
                 max_threshold: pulumi.Input[float],
                 min_threshold: pulumi.Input[float],
                 rule_type: pulumi.Input[str]):
        """
        A custom alert rule that checks if a value (depends on the custom alert type) is within the given range.
        :param pulumi.Input[bool] is_enabled: Status of the custom alert.
        :param pulumi.Input[float] max_threshold: The maximum threshold.
        :param pulumi.Input[float] min_threshold: The minimum threshold.
        :param pulumi.Input[str] rule_type: The type of the custom alert rule.
        """
        pulumi.set(__self__, "is_enabled", is_enabled)
        pulumi.set(__self__, "max_threshold", max_threshold)
        pulumi.set(__self__, "min_threshold", min_threshold)
        pulumi.set(__self__, "rule_type", rule_type)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Input[bool]:
        """
        Status of the custom alert.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="maxThreshold")
    def max_threshold(self) -> pulumi.Input[float]:
        """
        The maximum threshold.
        """
        return pulumi.get(self, "max_threshold")

    @max_threshold.setter
    def max_threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "max_threshold", value)

    @property
    @pulumi.getter(name="minThreshold")
    def min_threshold(self) -> pulumi.Input[float]:
        """
        The minimum threshold.
        """
        return pulumi.get(self, "min_threshold")

    @min_threshold.setter
    def min_threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "min_threshold", value)

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> pulumi.Input[str]:
        """
        The type of the custom alert rule.
        """
        return pulumi.get(self, "rule_type")

    @rule_type.setter
    def rule_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_type", value)


@pulumi.input_type
class TimeWindowCustomAlertRuleArgs:
    def __init__(__self__, *,
                 is_enabled: pulumi.Input[bool],
                 max_threshold: pulumi.Input[float],
                 min_threshold: pulumi.Input[float],
                 rule_type: pulumi.Input[str],
                 time_window_size: pulumi.Input[str]):
        """
        A custom alert rule that checks if the number of activities (depends on the custom alert type) in a time window is within the given range.
        :param pulumi.Input[bool] is_enabled: Status of the custom alert.
        :param pulumi.Input[float] max_threshold: The maximum threshold.
        :param pulumi.Input[float] min_threshold: The minimum threshold.
        :param pulumi.Input[str] rule_type: The type of the custom alert rule.
        :param pulumi.Input[str] time_window_size: The time window size in iso8601 format.
        """
        pulumi.set(__self__, "is_enabled", is_enabled)
        pulumi.set(__self__, "max_threshold", max_threshold)
        pulumi.set(__self__, "min_threshold", min_threshold)
        pulumi.set(__self__, "rule_type", rule_type)
        pulumi.set(__self__, "time_window_size", time_window_size)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Input[bool]:
        """
        Status of the custom alert.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="maxThreshold")
    def max_threshold(self) -> pulumi.Input[float]:
        """
        The maximum threshold.
        """
        return pulumi.get(self, "max_threshold")

    @max_threshold.setter
    def max_threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "max_threshold", value)

    @property
    @pulumi.getter(name="minThreshold")
    def min_threshold(self) -> pulumi.Input[float]:
        """
        The minimum threshold.
        """
        return pulumi.get(self, "min_threshold")

    @min_threshold.setter
    def min_threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "min_threshold", value)

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> pulumi.Input[str]:
        """
        The type of the custom alert rule.
        """
        return pulumi.get(self, "rule_type")

    @rule_type.setter
    def rule_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_type", value)

    @property
    @pulumi.getter(name="timeWindowSize")
    def time_window_size(self) -> pulumi.Input[str]:
        """
        The time window size in iso8601 format.
        """
        return pulumi.get(self, "time_window_size")

    @time_window_size.setter
    def time_window_size(self, value: pulumi.Input[str]):
        pulumi.set(self, "time_window_size", value)


@pulumi.input_type
class UserDefinedResourcesPropertiesArgs:
    def __init__(__self__, *,
                 query: pulumi.Input[str],
                 query_subscriptions: pulumi.Input[List[pulumi.Input[str]]]):
        """
        Properties of the IoT Security solution's user defined resources.
        :param pulumi.Input[str] query: Azure Resource Graph query which represents the security solution's user defined resources. Required to start with "where type != "Microsoft.Devices/IotHubs""
        :param pulumi.Input[List[pulumi.Input[str]]] query_subscriptions: List of Azure subscription ids on which the user defined resources query should be executed.
        """
        pulumi.set(__self__, "query", query)
        pulumi.set(__self__, "query_subscriptions", query_subscriptions)

    @property
    @pulumi.getter
    def query(self) -> pulumi.Input[str]:
        """
        Azure Resource Graph query which represents the security solution's user defined resources. Required to start with "where type != "Microsoft.Devices/IotHubs""
        """
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: pulumi.Input[str]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter(name="querySubscriptions")
    def query_subscriptions(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        List of Azure subscription ids on which the user defined resources query should be executed.
        """
        return pulumi.get(self, "query_subscriptions")

    @query_subscriptions.setter
    def query_subscriptions(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "query_subscriptions", value)


