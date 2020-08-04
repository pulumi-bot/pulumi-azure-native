# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class RegisteredPrefix(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    The properties that define a registered prefix.
      * `error_message` (`str`) - The error message associated with the validation state, if any.
      * `peering_service_prefix_key` (`str`) - The peering service prefix key that is to be shared with the customer.
      * `prefix` (`str`) - The customer's prefix from which traffic originates.
      * `prefix_validation_state` (`str`) - The prefix validation state.
      * `provisioning_state` (`str`) - The provisioning state of the resource.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, name=None, peering_name=None, prefix=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The customer's prefix that is registered by the peering service provider.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the registered prefix.
        :param pulumi.Input[str] peering_name: The name of the peering.
        :param pulumi.Input[str] prefix: The customer's prefix from which traffic originates.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if peering_name is None:
                raise TypeError("Missing required property 'peering_name'")
            __props__['peering_name'] = peering_name
            __props__['prefix'] = prefix
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['properties'] = None
            __props__['type'] = None
        super(RegisteredPrefix, __self__).__init__(
            'azurerm:peering/v20200401:RegisteredPrefix',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing RegisteredPrefix resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RegisteredPrefix(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop