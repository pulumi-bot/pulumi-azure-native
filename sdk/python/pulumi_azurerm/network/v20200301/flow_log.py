# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class FlowLog(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the flow log.
      * `enabled` (`bool`) - Flag to enable/disable flow logging.
      * `flow_analytics_configuration` (`dict`) - Parameters that define the configuration of traffic analytics.
        * `network_watcher_flow_analytics_configuration` (`dict`) - Parameters that define the configuration of traffic analytics.
          * `enabled` (`bool`) - Flag to enable/disable traffic analytics.
          * `traffic_analytics_interval` (`float`) - The interval in minutes which would decide how frequently TA service should do flow analytics.
          * `workspace_id` (`str`) - The resource guid of the attached workspace.
          * `workspace_region` (`str`) - The location of the attached workspace.
          * `workspace_resource_id` (`str`) - Resource Id of the attached workspace.

      * `format` (`dict`) - Parameters that define the flow log format.
        * `type` (`str`) - The file type of flow log.
        * `version` (`float`) - The version (revision) of the flow log.

      * `provisioning_state` (`str`) - The provisioning state of the flow log.
      * `retention_policy` (`dict`) - Parameters that define the retention policy for flow log.
        * `days` (`float`) - Number of days to retain flow log records.
        * `enabled` (`bool`) - Flag to enable/disable retention.

      * `storage_id` (`str`) - ID of the storage account which is used to store the flow log.
      * `target_resource_guid` (`str`) - Guid of network security group to which flow log will be applied.
      * `target_resource_id` (`str`) - ID of network security group to which flow log will be applied.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, id=None, location=None, name=None, network_watcher_name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        A flow log resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the flow log.
        :param pulumi.Input[str] network_watcher_name: The name of the network watcher.
        :param pulumi.Input[dict] properties: Properties of the flow log.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Flag to enable/disable flow logging.
          * `flow_analytics_configuration` (`pulumi.Input[dict]`) - Parameters that define the configuration of traffic analytics.
            * `network_watcher_flow_analytics_configuration` (`pulumi.Input[dict]`) - Parameters that define the configuration of traffic analytics.
              * `enabled` (`pulumi.Input[bool]`) - Flag to enable/disable traffic analytics.
              * `traffic_analytics_interval` (`pulumi.Input[float]`) - The interval in minutes which would decide how frequently TA service should do flow analytics.
              * `workspace_id` (`pulumi.Input[str]`) - The resource guid of the attached workspace.
              * `workspace_region` (`pulumi.Input[str]`) - The location of the attached workspace.
              * `workspace_resource_id` (`pulumi.Input[str]`) - Resource Id of the attached workspace.

          * `format` (`pulumi.Input[dict]`) - Parameters that define the flow log format.
            * `type` (`pulumi.Input[str]`) - The file type of flow log.
            * `version` (`pulumi.Input[float]`) - The version (revision) of the flow log.

          * `retention_policy` (`pulumi.Input[dict]`) - Parameters that define the retention policy for flow log.
            * `days` (`pulumi.Input[float]`) - Number of days to retain flow log records.
            * `enabled` (`pulumi.Input[bool]`) - Flag to enable/disable retention.

          * `storage_id` (`pulumi.Input[str]`) - ID of the storage account which is used to store the flow log.
          * `target_resource_id` (`pulumi.Input[str]`) - ID of network security group to which flow log will be applied.
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

            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if network_watcher_name is None:
                raise TypeError("Missing required property 'network_watcher_name'")
            __props__['network_watcher_name'] = network_watcher_name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['type'] = None
        super(FlowLog, __self__).__init__(
            'azurerm:network/v20200301:FlowLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing FlowLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return FlowLog(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
