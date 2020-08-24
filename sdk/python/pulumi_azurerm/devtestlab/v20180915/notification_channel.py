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

__all__ = ['NotificationChannel']


class NotificationChannel(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 email_recipient: Optional[pulumi.Input[str]] = None,
                 events: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EventArgs']]]]] = None,
                 lab_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_locale: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 web_hook_url: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A notification.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of notification.
        :param pulumi.Input[str] email_recipient: The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EventArgs']]]] events: The list of event for which this notification is enabled.
        :param pulumi.Input[str] lab_name: The name of the lab.
        :param pulumi.Input[str] location: The location of the resource.
        :param pulumi.Input[str] name: The name of the notification channel.
        :param pulumi.Input[str] notification_locale: The locale to use when sending a notification (fallback for unsupported languages is EN).
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
        :param pulumi.Input[str] web_hook_url: The webhook URL to send notifications to.
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

            __props__['description'] = description
            __props__['email_recipient'] = email_recipient
            __props__['events'] = events
            if lab_name is None:
                raise TypeError("Missing required property 'lab_name'")
            __props__['lab_name'] = lab_name
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['notification_locale'] = notification_locale
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['web_hook_url'] = web_hook_url
            __props__['created_date'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
            __props__['unique_identifier'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:devtestlab/v20160515:NotificationChannel")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(NotificationChannel, __self__).__init__(
            'azurerm:devtestlab/v20180915:NotificationChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NotificationChannel':
        """
        Get an existing NotificationChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return NotificationChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        The creation date of the notification channel.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of notification.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="emailRecipient")
    def email_recipient(self) -> Optional[str]:
        """
        The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
        """
        return pulumi.get(self, "email_recipient")

    @property
    @pulumi.getter
    def events(self) -> Optional[List['outputs.EventResponse']]:
        """
        The list of event for which this notification is enabled.
        """
        return pulumi.get(self, "events")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationLocale")
    def notification_locale(self) -> Optional[str]:
        """
        The locale to use when sending a notification (fallback for unsupported languages is EN).
        """
        return pulumi.get(self, "notification_locale")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> str:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")

    @property
    @pulumi.getter(name="webHookUrl")
    def web_hook_url(self) -> Optional[str]:
        """
        The webhook URL to send notifications to.
        """
        return pulumi.get(self, "web_hook_url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

