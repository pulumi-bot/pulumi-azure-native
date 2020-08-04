# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Connection(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Resource ETag
    """
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, etag=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        API connection

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: Resource ETag
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: Connection name
        :param pulumi.Input[str] resource_group_name: The resource group
        :param pulumi.Input[dict] tags: Resource tags

        The **properties** object supports the following:

          * `api` (`pulumi.Input[dict]`)
            * `brand_color` (`pulumi.Input[str]`) - Brand color
            * `description` (`pulumi.Input[str]`) - The custom API description
            * `display_name` (`pulumi.Input[str]`) - The display name
            * `icon_uri` (`pulumi.Input[str]`) - The icon URI
            * `id` (`pulumi.Input[str]`) - Resource reference id
            * `name` (`pulumi.Input[str]`) - The name of the API
            * `swagger` (`pulumi.Input[dict]`) - The JSON representation of the swagger
            * `type` (`pulumi.Input[str]`) - Resource reference type

          * `changed_time` (`pulumi.Input[str]`) - Timestamp of last connection change
          * `created_time` (`pulumi.Input[str]`) - Timestamp of the connection creation
          * `custom_parameter_values` (`pulumi.Input[dict]`) - Dictionary of custom parameter values
          * `display_name` (`pulumi.Input[str]`) - Display name
          * `non_secret_parameter_values` (`pulumi.Input[dict]`) - Dictionary of nonsecret parameter values
          * `parameter_values` (`pulumi.Input[dict]`) - Dictionary of parameter values
          * `statuses` (`pulumi.Input[list]`) - Status of the connection
            * `error` (`pulumi.Input[dict]`) - Connection error
              * `code` (`pulumi.Input[str]`) - Code of the status
              * `etag` (`pulumi.Input[str]`) - Resource ETag
              * `location` (`pulumi.Input[str]`) - Resource location
              * `message` (`pulumi.Input[str]`) - Description of the status
              * `tags` (`pulumi.Input[dict]`) - Resource tags

            * `status` (`pulumi.Input[str]`) - The gateway status
            * `target` (`pulumi.Input[str]`) - Target of the error

          * `test_links` (`pulumi.Input[list]`) - Links to test the API connection
            * `method` (`pulumi.Input[str]`) - HTTP Method
            * `request_uri` (`pulumi.Input[str]`) - Test link request URI
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

            __props__['etag'] = etag
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(Connection, __self__).__init__(
            'azurerm:web/v20160601:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Connection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop