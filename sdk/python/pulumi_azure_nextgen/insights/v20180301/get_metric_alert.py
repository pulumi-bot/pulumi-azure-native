# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetMetricAlertResult',
    'AwaitableGetMetricAlertResult',
    'get_metric_alert',
]

@pulumi.output_type
class GetMetricAlertResult:
    """
    The metric alert resource.
    """
    def __init__(__self__, actions=None, auto_mitigate=None, criteria=None, description=None, enabled=None, evaluation_frequency=None, id=None, last_updated_time=None, location=None, name=None, scopes=None, severity=None, tags=None, target_resource_region=None, target_resource_type=None, type=None, window_size=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        pulumi.set(__self__, "actions", actions)
        if auto_mitigate and not isinstance(auto_mitigate, bool):
            raise TypeError("Expected argument 'auto_mitigate' to be a bool")
        pulumi.set(__self__, "auto_mitigate", auto_mitigate)
        if criteria and not isinstance(criteria, dict):
            raise TypeError("Expected argument 'criteria' to be a dict")
        pulumi.set(__self__, "criteria", criteria)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if evaluation_frequency and not isinstance(evaluation_frequency, str):
            raise TypeError("Expected argument 'evaluation_frequency' to be a str")
        pulumi.set(__self__, "evaluation_frequency", evaluation_frequency)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_updated_time and not isinstance(last_updated_time, str):
            raise TypeError("Expected argument 'last_updated_time' to be a str")
        pulumi.set(__self__, "last_updated_time", last_updated_time)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if severity and not isinstance(severity, int):
            raise TypeError("Expected argument 'severity' to be a int")
        pulumi.set(__self__, "severity", severity)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if target_resource_region and not isinstance(target_resource_region, str):
            raise TypeError("Expected argument 'target_resource_region' to be a str")
        pulumi.set(__self__, "target_resource_region", target_resource_region)
        if target_resource_type and not isinstance(target_resource_type, str):
            raise TypeError("Expected argument 'target_resource_type' to be a str")
        pulumi.set(__self__, "target_resource_type", target_resource_type)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if window_size and not isinstance(window_size, str):
            raise TypeError("Expected argument 'window_size' to be a str")
        pulumi.set(__self__, "window_size", window_size)

    @property
    @pulumi.getter
    def actions(self) -> Optional[Sequence['outputs.MetricAlertActionResponse']]:
        """
        the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="autoMitigate")
    def auto_mitigate(self) -> Optional[bool]:
        """
        the flag that indicates whether the alert should be auto resolved or not. The default is true.
        """
        return pulumi.get(self, "auto_mitigate")

    @property
    @pulumi.getter
    def criteria(self) -> Any:
        """
        defines the specific alert criteria information.
        """
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        the description of the metric alert that will be included in the alert email.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        the flag that indicates whether the metric alert is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="evaluationFrequency")
    def evaluation_frequency(self) -> str:
        """
        how often the metric alert is evaluated represented in ISO 8601 duration format.
        """
        return pulumi.get(self, "evaluation_frequency")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Azure resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> str:
        """
        Last time the rule was updated in ISO8601 format.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def scopes(self) -> Optional[Sequence[str]]:
        """
        the list of resource id's that this metric alert is scoped to.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def severity(self) -> int:
        """
        Alert severity {0, 1, 2, 3, 4}
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetResourceRegion")
    def target_resource_region(self) -> Optional[str]:
        """
        the region of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        """
        return pulumi.get(self, "target_resource_region")

    @property
    @pulumi.getter(name="targetResourceType")
    def target_resource_type(self) -> Optional[str]:
        """
        the resource type of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        """
        return pulumi.get(self, "target_resource_type")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="windowSize")
    def window_size(self) -> str:
        """
        the period of time (in ISO 8601 duration format) that is used to monitor alert activity based on the threshold.
        """
        return pulumi.get(self, "window_size")


class AwaitableGetMetricAlertResult(GetMetricAlertResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMetricAlertResult(
            actions=self.actions,
            auto_mitigate=self.auto_mitigate,
            criteria=self.criteria,
            description=self.description,
            enabled=self.enabled,
            evaluation_frequency=self.evaluation_frequency,
            id=self.id,
            last_updated_time=self.last_updated_time,
            location=self.location,
            name=self.name,
            scopes=self.scopes,
            severity=self.severity,
            tags=self.tags,
            target_resource_region=self.target_resource_region,
            target_resource_type=self.target_resource_type,
            type=self.type,
            window_size=self.window_size)


def get_metric_alert(resource_group_name: Optional[str] = None,
                     rule_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMetricAlertResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group.
    :param str rule_name: The name of the rule.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['ruleName'] = rule_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:insights/v20180301:getMetricAlert', __args__, opts=opts, typ=GetMetricAlertResult).value

    return AwaitableGetMetricAlertResult(
        actions=__ret__.actions,
        auto_mitigate=__ret__.auto_mitigate,
        criteria=__ret__.criteria,
        description=__ret__.description,
        enabled=__ret__.enabled,
        evaluation_frequency=__ret__.evaluation_frequency,
        id=__ret__.id,
        last_updated_time=__ret__.last_updated_time,
        location=__ret__.location,
        name=__ret__.name,
        scopes=__ret__.scopes,
        severity=__ret__.severity,
        tags=__ret__.tags,
        target_resource_region=__ret__.target_resource_region,
        target_resource_type=__ret__.target_resource_type,
        type=__ret__.type,
        window_size=__ret__.window_size)
