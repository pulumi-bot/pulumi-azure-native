# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = ['Webhook']


class Webhook(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[Union[str, 'WebhookAction']]]]] = None,
                 custom_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 registry_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 service_uri: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[Union[str, 'WebhookStatus']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 webhook_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An object that represents a webhook for a container registry.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union[str, 'WebhookAction']]]] actions: The list of actions that trigger the webhook to post notifications.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_headers: Custom headers that will be added to the webhook notifications.
        :param pulumi.Input[str] location: The location of the webhook. This cannot be changed after the resource is created.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to which the container registry belongs.
        :param pulumi.Input[str] scope: The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
        :param pulumi.Input[str] service_uri: The service URI for the webhook to post notifications.
        :param pulumi.Input[Union[str, 'WebhookStatus']] status: The status of the webhook at the time the operation was called.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags for the webhook.
        :param pulumi.Input[str] webhook_name: The name of the webhook.
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

            if actions is None and not opts.urn:
                raise TypeError("Missing required property 'actions'")
            __props__['actions'] = actions
            __props__['custom_headers'] = custom_headers
            __props__['location'] = location
            if registry_name is None and not opts.urn:
                raise TypeError("Missing required property 'registry_name'")
            __props__['registry_name'] = registry_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['scope'] = scope
            if service_uri is None and not opts.urn:
                raise TypeError("Missing required property 'service_uri'")
            __props__['service_uri'] = service_uri
            __props__['status'] = status
            __props__['tags'] = tags
            if webhook_name is None and not opts.urn:
                raise TypeError("Missing required property 'webhook_name'")
            __props__['webhook_name'] = webhook_name
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:containerregistry/latest:Webhook"), pulumi.Alias(type_="azure-nextgen:containerregistry/v20171001:Webhook"), pulumi.Alias(type_="azure-nextgen:containerregistry/v20190501:Webhook"), pulumi.Alias(type_="azure-nextgen:containerregistry/v20191201preview:Webhook"), pulumi.Alias(type_="azure-nextgen:containerregistry/v20201101preview:Webhook")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Webhook, __self__).__init__(
            'azure-nextgen:containerregistry/v20170601preview:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Webhook':
        """
        Get an existing Webhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Webhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of actions that trigger the webhook to post notifications.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the resource. This cannot be changed after the resource is created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the webhook at the time the operation was called.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[str]]:
        """
        The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        The status of the webhook at the time the operation was called.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

