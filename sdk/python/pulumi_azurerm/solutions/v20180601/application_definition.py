# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ApplicationDefinition(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the resource.
      * `principal_id` (`str`) - The principal ID of resource identity.
      * `tenant_id` (`str`) - The tenant ID of resource.
      * `type` (`str`) - The identity type.
    """
    location: pulumi.Output[str]
    """
    Resource location
    """
    managed_by: pulumi.Output[str]
    """
    ID of the resource that manages this resource.
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    The managed application definition properties.
      * `artifacts` (`list`) - The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
        * `name` (`str`) - The managed application artifact name.
        * `type` (`str`) - The managed application artifact type.
        * `uri` (`str`) - The managed application artifact blob uri.

      * `authorizations` (`list`) - The managed application provider authorizations.
        * `principal_id` (`str`) - The provider's principal identifier. This is the identity that the provider will use to call ARM to manage the managed application resources.
        * `role_definition_id` (`str`) - The provider's role definition identifier. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.

      * `create_ui_definition` (`dict`) - The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
      * `description` (`str`) - The managed application definition description.
      * `display_name` (`str`) - The managed application definition display name.
      * `is_enabled` (`str`) - A value indicating whether the package is enabled or not.
      * `lock_level` (`str`) - The managed application lock level.
      * `main_template` (`dict`) - The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
      * `package_file_uri` (`str`) - The managed application definition package file Uri. Use this element
    """
    sku: pulumi.Output[dict]
    """
    The SKU of the resource.
      * `capacity` (`float`) - The SKU capacity.
      * `family` (`str`) - The SKU family.
      * `model` (`str`) - The SKU model.
      * `name` (`str`) - The SKU name.
      * `size` (`str`) - The SKU size.
      * `tier` (`str`) - The SKU tier.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, identity=None, location=None, managed_by=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Information about managed application definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: The identity of the resource.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] managed_by: ID of the resource that manages this resource.
        :param pulumi.Input[str] name: The name of the managed application definition.
        :param pulumi.Input[dict] properties: The managed application definition properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[dict] sku: The SKU of the resource.
        :param pulumi.Input[dict] tags: Resource tags

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.

        The **properties** object supports the following:

          * `artifacts` (`pulumi.Input[list]`) - The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
            * `name` (`pulumi.Input[str]`) - The managed application artifact name.
            * `type` (`pulumi.Input[str]`) - The managed application artifact type.
            * `uri` (`pulumi.Input[str]`) - The managed application artifact blob uri.

          * `authorizations` (`pulumi.Input[list]`) - The managed application provider authorizations.
            * `principal_id` (`pulumi.Input[str]`) - The provider's principal identifier. This is the identity that the provider will use to call ARM to manage the managed application resources.
            * `role_definition_id` (`pulumi.Input[str]`) - The provider's role definition identifier. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.

          * `create_ui_definition` (`pulumi.Input[dict]`) - The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
          * `description` (`pulumi.Input[str]`) - The managed application definition description.
          * `display_name` (`pulumi.Input[str]`) - The managed application definition display name.
          * `is_enabled` (`pulumi.Input[str]`) - A value indicating whether the package is enabled or not.
          * `lock_level` (`pulumi.Input[str]`) - The managed application lock level.
          * `main_template` (`pulumi.Input[dict]`) - The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
          * `package_file_uri` (`pulumi.Input[str]`) - The managed application definition package file Uri. Use this element

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - The SKU capacity.
          * `family` (`pulumi.Input[str]`) - The SKU family.
          * `model` (`pulumi.Input[str]`) - The SKU model.
          * `name` (`pulumi.Input[str]`) - The SKU name.
          * `size` (`pulumi.Input[str]`) - The SKU size.
          * `tier` (`pulumi.Input[str]`) - The SKU tier.
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

            __props__['identity'] = identity
            __props__['location'] = location
            __props__['managed_by'] = managed_by
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['type'] = None
        super(ApplicationDefinition, __self__).__init__(
            'azurerm:solutions/v20180601:ApplicationDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ApplicationDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ApplicationDefinition(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
