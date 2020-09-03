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

__all__ = ['UserSettingsWithLocation']


class UserSettingsWithLocation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[pulumi.InputType['UserPropertiesArgs']]] = None,
                 user_settings_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Response to get user settings

        ## Example Usage
        ### PutUserSettings

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        user_settings_with_location = azurerm.portal.latest.UserSettingsWithLocation("userSettingsWithLocation",
            location="eastus",
            user_settings_name="cloudconsole")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The provider location
        :param pulumi.Input[pulumi.InputType['UserPropertiesArgs']] properties: The cloud shell user settings properties.
        :param pulumi.Input[str] user_settings_name: The name of the user settings
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if user_settings_name is None:
                raise TypeError("Missing required property 'user_settings_name'")
            __props__['user_settings_name'] = user_settings_name
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:portal/v20181001:UserSettingsWithLocation")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(UserSettingsWithLocation, __self__).__init__(
            'azurerm:portal/latest:UserSettingsWithLocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'UserSettingsWithLocation':
        """
        Get an existing UserSettingsWithLocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return UserSettingsWithLocation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.UserPropertiesResponse']:
        """
        The cloud shell user settings properties.
        """
        return pulumi.get(self, "properties")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

