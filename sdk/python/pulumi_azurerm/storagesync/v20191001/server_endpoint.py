# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ServerEndpoint(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    Server Endpoint properties.
      * `cloud_tiering` (`str`) - Cloud Tiering.
      * `cloud_tiering_status` (`dict`) - Cloud tiering status. Only populated if cloud tiering is enabled.
        * `cache_performance` (`dict`) - Information regarding how well the local cache on the server is performing.
          * `cache_hit_bytes` (`float`) - Count of bytes that were served from the local server
          * `cache_hit_bytes_percent` (`float`) - Percentage of total bytes (hit + miss) that were served from the local server
          * `cache_miss_bytes` (`float`) - Count of bytes that were served from the cloud
          * `last_updated_timestamp` (`str`) - Last updated timestamp

        * `date_policy_status` (`dict`) - Status of the date policy
          * `last_updated_timestamp` (`str`) - Last updated timestamp
          * `tiered_files_most_recent_access_timestamp` (`str`) - Most recent access time of tiered files

        * `files_not_tiering` (`dict`) - Information regarding files that failed to be tiered
          * `errors` (`list`) - Array of tiering errors
            * `error_code` (`float`) - Error code (HResult)
            * `file_count` (`float`) - Count of files with this error

          * `last_updated_timestamp` (`str`) - Last updated timestamp
          * `total_file_count` (`float`) - Last cloud tiering result (HResult)

        * `health` (`str`) - Cloud tiering health state.
        * `health_last_updated_timestamp` (`str`) - The last updated timestamp of health state
        * `last_cloud_tiering_result` (`float`) - Last cloud tiering result (HResult)
        * `last_success_timestamp` (`str`) - Last cloud tiering success timestamp
        * `last_updated_timestamp` (`str`) - Last updated timestamp
        * `space_savings` (`dict`) - Information regarding how much local space cloud tiering is saving.
          * `cached_size_bytes` (`float`) - Cached content size on the server
          * `last_updated_timestamp` (`str`) - Last updated timestamp
          * `space_savings_bytes` (`float`) - Count of bytes saved on the server
          * `space_savings_percent` (`float`) - Percentage of cached size over total size
          * `total_size_cloud_bytes` (`float`) - Total size of content in the azure file share
          * `volume_size_bytes` (`float`) - Volume size

        * `volume_free_space_policy_status` (`dict`) - Status of the volume free space policy
          * `current_volume_free_space_percent` (`float`) - Current volume free space percentage.
          * `effective_volume_free_space_policy` (`float`) - In the case where multiple server endpoints are present in a volume, an effective free space policy is applied.
          * `last_updated_timestamp` (`str`) - Last updated timestamp

      * `friendly_name` (`str`) - Friendly Name
      * `last_operation_name` (`str`) - Resource Last Operation Name
      * `last_workflow_id` (`str`) - ServerEndpoint lastWorkflowId
      * `offline_data_transfer` (`str`) - Offline data transfer
      * `offline_data_transfer_share_name` (`str`) - Offline data transfer share name
      * `offline_data_transfer_storage_account_resource_id` (`str`) - Offline data transfer storage account resource ID
      * `offline_data_transfer_storage_account_tenant_id` (`str`) - Offline data transfer storage account tenant ID
      * `provisioning_state` (`str`) - ServerEndpoint Provisioning State
      * `recall_status` (`dict`) - Recall status. Only populated if cloud tiering is enabled.
        * `last_updated_timestamp` (`str`) - Last updated timestamp
        * `recall_errors` (`list`) - Array of recall errors
          * `count` (`float`) - Count of occurences of the error
          * `error_code` (`float`) - Error code (HResult)

        * `total_recall_errors_count` (`float`) - Total count of recall errors.

      * `server_local_path` (`str`) - Server Local path.
      * `server_resource_id` (`str`) - Server Resource Id.
      * `sync_status` (`dict`) - Server Endpoint sync status
        * `combined_health` (`str`) - Combined Health Status.
        * `download_activity` (`dict`) - Download sync activity
          * `applied_bytes` (`float`) - Applied bytes
          * `applied_item_count` (`float`) - Applied item count.
          * `per_item_error_count` (`float`) - Per item error count
          * `timestamp` (`str`) - Timestamp when properties were updated
          * `total_bytes` (`float`) - Total bytes (if available)
          * `total_item_count` (`float`) - Total item count (if available)

        * `download_health` (`str`) - Download Health Status.
        * `download_status` (`dict`) - Download Status
          * `files_not_syncing_errors` (`list`) - Array of per-item errors coming from the last sync session.
            * `error_code` (`float`) - Error code (HResult)
            * `persistent_count` (`float`) - Count of persistent files not syncing with the specified error code
            * `transient_count` (`float`) - Count of transient files not syncing with the specified error code

          * `last_sync_per_item_error_count` (`float`) - Last sync per item error count.
          * `last_sync_result` (`float`) - Last sync result (HResult)
          * `last_sync_success_timestamp` (`str`) - Last sync success timestamp
          * `last_sync_timestamp` (`str`) - Last sync timestamp
          * `persistent_files_not_syncing_count` (`float`) - Count of persistent files not syncing.
          * `transient_files_not_syncing_count` (`float`) - Count of transient files not syncing.

        * `last_updated_timestamp` (`str`) - Last Updated Timestamp
        * `offline_data_transfer_status` (`str`) - Offline Data Transfer State
        * `sync_activity` (`str`) - Sync activity
        * `total_persistent_files_not_syncing_count` (`float`) - Total count of persistent files not syncing (combined upload + download).
        * `upload_activity` (`dict`) - Upload sync activity
        * `upload_health` (`str`) - Upload Health Status.
        * `upload_status` (`dict`) - Upload Status

      * `tier_files_older_than_days` (`float`) - Tier files older than days.
      * `volume_free_space_percent` (`float`) - Level of free space to be maintained by Cloud Tiering if it is enabled.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, storage_sync_service_name=None, sync_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Server Endpoint object.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of Server Endpoint object.
        :param pulumi.Input[dict] properties: The parameters used to create the server endpoint.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] storage_sync_service_name: Name of Storage Sync Service resource.
        :param pulumi.Input[str] sync_group_name: Name of Sync Group resource.

        The **properties** object supports the following:

          * `cloud_tiering` (`pulumi.Input[str]`) - Cloud Tiering.
          * `friendly_name` (`pulumi.Input[str]`) - Friendly Name
          * `offline_data_transfer` (`pulumi.Input[str]`) - Offline data transfer
          * `offline_data_transfer_share_name` (`pulumi.Input[str]`) - Offline data transfer share name
          * `server_local_path` (`pulumi.Input[str]`) - Server Local path.
          * `server_resource_id` (`pulumi.Input[str]`) - Server Resource Id.
          * `tier_files_older_than_days` (`pulumi.Input[float]`) - Tier files older than days.
          * `volume_free_space_percent` (`pulumi.Input[float]`) - Level of free space to be maintained by Cloud Tiering if it is enabled.
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
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if storage_sync_service_name is None:
                raise TypeError("Missing required property 'storage_sync_service_name'")
            __props__['storage_sync_service_name'] = storage_sync_service_name
            if sync_group_name is None:
                raise TypeError("Missing required property 'sync_group_name'")
            __props__['sync_group_name'] = sync_group_name
            __props__['type'] = None
        super(ServerEndpoint, __self__).__init__(
            'azurerm:storagesync/v20191001:ServerEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ServerEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ServerEndpoint(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
