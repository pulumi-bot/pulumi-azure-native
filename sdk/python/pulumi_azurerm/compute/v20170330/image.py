# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Image(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Describes the properties of an Image.
      * `provisioning_state` (`str`) - The provisioning state.
      * `source_virtual_machine` (`dict`) - The source virtual machine from which Image is created.
        * `id` (`str`) - Resource Id

      * `storage_profile` (`dict`) - Specifies the storage settings for the virtual machine disks.
        * `data_disks` (`list`) - Specifies the parameters that are used to add a data disk to a virtual machine. <br><br> For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
          * `blob_uri` (`str`) - The Virtual Hard Disk.
          * `caching` (`str`) - Specifies the caching requirements. <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage. ReadOnly for Premium storage**
          * `disk_size_gb` (`float`) - Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the disk in a virtual machine image. <br><br> This value cannot be larger than 1023 GB
          * `lun` (`float`) - Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
          * `managed_disk` (`dict`) - The managedDisk.
          * `snapshot` (`dict`) - The snapshot.
          * `storage_account_type` (`str`) - Specifies the storage account type for the managed disk. Possible values are: Standard_LRS or Premium_LRS.

        * `os_disk` (`dict`) - Specifies information about the operating system disk used by the virtual machine. <br><br> For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
          * `blob_uri` (`str`) - The Virtual Hard Disk.
          * `caching` (`str`) - Specifies the caching requirements. <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage. ReadOnly for Premium storage**
          * `disk_size_gb` (`float`) - Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the disk in a virtual machine image. <br><br> This value cannot be larger than 1023 GB
          * `managed_disk` (`dict`) - The managedDisk.
          * `os_state` (`str`) - The OS State.
          * `os_type` (`str`) - This property allows you to specify the type of the OS that is included in the disk if creating a VM from a custom image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
          * `snapshot` (`dict`) - The snapshot.
          * `storage_account_type` (`str`) - Specifies the storage account type for the managed disk. Possible values are: Standard_LRS or Premium_LRS.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, source_virtual_machine=None, storage_profile=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the image.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] source_virtual_machine: The source virtual machine from which Image is created.
        :param pulumi.Input[dict] storage_profile: Specifies the storage settings for the virtual machine disks.
        :param pulumi.Input[dict] tags: Resource tags

        The **source_virtual_machine** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource Id

        The **storage_profile** object supports the following:

          * `data_disks` (`pulumi.Input[list]`) - Specifies the parameters that are used to add a data disk to a virtual machine. <br><br> For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
            * `blob_uri` (`pulumi.Input[str]`) - The Virtual Hard Disk.
            * `caching` (`pulumi.Input[str]`) - Specifies the caching requirements. <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage. ReadOnly for Premium storage**
            * `disk_size_gb` (`pulumi.Input[float]`) - Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the disk in a virtual machine image. <br><br> This value cannot be larger than 1023 GB
            * `lun` (`pulumi.Input[float]`) - Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
            * `managed_disk` (`pulumi.Input[dict]`) - The managedDisk.
            * `snapshot` (`pulumi.Input[dict]`) - The snapshot.
            * `storage_account_type` (`pulumi.Input[str]`) - Specifies the storage account type for the managed disk. Possible values are: Standard_LRS or Premium_LRS.

          * `os_disk` (`pulumi.Input[dict]`) - Specifies information about the operating system disk used by the virtual machine. <br><br> For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
            * `blob_uri` (`pulumi.Input[str]`) - The Virtual Hard Disk.
            * `caching` (`pulumi.Input[str]`) - Specifies the caching requirements. <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage. ReadOnly for Premium storage**
            * `disk_size_gb` (`pulumi.Input[float]`) - Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the disk in a virtual machine image. <br><br> This value cannot be larger than 1023 GB
            * `managed_disk` (`pulumi.Input[dict]`) - The managedDisk.
            * `os_state` (`pulumi.Input[str]`) - The OS State.
            * `os_type` (`pulumi.Input[str]`) - This property allows you to specify the type of the OS that is included in the disk if creating a VM from a custom image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
            * `snapshot` (`pulumi.Input[dict]`) - The snapshot.
            * `storage_account_type` (`pulumi.Input[str]`) - Specifies the storage account type for the managed disk. Possible values are: Standard_LRS or Premium_LRS.
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
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['source_virtual_machine'] = source_virtual_machine
            __props__['storage_profile'] = storage_profile
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(Image, __self__).__init__(
            'azurerm:compute/v20170330:Image',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Image resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Image(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
