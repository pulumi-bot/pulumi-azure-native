# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PrivateEndpointConnection(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the resource.
      * `principal_id` (`str`) - The principal ID of resource identity.
      * `tenant_id` (`str`) - The tenant ID of resource.
      * `type` (`str`) - The identity type.
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
    Resource properties.
      * `private_endpoint` (`dict`) - The resource of private end point.
        * `id` (`str`) - The ARM identifier for Private Endpoint

      * `private_link_service_connection_state` (`dict`) - A collection of information about the state of the connection between service consumer and provider.
        * `actions_required` (`str`) - A message indicating if changes on the service provider require any updates on the consumer.
        * `description` (`str`) - The reason for approval/rejection of the connection.
        * `status` (`str`) - Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.

      * `provisioning_state` (`str`) - The provisioning state of the private endpoint connection resource.
    """
    sku: pulumi.Output[dict]
    """
    The sku of the workspace.
      * `name` (`str`) - Name of the sku
      * `tier` (`str`) - Tier of the sku like Basic or Enterprise
    """
    tags: pulumi.Output[dict]
    """
    Contains resource tags defined as key/value pairs.
    """
    type: pulumi.Output[str]
    """
    Specifies the type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, identity=None, location=None, name=None, private_link_service_connection_state=None, provisioning_state=None, resource_group_name=None, sku=None, tags=None, workspace_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The Private Endpoint Connection resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: The identity of the resource.
        :param pulumi.Input[str] location: Specifies the location of the resource.
        :param pulumi.Input[str] name: The name of the private endpoint connection associated with the workspace
        :param pulumi.Input[dict] private_link_service_connection_state: A collection of information about the state of the connection between service consumer and provider.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the private endpoint connection resource.
        :param pulumi.Input[str] resource_group_name: Name of the resource group in which workspace is located.
        :param pulumi.Input[dict] sku: The sku of the workspace.
        :param pulumi.Input[dict] tags: Contains resource tags defined as key/value pairs.
        :param pulumi.Input[str] workspace_name: Name of Azure Machine Learning workspace.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.

        The **private_link_service_connection_state** object supports the following:

          * `actions_required` (`pulumi.Input[str]`) - A message indicating if changes on the service provider require any updates on the consumer.
          * `description` (`pulumi.Input[str]`) - The reason for approval/rejection of the connection.
          * `status` (`pulumi.Input[str]`) - Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - Name of the sku
          * `tier` (`pulumi.Input[str]`) - Tier of the sku like Basic or Enterprise
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

            __props__['identity'] = identity
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if private_link_service_connection_state is None:
                raise TypeError("Missing required property 'private_link_service_connection_state'")
            __props__['private_link_service_connection_state'] = private_link_service_connection_state
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['properties'] = None
            __props__['type'] = None
        super(PrivateEndpointConnection, __self__).__init__(
            'azurerm:machinelearningservices/v20200301:PrivateEndpointConnection',
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