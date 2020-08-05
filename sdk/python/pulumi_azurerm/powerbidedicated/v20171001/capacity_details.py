# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class CapacityDetails(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Location of the PowerBI Dedicated resource.
    """
    name: pulumi.Output[str]
    """
    The name of the PowerBI Dedicated resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the provision operation request.
      * `administration` (`dict`) - A collection of Dedicated capacity administrators
        * `members` (`list`) - An array of administrator user identities.

      * `provisioning_state` (`str`) - The current deployment state of PowerBI Dedicated resource. The provisioningState is to indicate states for resource provisioning.
      * `state` (`str`) - The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning.
    """
    sku: pulumi.Output[dict]
    """
    The SKU of the PowerBI Dedicated resource.
      * `name` (`str`) - Name of the SKU level.
      * `tier` (`str`) - The name of the Azure pricing tier to which the SKU applies.
    """
    tags: pulumi.Output[dict]
    """
    Key-value pairs of additional resource provisioning properties.
    """
    type: pulumi.Output[str]
    """
    The type of the PowerBI Dedicated resource.
    """
    def __init__(__self__, resource_name, opts=None, administration=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents an instance of a Dedicated Capacity resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] administration: A collection of Dedicated capacity administrators
        :param pulumi.Input[str] location: Location of the PowerBI Dedicated resource.
        :param pulumi.Input[str] name: The name of the Dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
        :param pulumi.Input[str] resource_group_name: The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
        :param pulumi.Input[dict] sku: The SKU of the PowerBI Dedicated resource.
        :param pulumi.Input[dict] tags: Key-value pairs of additional resource provisioning properties.

        The **administration** object supports the following:

          * `members` (`pulumi.Input[list]`) - An array of administrator user identities.

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - Name of the SKU level.
          * `tier` (`pulumi.Input[str]`) - The name of the Azure pricing tier to which the SKU applies.
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

            __props__['administration'] = administration
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
            __props__['properties'] = None
            __props__['type'] = None
        super(CapacityDetails, __self__).__init__(
            'azurerm:powerbidedicated/v20171001:CapacityDetails',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing CapacityDetails resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return CapacityDetails(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
