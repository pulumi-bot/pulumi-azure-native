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
    managed_by_extended: pulumi.Output[list]
    """
    List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
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
        * `gallery_image_reference` (`dict`) - Required if creating from a Gallery Image. The id of the ImageDiskReference will be the ARM id of the shared galley image version from which to create a disk.
          * `id` (`str`) - A relative uri containing either a Platform Image Repository or user image reference.
          * `lun` (`float`) - If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null.

        * `image_reference` (`dict`) - Disk source information.
        * `source_resource_id` (`str`) - If createOption is Copy, this is the ARM id of the source snapshot or disk.
        * `source_unique_id` (`str`) - If this field is set, this is the unique id identifying the source of this resource.
        * `source_uri` (`str`) - If createOption is Import, this is the URI of a blob to be imported into a managed disk.
        * `storage_account_id` (`str`) - Required if createOption is Import. The Azure Resource Manager identifier of the storage account containing the blob to import as a disk.
        * `upload_size_bytes` (`float`) - If createOption is Upload, this is the size of the contents of the upload including the VHD footer. This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer).

      * `disk_access_id` (`str`) - ARM id of the DiskAccess resource for using private endpoints on disks.
      * `disk_iops_read_only` (`float`) - The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
      * `disk_iops_read_write` (`float`) - The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
      * `disk_m_bps_read_only` (`float`) - The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
      * `disk_m_bps_read_write` (`float`) - The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
      * `disk_size_bytes` (`float`) - The size of the disk in bytes. This field is read only.
      * `disk_size_gb` (`float`) - If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
      * `disk_state` (`str`) - The state of the disk.
      * `encryption` (`dict`) - Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
        * `disk_encryption_set_id` (`str`) - ResourceId of the disk encryption set to use for enabling encryption at rest.
        * `type` (`str`) - The type of key used to encrypt the data of the disk.

      * `encryption_settings_collection` (`dict`) - Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
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
      * `max_shares` (`float`) - The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
      * `network_access_policy` (`str`) - Policy for accessing the disk via network.
      * `os_type` (`str`) - The Operating System type.
      * `provisioning_state` (`str`) - The disk provisioning state.
      * `share_info` (`list`) - Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
        * `vm_uri` (`str`) - A relative URI containing the ID of the VM that has the disk attached.

      * `time_created` (`str`) - The time when the disk was created.
      * `unique_id` (`str`) - Unique Guid identifying the resource.
    """
    sku: pulumi.Output[dict]
    """
    The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
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
    def __init__(__self__, resource_name, opts=None, creation_data=None, disk_access_id=None, disk_iops_read_only=None, disk_iops_read_write=None, disk_m_bps_read_only=None, disk_m_bps_read_write=None, disk_size_gb=None, encryption=None, encryption_settings_collection=None, hyper_v_generation=None, location=None, max_shares=None, name=None, network_access_policy=None, os_type=None, resource_group_name=None, sku=None, tags=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Disk resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] creation_data: Disk source information. CreationData information cannot be changed after the disk has been created.
        :param pulumi.Input[str] disk_access_id: ARM id of the DiskAccess resource for using private endpoints on disks.
        :param pulumi.Input[float] disk_iops_read_only: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
        :param pulumi.Input[float] disk_iops_read_write: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        :param pulumi.Input[float] disk_m_bps_read_only: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
        :param pulumi.Input[float] disk_m_bps_read_write: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
        :param pulumi.Input[float] disk_size_gb: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        :param pulumi.Input[dict] encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
        :param pulumi.Input[dict] encryption_settings_collection: Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
        :param pulumi.Input[str] hyper_v_generation: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[float] max_shares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
        :param pulumi.Input[str] name: The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
        :param pulumi.Input[str] network_access_policy: Policy for accessing the disk via network.
        :param pulumi.Input[str] os_type: The Operating System type.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[list] zones: The Logical zone list for Disk.

        The **creation_data** object supports the following:

          * `create_option` (`pulumi.Input[str]`) - This enumerates the possible sources of a disk's creation.
          * `gallery_image_reference` (`pulumi.Input[dict]`) - Required if creating from a Gallery Image. The id of the ImageDiskReference will be the ARM id of the shared galley image version from which to create a disk.
            * `id` (`pulumi.Input[str]`) - A relative uri containing either a Platform Image Repository or user image reference.
            * `lun` (`pulumi.Input[float]`) - If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use. For OS disks, this field is null.

          * `image_reference` (`pulumi.Input[dict]`) - Disk source information.
          * `source_resource_id` (`pulumi.Input[str]`) - If createOption is Copy, this is the ARM id of the source snapshot or disk.
          * `source_uri` (`pulumi.Input[str]`) - If createOption is Import, this is the URI of a blob to be imported into a managed disk.
          * `storage_account_id` (`pulumi.Input[str]`) - Required if createOption is Import. The Azure Resource Manager identifier of the storage account containing the blob to import as a disk.
          * `upload_size_bytes` (`pulumi.Input[float]`) - If createOption is Upload, this is the size of the contents of the upload including the VHD footer. This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer).

        The **encryption** object supports the following:

          * `disk_encryption_set_id` (`pulumi.Input[str]`) - ResourceId of the disk encryption set to use for enabling encryption at rest.
          * `type` (`pulumi.Input[str]`) - The type of key used to encrypt the data of the disk.

        The **encryption_settings_collection** object supports the following:

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
            __props__['disk_access_id'] = disk_access_id
            __props__['disk_iops_read_only'] = disk_iops_read_only
            __props__['disk_iops_read_write'] = disk_iops_read_write
            __props__['disk_m_bps_read_only'] = disk_m_bps_read_only
            __props__['disk_m_bps_read_write'] = disk_m_bps_read_write
            __props__['disk_size_gb'] = disk_size_gb
            __props__['encryption'] = encryption
            __props__['encryption_settings_collection'] = encryption_settings_collection
            __props__['hyper_v_generation'] = hyper_v_generation
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['max_shares'] = max_shares
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_access_policy'] = network_access_policy
            __props__['os_type'] = os_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['managed_by'] = None
            __props__['managed_by_extended'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(Disk, __self__).__init__(
            'azurerm:compute/v20200501:Disk',
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