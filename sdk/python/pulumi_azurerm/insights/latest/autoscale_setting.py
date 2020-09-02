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

__all__ = ['AutoscaleSetting']


class AutoscaleSetting(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscale_setting_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AutoscaleNotificationArgs']]]]] = None,
                 profiles: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AutoscaleProfileArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_resource_uri: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The autoscale setting resource.

        ## Create or update an autoscale setting

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        autoscale_setting = azurerm.insights.latest.AutoscaleSetting("autoscaleSetting",
            autoscale_setting_name="MySetting",
            enabled=True,
            location="West US",
            notifications=[{
                "email": {
                    "customEmails": [
                        "gu@ms.com",
                        "ge@ns.net",
                    ],
                    "sendToSubscriptionAdministrator": True,
                    "sendToSubscriptionCoAdministrators": True,
                },
                "operation": "Scale",
                "webhooks": [{
                    "properties": {},
                    "serviceUri": "http://myservice.com",
                }],
            }],
            profiles=[
                {
                    "capacity": {
                        "default": "1",
                        "maximum": "10",
                        "minimum": "1",
                    },
                    "fixedDate": {
                        "end": "2015-03-05T14:30:00Z",
                        "start": "2015-03-05T14:00:00Z",
                        "timeZone": "UTC",
                    },
                    "name": "adios",
                    "rules": [
                        {
                            "metricTrigger": {
                                "metricName": "Percentage CPU",
                                "metricResourceUri": "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc",
                                "operator": "GreaterThan",
                                "statistic": "Average",
                                "threshold": 10,
                                "timeAggregation": "Average",
                                "timeGrain": "PT1M",
                                "timeWindow": "PT5M",
                            },
                            "scaleAction": {
                                "cooldown": "PT5M",
                                "direction": "Increase",
                                "type": "ChangeCount",
                                "value": "1",
                            },
                        },
                        {
                            "metricTrigger": {
                                "metricName": "Percentage CPU",
                                "metricResourceUri": "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc",
                                "operator": "GreaterThan",
                                "statistic": "Average",
                                "threshold": 15,
                                "timeAggregation": "Average",
                                "timeGrain": "PT2M",
                                "timeWindow": "PT5M",
                            },
                            "scaleAction": {
                                "cooldown": "PT6M",
                                "direction": "Decrease",
                                "type": "ChangeCount",
                                "value": "2",
                            },
                        },
                    ],
                },
                {
                    "capacity": {
                        "default": "1",
                        "maximum": "10",
                        "minimum": "1",
                    },
                    "name": "saludos",
                    "recurrence": {
                        "frequency": "Week",
                        "schedule": {
                            "days": ["1"],
                            "hours": [5],
                            "minutes": [15],
                            "timeZone": "UTC",
                        },
                    },
                    "rules": [
                        {
                            "metricTrigger": {
                                "metricName": "Percentage CPU",
                                "metricResourceUri": "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc",
                                "operator": "GreaterThan",
                                "statistic": "Average",
                                "threshold": 10,
                                "timeAggregation": "Average",
                                "timeGrain": "PT1M",
                                "timeWindow": "PT5M",
                            },
                            "scaleAction": {
                                "cooldown": "PT5M",
                                "direction": "Increase",
                                "type": "ChangeCount",
                                "value": "1",
                            },
                        },
                        {
                            "metricTrigger": {
                                "metricName": "Percentage CPU",
                                "metricResourceUri": "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc",
                                "operator": "GreaterThan",
                                "statistic": "Average",
                                "threshold": 15,
                                "timeAggregation": "Average",
                                "timeGrain": "PT2M",
                                "timeWindow": "PT5M",
                            },
                            "scaleAction": {
                                "cooldown": "PT6M",
                                "direction": "Decrease",
                                "type": "ChangeCount",
                                "value": "2",
                            },
                        },
                    ],
                },
            ],
            resource_group_name="TestingMetricsScaleSet",
            tags={},
            target_resource_uri="/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] autoscale_setting_name: The autoscale setting name.
        :param pulumi.Input[bool] enabled: the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'true'.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: the name of the autoscale setting.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AutoscaleNotificationArgs']]]] notifications: the collection of notifications.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AutoscaleProfileArgs']]]] profiles: the collection of automatic scaling profiles that specify different scaling parameters for different time periods. A maximum of 20 profiles can be specified.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[str] target_resource_uri: the resource identifier of the resource that the autoscale setting should be added to.
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

            if autoscale_setting_name is None:
                raise TypeError("Missing required property 'autoscale_setting_name'")
            __props__['autoscale_setting_name'] = autoscale_setting_name
            __props__['enabled'] = enabled
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['name'] = name
            __props__['notifications'] = notifications
            if profiles is None:
                raise TypeError("Missing required property 'profiles'")
            __props__['profiles'] = profiles
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['target_resource_uri'] = target_resource_uri
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:insights/v20150401:AutoscaleSetting")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AutoscaleSetting, __self__).__init__(
            'azurerm:insights/latest:AutoscaleSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AutoscaleSetting':
        """
        Get an existing AutoscaleSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AutoscaleSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'true'.
        """
        return pulumi.get(self, "enabled")

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
    def notifications(self) -> pulumi.Output[Optional[List['outputs.AutoscaleNotificationResponse']]]:
        """
        the collection of notifications.
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter
    def profiles(self) -> pulumi.Output[List['outputs.AutoscaleProfileResponse']]:
        """
        the collection of automatic scaling profiles that specify different scaling parameters for different time periods. A maximum of 20 profiles can be specified.
        """
        return pulumi.get(self, "profiles")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetResourceUri")
    def target_resource_uri(self) -> pulumi.Output[Optional[str]]:
        """
        the resource identifier of the resource that the autoscale setting should be added to.
        """
        return pulumi.get(self, "target_resource_uri")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
