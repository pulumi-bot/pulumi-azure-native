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
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    The properties of a private endpoint.
      * `private_endpoint` (`dict`) - The resource of private endpoint.
        * `id` (`str`) - The resource Id for private endpoint

      * `private_link_service_connection_state` (`dict`) - A collection of information about the state of the connection between service consumer and provider.
        * `actions_required` (`str`) - Any action that is required beyond basic workflow (approve/ reject/ disconnect)
        * `description` (`str`) - The private link service connection description.
        * `status` (`str`) - The private link service connection status.

      * `provisioning_state` (`str`) - The provisioning status of the private endpoint connection.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, config_store_name=None, name=None, private_endpoint=None, private_link_service_connection_state=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        A private endpoint connection

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_store_name: The name of the configuration store.
        :param pulumi.Input[str] name: Private endpoint connection name
        :param pulumi.Input[dict] private_endpoint: The resource of private endpoint.
        :param pulumi.Input[dict] private_link_service_connection_state: A collection of information about the state of the connection between service consumer and provider.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to which the container registry belongs.

        The **private_endpoint** object supports the following:

          * `id` (`pulumi.Input[str]`) - The resource Id for private endpoint

        The **private_link_service_connection_state** object supports the following:

          * `description` (`pulumi.Input[str]`) - The private link service connection description.
          * `status` (`pulumi.Input[str]`) - The private link service connection status.
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

            if config_store_name is None:
                raise TypeError("Missing required property 'config_store_name'")
            __props__['config_store_name'] = config_store_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['private_endpoint'] = private_endpoint
            if private_link_service_connection_state is None:
                raise TypeError("Missing required property 'private_link_service_connection_state'")
            __props__['private_link_service_connection_state'] = private_link_service_connection_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['properties'] = None
            __props__['type'] = None
        super(PrivateEndpointConnection, __self__).__init__(
            'azurerm:appconfiguration/v20200601:PrivateEndpointConnection',
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
