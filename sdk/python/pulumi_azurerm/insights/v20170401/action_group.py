# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ActionGroup(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Azure resource name
    """
    properties: pulumi.Output[dict]
    """
    The action groups properties of the resource.
      * `automation_runbook_receivers` (`list`) - The list of AutomationRunbook receivers that are part of this action group.
        * `automation_account_id` (`str`) - The Azure automation account Id which holds this runbook and authenticate to Azure resource.
        * `is_global_runbook` (`bool`) - Indicates whether this instance is global runbook.
        * `name` (`str`) - Indicates name of the webhook.
        * `runbook_name` (`str`) - The name for this runbook.
        * `service_uri` (`str`) - The URI where webhooks should be sent.
        * `webhook_resource_id` (`str`) - The resource id for webhook linked to this runbook.

      * `azure_app_push_receivers` (`list`) - The list of AzureAppPush receivers that are part of this action group.
        * `email_address` (`str`) - The email address registered for the Azure mobile app.
        * `name` (`str`) - The name of the Azure mobile app push receiver. Names must be unique across all receivers within an action group.

      * `email_receivers` (`list`) - The list of email receivers that are part of this action group.
        * `email_address` (`str`) - The email address of this receiver.
        * `name` (`str`) - The name of the email receiver. Names must be unique across all receivers within an action group.
        * `status` (`str`) - The receiver status of the e-mail.

      * `enabled` (`bool`) - Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
      * `group_short_name` (`str`) - The short name of the action group. This will be used in SMS messages.
      * `itsm_receivers` (`list`) - The list of ITSM receivers that are part of this action group.
        * `connection_id` (`str`) - Unique identification of ITSM connection among multiple defined in above workspace.
        * `name` (`str`) - The name of the Itsm receiver. Names must be unique across all receivers within an action group.
        * `region` (`str`) - Region in which workspace resides. Supported values:'centralindia','japaneast','southeastasia','australiasoutheast','uksouth','westcentralus','canadacentral','eastus','westeurope'
        * `ticket_configuration` (`str`) - JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
        * `workspace_id` (`str`) - OMS LA instance identifier.

      * `sms_receivers` (`list`) - The list of SMS receivers that are part of this action group.
        * `country_code` (`str`) - The country code of the SMS receiver.
        * `name` (`str`) - The name of the SMS receiver. Names must be unique across all receivers within an action group.
        * `phone_number` (`str`) - The phone number of the SMS receiver.
        * `status` (`str`) - The status of the receiver.

      * `webhook_receivers` (`list`) - The list of webhook receivers that are part of this action group.
        * `name` (`str`) - The name of the webhook receiver. Names must be unique across all receivers within an action group.
        * `service_uri` (`str`) - The URI where webhooks should be sent.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Azure resource type
    """
    def __init__(__self__, resource_name, opts=None, automation_runbook_receivers=None, azure_app_push_receivers=None, email_receivers=None, enabled=None, group_short_name=None, itsm_receivers=None, location=None, name=None, resource_group_name=None, sms_receivers=None, tags=None, webhook_receivers=None, __props__=None, __name__=None, __opts__=None):
        """
        An action group resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] automation_runbook_receivers: The list of AutomationRunbook receivers that are part of this action group.
        :param pulumi.Input[list] azure_app_push_receivers: The list of AzureAppPush receivers that are part of this action group.
        :param pulumi.Input[list] email_receivers: The list of email receivers that are part of this action group.
        :param pulumi.Input[bool] enabled: Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
        :param pulumi.Input[str] group_short_name: The short name of the action group. This will be used in SMS messages.
        :param pulumi.Input[list] itsm_receivers: The list of ITSM receivers that are part of this action group.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the action group.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[list] sms_receivers: The list of SMS receivers that are part of this action group.
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[list] webhook_receivers: The list of webhook receivers that are part of this action group.

        The **automation_runbook_receivers** object supports the following:

          * `automation_account_id` (`pulumi.Input[str]`) - The Azure automation account Id which holds this runbook and authenticate to Azure resource.
          * `is_global_runbook` (`pulumi.Input[bool]`) - Indicates whether this instance is global runbook.
          * `name` (`pulumi.Input[str]`) - Indicates name of the webhook.
          * `runbook_name` (`pulumi.Input[str]`) - The name for this runbook.
          * `service_uri` (`pulumi.Input[str]`) - The URI where webhooks should be sent.
          * `webhook_resource_id` (`pulumi.Input[str]`) - The resource id for webhook linked to this runbook.

        The **azure_app_push_receivers** object supports the following:

          * `email_address` (`pulumi.Input[str]`) - The email address registered for the Azure mobile app.
          * `name` (`pulumi.Input[str]`) - The name of the Azure mobile app push receiver. Names must be unique across all receivers within an action group.

        The **email_receivers** object supports the following:

          * `email_address` (`pulumi.Input[str]`) - The email address of this receiver.
          * `name` (`pulumi.Input[str]`) - The name of the email receiver. Names must be unique across all receivers within an action group.

        The **itsm_receivers** object supports the following:

          * `connection_id` (`pulumi.Input[str]`) - Unique identification of ITSM connection among multiple defined in above workspace.
          * `name` (`pulumi.Input[str]`) - The name of the Itsm receiver. Names must be unique across all receivers within an action group.
          * `region` (`pulumi.Input[str]`) - Region in which workspace resides. Supported values:'centralindia','japaneast','southeastasia','australiasoutheast','uksouth','westcentralus','canadacentral','eastus','westeurope'
          * `ticket_configuration` (`pulumi.Input[str]`) - JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
          * `workspace_id` (`pulumi.Input[str]`) - OMS LA instance identifier.

        The **sms_receivers** object supports the following:

          * `country_code` (`pulumi.Input[str]`) - The country code of the SMS receiver.
          * `name` (`pulumi.Input[str]`) - The name of the SMS receiver. Names must be unique across all receivers within an action group.
          * `phone_number` (`pulumi.Input[str]`) - The phone number of the SMS receiver.

        The **webhook_receivers** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the webhook receiver. Names must be unique across all receivers within an action group.
          * `service_uri` (`pulumi.Input[str]`) - The URI where webhooks should be sent.
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

            __props__['automation_runbook_receivers'] = automation_runbook_receivers
            __props__['azure_app_push_receivers'] = azure_app_push_receivers
            __props__['email_receivers'] = email_receivers
            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            if group_short_name is None:
                raise TypeError("Missing required property 'group_short_name'")
            __props__['group_short_name'] = group_short_name
            __props__['itsm_receivers'] = itsm_receivers
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sms_receivers'] = sms_receivers
            __props__['tags'] = tags
            __props__['webhook_receivers'] = webhook_receivers
            __props__['properties'] = None
            __props__['type'] = None
        super(ActionGroup, __self__).__init__(
            'azurerm:insights/v20170401:ActionGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ActionGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ActionGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop