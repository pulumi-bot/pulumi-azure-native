# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['MetricAlert']


class MetricAlert(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['MetricAlertActionArgs']]]]] = None,
                 auto_mitigate: Optional[pulumi.Input[bool]] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['MetricAlertCriteriaArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 evaluation_frequency: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 severity: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_resource_region: Optional[pulumi.Input[str]] = None,
                 target_resource_type: Optional[pulumi.Input[str]] = None,
                 window_size: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The metric alert resource.

        ## Example Usage
        ### Create or update a dynamic alert rule for Multiple Resources

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        metric_alert = azurerm.insights.latest.MetricAlert("metricAlert",
            actions=[{
                "actionGroupId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
                "webHookProperties": {
                    "key11": "value11",
                    "key12": "value12",
                },
            }],
            auto_mitigate=False,
            criteria={
                "odataType": "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
            },
            description="This is the description of the rule1",
            enabled=True,
            evaluation_frequency="PT1M",
            location="global",
            resource_group_name="gigtest",
            rule_name="MetricAlertOnMultipleResources",
            scopes=[
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1",
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2",
            ],
            severity=3,
            tags={},
            target_resource_region="southcentralus",
            target_resource_type="Microsoft.Compute/virtualMachines",
            window_size="PT15M")

        ```
        ### Create or update a dynamic alert rule for Single Resource

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        metric_alert = azurerm.insights.latest.MetricAlert("metricAlert",
            actions=[{
                "actionGroupId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
                "webHookProperties": {
                    "key11": "value11",
                    "key12": "value12",
                },
            }],
            auto_mitigate=False,
            criteria={
                "odataType": "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
            },
            description="This is the description of the rule1",
            enabled=True,
            evaluation_frequency="PT1M",
            location="global",
            resource_group_name="gigtest",
            rule_name="chiricutin",
            scopes=["/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme"],
            severity=3,
            tags={},
            target_resource_region="southcentralus",
            target_resource_type="Microsoft.Compute/virtualMachines",
            window_size="PT15M")

        ```
        ### Create or update a web test alert rule

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        metric_alert = azurerm.insights.latest.MetricAlert("metricAlert",
            actions=[],
            criteria={
                "odataType": "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria",
            },
            description="Automatically created alert rule for availability test \"component-example\" a",
            enabled=True,
            evaluation_frequency="PT1M",
            location="global",
            resource_group_name="rg-example",
            rule_name="webtest-name-example",
            scopes=[
                "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example",
                "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example",
            ],
            severity=4,
            tags={
                "hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example": "Resource",
                "hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example": "Resource",
            },
            window_size="PT15M")

        ```
        ### Create or update an alert rule for Multiple Resource

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        metric_alert = azurerm.insights.latest.MetricAlert("metricAlert",
            actions=[{
                "actionGroupId": "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
                "webHookProperties": {
                    "key11": "value11",
                    "key12": "value12",
                },
            }],
            auto_mitigate=False,
            criteria={
                "odataType": "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
            },
            description="This is the description of the rule1",
            enabled=True,
            evaluation_frequency="PT1M",
            location="global",
            resource_group_name="gigtest",
            rule_name="MetricAlertOnMultipleResources",
            scopes=[
                "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1",
                "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2",
            ],
            severity=3,
            tags={},
            target_resource_region="southcentralus",
            target_resource_type="Microsoft.Compute/virtualMachines",
            window_size="PT15M")

        ```
        ### Create or update an alert rule for Single Resource

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        metric_alert = azurerm.insights.latest.MetricAlert("metricAlert",
            actions=[{
                "actionGroupId": "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
                "webHookProperties": {
                    "key11": "value11",
                    "key12": "value12",
                },
            }],
            auto_mitigate=False,
            criteria={
                "odataType": "Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria",
            },
            description="This is the description of the rule1",
            enabled=True,
            evaluation_frequency="Pt1m",
            location="global",
            resource_group_name="gigtest",
            rule_name="chiricutin",
            scopes=["/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme"],
            severity=3,
            tags={},
            window_size="Pt15m")

        ```
        ### Create or update an alert rule on Resource group(s)

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        metric_alert = azurerm.insights.latest.MetricAlert("metricAlert",
            actions=[{
                "actionGroupId": "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
                "webHookProperties": {
                    "key11": "value11",
                    "key12": "value12",
                },
            }],
            auto_mitigate=False,
            criteria={
                "odataType": "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
            },
            description="This is the description of the rule1",
            enabled=True,
            evaluation_frequency="PT1M",
            location="global",
            resource_group_name="gigtest1",
            rule_name="MetricAlertAtResourceGroupLevel",
            scopes=[
                "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest1",
                "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest2",
            ],
            severity=3,
            tags={},
            target_resource_region="southcentralus",
            target_resource_type="Microsoft.Compute/virtualMachines",
            window_size="PT15M")

        ```
        ### Create or update an alert rule on Subscription

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        metric_alert = azurerm.insights.latest.MetricAlert("metricAlert",
            actions=[{
                "actionGroupId": "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
                "webHookProperties": {
                    "key11": "value11",
                    "key12": "value12",
                },
            }],
            auto_mitigate=False,
            criteria={
                "odataType": "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
            },
            description="This is the description of the rule1",
            enabled=True,
            evaluation_frequency="PT1M",
            location="global",
            resource_group_name="gigtest",
            rule_name="MetricAlertAtSubscriptionLevel",
            scopes=["/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7"],
            severity=3,
            tags={},
            target_resource_region="southcentralus",
            target_resource_type="Microsoft.Compute/virtualMachines",
            window_size="PT15M")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['MetricAlertActionArgs']]]] actions: the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
        :param pulumi.Input[bool] auto_mitigate: the flag that indicates whether the alert should be auto resolved or not. The default is true.
        :param pulumi.Input[pulumi.InputType['MetricAlertCriteriaArgs']] criteria: defines the specific alert criteria information.
        :param pulumi.Input[str] description: the description of the metric alert that will be included in the alert email.
        :param pulumi.Input[bool] enabled: the flag that indicates whether the metric alert is enabled.
        :param pulumi.Input[str] evaluation_frequency: how often the metric alert is evaluated represented in ISO 8601 duration format.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] rule_name: The name of the rule.
        :param pulumi.Input[List[pulumi.Input[str]]] scopes: the list of resource id's that this metric alert is scoped to.
        :param pulumi.Input[float] severity: Alert severity {0, 1, 2, 3, 4}
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[str] target_resource_region: the region of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        :param pulumi.Input[str] target_resource_type: the resource type of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        :param pulumi.Input[str] window_size: the period of time (in ISO 8601 duration format) that is used to monitor alert activity based on the threshold.
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

            __props__['actions'] = actions
            __props__['auto_mitigate'] = auto_mitigate
            if criteria is None:
                raise TypeError("Missing required property 'criteria'")
            __props__['criteria'] = criteria
            if description is None:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            if evaluation_frequency is None:
                raise TypeError("Missing required property 'evaluation_frequency'")
            __props__['evaluation_frequency'] = evaluation_frequency
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if rule_name is None:
                raise TypeError("Missing required property 'rule_name'")
            __props__['rule_name'] = rule_name
            __props__['scopes'] = scopes
            if severity is None:
                raise TypeError("Missing required property 'severity'")
            __props__['severity'] = severity
            __props__['tags'] = tags
            __props__['target_resource_region'] = target_resource_region
            __props__['target_resource_type'] = target_resource_type
            if window_size is None:
                raise TypeError("Missing required property 'window_size'")
            __props__['window_size'] = window_size
            __props__['last_updated_time'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:insights/v20180301:MetricAlert")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(MetricAlert, __self__).__init__(
            'azurerm:insights/latest:MetricAlert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MetricAlert':
        """
        Get an existing MetricAlert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return MetricAlert(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Optional[List['outputs.MetricAlertActionResponse']]]:
        """
        the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="autoMitigate")
    def auto_mitigate(self) -> pulumi.Output[Optional[bool]]:
        """
        the flag that indicates whether the alert should be auto resolved or not. The default is true.
        """
        return pulumi.get(self, "auto_mitigate")

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Output['outputs.MetricAlertCriteriaResponse']:
        """
        defines the specific alert criteria information.
        """
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        the description of the metric alert that will be included in the alert email.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        the flag that indicates whether the metric alert is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="evaluationFrequency")
    def evaluation_frequency(self) -> pulumi.Output[str]:
        """
        how often the metric alert is evaluated represented in ISO 8601 duration format.
        """
        return pulumi.get(self, "evaluation_frequency")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        """
        Last time the rule was updated in ISO8601 format.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Optional[List[str]]]:
        """
        the list of resource id's that this metric alert is scoped to.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[float]:
        """
        Alert severity {0, 1, 2, 3, 4}
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetResourceRegion")
    def target_resource_region(self) -> pulumi.Output[Optional[str]]:
        """
        the region of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        """
        return pulumi.get(self, "target_resource_region")

    @property
    @pulumi.getter(name="targetResourceType")
    def target_resource_type(self) -> pulumi.Output[Optional[str]]:
        """
        the resource type of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        """
        return pulumi.get(self, "target_resource_type")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="windowSize")
    def window_size(self) -> pulumi.Output[str]:
        """
        the period of time (in ISO 8601 duration format) that is used to monitor alert activity based on the threshold.
        """
        return pulumi.get(self, "window_size")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

