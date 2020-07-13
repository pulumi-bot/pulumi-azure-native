# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Catalog(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Resource etag
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
    """
    Azure Data Catalog properties.
      * `admins` (`list`) - Azure data catalog admin list.
        * `object_id` (`str`) - Object Id for the user
        * `upn` (`str`) - UPN of the user.

      * `enable_automatic_unit_adjustment` (`bool`) - Automatic unit adjustment enabled or not.
      * `sku` (`str`) - Azure data catalog SKU.
      * `successfully_provisioned` (`bool`) - Azure data catalog provision status.
      * `units` (`float`) - Azure data catalog units.
      * `users` (`list`) - Azure data catalog user list.
    """
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
        Azure Data Catalog.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: Resource etag
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the data catalog in the specified subscription and resource group.
        :param pulumi.Input[dict] properties: Azure Data Catalog properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
        :param pulumi.Input[dict] tags: Resource tags

        The **properties** object supports the following:

          * `admins` (`pulumi.Input[list]`) - Azure data catalog admin list.
            * `object_id` (`pulumi.Input[str]`) - Object Id for the user
            * `upn` (`pulumi.Input[str]`) - UPN of the user.

          * `enable_automatic_unit_adjustment` (`pulumi.Input[bool]`) - Automatic unit adjustment enabled or not.
          * `sku` (`pulumi.Input[str]`) - Azure data catalog SKU.
          * `successfully_provisioned` (`pulumi.Input[bool]`) - Azure data catalog provision status.
          * `units` (`pulumi.Input[float]`) - Azure data catalog units.
          * `users` (`pulumi.Input[list]`) - Azure data catalog user list.
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
        super(Catalog, __self__).__init__(
            'azurerm:datacatalog:Catalog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Catalog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Catalog(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
