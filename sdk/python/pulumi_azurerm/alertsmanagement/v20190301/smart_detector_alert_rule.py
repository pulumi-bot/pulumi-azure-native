# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class SmartDetectorAlertRule(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The resource name.
    """
    properties: pulumi.Output[dict]
    """
    The properties of the alert rule.
      * `action_groups` (`dict`) - The alert rule actions.
        * `custom_email_subject` (`str`) - An optional custom email subject to use in email notifications.
        * `custom_webhook_payload` (`str`) - An optional custom web-hook payload to use in web-hook notifications.
        * `group_ids` (`list`) - The Action Group resource IDs.

      * `description` (`str`) - The alert rule description.
      * `detector` (`dict`) - The alert rule's detector.
        * `description` (`str`) - The Smart Detector description. By default this is not populated, unless it's specified in expandDetector
        * `id` (`str`) - The detector id.
        * `image_paths` (`list`) - The Smart Detector image path. By default this is not populated, unless it's specified in expandDetector
        * `name` (`str`) - The Smart Detector name. By default this is not populated, unless it's specified in expandDetector
        * `parameters` (`dict`) - The detector's parameters.'
        * `supported_resource_types` (`list`) - The Smart Detector supported resource types. By default this is not populated, unless it's specified in expandDetector

      * `frequency` (`str`) - The alert rule frequency in ISO8601 format. The time granularity must be in minutes and minimum value is 5 minutes.
      * `scope` (`list`) - The alert rule resources scope.
      * `severity` (`str`) - The alert rule severity.
      * `state` (`str`) - The alert rule state.
      * `throttling` (`dict`) - The alert rule throttling information.
        * `duration` (`str`) - The required duration (in ISO8601 format) to wait before notifying on the alert rule again. The time granularity must be in minutes and minimum value is 0 minutes
    """
    type: pulumi.Output[str]
    """
    The resource type.
    """
    def __init__(__self__, resource_name, opts=None, action_groups=None, description=None, detector=None, frequency=None, name=None, resource_group_name=None, scope=None, severity=None, state=None, throttling=None, __props__=None, __name__=None, __opts__=None):
        """
        The alert rule information

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] action_groups: The alert rule actions.
        :param pulumi.Input[str] description: The alert rule description.
        :param pulumi.Input[dict] detector: The alert rule's detector.
        :param pulumi.Input[str] frequency: The alert rule frequency in ISO8601 format. The time granularity must be in minutes and minimum value is 5 minutes.
        :param pulumi.Input[str] name: The name of the alert rule.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[list] scope: The alert rule resources scope.
        :param pulumi.Input[str] severity: The alert rule severity.
        :param pulumi.Input[str] state: The alert rule state.
        :param pulumi.Input[dict] throttling: The alert rule throttling information.

        The **action_groups** object supports the following:

          * `custom_email_subject` (`pulumi.Input[str]`) - An optional custom email subject to use in email notifications.
          * `custom_webhook_payload` (`pulumi.Input[str]`) - An optional custom web-hook payload to use in web-hook notifications.
          * `group_ids` (`pulumi.Input[list]`) - The Action Group resource IDs.

        The **detector** object supports the following:

          * `description` (`pulumi.Input[str]`) - The Smart Detector description. By default this is not populated, unless it's specified in expandDetector
          * `id` (`pulumi.Input[str]`) - The detector id.
          * `image_paths` (`pulumi.Input[list]`) - The Smart Detector image path. By default this is not populated, unless it's specified in expandDetector
          * `name` (`pulumi.Input[str]`) - The Smart Detector name. By default this is not populated, unless it's specified in expandDetector
          * `parameters` (`pulumi.Input[dict]`) - The detector's parameters.'
          * `supported_resource_types` (`pulumi.Input[list]`) - The Smart Detector supported resource types. By default this is not populated, unless it's specified in expandDetector

        The **throttling** object supports the following:

          * `duration` (`pulumi.Input[str]`) - The required duration (in ISO8601 format) to wait before notifying on the alert rule again. The time granularity must be in minutes and minimum value is 0 minutes
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

            if action_groups is None:
                raise TypeError("Missing required property 'action_groups'")
            __props__['action_groups'] = action_groups
            __props__['description'] = description
            if detector is None:
                raise TypeError("Missing required property 'detector'")
            __props__['detector'] = detector
            if frequency is None:
                raise TypeError("Missing required property 'frequency'")
            __props__['frequency'] = frequency
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            if severity is None:
                raise TypeError("Missing required property 'severity'")
            __props__['severity'] = severity
            if state is None:
                raise TypeError("Missing required property 'state'")
            __props__['state'] = state
            __props__['throttling'] = throttling
            __props__['properties'] = None
            __props__['type'] = None
        super(SmartDetectorAlertRule, __self__).__init__(
            'azurerm:alertsmanagement/v20190301:SmartDetectorAlertRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing SmartDetectorAlertRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return SmartDetectorAlertRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop