# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PrivateEndpointConnection(pulumi.CustomResource):
    e_tag: pulumi.Output[str]
    """
    Optional ETag.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name associated with the resource.
    """
    properties: pulumi.Output[dict]
    """
    PrivateEndpointConnectionResource properties
      * `private_endpoint` (`dict`) - Gets or sets private endpoint associated with the private endpoint connection
        * `id` (`str`) - Gets or sets id

      * `private_link_service_connection_state` (`dict`) - Gets or sets private link service connection state
        * `action_required` (`str`) - Gets or sets actions required
        * `description` (`str`) - Gets or sets description
        * `status` (`str`) - Gets or sets the status

      * `provisioning_state` (`str`) - Gets or sets provisioning state of the private endpoint connection
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
    """
    def __init__(__self__, resource_name, opts=None, e_tag=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, vault_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Private Endpoint Connection Response Properties

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] e_tag: Optional ETag.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the private endpoint connection.
        :param pulumi.Input[dict] properties: PrivateEndpointConnectionResource properties
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the recovery services vault is present.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[str] vault_name: The name of the recovery services vault.

        The **properties** object supports the following:

          * `private_endpoint` (`pulumi.Input[dict]`) - Gets or sets private endpoint associated with the private endpoint connection
            * `id` (`pulumi.Input[str]`) - Gets or sets id

          * `private_link_service_connection_state` (`pulumi.Input[dict]`) - Gets or sets private link service connection state
            * `action_required` (`pulumi.Input[str]`) - Gets or sets actions required
            * `description` (`pulumi.Input[str]`) - Gets or sets description
            * `status` (`pulumi.Input[str]`) - Gets or sets the status

          * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets provisioning state of the private endpoint connection
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

            __props__['e_tag'] = e_tag
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if vault_name is None:
                raise TypeError("Missing required property 'vault_name'")
            __props__['vault_name'] = vault_name
            __props__['type'] = None
        super(PrivateEndpointConnection, __self__).__init__(
            'azurerm:recoveryservices/v20200202:PrivateEndpointConnection',
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
