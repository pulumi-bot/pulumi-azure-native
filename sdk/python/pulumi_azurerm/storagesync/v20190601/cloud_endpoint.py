# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class CloudEndpoint(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    Cloud Endpoint properties.
      * `azure_file_share_name` (`str`) - Azure file share name
      * `backup_enabled` (`str`) - Backup Enabled
      * `friendly_name` (`str`) - Friendly Name
      * `last_operation_name` (`str`) - Resource Last Operation Name
      * `last_workflow_id` (`str`) - CloudEndpoint lastWorkflowId
      * `partnership_id` (`str`) - Partnership Id
      * `provisioning_state` (`str`) - CloudEndpoint Provisioning State
      * `storage_account_resource_id` (`str`) - Storage Account Resource Id
      * `storage_account_tenant_id` (`str`) - Storage Account Tenant Id
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, storage_sync_service_name=None, sync_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Cloud Endpoint object.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of Cloud Endpoint object.
        :param pulumi.Input[dict] properties: The parameters used to create the cloud endpoint.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] storage_sync_service_name: Name of Storage Sync Service resource.
        :param pulumi.Input[str] sync_group_name: Name of Sync Group resource.

        The **properties** object supports the following:

          * `azure_file_share_name` (`pulumi.Input[str]`) - Azure file share name
          * `friendly_name` (`pulumi.Input[str]`) - Friendly Name
          * `storage_account_resource_id` (`pulumi.Input[str]`) - Storage Account Resource Id
          * `storage_account_tenant_id` (`pulumi.Input[str]`) - Storage Account Tenant Id
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
        super(CloudEndpoint, __self__).__init__(
            'azurerm:storagesync/v20190601:CloudEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing CloudEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return CloudEndpoint(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
