# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Snapshot(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    managed_by: pulumi.Output[str]
    """
    Unused. Always Null.
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Snapshot resource properties.
      * `creation_data` (`dict`) - Disk source information. CreationData information cannot be changed after the disk has been created.
        * `create_option` (`str`) - This enumerates the possible sources of a disk's creation.
        * `image_reference` (`dict`) - Disk source information.
          * `id` (`str`) - A relative uri containing either a Platform Image Repository or user image reference.
          * `lun` (`float`) - If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null.

        * `source_resource_id` (`str`) - If createOption is Copy, this is the ARM id of the source snapshot or disk.
        * `source_unique_id` (`str`) - If this field is set, this is the unique id identifying the source of this resource.
        * `source_uri` (`str`) - If createOption is Import, this is the URI of a blob to be imported into a managed disk.
        * `storage_account_id` (`str`) - If createOption is Import, the Azure Resource Manager identifier of the storage account containing the blob to import as a disk. Required only if the blob is in a different subscription
        * `upload_size_bytes` (`float`) - If createOption is Upload, this is the size of the contents of the upload including the VHD footer. This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer).

      * `disk_size_bytes` (`float`) - The size of the disk in bytes. This field is read only.
      * `disk_size_gb` (`float`) - If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
      * `encryption_settings_collection` (`dict`) - Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
        * `enabled` (`bool`) - Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is null in the request object, the existing settings remain unchanged.
        * `encryption_settings` (`list`) - A collection of encryption settings, one for each disk volume.
          * `disk_encryption_key` (`dict`) - Key Vault Secret Url and vault id of the disk encryption key
            * `secret_url` (`str`) - Url pointing to a key or secret in KeyVault
            * `source_vault` (`dict`) - Resource id of the KeyVault containing the key or secret
              * `id` (`str`) - Resource Id

          * `key_encryption_key` (`dict`) - Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when provided is used to unwrap the disk encryption key.
            * `key_url` (`str`) - Url pointing to a key or secret in KeyVault
            * `source_vault` (`dict`) - Resource id of the KeyVault containing the key or secret

        * `encryption_settings_version` (`str`) - Describes what type of encryption is used for the disks. Once this field is set, it cannot be overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.

      * `hyper_v_generation` (`str`) - The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
      * `incremental` (`bool`) - Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
      * `os_type` (`str`) - The Operating System type.
      * `provisioning_state` (`str`) - The disk provisioning state.
      * `time_created` (`str`) - The time when the disk was created.
      * `unique_id` (`str`) - Unique Guid identifying the resource.
    """
    sku: pulumi.Output[dict]
    """
    The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
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
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Snapshot resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.
        :param pulumi.Input[dict] properties: Snapshot resource properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
        :param pulumi.Input[dict] tags: Resource tags

        The **properties** object supports the following:

          * `creation_data` (`pulumi.Input[dict]`) - Disk source information. CreationData information cannot be changed after the disk has been created.
            * `create_option` (`pulumi.Input[str]`) - This enumerates the possible sources of a disk's creation.
            * `image_reference` (`pulumi.Input[dict]`) - Disk source information.
              * `id` (`pulumi.Input[str]`) - A relative uri containing either a Platform Image Repository or user image reference.
              * `lun` (`pulumi.Input[float]`) - If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null.

            * `source_resource_id` (`pulumi.Input[str]`) - If createOption is Copy, this is the ARM id of the source snapshot or disk.
            * `source_uri` (`pulumi.Input[str]`) - If createOption is Import, this is the URI of a blob to be imported into a managed disk.
            * `storage_account_id` (`pulumi.Input[str]`) - If createOption is Import, the Azure Resource Manager identifier of the storage account containing the blob to import as a disk. Required only if the blob is in a different subscription
            * `upload_size_bytes` (`pulumi.Input[float]`) - If createOption is Upload, this is the size of the contents of the upload including the VHD footer. This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer).

          * `disk_size_gb` (`pulumi.Input[float]`) - If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
          * `encryption_settings_collection` (`pulumi.Input[dict]`) - Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
            * `enabled` (`pulumi.Input[bool]`) - Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is null in the request object, the existing settings remain unchanged.
            * `encryption_settings` (`pulumi.Input[list]`) - A collection of encryption settings, one for each disk volume.
              * `disk_encryption_key` (`pulumi.Input[dict]`) - Key Vault Secret Url and vault id of the disk encryption key
                * `secret_url` (`pulumi.Input[str]`) - Url pointing to a key or secret in KeyVault
                * `source_vault` (`pulumi.Input[dict]`) - Resource id of the KeyVault containing the key or secret
                  * `id` (`pulumi.Input[str]`) - Resource Id

              * `key_encryption_key` (`pulumi.Input[dict]`) - Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when provided is used to unwrap the disk encryption key.
                * `key_url` (`pulumi.Input[str]`) - Url pointing to a key or secret in KeyVault
                * `source_vault` (`pulumi.Input[dict]`) - Resource id of the KeyVault containing the key or secret

            * `encryption_settings_version` (`pulumi.Input[str]`) - Describes what type of encryption is used for the disks. Once this field is set, it cannot be overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.

          * `hyper_v_generation` (`pulumi.Input[str]`) - The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
          * `incremental` (`pulumi.Input[bool]`) - Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
          * `os_type` (`pulumi.Input[str]`) - The Operating System type.

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
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['managed_by'] = None
            __props__['type'] = None
        super(Snapshot, __self__).__init__(
            'azurerm:compute/v20190301:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Snapshot(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
