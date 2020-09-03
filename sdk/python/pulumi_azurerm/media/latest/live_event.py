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

__all__ = ['LiveEvent']


class LiveEvent(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 auto_start: Optional[pulumi.Input[bool]] = None,
                 cross_site_access_policies: Optional[pulumi.Input[pulumi.InputType['CrossSiteAccessPoliciesArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encoding: Optional[pulumi.Input[pulumi.InputType['LiveEventEncodingArgs']]] = None,
                 input: Optional[pulumi.Input[pulumi.InputType['LiveEventInputArgs']]] = None,
                 live_event_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 preview: Optional[pulumi.Input[pulumi.InputType['LiveEventPreviewArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 stream_options: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 use_static_hostname: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The Live Event.

        ## Example Usage
        ### Create a LiveEvent

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        live_event = azurerm.media.latest.LiveEvent("liveEvent",
            account_name="slitestmedia10",
            description="test event 1",
            input={
                "keyFrameIntervalDuration": "PT6S",
                "streamingProtocol": "RTMP",
            },
            live_event_name="myLiveEvent1",
            location="West US",
            preview={
                "accessControl": {
                    "ip": {
                        "allow": [{
                            "address": "0.0.0.0",
                            "name": "AllowAll",
                            "subnetPrefixLength": 0,
                        }],
                    },
                },
            },
            resource_group_name="mediaresources",
            tags={
                "tag1": "value1",
                "tag2": "value2",
            })

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[bool] auto_start: The flag indicates if the resource should be automatically started on creation.
        :param pulumi.Input[pulumi.InputType['CrossSiteAccessPoliciesArgs']] cross_site_access_policies: The Live Event access policies.
        :param pulumi.Input[str] description: The Live Event description.
        :param pulumi.Input[pulumi.InputType['LiveEventEncodingArgs']] encoding: The Live Event encoding.
        :param pulumi.Input[pulumi.InputType['LiveEventInputArgs']] input: The Live Event input.
        :param pulumi.Input[str] live_event_name: The name of the Live Event.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[pulumi.InputType['LiveEventPreviewArgs']] preview: The Live Event preview.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[List[pulumi.Input[str]]] stream_options: The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[bool] use_static_hostname: Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['auto_start'] = auto_start
            __props__['cross_site_access_policies'] = cross_site_access_policies
            __props__['description'] = description
            __props__['encoding'] = encoding
            if input is None:
                raise TypeError("Missing required property 'input'")
            __props__['input'] = input
            if live_event_name is None:
                raise TypeError("Missing required property 'live_event_name'")
            __props__['live_event_name'] = live_event_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['preview'] = preview
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['stream_options'] = stream_options
            __props__['tags'] = tags
            __props__['use_static_hostname'] = use_static_hostname
            __props__['created'] = None
            __props__['last_modified'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['resource_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:media/v20180330preview:LiveEvent"), pulumi.Alias(type_="azurerm:media/v20180601preview:LiveEvent"), pulumi.Alias(type_="azurerm:media/v20180701:LiveEvent"), pulumi.Alias(type_="azurerm:media/v20190501preview:LiveEvent"), pulumi.Alias(type_="azurerm:media/v20200501:LiveEvent")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(LiveEvent, __self__).__init__(
            'azurerm:media/latest:LiveEvent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LiveEvent':
        """
        Get an existing LiveEvent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return LiveEvent(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        The exact time the Live Event was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="crossSiteAccessPolicies")
    def cross_site_access_policies(self) -> pulumi.Output[Optional['outputs.CrossSiteAccessPoliciesResponse']]:
        """
        The Live Event access policies.
        """
        return pulumi.get(self, "cross_site_access_policies")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The Live Event description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def encoding(self) -> pulumi.Output[Optional['outputs.LiveEventEncodingResponse']]:
        """
        The Live Event encoding.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter
    def input(self) -> pulumi.Output['outputs.LiveEventInputResponse']:
        """
        The Live Event input.
        """
        return pulumi.get(self, "input")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[str]:
        """
        The exact time the Live Event was last modified.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def preview(self) -> pulumi.Output[Optional['outputs.LiveEventPreviewResponse']]:
        """
        The Live Event preview.
        """
        return pulumi.get(self, "preview")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the Live Event.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceState")
    def resource_state(self) -> pulumi.Output[str]:
        """
        The resource state of the Live Event.
        """
        return pulumi.get(self, "resource_state")

    @property
    @pulumi.getter(name="streamOptions")
    def stream_options(self) -> pulumi.Output[Optional[List[str]]]:
        """
        The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
        """
        return pulumi.get(self, "stream_options")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="useStaticHostname")
    def use_static_hostname(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
        """
        return pulumi.get(self, "use_static_hostname")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

