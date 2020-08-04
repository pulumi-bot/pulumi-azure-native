# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Share(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The object name.
    """
    properties: pulumi.Output[dict]
    """
    The share properties.
      * `access_protocol` (`str`) - Access protocol to be used by the share.
      * `azure_container_info` (`dict`) - Azure container mapping for the share.
        * `container_name` (`str`) - Container name (Based on the data format specified, this represents the name of Azure Files/Page blob/Block blob).
        * `data_format` (`str`) - Storage format used for the file represented by the share.
        * `storage_account_credential_id` (`str`) - ID of the storage account credential used to access storage.

      * `client_access_rights` (`list`) - List of IP addresses and corresponding access rights on the share(required for NFS protocol).
        * `access_permission` (`str`) - Type of access to be allowed for the client.
        * `client` (`str`) - IP of the client.

      * `data_policy` (`str`) - Data policy of the share.
      * `description` (`str`) - Description for the share.
      * `monitoring_status` (`str`) - Current monitoring status of the share.
      * `refresh_details` (`dict`) - Details of the refresh job on this share.
        * `error_manifest_file` (`str`) - Indicates the relative path of the error xml for the last refresh job on this particular share or container, if any. This could be a failed job or a successful job.
        * `in_progress_refresh_job_id` (`str`) - If a refresh job is currently in progress on this share or container, this field indicates the ARM resource ID of that job. The field is empty if no job is in progress.
        * `last_completed_refresh_job_time_in_utc` (`str`) - Indicates the completed time for the last refresh job on this particular share or container, if any.This could be a failed job or a successful job.
        * `last_job` (`str`) - Indicates the id of the last refresh job on this particular share or container,if any. This could be a failed job or a successful job.

      * `share_mappings` (`list`) - Share mount point to the role.
        * `mount_point` (`str`) - Mount point for the share.
        * `role_id` (`str`) - ID of the role to which share is mounted.
        * `role_type` (`str`) - Role type.
        * `share_id` (`str`) - ID of the share mounted to the role VM.

      * `share_status` (`str`) - Current status of the share.
      * `user_access_rights` (`list`) - Mapping of users and corresponding access rights on the share (required for SMB protocol).
        * `access_type` (`str`) - Type of access to be allowed for the user.
        * `user_id` (`str`) - User ID (already existing in the device).
    """
    type: pulumi.Output[str]
    """
    The hierarchical type of the object.
    """
    def __init__(__self__, resource_name, opts=None, access_protocol=None, azure_container_info=None, client_access_rights=None, data_policy=None, description=None, device_name=None, monitoring_status=None, name=None, refresh_details=None, resource_group_name=None, share_status=None, user_access_rights=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a share on the  Data Box Edge/Gateway device.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_protocol: Access protocol to be used by the share.
        :param pulumi.Input[dict] azure_container_info: Azure container mapping for the share.
        :param pulumi.Input[list] client_access_rights: List of IP addresses and corresponding access rights on the share(required for NFS protocol).
        :param pulumi.Input[str] data_policy: Data policy of the share.
        :param pulumi.Input[str] description: Description for the share.
        :param pulumi.Input[str] device_name: The device name.
        :param pulumi.Input[str] monitoring_status: Current monitoring status of the share.
        :param pulumi.Input[str] name: The share name.
        :param pulumi.Input[dict] refresh_details: Details of the refresh job on this share.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[str] share_status: Current status of the share.
        :param pulumi.Input[list] user_access_rights: Mapping of users and corresponding access rights on the share (required for SMB protocol).

        The **azure_container_info** object supports the following:

          * `container_name` (`pulumi.Input[str]`) - Container name (Based on the data format specified, this represents the name of Azure Files/Page blob/Block blob).
          * `data_format` (`pulumi.Input[str]`) - Storage format used for the file represented by the share.
          * `storage_account_credential_id` (`pulumi.Input[str]`) - ID of the storage account credential used to access storage.

        The **client_access_rights** object supports the following:

          * `access_permission` (`pulumi.Input[str]`) - Type of access to be allowed for the client.
          * `client` (`pulumi.Input[str]`) - IP of the client.

        The **refresh_details** object supports the following:

          * `error_manifest_file` (`pulumi.Input[str]`) - Indicates the relative path of the error xml for the last refresh job on this particular share or container, if any. This could be a failed job or a successful job.
          * `in_progress_refresh_job_id` (`pulumi.Input[str]`) - If a refresh job is currently in progress on this share or container, this field indicates the ARM resource ID of that job. The field is empty if no job is in progress.
          * `last_completed_refresh_job_time_in_utc` (`pulumi.Input[str]`) - Indicates the completed time for the last refresh job on this particular share or container, if any.This could be a failed job or a successful job.
          * `last_job` (`pulumi.Input[str]`) - Indicates the id of the last refresh job on this particular share or container,if any. This could be a failed job or a successful job.

        The **user_access_rights** object supports the following:

          * `access_type` (`pulumi.Input[str]`) - Type of access to be allowed for the user.
          * `user_id` (`pulumi.Input[str]`) - User ID (already existing in the device).
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

            if access_protocol is None:
                raise TypeError("Missing required property 'access_protocol'")
            __props__['access_protocol'] = access_protocol
            __props__['azure_container_info'] = azure_container_info
            __props__['client_access_rights'] = client_access_rights
            __props__['data_policy'] = data_policy
            __props__['description'] = description
            if device_name is None:
                raise TypeError("Missing required property 'device_name'")
            __props__['device_name'] = device_name
            if monitoring_status is None:
                raise TypeError("Missing required property 'monitoring_status'")
            __props__['monitoring_status'] = monitoring_status
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['refresh_details'] = refresh_details
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if share_status is None:
                raise TypeError("Missing required property 'share_status'")
            __props__['share_status'] = share_status
            __props__['user_access_rights'] = user_access_rights
            __props__['properties'] = None
            __props__['type'] = None
        super(Share, __self__).__init__(
            'azurerm:databoxedge/v20190801:Share',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Share resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Share(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop