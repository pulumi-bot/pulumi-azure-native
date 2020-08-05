# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Controller(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Region where the Azure resource is located.
    """
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    sku: pulumi.Output[dict]
    """
    Model representing SKU for Azure Dev Spaces Controller.
      * `name` (`str`) - The name of the SKU for Azure Dev Spaces Controller.
      * `tier` (`str`) - The tier of the SKU for Azure Dev Spaces Controller.
    """
    tags: pulumi.Output[dict]
    """
    Tags for the Azure resource.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, target_container_host_credentials_base64=None, target_container_host_resource_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Controller resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Region where the Azure resource is located.
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[str] resource_group_name: Resource group to which the resource belongs.
        :param pulumi.Input[dict] sku: Model representing SKU for Azure Dev Spaces Controller.
        :param pulumi.Input[dict] tags: Tags for the Azure resource.
        :param pulumi.Input[str] target_container_host_credentials_base64: Credentials of the target container host (base64).
        :param pulumi.Input[str] target_container_host_resource_id: Resource ID of the target container host

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the SKU for Azure Dev Spaces Controller.
          * `tier` (`pulumi.Input[str]`) - The tier of the SKU for Azure Dev Spaces Controller.
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
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            if target_container_host_credentials_base64 is None:
                raise TypeError("Missing required property 'target_container_host_credentials_base64'")
            __props__['target_container_host_credentials_base64'] = target_container_host_credentials_base64
            if target_container_host_resource_id is None:
                raise TypeError("Missing required property 'target_container_host_resource_id'")
            __props__['target_container_host_resource_id'] = target_container_host_resource_id
            __props__['properties'] = None
            __props__['type'] = None
        super(Controller, __self__).__init__(
            'azurerm:devspaces/v20190401:Controller',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Controller resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Controller(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
