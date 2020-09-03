# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['View']


class View(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 hub_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 view_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The view resource format.

        ## Example Usage
        ### Views_CreateOrUpdate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        view = azurerm.customerinsights.v20170426.View("view",
            definition="{\\\"isProfileType\\\":false,\\\"profileTypes\\\":[],\\\"widgets\\\":[],\\\"style\\\":[]}",
            display_name={
                "en": "some name",
            },
            hub_name="sdkTestHub",
            resource_group_name="TestHubRG",
            user_id="testUser",
            view_name="testView")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] definition: View definition.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] display_name: Localized display name for the view.
        :param pulumi.Input[str] hub_name: The name of the hub.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] user_id: the user ID.
        :param pulumi.Input[str] view_name: The name of the view.
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

            if definition is None:
                raise TypeError("Missing required property 'definition'")
            __props__['definition'] = definition
            __props__['display_name'] = display_name
            if hub_name is None:
                raise TypeError("Missing required property 'hub_name'")
            __props__['hub_name'] = hub_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['user_id'] = user_id
            if view_name is None:
                raise TypeError("Missing required property 'view_name'")
            __props__['view_name'] = view_name
            __props__['changed'] = None
            __props__['created'] = None
            __props__['name'] = None
            __props__['tenant_id'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:customerinsights/latest:View"), pulumi.Alias(type_="azurerm:customerinsights/v20170101:View")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(View, __self__).__init__(
            'azurerm:customerinsights/v20170426:View',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'View':
        """
        Get an existing View resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return View(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def changed(self) -> pulumi.Output[str]:
        """
        Date time when view was last modified.
        """
        return pulumi.get(self, "changed")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        Date time when view was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[str]:
        """
        View definition.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Localized display name for the view.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        the hub name.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[Optional[str]]:
        """
        the user ID.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="viewName")
    def view_name(self) -> pulumi.Output[str]:
        """
        Name of the view.
        """
        return pulumi.get(self, "view_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

