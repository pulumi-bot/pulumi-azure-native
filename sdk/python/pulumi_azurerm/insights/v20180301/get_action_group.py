# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetActionGroupResult',
    'AwaitableGetActionGroupResult',
    'get_action_group',
]

@pulumi.output_type
class GetActionGroupResult:
    """
    An action group resource.
    """
    def __init__(__self__, automation_runbook_receivers=None, azure_app_push_receivers=None, azure_function_receivers=None, email_receivers=None, enabled=None, group_short_name=None, itsm_receivers=None, location=None, logic_app_receivers=None, name=None, sms_receivers=None, tags=None, type=None, voice_receivers=None, webhook_receivers=None):
        if automation_runbook_receivers and not isinstance(automation_runbook_receivers, list):
            raise TypeError("Expected argument 'automation_runbook_receivers' to be a list")
        pulumi.set(__self__, "automation_runbook_receivers", automation_runbook_receivers)
        if azure_app_push_receivers and not isinstance(azure_app_push_receivers, list):
            raise TypeError("Expected argument 'azure_app_push_receivers' to be a list")
        pulumi.set(__self__, "azure_app_push_receivers", azure_app_push_receivers)
        if azure_function_receivers and not isinstance(azure_function_receivers, list):
            raise TypeError("Expected argument 'azure_function_receivers' to be a list")
        pulumi.set(__self__, "azure_function_receivers", azure_function_receivers)
        if email_receivers and not isinstance(email_receivers, list):
            raise TypeError("Expected argument 'email_receivers' to be a list")
        pulumi.set(__self__, "email_receivers", email_receivers)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if group_short_name and not isinstance(group_short_name, str):
            raise TypeError("Expected argument 'group_short_name' to be a str")
        pulumi.set(__self__, "group_short_name", group_short_name)
        if itsm_receivers and not isinstance(itsm_receivers, list):
            raise TypeError("Expected argument 'itsm_receivers' to be a list")
        pulumi.set(__self__, "itsm_receivers", itsm_receivers)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if logic_app_receivers and not isinstance(logic_app_receivers, list):
            raise TypeError("Expected argument 'logic_app_receivers' to be a list")
        pulumi.set(__self__, "logic_app_receivers", logic_app_receivers)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sms_receivers and not isinstance(sms_receivers, list):
            raise TypeError("Expected argument 'sms_receivers' to be a list")
        pulumi.set(__self__, "sms_receivers", sms_receivers)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if voice_receivers and not isinstance(voice_receivers, list):
            raise TypeError("Expected argument 'voice_receivers' to be a list")
        pulumi.set(__self__, "voice_receivers", voice_receivers)
        if webhook_receivers and not isinstance(webhook_receivers, list):
            raise TypeError("Expected argument 'webhook_receivers' to be a list")
        pulumi.set(__self__, "webhook_receivers", webhook_receivers)

    @property
    @pulumi.getter(name="automationRunbookReceivers")
    def automation_runbook_receivers(self) -> Optional[List['outputs.AutomationRunbookReceiverResponse']]:
        """
        The list of AutomationRunbook receivers that are part of this action group.
        """
        return pulumi.get(self, "automation_runbook_receivers")

    @property
    @pulumi.getter(name="azureAppPushReceivers")
    def azure_app_push_receivers(self) -> Optional[List['outputs.AzureAppPushReceiverResponse']]:
        """
        The list of AzureAppPush receivers that are part of this action group.
        """
        return pulumi.get(self, "azure_app_push_receivers")

    @property
    @pulumi.getter(name="azureFunctionReceivers")
    def azure_function_receivers(self) -> Optional[List['outputs.AzureFunctionReceiverResponse']]:
        """
        The list of azure function receivers that are part of this action group.
        """
        return pulumi.get(self, "azure_function_receivers")

    @property
    @pulumi.getter(name="emailReceivers")
    def email_receivers(self) -> Optional[List['outputs.EmailReceiverResponse']]:
        """
        The list of email receivers that are part of this action group.
        """
        return pulumi.get(self, "email_receivers")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="groupShortName")
    def group_short_name(self) -> str:
        """
        The short name of the action group. This will be used in SMS messages.
        """
        return pulumi.get(self, "group_short_name")

    @property
    @pulumi.getter(name="itsmReceivers")
    def itsm_receivers(self) -> Optional[List['outputs.ItsmReceiverResponse']]:
        """
        The list of ITSM receivers that are part of this action group.
        """
        return pulumi.get(self, "itsm_receivers")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logicAppReceivers")
    def logic_app_receivers(self) -> Optional[List['outputs.LogicAppReceiverResponse']]:
        """
        The list of logic app receivers that are part of this action group.
        """
        return pulumi.get(self, "logic_app_receivers")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="smsReceivers")
    def sms_receivers(self) -> Optional[List['outputs.SmsReceiverResponse']]:
        """
        The list of SMS receivers that are part of this action group.
        """
        return pulumi.get(self, "sms_receivers")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="voiceReceivers")
    def voice_receivers(self) -> Optional[List['outputs.VoiceReceiverResponse']]:
        """
        The list of voice receivers that are part of this action group.
        """
        return pulumi.get(self, "voice_receivers")

    @property
    @pulumi.getter(name="webhookReceivers")
    def webhook_receivers(self) -> Optional[List['outputs.WebhookReceiverResponse']]:
        """
        The list of webhook receivers that are part of this action group.
        """
        return pulumi.get(self, "webhook_receivers")


class AwaitableGetActionGroupResult(GetActionGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActionGroupResult(
            automation_runbook_receivers=self.automation_runbook_receivers,
            azure_app_push_receivers=self.azure_app_push_receivers,
            azure_function_receivers=self.azure_function_receivers,
            email_receivers=self.email_receivers,
            enabled=self.enabled,
            group_short_name=self.group_short_name,
            itsm_receivers=self.itsm_receivers,
            location=self.location,
            logic_app_receivers=self.logic_app_receivers,
            name=self.name,
            sms_receivers=self.sms_receivers,
            tags=self.tags,
            type=self.type,
            voice_receivers=self.voice_receivers,
            webhook_receivers=self.webhook_receivers)


def get_action_group(name: Optional[str] = None,
                     resource_group_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActionGroupResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the action group.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:insights/v20180301:getActionGroup', __args__, opts=opts, typ=GetActionGroupResult).value

    return AwaitableGetActionGroupResult(
        automation_runbook_receivers=__ret__.automation_runbook_receivers,
        azure_app_push_receivers=__ret__.azure_app_push_receivers,
        azure_function_receivers=__ret__.azure_function_receivers,
        email_receivers=__ret__.email_receivers,
        enabled=__ret__.enabled,
        group_short_name=__ret__.group_short_name,
        itsm_receivers=__ret__.itsm_receivers,
        location=__ret__.location,
        logic_app_receivers=__ret__.logic_app_receivers,
        name=__ret__.name,
        sms_receivers=__ret__.sms_receivers,
        tags=__ret__.tags,
        type=__ret__.type,
        voice_receivers=__ret__.voice_receivers,
        webhook_receivers=__ret__.webhook_receivers)
