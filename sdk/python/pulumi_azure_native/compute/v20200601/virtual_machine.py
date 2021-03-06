# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['VirtualMachine']


class VirtualMachine(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_capabilities: Optional[pulumi.Input[pulumi.InputType['AdditionalCapabilitiesArgs']]] = None,
                 availability_set: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 billing_profile: Optional[pulumi.Input[pulumi.InputType['BillingProfileArgs']]] = None,
                 diagnostics_profile: Optional[pulumi.Input[pulumi.InputType['DiagnosticsProfileArgs']]] = None,
                 eviction_policy: Optional[pulumi.Input[Union[str, 'VirtualMachineEvictionPolicyTypes']]] = None,
                 extensions_time_budget: Optional[pulumi.Input[str]] = None,
                 hardware_profile: Optional[pulumi.Input[pulumi.InputType['HardwareProfileArgs']]] = None,
                 host: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 host_group: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']]] = None,
                 license_type: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 network_profile: Optional[pulumi.Input[pulumi.InputType['NetworkProfileArgs']]] = None,
                 os_profile: Optional[pulumi.Input[pulumi.InputType['OSProfileArgs']]] = None,
                 plan: Optional[pulumi.Input[pulumi.InputType['PlanArgs']]] = None,
                 priority: Optional[pulumi.Input[Union[str, 'VirtualMachinePriorityTypes']]] = None,
                 proximity_placement_group: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 security_profile: Optional[pulumi.Input[pulumi.InputType['SecurityProfileArgs']]] = None,
                 storage_profile: Optional[pulumi.Input[pulumi.InputType['StorageProfileArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_machine_scale_set: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 vm_name: Optional[pulumi.Input[str]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Describes a Virtual Machine.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AdditionalCapabilitiesArgs']] additional_capabilities: Specifies additional capabilities enabled or disabled on the virtual machine.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] availability_set: Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
        :param pulumi.Input[pulumi.InputType['BillingProfileArgs']] billing_profile: Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
        :param pulumi.Input[pulumi.InputType['DiagnosticsProfileArgs']] diagnostics_profile: Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
        :param pulumi.Input[Union[str, 'VirtualMachineEvictionPolicyTypes']] eviction_policy: Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
        :param pulumi.Input[str] extensions_time_budget: Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). <br><br> Minimum api-version: 2020-06-01
        :param pulumi.Input[pulumi.InputType['HardwareProfileArgs']] hardware_profile: Specifies the hardware settings for the virtual machine.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] host: Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] host_group: Specifies information about the dedicated host group that the virtual machine resides in. <br><br>Minimum api-version: 2020-06-01. <br><br>NOTE: User cannot specify both host and hostGroup properties.
        :param pulumi.Input[pulumi.InputType['VirtualMachineIdentityArgs']] identity: The identity of the virtual machine, if configured.
        :param pulumi.Input[str] license_type: Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[pulumi.InputType['NetworkProfileArgs']] network_profile: Specifies the network interfaces of the virtual machine.
        :param pulumi.Input[pulumi.InputType['OSProfileArgs']] os_profile: Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
        :param pulumi.Input[pulumi.InputType['PlanArgs']] plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
        :param pulumi.Input[Union[str, 'VirtualMachinePriorityTypes']] priority: Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] proximity_placement_group: Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['SecurityProfileArgs']] security_profile: Specifies the Security related profile settings for the virtual machine.
        :param pulumi.Input[pulumi.InputType['StorageProfileArgs']] storage_profile: Specifies the storage settings for the virtual machine disks.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] virtual_machine_scale_set: Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
        :param pulumi.Input[str] vm_name: The name of the virtual machine.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: The virtual machine zones.
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

            __props__['additional_capabilities'] = additional_capabilities
            __props__['availability_set'] = availability_set
            __props__['billing_profile'] = billing_profile
            __props__['diagnostics_profile'] = diagnostics_profile
            __props__['eviction_policy'] = eviction_policy
            __props__['extensions_time_budget'] = extensions_time_budget
            __props__['hardware_profile'] = hardware_profile
            __props__['host'] = host
            __props__['host_group'] = host_group
            __props__['identity'] = identity
            __props__['license_type'] = license_type
            __props__['location'] = location
            __props__['network_profile'] = network_profile
            __props__['os_profile'] = os_profile
            __props__['plan'] = plan
            __props__['priority'] = priority
            __props__['proximity_placement_group'] = proximity_placement_group
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['security_profile'] = security_profile
            __props__['storage_profile'] = storage_profile
            __props__['tags'] = tags
            __props__['virtual_machine_scale_set'] = virtual_machine_scale_set
            __props__['vm_name'] = vm_name
            __props__['zones'] = zones
            __props__['instance_view'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['resources'] = None
            __props__['type'] = None
            __props__['vm_id'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:compute/v20200601:VirtualMachine"), pulumi.Alias(type_="azure-native:compute:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/latest:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/latest:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20150615:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20150615:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20160330:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20160330:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20160430preview:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20160430preview:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20170330:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20170330:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20171201:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20171201:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20180401:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20180401:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20180601:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20180601:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20181001:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20181001:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20190301:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20190301:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20190701:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20190701:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20191201:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20191201:VirtualMachine"), pulumi.Alias(type_="azure-native:compute/v20201201:VirtualMachine"), pulumi.Alias(type_="azure-nextgen:compute/v20201201:VirtualMachine")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualMachine, __self__).__init__(
            'azure-native:compute/v20200601:VirtualMachine',
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

        __props__["additional_capabilities"] = None
        __props__["availability_set"] = None
        __props__["billing_profile"] = None
        __props__["diagnostics_profile"] = None
        __props__["eviction_policy"] = None
        __props__["extensions_time_budget"] = None
        __props__["hardware_profile"] = None
        __props__["host"] = None
        __props__["host_group"] = None
        __props__["identity"] = None
        __props__["instance_view"] = None
        __props__["license_type"] = None
        __props__["location"] = None
        __props__["name"] = None
        __props__["network_profile"] = None
        __props__["os_profile"] = None
        __props__["plan"] = None
        __props__["priority"] = None
        __props__["provisioning_state"] = None
        __props__["proximity_placement_group"] = None
        __props__["resources"] = None
        __props__["security_profile"] = None
        __props__["storage_profile"] = None
        __props__["tags"] = None
        __props__["type"] = None
        __props__["virtual_machine_scale_set"] = None
        __props__["vm_id"] = None
        __props__["zones"] = None
        return VirtualMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalCapabilities")
    def additional_capabilities(self) -> pulumi.Output[Optional['outputs.AdditionalCapabilitiesResponse']]:
        """
        Specifies additional capabilities enabled or disabled on the virtual machine.
        """
        return pulumi.get(self, "additional_capabilities")

    @property
    @pulumi.getter(name="availabilitySet")
    def availability_set(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
        """
        return pulumi.get(self, "availability_set")

    @property
    @pulumi.getter(name="billingProfile")
    def billing_profile(self) -> pulumi.Output[Optional['outputs.BillingProfileResponse']]:
        """
        Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
        """
        return pulumi.get(self, "billing_profile")

    @property
    @pulumi.getter(name="diagnosticsProfile")
    def diagnostics_profile(self) -> pulumi.Output[Optional['outputs.DiagnosticsProfileResponse']]:
        """
        Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
        """
        return pulumi.get(self, "diagnostics_profile")

    @property
    @pulumi.getter(name="evictionPolicy")
    def eviction_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
        """
        return pulumi.get(self, "eviction_policy")

    @property
    @pulumi.getter(name="extensionsTimeBudget")
    def extensions_time_budget(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). <br><br> Minimum api-version: 2020-06-01
        """
        return pulumi.get(self, "extensions_time_budget")

    @property
    @pulumi.getter(name="hardwareProfile")
    def hardware_profile(self) -> pulumi.Output[Optional['outputs.HardwareProfileResponse']]:
        """
        Specifies the hardware settings for the virtual machine.
        """
        return pulumi.get(self, "hardware_profile")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="hostGroup")
    def host_group(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        Specifies information about the dedicated host group that the virtual machine resides in. <br><br>Minimum api-version: 2020-06-01. <br><br>NOTE: User cannot specify both host and hostGroup properties.
        """
        return pulumi.get(self, "host_group")

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
        Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
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
        Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
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
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="proximityPlacementGroup")
    def proximity_placement_group(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
        """
        return pulumi.get(self, "proximity_placement_group")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[Sequence['outputs.VirtualMachineExtensionResponse']]:
        """
        The virtual machine child extension resources.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="securityProfile")
    def security_profile(self) -> pulumi.Output[Optional['outputs.SecurityProfileResponse']]:
        """
        Specifies the Security related profile settings for the virtual machine.
        """
        return pulumi.get(self, "security_profile")

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
    @pulumi.getter(name="virtualMachineScaleSet")
    def virtual_machine_scale_set(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
        """
        return pulumi.get(self, "virtual_machine_scale_set")

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> pulumi.Output[str]:
        """
        Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
        """
        return pulumi.get(self, "vm_id")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The virtual machine zones.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

