# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class BotServiceChannel(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Entity Tag
    """
    kind: pulumi.Output[str]
    """
    Required. Gets or sets the Kind of the resource.
    """
    location: pulumi.Output[str]
    """
    Specifies the location of the resource.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    The set of properties specific to bot channel resource
      * `channel_name` (`str`) - The channel name
    """
    sku: pulumi.Output[dict]
    """
    Gets or sets the SKU of the resource.
      * `name` (`str`) - The sku name
      * `tier` (`str`) - Gets the sku tier. This is based on the SKU name.
    """
    tags: pulumi.Output[dict]
    """
    Contains resource tags defined as key/value pairs.
    """
    type: pulumi.Output[str]
    """
    Specifies the type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, etag=None, kind=None, location=None, name=None, properties=None, resource_group_name=None, resource_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Bot channel resource definition

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: Entity Tag
        :param pulumi.Input[str] kind: Required. Gets or sets the Kind of the resource.
        :param pulumi.Input[str] location: Specifies the location of the resource.
        :param pulumi.Input[str] name: The name of the Channel resource.
        :param pulumi.Input[dict] properties: The set of properties specific to bot channel resource
        :param pulumi.Input[str] resource_group_name: The name of the Bot resource group in the user subscription.
        :param pulumi.Input[str] resource_name: The name of the Bot resource.
        :param pulumi.Input[dict] sku: Gets or sets the SKU of the resource.
        :param pulumi.Input[dict] tags: Contains resource tags defined as key/value pairs.

        The **properties** object supports the following:

          * `channel_name` (`pulumi.Input[str]`) - The channel name

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - The sku name
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['etag'] = etag
            __props__['kind'] = kind
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name is None:
                raise TypeError("Missing required property 'resource_name'")
            __props__['resource_name'] = resource_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['type'] = None
        super(BotServiceChannel, __self__).__init__(
            'azurerm:botservice:BotServiceChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing BotServiceChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return BotServiceChannel(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
