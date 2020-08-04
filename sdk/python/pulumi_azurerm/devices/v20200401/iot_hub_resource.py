# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class IotHubResource(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
    """
    location: pulumi.Output[str]
    """
    The resource location.
    """
    name: pulumi.Output[str]
    """
    The resource name.
    """
    properties: pulumi.Output[dict]
    """
    IotHub properties
      * `authorization_policies` (`list`) - The shared access policies you can use to secure a connection to the IoT hub.
        * `key_name` (`str`) - The name of the shared access policy.
        * `primary_key` (`str`) - The primary key.
        * `rights` (`str`) - The permissions assigned to the shared access policy.
        * `secondary_key` (`str`) - The secondary key.

      * `cloud_to_device` (`dict`) - The IoT hub cloud-to-device messaging properties.
        * `default_ttl_as_iso8601` (`str`) - The default time to live for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        * `feedback` (`dict`) - The properties of the feedback queue for cloud-to-device messages.
          * `lock_duration_as_iso8601` (`str`) - The lock duration for the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
          * `max_delivery_count` (`float`) - The number of times the IoT hub attempts to deliver a message on the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
          * `ttl_as_iso8601` (`str`) - The period of time for which a message is available to consume before it is expired by the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.

        * `max_delivery_count` (`float`) - The max delivery count for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.

      * `comments` (`str`) - IoT hub comments.
      * `enable_file_upload_notifications` (`bool`) - If True, file upload notifications are enabled.
      * `event_hub_endpoints` (`dict`) - The Event Hub-compatible endpoint properties. The only possible keys to this dictionary is events. This key has to be present in the dictionary while making create or update calls for the IoT hub.
      * `features` (`str`) - The capabilities and features enabled for the IoT hub.
      * `host_name` (`str`) - The name of the host.
      * `ip_filter_rules` (`list`) - The IP filter rules.
        * `action` (`str`) - The desired action for requests captured by this rule.
        * `filter_name` (`str`) - The name of the IP filter rule.
        * `ip_mask` (`str`) - A string that contains the IP address range in CIDR notation for the rule.

      * `locations` (`list`) - Primary and secondary location for iot hub
        * `location` (`str`) - The name of the Azure region
        * `role` (`str`) - The role of the region, can be either primary or secondary. The primary region is where the IoT hub is currently provisioned. The secondary region is the Azure disaster recovery (DR) paired region and also the region where the IoT hub can failover to.

      * `messaging_endpoints` (`dict`) - The messaging endpoint properties for the file upload notification queue.
      * `min_tls_version` (`str`) - Specifies the minimum TLS version to support for this hub. Can be set to "1.2" to have clients that use a TLS version below 1.2 to be rejected.
      * `private_endpoint_connections` (`list`) - Private endpoint connections created on this IotHub
        * `id` (`str`) - The resource identifier.
        * `name` (`str`) - The resource name.
        * `properties` (`dict`) - The properties of a private endpoint connection
          * `private_endpoint` (`dict`) - The private endpoint property of a private endpoint connection
            * `id` (`str`) - The resource identifier.

          * `private_link_service_connection_state` (`dict`) - The current state of a private endpoint connection
            * `actions_required` (`str`) - Actions required for a private endpoint connection
            * `description` (`str`) - The description for the current state of a private endpoint connection
            * `status` (`str`) - The status of a private endpoint connection

        * `type` (`str`) - The resource type.

      * `provisioning_state` (`str`) - The provisioning state.
      * `public_network_access` (`str`) - Whether requests from Public Network are allowed
      * `routing` (`dict`) - The routing related properties of the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging
        * `endpoints` (`dict`) - The properties related to the custom endpoints to which your IoT hub routes messages based on the routing rules. A maximum of 10 custom endpoints are allowed across all endpoint types for paid hubs and only 1 custom endpoint is allowed across all endpoint types for free hubs.
          * `event_hubs` (`list`) - The list of Event Hubs endpoints that IoT hub routes messages to, based on the routing rules. This list does not include the built-in Event Hubs endpoint.
            * `authentication_type` (`str`) - Method used to authenticate against the event hub endpoint
            * `connection_string` (`str`) - The connection string of the event hub endpoint. 
            * `endpoint_uri` (`str`) - The url of the event hub endpoint. It must include the protocol sb://
            * `entity_path` (`str`) - Event hub name on the event hub namespace
            * `id` (`str`) - Id of the event hub endpoint
            * `name` (`str`) - The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types.
            * `resource_group` (`str`) - The name of the resource group of the event hub endpoint.
            * `subscription_id` (`str`) - The subscription identifier of the event hub endpoint.

          * `service_bus_queues` (`list`) - The list of Service Bus queue endpoints that IoT hub routes the messages to, based on the routing rules.
            * `authentication_type` (`str`) - Method used to authenticate against the service bus queue endpoint
            * `connection_string` (`str`) - The connection string of the service bus queue endpoint.
            * `endpoint_uri` (`str`) - The url of the service bus queue endpoint. It must include the protocol sb://
            * `entity_path` (`str`) - Queue name on the service bus namespace
            * `id` (`str`) - Id of the service bus queue endpoint
            * `name` (`str`) - The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types. The name need not be the same as the actual queue name.
            * `resource_group` (`str`) - The name of the resource group of the service bus queue endpoint.
            * `subscription_id` (`str`) - The subscription identifier of the service bus queue endpoint.

          * `service_bus_topics` (`list`) - The list of Service Bus topic endpoints that the IoT hub routes the messages to, based on the routing rules.
            * `authentication_type` (`str`) - Method used to authenticate against the service bus topic endpoint
            * `connection_string` (`str`) - The connection string of the service bus topic endpoint.
            * `endpoint_uri` (`str`) - The url of the service bus topic endpoint. It must include the protocol sb://
            * `entity_path` (`str`) - Queue name on the service bus topic
            * `id` (`str`) - Id of the service bus topic endpoint
            * `name` (`str`) - The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types.  The name need not be the same as the actual topic name.
            * `resource_group` (`str`) - The name of the resource group of the service bus topic endpoint.
            * `subscription_id` (`str`) - The subscription identifier of the service bus topic endpoint.

          * `storage_containers` (`list`) - The list of storage container endpoints that IoT hub routes messages to, based on the routing rules.
            * `authentication_type` (`str`) - Method used to authenticate against the storage endpoint
            * `batch_frequency_in_seconds` (`float`) - Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
            * `connection_string` (`str`) - The connection string of the storage account.
            * `container_name` (`str`) - The name of storage container in the storage account.
            * `encoding` (`str`) - Encoding that is used to serialize messages to blobs. Supported values are 'avro', 'avrodeflate', and 'JSON'. Default value is 'avro'.
            * `endpoint_uri` (`str`) - The url of the storage endpoint. It must include the protocol https://
            * `file_name_format` (`str`) - File name format for the blob. Default format is {iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}. All parameters are mandatory but can be reordered.
            * `id` (`str`) - Id of the storage container endpoint
            * `max_chunk_size_in_bytes` (`float`) - Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
            * `name` (`str`) - The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types.
            * `resource_group` (`str`) - The name of the resource group of the storage account.
            * `subscription_id` (`str`) - The subscription identifier of the storage account.

        * `enrichments` (`list`) - The list of user-provided enrichments that the IoT hub applies to messages to be delivered to built-in and custom endpoints. See: https://aka.ms/telemetryoneventgrid
          * `endpoint_names` (`list`) - The list of endpoints for which the enrichment is applied to the message.
          * `key` (`str`) - The key or name for the enrichment property.
          * `value` (`str`) - The value for the enrichment property.

        * `fallback_route` (`dict`) - The properties of the route that is used as a fall-back route when none of the conditions specified in the 'routes' section are met. This is an optional parameter. When this property is not set, the messages which do not meet any of the conditions specified in the 'routes' section get routed to the built-in eventhub endpoint.
          * `condition` (`str`) - The condition which is evaluated in order to apply the fallback route. If the condition is not provided it will evaluate to true by default. For grammar, See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language
          * `endpoint_names` (`list`) - The list of endpoints to which the messages that satisfy the condition are routed to. Currently only 1 endpoint is allowed.
          * `is_enabled` (`bool`) - Used to specify whether the fallback route is enabled.
          * `name` (`str`) - The name of the route. The name can only include alphanumeric characters, periods, underscores, hyphens, has a maximum length of 64 characters, and must be unique.
          * `source` (`str`) - The source to which the routing rule is to be applied to. For example, DeviceMessages

        * `routes` (`list`) - The list of user-provided routing rules that the IoT hub uses to route messages to built-in and custom endpoints. A maximum of 100 routing rules are allowed for paid hubs and a maximum of 5 routing rules are allowed for free hubs.
          * `condition` (`str`) - The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language
          * `endpoint_names` (`list`) - The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
          * `is_enabled` (`bool`) - Used to specify whether a route is enabled.
          * `name` (`str`) - The name of the route. The name can only include alphanumeric characters, periods, underscores, hyphens, has a maximum length of 64 characters, and must be unique.
          * `source` (`str`) - The source that the routing rule is to be applied to, such as DeviceMessages.

      * `state` (`str`) - The hub state.
      * `storage_endpoints` (`dict`) - The list of Azure Storage endpoints where you can upload files. Currently you can configure only one Azure Storage account and that MUST have its key as $default. Specifying more than one storage account causes an error to be thrown. Not specifying a value for this property when the enableFileUploadNotifications property is set to True, causes an error to be thrown.
    """
    sku: pulumi.Output[dict]
    """
    IotHub SKU info
      * `capacity` (`float`) - The number of provisioned IoT Hub units. See: https://docs.microsoft.com/azure/azure-subscription-service-limits#iot-hub-limits.
      * `name` (`str`) - The name of the SKU.
      * `tier` (`str`) - The billing tier for the IoT hub.
    """
    tags: pulumi.Output[dict]
    """
    The resource tags.
    """
    type: pulumi.Output[str]
    """
    The resource type.
    """
    def __init__(__self__, resource_name, opts=None, etag=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The description of the IoT hub.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[str] name: The name of the IoT hub.
        :param pulumi.Input[dict] properties: IotHub properties
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the IoT hub.
        :param pulumi.Input[dict] sku: IotHub SKU info
        :param pulumi.Input[dict] tags: The resource tags.

        The **properties** object supports the following:

          * `authorization_policies` (`pulumi.Input[list]`) - The shared access policies you can use to secure a connection to the IoT hub.
            * `key_name` (`pulumi.Input[str]`) - The name of the shared access policy.
            * `primary_key` (`pulumi.Input[str]`) - The primary key.
            * `rights` (`pulumi.Input[str]`) - The permissions assigned to the shared access policy.
            * `secondary_key` (`pulumi.Input[str]`) - The secondary key.

          * `cloud_to_device` (`pulumi.Input[dict]`) - The IoT hub cloud-to-device messaging properties.
            * `default_ttl_as_iso8601` (`pulumi.Input[str]`) - The default time to live for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
            * `feedback` (`pulumi.Input[dict]`) - The properties of the feedback queue for cloud-to-device messages.
              * `lock_duration_as_iso8601` (`pulumi.Input[str]`) - The lock duration for the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
              * `max_delivery_count` (`pulumi.Input[float]`) - The number of times the IoT hub attempts to deliver a message on the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
              * `ttl_as_iso8601` (`pulumi.Input[str]`) - The period of time for which a message is available to consume before it is expired by the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.

            * `max_delivery_count` (`pulumi.Input[float]`) - The max delivery count for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.

          * `comments` (`pulumi.Input[str]`) - IoT hub comments.
          * `enable_file_upload_notifications` (`pulumi.Input[bool]`) - If True, file upload notifications are enabled.
          * `event_hub_endpoints` (`pulumi.Input[dict]`) - The Event Hub-compatible endpoint properties. The only possible keys to this dictionary is events. This key has to be present in the dictionary while making create or update calls for the IoT hub.
          * `features` (`pulumi.Input[str]`) - The capabilities and features enabled for the IoT hub.
          * `ip_filter_rules` (`pulumi.Input[list]`) - The IP filter rules.
            * `action` (`pulumi.Input[str]`) - The desired action for requests captured by this rule.
            * `filter_name` (`pulumi.Input[str]`) - The name of the IP filter rule.
            * `ip_mask` (`pulumi.Input[str]`) - A string that contains the IP address range in CIDR notation for the rule.

          * `messaging_endpoints` (`pulumi.Input[dict]`) - The messaging endpoint properties for the file upload notification queue.
          * `min_tls_version` (`pulumi.Input[str]`) - Specifies the minimum TLS version to support for this hub. Can be set to "1.2" to have clients that use a TLS version below 1.2 to be rejected.
          * `private_endpoint_connections` (`pulumi.Input[list]`) - Private endpoint connections created on this IotHub
            * `name` (`pulumi.Input[str]`) - The resource name.
            * `properties` (`pulumi.Input[dict]`) - The properties of a private endpoint connection
              * `private_endpoint` (`pulumi.Input[dict]`) - The private endpoint property of a private endpoint connection
                * `id` (`pulumi.Input[str]`) - The resource identifier.

              * `private_link_service_connection_state` (`pulumi.Input[dict]`) - The current state of a private endpoint connection
                * `actions_required` (`pulumi.Input[str]`) - Actions required for a private endpoint connection
                * `description` (`pulumi.Input[str]`) - The description for the current state of a private endpoint connection
                * `status` (`pulumi.Input[str]`) - The status of a private endpoint connection

            * `type` (`pulumi.Input[str]`) - The resource type.

          * `public_network_access` (`pulumi.Input[str]`) - Whether requests from Public Network are allowed
          * `routing` (`pulumi.Input[dict]`) - The routing related properties of the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging
            * `endpoints` (`pulumi.Input[dict]`) - The properties related to the custom endpoints to which your IoT hub routes messages based on the routing rules. A maximum of 10 custom endpoints are allowed across all endpoint types for paid hubs and only 1 custom endpoint is allowed across all endpoint types for free hubs.
              * `event_hubs` (`pulumi.Input[list]`) - The list of Event Hubs endpoints that IoT hub routes messages to, based on the routing rules. This list does not include the built-in Event Hubs endpoint.
                * `authentication_type` (`pulumi.Input[str]`) - Method used to authenticate against the event hub endpoint
                * `connection_string` (`pulumi.Input[str]`) - The connection string of the event hub endpoint. 
                * `endpoint_uri` (`pulumi.Input[str]`) - The url of the event hub endpoint. It must include the protocol sb://
                * `entity_path` (`pulumi.Input[str]`) - Event hub name on the event hub namespace
                * `id` (`pulumi.Input[str]`) - Id of the event hub endpoint
                * `name` (`pulumi.Input[str]`) - The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types.
                * `resource_group` (`pulumi.Input[str]`) - The name of the resource group of the event hub endpoint.
                * `subscription_id` (`pulumi.Input[str]`) - The subscription identifier of the event hub endpoint.

              * `service_bus_queues` (`pulumi.Input[list]`) - The list of Service Bus queue endpoints that IoT hub routes the messages to, based on the routing rules.
                * `authentication_type` (`pulumi.Input[str]`) - Method used to authenticate against the service bus queue endpoint
                * `connection_string` (`pulumi.Input[str]`) - The connection string of the service bus queue endpoint.
                * `endpoint_uri` (`pulumi.Input[str]`) - The url of the service bus queue endpoint. It must include the protocol sb://
                * `entity_path` (`pulumi.Input[str]`) - Queue name on the service bus namespace
                * `id` (`pulumi.Input[str]`) - Id of the service bus queue endpoint
                * `name` (`pulumi.Input[str]`) - The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types. The name need not be the same as the actual queue name.
                * `resource_group` (`pulumi.Input[str]`) - The name of the resource group of the service bus queue endpoint.
                * `subscription_id` (`pulumi.Input[str]`) - The subscription identifier of the service bus queue endpoint.

              * `service_bus_topics` (`pulumi.Input[list]`) - The list of Service Bus topic endpoints that the IoT hub routes the messages to, based on the routing rules.
                * `authentication_type` (`pulumi.Input[str]`) - Method used to authenticate against the service bus topic endpoint
                * `connection_string` (`pulumi.Input[str]`) - The connection string of the service bus topic endpoint.
                * `endpoint_uri` (`pulumi.Input[str]`) - The url of the service bus topic endpoint. It must include the protocol sb://
                * `entity_path` (`pulumi.Input[str]`) - Queue name on the service bus topic
                * `id` (`pulumi.Input[str]`) - Id of the service bus topic endpoint
                * `name` (`pulumi.Input[str]`) - The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types.  The name need not be the same as the actual topic name.
                * `resource_group` (`pulumi.Input[str]`) - The name of the resource group of the service bus topic endpoint.
                * `subscription_id` (`pulumi.Input[str]`) - The subscription identifier of the service bus topic endpoint.

              * `storage_containers` (`pulumi.Input[list]`) - The list of storage container endpoints that IoT hub routes messages to, based on the routing rules.
                * `authentication_type` (`pulumi.Input[str]`) - Method used to authenticate against the storage endpoint
                * `batch_frequency_in_seconds` (`pulumi.Input[float]`) - Time interval at which blobs are written to storage. Value should be between 60 and 720 seconds. Default value is 300 seconds.
                * `connection_string` (`pulumi.Input[str]`) - The connection string of the storage account.
                * `container_name` (`pulumi.Input[str]`) - The name of storage container in the storage account.
                * `encoding` (`pulumi.Input[str]`) - Encoding that is used to serialize messages to blobs. Supported values are 'avro', 'avrodeflate', and 'JSON'. Default value is 'avro'.
                * `endpoint_uri` (`pulumi.Input[str]`) - The url of the storage endpoint. It must include the protocol https://
                * `file_name_format` (`pulumi.Input[str]`) - File name format for the blob. Default format is {iothub}/{partition}/{YYYY}/{MM}/{DD}/{HH}/{mm}. All parameters are mandatory but can be reordered.
                * `id` (`pulumi.Input[str]`) - Id of the storage container endpoint
                * `max_chunk_size_in_bytes` (`pulumi.Input[float]`) - Maximum number of bytes for each blob written to storage. Value should be between 10485760(10MB) and 524288000(500MB). Default value is 314572800(300MB).
                * `name` (`pulumi.Input[str]`) - The name that identifies this endpoint. The name can only include alphanumeric characters, periods, underscores, hyphens and has a maximum length of 64 characters. The following names are reserved:  events, fileNotifications, $default. Endpoint names must be unique across endpoint types.
                * `resource_group` (`pulumi.Input[str]`) - The name of the resource group of the storage account.
                * `subscription_id` (`pulumi.Input[str]`) - The subscription identifier of the storage account.

            * `enrichments` (`pulumi.Input[list]`) - The list of user-provided enrichments that the IoT hub applies to messages to be delivered to built-in and custom endpoints. See: https://aka.ms/telemetryoneventgrid
              * `endpoint_names` (`pulumi.Input[list]`) - The list of endpoints for which the enrichment is applied to the message.
              * `key` (`pulumi.Input[str]`) - The key or name for the enrichment property.
              * `value` (`pulumi.Input[str]`) - The value for the enrichment property.

            * `fallback_route` (`pulumi.Input[dict]`) - The properties of the route that is used as a fall-back route when none of the conditions specified in the 'routes' section are met. This is an optional parameter. When this property is not set, the messages which do not meet any of the conditions specified in the 'routes' section get routed to the built-in eventhub endpoint.
              * `condition` (`pulumi.Input[str]`) - The condition which is evaluated in order to apply the fallback route. If the condition is not provided it will evaluate to true by default. For grammar, See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language
              * `endpoint_names` (`pulumi.Input[list]`) - The list of endpoints to which the messages that satisfy the condition are routed to. Currently only 1 endpoint is allowed.
              * `is_enabled` (`pulumi.Input[bool]`) - Used to specify whether the fallback route is enabled.
              * `name` (`pulumi.Input[str]`) - The name of the route. The name can only include alphanumeric characters, periods, underscores, hyphens, has a maximum length of 64 characters, and must be unique.
              * `source` (`pulumi.Input[str]`) - The source to which the routing rule is to be applied to. For example, DeviceMessages

            * `routes` (`pulumi.Input[list]`) - The list of user-provided routing rules that the IoT hub uses to route messages to built-in and custom endpoints. A maximum of 100 routing rules are allowed for paid hubs and a maximum of 5 routing rules are allowed for free hubs.
              * `condition` (`pulumi.Input[str]`) - The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language
              * `endpoint_names` (`pulumi.Input[list]`) - The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
              * `is_enabled` (`pulumi.Input[bool]`) - Used to specify whether a route is enabled.
              * `name` (`pulumi.Input[str]`) - The name of the route. The name can only include alphanumeric characters, periods, underscores, hyphens, has a maximum length of 64 characters, and must be unique.
              * `source` (`pulumi.Input[str]`) - The source that the routing rule is to be applied to, such as DeviceMessages.

          * `storage_endpoints` (`pulumi.Input[dict]`) - The list of Azure Storage endpoints where you can upload files. Currently you can configure only one Azure Storage account and that MUST have its key as $default. Specifying more than one storage account causes an error to be thrown. Not specifying a value for this property when the enableFileUploadNotifications property is set to True, causes an error to be thrown.

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - The number of provisioned IoT Hub units. See: https://docs.microsoft.com/azure/azure-subscription-service-limits#iot-hub-limits.
          * `name` (`pulumi.Input[str]`) - The name of the SKU.
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
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['type'] = None
        super(IotHubResource, __self__).__init__(
            'azurerm:devices/v20200401:IotHubResource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing IotHubResource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IotHubResource(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
