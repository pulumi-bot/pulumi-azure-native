# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Disk(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    managed_by: pulumi.Output[str]
    """
    A relative URI containing the ID of the VM that has the disk attached.
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Disk resource properties.
      * `creation_data` (`dict`) - Disk source information. CreationData information cannot be changed after the disk has been created.
        * `create_option` (`str`) - This enumerates the possible sources of a disk's creation.
        * `image_reference` (`dict`) - Disk source information.
          * `id` (`str`) - A relative uri containing either a Platform Image Repository or user image reference.
          * `lun` (`float`) - If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null.

        * `source_resource_id` (`str`) - If createOption is Copy, this is the ARM id of the source snapshot or disk.
        * `source_uri` (`str`) - If createOption is Import, this is the URI of a blob to be imported into a managed disk.
        * `storage_account_id` (`str`) - If createOption is Import, the Azure Resource Manager identifier of the storage account containing the blob to import as a disk. Required only if the blob is in a different subscription

      * `disk_size_gb` (`float`) - If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
      * `encryption_settings` (`dict`) - Encryption settings for disk or snapshot
        * `disk_encryption_key` (`dict`) - Key Vault Secret Url and vault id of the disk encryption key
          * `secret_url` (`str`) - Url pointing to a key or secret in KeyVault
          * `source_vault` (`dict`) - Resource id of the KeyVault containing the key or secret
            * `id` (`str`) - Resource Id

        * `enabled` (`bool`) - Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is null in the request object, the existing settings remain unchanged.
        * `key_encryption_key` (`dict`) - Key Vault Key Url and vault id of the key encryption key
          * `key_url` (`str`) - Url pointing to a key or secret in KeyVault
          * `source_vault` (`dict`) - Resource id of the KeyVault containing the key or secret

      * `os_type` (`str`) - The Operating System type.
      * `provisioning_state` (`str`) - The disk provisioning state.
      * `time_created` (`str`) - The time when the disk was created.
    """
    sku: pulumi.Output[dict]
    """
    The disks and snapshots sku name. Can be Standard_LRS or Premium_LRS.
      * `name` (`str`) - The sku name.
      * `tier` (`str`) - The sku tier.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    zones: pulumi.Output[list]
    """
    The Logical zone list for Disk.
    """
    def __init__(__self__, resource_name, opts=None, creation_data=None, disk_size_gb=None, encryption_settings=None, location=None, name=None, os_type=None, resource_group_name=None, sku=None, tags=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Disk resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] creation_data: Disk source information. CreationData information cannot be changed after the disk has been created.
        :param pulumi.Input[float] disk_size_gb: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        :param pulumi.Input[dict] encryption_settings: Encryption settings for disk or snapshot
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
        :param pulumi.Input[str] os_type: The Operating System type.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: The disks and snapshots sku name. Can be Standard_LRS or Premium_LRS.
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[list] zones: The Logical zone list for Disk.

        The **creation_data** object supports the following:

          * `create_option` (`pulumi.Input[str]`) - This enumerates the possible sources of a disk's creation.
          * `image_reference` (`pulumi.Input[dict]`) - Disk source information.
            * `id` (`pulumi.Input[str]`) - A relative uri containing either a Platform Image Repository or user image reference.
            * `lun` (`pulumi.Input[float]`) - If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null.

          * `source_resource_id` (`pulumi.Input[str]`) - If createOption is Copy, this is the ARM id of the source snapshot or disk.
          * `source_uri` (`pulumi.Input[str]`) - If createOption is Import, this is the URI of a blob to be imported into a managed disk.
          * `storage_account_id` (`pulumi.Input[str]`) - If createOption is Import, the Azure Resource Manager identifier of the storage account containing the blob to import as a disk. Required only if the blob is in a different subscription

        The **encryption_settings** object supports the following:

          * `disk_encryption_key` (`pulumi.Input[dict]`) - Key Vault Secret Url and vault id of the disk encryption key
            * `secret_url` (`pulumi.Input[str]`) - Url pointing to a key or secret in KeyVault
            * `source_vault` (`pulumi.Input[dict]`) - Resource id of the KeyVault containing the key or secret
              * `id` (`pulumi.Input[str]`) - Resource Id

          * `enabled` (`pulumi.Input[bool]`) - Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is null in the request object, the existing settings remain unchanged.
          * `key_encryption_key` (`pulumi.Input[dict]`) - Key Vault Key Url and vault id of the key encryption key
            * `key_url` (`pulumi.Input[str]`) - Url pointing to a key or secret in KeyVault
            * `source_vault` (`pulumi.Input[dict]`) - Resource id of the KeyVault containing the key or secret

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - The sku name.
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

            if creation_data is None:
                raise TypeError("Missing required property 'creation_data'")
            __props__['creation_data'] = creation_data
            __props__['disk_size_gb'] = disk_size_gb
            __props__['encryption_settings'] = encryption_settings
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['os_type'] = os_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['managed_by'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(Disk, __self__).__init__(
            'azurerm:compute/v20170330:Disk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Disk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Disk(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
