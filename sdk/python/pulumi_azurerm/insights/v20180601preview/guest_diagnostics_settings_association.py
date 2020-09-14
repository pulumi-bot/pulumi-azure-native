# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['GuestDiagnosticsSettingsAssociation']


class GuestDiagnosticsSettingsAssociation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 association_name: Optional[pulumi.Input[str]] = None,
                 guest_diagnostic_settings_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_uri: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Virtual machine guest diagnostic settings resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] association_name: The name of the diagnostic settings association.
        :param pulumi.Input[str] guest_diagnostic_settings_name: The guest diagnostic settings name.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] resource_uri: The fully qualified ID of the resource, including the resource name and resource type.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
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

            if association_name is None:
                raise TypeError("Missing required property 'association_name'")
            __props__['association_name'] = association_name
            if guest_diagnostic_settings_name is None:
                raise TypeError("Missing required property 'guest_diagnostic_settings_name'")
            __props__['guest_diagnostic_settings_name'] = guest_diagnostic_settings_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if resource_uri is None:
                raise TypeError("Missing required property 'resource_uri'")
            __props__['resource_uri'] = resource_uri
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:insights/latest:GuestDiagnosticsSettingsAssociation")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GuestDiagnosticsSettingsAssociation, __self__).__init__(
            'azurerm:insights/v20180601preview:GuestDiagnosticsSettingsAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GuestDiagnosticsSettingsAssociation':
        """
        Get an existing GuestDiagnosticsSettingsAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return GuestDiagnosticsSettingsAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="guestDiagnosticSettingsName")
    def guest_diagnostic_settings_name(self) -> pulumi.Output[str]:
        """
        The guest diagnostic settings name.
        """
        return pulumi.get(self, "guest_diagnostic_settings_name")

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
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

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

