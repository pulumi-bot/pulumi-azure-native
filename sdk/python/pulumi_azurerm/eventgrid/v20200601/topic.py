# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Topic(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Location of the resource.
    """
    name: pulumi.Output[str]
    """
    Name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the topic.
      * `endpoint` (`str`) - Endpoint for the topic.
      * `inbound_ip_rules` (`list`) - This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
        * `action` (`str`) - Action to perform based on the match or no match of the IpMask.
        * `ip_mask` (`str`) - IP Address in CIDR notation e.g., 10.0.0.0/8.

      * `input_schema` (`str`) - This determines the format that Event Grid should expect for incoming events published to the topic.
      * `input_schema_mapping` (`dict`) - This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
        * `input_schema_mapping_type` (`str`) - Type of the custom mapping

      * `metric_resource_id` (`str`) - Metric resource id for the topic.
      * `private_endpoint_connections` (`list`)
        * `id` (`str`) - Fully qualified identifier of the resource.
        * `name` (`str`) - Name of the resource.
        * `properties` (`dict`) - Properties of the PrivateEndpointConnection.
          * `group_ids` (`list`) - GroupIds from the private link service resource.
          * `private_endpoint` (`dict`) - The Private Endpoint resource for this Connection.
            * `id` (`str`) - The ARM identifier for Private Endpoint.

          * `private_link_service_connection_state` (`dict`) - Details about the state of the connection.
            * `actions_required` (`str`) - Actions required (if any).
            * `description` (`str`) - Description of the connection state.
            * `status` (`str`) - Status of the connection.

          * `provisioning_state` (`str`) - Provisioning state of the Private Endpoint Connection.

        * `type` (`str`) - Type of the resource.

      * `provisioning_state` (`str`) - Provisioning state of the topic.
      * `public_network_access` (`str`) - This determines if traffic is allowed over public network. By default it is enabled. 
        You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
    """
    tags: pulumi.Output[dict]
    """
    Tags of the resource.
    """
    type: pulumi.Output[str]
    """
    Type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        EventGrid Topic

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Location of the resource.
        :param pulumi.Input[str] name: Name of the topic.
        :param pulumi.Input[dict] properties: Properties of the topic.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription.
        :param pulumi.Input[dict] tags: Tags of the resource.

        The **properties** object supports the following:

          * `inbound_ip_rules` (`pulumi.Input[list]`) - This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
            * `action` (`pulumi.Input[str]`) - Action to perform based on the match or no match of the IpMask.
            * `ip_mask` (`pulumi.Input[str]`) - IP Address in CIDR notation e.g., 10.0.0.0/8.

          * `input_schema` (`pulumi.Input[str]`) - This determines the format that Event Grid should expect for incoming events published to the topic.
          * `input_schema_mapping` (`pulumi.Input[dict]`) - This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
            * `input_schema_mapping_type` (`pulumi.Input[str]`) - Type of the custom mapping

          * `private_endpoint_connections` (`pulumi.Input[list]`)
            * `name` (`pulumi.Input[str]`) - Name of the resource.
            * `properties` (`pulumi.Input[dict]`) - Properties of the PrivateEndpointConnection.
              * `group_ids` (`pulumi.Input[list]`) - GroupIds from the private link service resource.
              * `private_endpoint` (`pulumi.Input[dict]`) - The Private Endpoint resource for this Connection.
                * `id` (`pulumi.Input[str]`) - The ARM identifier for Private Endpoint.

              * `private_link_service_connection_state` (`pulumi.Input[dict]`) - Details about the state of the connection.
                * `actions_required` (`pulumi.Input[str]`) - Actions required (if any).
                * `description` (`pulumi.Input[str]`) - Description of the connection state.
                * `status` (`pulumi.Input[str]`) - Status of the connection.

              * `provisioning_state` (`pulumi.Input[str]`) - Provisioning state of the Private Endpoint Connection.

            * `type` (`pulumi.Input[str]`) - Type of the resource.

          * `public_network_access` (`pulumi.Input[str]`) - This determines if traffic is allowed over public network. By default it is enabled. 
            You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.TopicProperties.InboundIpRules" />
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

            if location is None:
                raise TypeError("Missing required property 'location'")
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
        super(Topic, __self__).__init__(
            'azurerm:eventgrid/v20200601:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Topic(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
