# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ProtectedItem(pulumi.CustomResource):
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
    ProtectedItemResource properties
      * `backup_management_type` (`str`) - Type of backup management for the backed up item.
      * `backup_set_name` (`str`) - Name of the backup set the backup item belongs to
      * `container_name` (`str`) - Unique name of container
      * `create_mode` (`str`) - Create mode to indicate recovery of existing soft deleted data source or creation of new data source.
      * `deferred_delete_time_in_utc` (`str`) - Time for deferred deletion in UTC
      * `deferred_delete_time_remaining` (`str`) - Time remaining before the DS marked for deferred delete is permanently deleted
      * `is_deferred_delete_schedule_upcoming` (`bool`) - Flag to identify whether the deferred deleted DS is to be purged soon
      * `is_rehydrate` (`bool`) - Flag to identify that deferred deleted DS is to be moved into Pause state
      * `is_scheduled_for_deferred_delete` (`bool`) - Flag to identify whether the DS is scheduled for deferred delete
      * `last_recovery_point` (`str`) - Timestamp when the last (latest) backup copy was created for this backup item.
      * `policy_id` (`str`) - ID of the backup policy with which this item is backed up.
      * `protected_item_type` (`str`) - backup item type.
      * `source_resource_id` (`str`) - ARM ID of the resource to be backed up.
      * `workload_type` (`str`) - Type of workload this item represents.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
    """
    def __init__(__self__, resource_name, opts=None, container_name=None, e_tag=None, fabric_name=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, vault_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Base class for backup items.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_name: Container name associated with the backup item.
        :param pulumi.Input[str] e_tag: Optional ETag.
        :param pulumi.Input[str] fabric_name: Fabric name associated with the backup item.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: Item name to be backed up.
        :param pulumi.Input[dict] properties: ProtectedItemResource properties
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the recovery services vault is present.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[str] vault_name: The name of the recovery services vault.

        The **properties** object supports the following:

          * `backup_management_type` (`pulumi.Input[str]`) - Type of backup management for the backed up item.
          * `backup_set_name` (`pulumi.Input[str]`) - Name of the backup set the backup item belongs to
          * `container_name` (`pulumi.Input[str]`) - Unique name of container
          * `create_mode` (`pulumi.Input[str]`) - Create mode to indicate recovery of existing soft deleted data source or creation of new data source.
          * `deferred_delete_time_in_utc` (`pulumi.Input[str]`) - Time for deferred deletion in UTC
          * `deferred_delete_time_remaining` (`pulumi.Input[str]`) - Time remaining before the DS marked for deferred delete is permanently deleted
          * `is_deferred_delete_schedule_upcoming` (`pulumi.Input[bool]`) - Flag to identify whether the deferred deleted DS is to be purged soon
          * `is_rehydrate` (`pulumi.Input[bool]`) - Flag to identify that deferred deleted DS is to be moved into Pause state
          * `is_scheduled_for_deferred_delete` (`pulumi.Input[bool]`) - Flag to identify whether the DS is scheduled for deferred delete
          * `last_recovery_point` (`pulumi.Input[str]`) - Timestamp when the last (latest) backup copy was created for this backup item.
          * `policy_id` (`pulumi.Input[str]`) - ID of the backup policy with which this item is backed up.
          * `protected_item_type` (`pulumi.Input[str]`) - backup item type.
          * `source_resource_id` (`pulumi.Input[str]`) - ARM ID of the resource to be backed up.
          * `workload_type` (`pulumi.Input[str]`) - Type of workload this item represents.
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

            if container_name is None:
                raise TypeError("Missing required property 'container_name'")
            __props__['container_name'] = container_name
            __props__['e_tag'] = e_tag
            if fabric_name is None:
                raise TypeError("Missing required property 'fabric_name'")
            __props__['fabric_name'] = fabric_name
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
        super(ProtectedItem, __self__).__init__(
            'azurerm:recoveryservices/v20190513:ProtectedItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ProtectedItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ProtectedItem(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
