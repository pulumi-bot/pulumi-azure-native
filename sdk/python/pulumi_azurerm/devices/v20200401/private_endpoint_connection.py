# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PrivateEndpointConnection(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The resource name.
    """
    properties: pulumi.Output[dict]
    """
    The properties of a private endpoint connection
      * `private_endpoint` (`dict`) - The private endpoint property of a private endpoint connection
        * `id` (`str`) - The resource identifier.

      * `private_link_service_connection_state` (`dict`) - The current state of a private endpoint connection
        * `actions_required` (`str`) - Actions required for a private endpoint connection
        * `description` (`str`) - The description for the current state of a private endpoint connection
        * `status` (`str`) - The status of a private endpoint connection
    """
    type: pulumi.Output[str]
    """
    The resource type.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, resource_name_=None, __props__=None, __name__=None, __opts__=None):
        """
        The private endpoint connection of an IotHub

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the private endpoint connection
        :param pulumi.Input[dict] properties: The properties of a private endpoint connection
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the IoT hub.
        :param pulumi.Input[str] resource_name_: The name of the IoT hub.

        The **properties** object supports the following:

          * `private_endpoint` (`pulumi.Input[dict]`) - The private endpoint property of a private endpoint connection
          * `private_link_service_connection_state` (`pulumi.Input[dict]`) - The current state of a private endpoint connection
            * `actions_required` (`pulumi.Input[str]`) - Actions required for a private endpoint connection
            * `description` (`pulumi.Input[str]`) - The description for the current state of a private endpoint connection
            * `status` (`pulumi.Input[str]`) - The status of a private endpoint connection
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
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['type'] = None
        super(PrivateEndpointConnection, __self__).__init__(
            'azurerm:devices/v20200401:PrivateEndpointConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing PrivateEndpointConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PrivateEndpointConnection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
