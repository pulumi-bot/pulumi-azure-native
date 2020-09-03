# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VirtualMachine']


class VirtualMachine(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_set: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 diagnostics_profile: Optional[pulumi.Input[pulumi.InputType['DiagnosticsProfileArgs']]] = None,
                 hardware_profile: Optional[pulumi.Input[pulumi.InputType['HardwareProfileArgs']]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']]] = None,
                 license_type: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 network_profile: Optional[pulumi.Input[pulumi.InputType['NetworkProfileArgs']]] = None,
                 os_profile: Optional[pulumi.Input[pulumi.InputType['OSProfileArgs']]] = None,
                 plan: Optional[pulumi.Input[pulumi.InputType['PlanArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_profile: Optional[pulumi.Input[pulumi.InputType['StorageProfileArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vm_name: Optional[pulumi.Input[str]] = None,
                 zones: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Describes a Virtual Machine.

        ## Example Usage
        ### Create a custom-image vm from an unmanaged generalized os image.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            hardware_profile={
                "vmSize": "Standard_D1_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminPassword": "{your-password}",
                "adminUsername": "{your-username}",
                "computerName": "myVM",
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "image": {
                        "uri": "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/{existing-generalized-os-image-blob-name}.vhd",
                    },
                    "name": "myVMosdisk",
                    "osType": "Windows",
                    "vhd": {
                        "uri": "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd",
                    },
                },
            },
            vm_name="{vm-name}")

        ```
        ### Create a platform-image vm with unmanaged os and data disks.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            hardware_profile={
                "vmSize": "Standard_D2_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminPassword": "{your-password}",
                "adminUsername": "{your-username}",
                "computerName": "myVM",
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "dataDisks": [
                    {
                        "createOption": "Empty",
                        "diskSizeGB": 1023,
                        "lun": 0,
                        "vhd": {
                            "uri": "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk0.vhd",
                        },
                    },
                    {
                        "createOption": "Empty",
                        "diskSizeGB": 1023,
                        "lun": 1,
                        "vhd": {
                            "uri": "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk1.vhd",
                        },
                    },
                ],
                "imageReference": {
                    "offer": "WindowsServer",
                    "publisher": "MicrosoftWindowsServer",
                    "sku": "2016-Datacenter",
                    "version": "latest",
                },
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "name": "myVMosdisk",
                    "vhd": {
                        "uri": "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd",
                    },
                },
            },
            vm_name="{vm-name}")

        ```
        ### Create a vm from a custom image.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            hardware_profile={
                "vmSize": "Standard_D1_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminPassword": "{your-password}",
                "adminUsername": "{your-username}",
                "computerName": "myVM",
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "imageReference": {
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}",
                },
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "managedDisk": {
                        "storageAccountType": "Standard_LRS",
                    },
                    "name": "myVMosdisk",
                },
            },
            vm_name="myVM")

        ```
        ### Create a vm in an availability set.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            availability_set={
                "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/{existing-availability-set-name}",
            },
            hardware_profile={
                "vmSize": "Standard_D1_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminPassword": "{your-password}",
                "adminUsername": "{your-username}",
                "computerName": "myVM",
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "imageReference": {
                    "offer": "WindowsServer",
                    "publisher": "MicrosoftWindowsServer",
                    "sku": "2016-Datacenter",
                    "version": "latest",
                },
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "managedDisk": {
                        "storageAccountType": "Standard_LRS",
                    },
                    "name": "myVMosdisk",
                },
            },
            vm_name="myVM")

        ```
        ### Create a vm with a marketplace image plan.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            hardware_profile={
                "vmSize": "Standard_D1_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminPassword": "{your-password}",
                "adminUsername": "{your-username}",
                "computerName": "myVM",
            },
            plan={
                "name": "windows2016",
                "product": "windows-data-science-vm",
                "publisher": "microsoft-ads",
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "imageReference": {
                    "offer": "windows-data-science-vm",
                    "publisher": "microsoft-ads",
                    "sku": "windows2016",
                    "version": "latest",
                },
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "managedDisk": {
                        "storageAccountType": "Standard_LRS",
                    },
                    "name": "myVMosdisk",
                },
            },
            vm_name="myVM")

        ```
        ### Create a vm with boot diagnostics.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            diagnostics_profile={
                "bootDiagnostics": {
                    "enabled": True,
                    "storageUri": "http://{existing-storage-account-name}.blob.core.windows.net",
                },
            },
            hardware_profile={
                "vmSize": "Standard_D1_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminPassword": "{your-password}",
                "adminUsername": "{your-username}",
                "computerName": "myVM",
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "imageReference": {
                    "offer": "WindowsServer",
                    "publisher": "MicrosoftWindowsServer",
                    "sku": "2016-Datacenter",
                    "version": "latest",
                },
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "managedDisk": {
                        "storageAccountType": "Standard_LRS",
                    },
                    "name": "myVMosdisk",
                },
            },
            vm_name="myVM")

        ```
        ### Create a vm with empty data disks.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            hardware_profile={
                "vmSize": "Standard_D2_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminPassword": "{your-password}",
                "adminUsername": "{your-username}",
                "computerName": "myVM",
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "dataDisks": [
                    {
                        "createOption": "Empty",
                        "diskSizeGB": 1023,
                        "lun": 0,
                    },
                    {
                        "createOption": "Empty",
                        "diskSizeGB": 1023,
                        "lun": 1,
                    },
                ],
                "imageReference": {
                    "offer": "WindowsServer",
                    "publisher": "MicrosoftWindowsServer",
                    "sku": "2016-Datacenter",
                    "version": "latest",
                },
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "managedDisk": {
                        "storageAccountType": "Standard_LRS",
                    },
                    "name": "myVMosdisk",
                },
            },
            vm_name="myVM")

        ```
        ### Create a vm with password authentication.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            hardware_profile={
                "vmSize": "Standard_D1_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminPassword": "{your-password}",
                "adminUsername": "{your-username}",
                "computerName": "myVM",
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "imageReference": {
                    "offer": "WindowsServer",
                    "publisher": "MicrosoftWindowsServer",
                    "sku": "2016-Datacenter",
                    "version": "latest",
                },
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "managedDisk": {
                        "storageAccountType": "Standard_LRS",
                    },
                    "name": "myVMosdisk",
                },
            },
            vm_name="myVM")

        ```
        ### Create a vm with premium storage.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            hardware_profile={
                "vmSize": "Standard_D1_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminPassword": "{your-password}",
                "adminUsername": "{your-username}",
                "computerName": "myVM",
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "imageReference": {
                    "offer": "WindowsServer",
                    "publisher": "MicrosoftWindowsServer",
                    "sku": "2016-Datacenter",
                    "version": "latest",
                },
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "managedDisk": {
                        "storageAccountType": "Premium_LRS",
                    },
                    "name": "myVMosdisk",
                },
            },
            vm_name="myVM")

        ```
        ### Create a vm with ssh authentication.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_machine = azurerm.compute.v20171201.VirtualMachine("virtualMachine",
            hardware_profile={
                "vmSize": "Standard_D1_v2",
            },
            location="westus",
            network_profile={
                "networkInterfaces": [{
                    "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
                }],
            },
            os_profile={
                "adminUsername": "{your-username}",
                "computerName": "myVM",
                "linuxConfiguration": {
                    "disablePasswordAuthentication": True,
                    "ssh": {
                        "publicKeys": [{
                            "keyData": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1",
                            "path": "/home/{your-username}/.ssh/authorized_keys",
                        }],
                    },
                },
            },
            resource_group_name="myResourceGroup",
            storage_profile={
                "imageReference": {
                    "offer": "{image_offer}",
                    "publisher": "{image_publisher}",
                    "sku": "{image_sku}",
                    "version": "latest",
                },
                "osDisk": {
                    "caching": "ReadWrite",
                    "createOption": "FromImage",
                    "managedDisk": {
                        "storageAccountType": "Standard_LRS",
                    },
                    "name": "myVMosdisk",
                },
            },
            vm_name="myVM")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] availability_set: Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
        :param pulumi.Input[pulumi.InputType['DiagnosticsProfileArgs']] diagnostics_profile: Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
        :param pulumi.Input[pulumi.InputType['HardwareProfileArgs']] hardware_profile: Specifies the hardware settings for the virtual machine.
        :param pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']] identity: The identity of the virtual machine, if configured.
        :param pulumi.Input[str] license_type: Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[pulumi.InputType['NetworkProfileArgs']] network_profile: Specifies the network interfaces of the virtual machine.
        :param pulumi.Input[pulumi.InputType['OSProfileArgs']] os_profile: Specifies the operating system settings for the virtual machine.
        :param pulumi.Input[pulumi.InputType['PlanArgs']] plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['StorageProfileArgs']] storage_profile: Specifies the storage settings for the virtual machine disks.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[str] vm_name: The name of the virtual machine.
        :param pulumi.Input[List[pulumi.Input[str]]] zones: The virtual machine zones.
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

            __props__['availability_set'] = availability_set
            __props__['diagnostics_profile'] = diagnostics_profile
            __props__['hardware_profile'] = hardware_profile
            __props__['identity'] = identity
            __props__['license_type'] = license_type
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['network_profile'] = network_profile
            __props__['os_profile'] = os_profile
            __props__['plan'] = plan
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['storage_profile'] = storage_profile
            __props__['tags'] = tags
            if vm_name is None:
                raise TypeError("Missing required property 'vm_name'")
            __props__['vm_name'] = vm_name
            __props__['zones'] = zones
            __props__['instance_view'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['resources'] = None
            __props__['type'] = None
            __props__['vm_id'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:compute/latest:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20150615:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20160330:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20160430preview:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20170330:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20180401:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20180601:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20181001:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20190301:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20190701:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20191201:VirtualMachine"), pulumi.Alias(type_="azurerm:compute/v20200601:VirtualMachine")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualMachine, __self__).__init__(
            'azurerm:compute/v20171201:VirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualMachine':
        """
        Get an existing VirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilitySet")
    def availability_set(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
        """
        return pulumi.get(self, "availability_set")

    @property
    @pulumi.getter(name="diagnosticsProfile")
    def diagnostics_profile(self) -> pulumi.Output[Optional['outputs.DiagnosticsProfileResponse']]:
        """
        Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
        """
        return pulumi.get(self, "diagnostics_profile")

    @property
    @pulumi.getter(name="hardwareProfile")
    def hardware_profile(self) -> pulumi.Output[Optional['outputs.HardwareProfileResponse']]:
        """
        Specifies the hardware settings for the virtual machine.
        """
        return pulumi.get(self, "hardware_profile")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.VirtualMachineIdentityResponse']]:
        """
        The identity of the virtual machine, if configured.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="instanceView")
    def instance_view(self) -> pulumi.Output['outputs.VirtualMachineInstanceViewResponse']:
        """
        The virtual machine instance view.
        """
        return pulumi.get(self, "instance_view")

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
        """
        return pulumi.get(self, "license_type")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkProfile")
    def network_profile(self) -> pulumi.Output[Optional['outputs.NetworkProfileResponse']]:
        """
        Specifies the network interfaces of the virtual machine.
        """
        return pulumi.get(self, "network_profile")

    @property
    @pulumi.getter(name="osProfile")
    def os_profile(self) -> pulumi.Output[Optional['outputs.OSProfileResponse']]:
        """
        Specifies the operating system settings for the virtual machine.
        """
        return pulumi.get(self, "os_profile")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[Optional['outputs.PlanResponse']]:
        """
        Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[List['outputs.VirtualMachineExtensionResponse']]:
        """
        The virtual machine child extension resources.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> pulumi.Output[Optional['outputs.StorageProfileResponse']]:
        """
        Specifies the storage settings for the virtual machine disks.
        """
        return pulumi.get(self, "storage_profile")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> pulumi.Output[str]:
        """
        Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
        """
        return pulumi.get(self, "vm_id")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[List[str]]]:
        """
        The virtual machine zones.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

