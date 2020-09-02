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

__all__ = ['VirtualMachineScaleSet']


class VirtualMachineScaleSet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['VirtualMachineScaleSetIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 overprovision: Optional[pulumi.Input[bool]] = None,
                 plan: Optional[pulumi.Input[pulumi.InputType['PlanArgs']]] = None,
                 platform_fault_domain_count: Optional[pulumi.Input[float]] = None,
                 proximity_placement_group: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 single_placement_group: Optional[pulumi.Input[bool]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['SkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upgrade_policy: Optional[pulumi.Input[pulumi.InputType['UpgradePolicyArgs']]] = None,
                 virtual_machine_profile: Optional[pulumi.Input[pulumi.InputType['VirtualMachineScaleSetVMProfileArgs']]] = None,
                 vm_scale_set_name: Optional[pulumi.Input[str]] = None,
                 zone_balance: Optional[pulumi.Input[bool]] = None,
                 zones: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Describes a Virtual Machine Scale Set.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VirtualMachineScaleSetIdentityArgs']] identity: The identity of the virtual machine scale set, if configured.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[bool] overprovision: Specifies whether the Virtual Machine Scale Set should be overprovisioned.
        :param pulumi.Input[pulumi.InputType['PlanArgs']] plan: Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
        :param pulumi.Input[float] platform_fault_domain_count: Fault Domain count for each placement group.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] proximity_placement_group: Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[bool] single_placement_group: When true this limits the scale set to a single placement group, of max size 100 virtual machines.
        :param pulumi.Input[pulumi.InputType['SkuArgs']] sku: The virtual machine scale set sku.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[pulumi.InputType['UpgradePolicyArgs']] upgrade_policy: The upgrade policy.
        :param pulumi.Input[pulumi.InputType['VirtualMachineScaleSetVMProfileArgs']] virtual_machine_profile: The virtual machine profile.
        :param pulumi.Input[str] vm_scale_set_name: The name of the VM scale set to create or update.
        :param pulumi.Input[bool] zone_balance: Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
        :param pulumi.Input[List[pulumi.Input[str]]] zones: The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
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

            __props__['identity'] = identity
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['overprovision'] = overprovision
            __props__['plan'] = plan
            __props__['platform_fault_domain_count'] = platform_fault_domain_count
            __props__['proximity_placement_group'] = proximity_placement_group
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['single_placement_group'] = single_placement_group
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['upgrade_policy'] = upgrade_policy
            __props__['virtual_machine_profile'] = virtual_machine_profile
            if vm_scale_set_name is None:
                raise TypeError("Missing required property 'vm_scale_set_name'")
            __props__['vm_scale_set_name'] = vm_scale_set_name
            __props__['zone_balance'] = zone_balance
            __props__['zones'] = zones
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
            __props__['unique_id'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:compute/latest:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20150615:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20160330:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20160430preview:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20170330:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20171201:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20180601:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20181001:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20190301:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20190701:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20191201:VirtualMachineScaleSet"), pulumi.Alias(type_="azurerm:compute/v20200601:VirtualMachineScaleSet")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualMachineScaleSet, __self__).__init__(
            'azurerm:compute/v20180401:VirtualMachineScaleSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualMachineScaleSet':
        """
        Get an existing VirtualMachineScaleSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualMachineScaleSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.VirtualMachineScaleSetIdentityResponse']]:
        """
        The identity of the virtual machine scale set, if configured.
        """
        return pulumi.get(self, "identity")

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
    @pulumi.getter
    def overprovision(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the Virtual Machine Scale Set should be overprovisioned.
        """
        return pulumi.get(self, "overprovision")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output[Optional['outputs.PlanResponse']]:
        """
        Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="platformFaultDomainCount")
    def platform_fault_domain_count(self) -> pulumi.Output[Optional[float]]:
        """
        Fault Domain count for each placement group.
        """
        return pulumi.get(self, "platform_fault_domain_count")

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
        Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
        """
        return pulumi.get(self, "proximity_placement_group")

    @property
    @pulumi.getter(name="singlePlacementGroup")
    def single_placement_group(self) -> pulumi.Output[Optional[bool]]:
        """
        When true this limits the scale set to a single placement group, of max size 100 virtual machines.
        """
        return pulumi.get(self, "single_placement_group")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        The virtual machine scale set sku.
        """
        return pulumi.get(self, "sku")

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
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID which uniquely identifies a Virtual Machine Scale Set.
        """
        return pulumi.get(self, "unique_id")

    @property
    @pulumi.getter(name="upgradePolicy")
    def upgrade_policy(self) -> pulumi.Output[Optional['outputs.UpgradePolicyResponse']]:
        """
        The upgrade policy.
        """
        return pulumi.get(self, "upgrade_policy")

    @property
    @pulumi.getter(name="virtualMachineProfile")
    def virtual_machine_profile(self) -> pulumi.Output[Optional['outputs.VirtualMachineScaleSetVMProfileResponse']]:
        """
        The virtual machine profile.
        """
        return pulumi.get(self, "virtual_machine_profile")

    @property
    @pulumi.getter(name="zoneBalance")
    def zone_balance(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
        """
        return pulumi.get(self, "zone_balance")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[List[str]]]:
        """
        The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

