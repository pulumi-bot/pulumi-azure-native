# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class WebAppPremierAddOnSlot(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource.
    """
    location: pulumi.Output[str]
    """
    Resource Location.
    """
    name: pulumi.Output[str]
    """
    Resource Name.
    """
    properties: pulumi.Output[dict]
    """
    PremierAddOn resource specific properties
      * `location` (`str`) - Premier add on Location.
      * `marketplace_offer` (`str`) - Premier add on Marketplace offer.
      * `marketplace_publisher` (`str`) - Premier add on Marketplace publisher.
      * `name` (`str`) - Premier add on Name.
      * `product` (`str`) - Premier add on Product.
      * `sku` (`str`) - Premier add on SKU.
      * `tags` (`dict`) - Premier add on Tags.
      * `vendor` (`str`) - Premier add on Vendor.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, kind=None, location=None, marketplace_offer=None, marketplace_publisher=None, name=None, product=None, resource_group_name=None, sku=None, slot=None, tags=None, vendor=None, __props__=None, __name__=None, __opts__=None):
        """
        Premier add-on.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] location: Resource Location.
        :param pulumi.Input[str] marketplace_offer: Premier add on Marketplace offer.
        :param pulumi.Input[str] marketplace_publisher: Premier add on Marketplace publisher.
        :param pulumi.Input[str] name: Premier add on Name.
        :param pulumi.Input[str] product: Premier add on Product.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[str] sku: Premier add on SKU.
        :param pulumi.Input[str] slot: Name of the deployment slot. If a slot is not specified, the API will update the named add-on for the production slot.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[str] vendor: Premier add on Vendor.
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

            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['marketplace_offer'] = marketplace_offer
            __props__['marketplace_publisher'] = marketplace_publisher
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['product'] = product
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            if slot is None:
                raise TypeError("Missing required property 'slot'")
            __props__['slot'] = slot
            __props__['tags'] = tags
            __props__['vendor'] = vendor
            __props__['properties'] = None
            __props__['type'] = None
        super(WebAppPremierAddOnSlot, __self__).__init__(
            'azurerm:web/v20160801:WebAppPremierAddOnSlot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WebAppPremierAddOnSlot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebAppPremierAddOnSlot(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
