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

__all__ = ['HanaInstance']


class HanaInstance(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hana_instance_id: Optional[pulumi.Input[str]] = None,
                 hana_instance_name: Optional[pulumi.Input[str]] = None,
                 hardware_profile: Optional[pulumi.Input[pulumi.InputType['HardwareProfileArgs']]] = None,
                 hw_revision: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 network_profile: Optional[pulumi.Input[pulumi.InputType['NetworkProfileArgs']]] = None,
                 os_profile: Optional[pulumi.Input[pulumi.InputType['OSProfileArgs']]] = None,
                 partner_node_id: Optional[pulumi.Input[str]] = None,
                 power_state: Optional[pulumi.Input[Union[str, 'HanaInstancePowerStateEnum']]] = None,
                 provisioning_state: Optional[pulumi.Input[Union[str, 'HanaProvisioningStatesEnum']]] = None,
                 proximity_placement_group: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_profile: Optional[pulumi.Input[pulumi.InputType['StorageProfileArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        HANA instance info on Azure (ARM properties and HANA properties)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hana_instance_id: Specifies the HANA instance unique ID.
        :param pulumi.Input[str] hana_instance_name: Name of the SAP HANA on Azure instance.
        :param pulumi.Input[pulumi.InputType['HardwareProfileArgs']] hardware_profile: Specifies the hardware settings for the HANA instance.
        :param pulumi.Input[str] hw_revision: Hardware revision of a HANA instance
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[pulumi.InputType['NetworkProfileArgs']] network_profile: Specifies the network settings for the HANA instance.
        :param pulumi.Input[pulumi.InputType['OSProfileArgs']] os_profile: Specifies the operating system settings for the HANA instance.
        :param pulumi.Input[str] partner_node_id: ARM ID of another HanaInstance that will share a network with this HanaInstance
        :param pulumi.Input[Union[str, 'HanaInstancePowerStateEnum']] power_state: Resource power state
        :param pulumi.Input[Union[str, 'HanaProvisioningStatesEnum']] provisioning_state: State of provisioning of the HanaInstance
        :param pulumi.Input[str] proximity_placement_group: Resource proximity placement group
        :param pulumi.Input[str] resource_group_name: Name of the resource group.
        :param pulumi.Input[pulumi.InputType['StorageProfileArgs']] storage_profile: Specifies the storage settings for the HANA instance disks.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
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

            __props__['hana_instance_id'] = hana_instance_id
            __props__['hana_instance_name'] = hana_instance_name
            __props__['hardware_profile'] = hardware_profile
            __props__['hw_revision'] = hw_revision
            __props__['location'] = location
            __props__['network_profile'] = network_profile
            __props__['os_profile'] = os_profile
            __props__['partner_node_id'] = partner_node_id
            __props__['power_state'] = power_state
            __props__['provisioning_state'] = provisioning_state
            __props__['proximity_placement_group'] = proximity_placement_group
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['storage_profile'] = storage_profile
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:hanaonazure/v20171103preview:HanaInstance"), pulumi.Alias(type_="azure-native:hanaonazure:HanaInstance"), pulumi.Alias(type_="azure-nextgen:hanaonazure:HanaInstance")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(HanaInstance, __self__).__init__(
            'azure-native:hanaonazure/v20171103preview:HanaInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HanaInstance':
        """
        Get an existing HanaInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["hana_instance_id"] = None
        __props__["hardware_profile"] = None
        __props__["hw_revision"] = None
        __props__["location"] = None
        __props__["name"] = None
        __props__["network_profile"] = None
        __props__["os_profile"] = None
        __props__["partner_node_id"] = None
        __props__["power_state"] = None
        __props__["provisioning_state"] = None
        __props__["proximity_placement_group"] = None
        __props__["storage_profile"] = None
        __props__["tags"] = None
        __props__["type"] = None
        return HanaInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="hanaInstanceId")
    def hana_instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the HANA instance unique ID.
        """
        return pulumi.get(self, "hana_instance_id")

    @property
    @pulumi.getter(name="hardwareProfile")
    def hardware_profile(self) -> pulumi.Output[Optional['outputs.HardwareProfileResponse']]:
        """
        Specifies the hardware settings for the HANA instance.
        """
        return pulumi.get(self, "hardware_profile")

    @property
    @pulumi.getter(name="hwRevision")
    def hw_revision(self) -> pulumi.Output[Optional[str]]:
        """
        Hardware revision of a HANA instance
        """
        return pulumi.get(self, "hw_revision")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
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
        Specifies the network settings for the HANA instance.
        """
        return pulumi.get(self, "network_profile")

    @property
    @pulumi.getter(name="osProfile")
    def os_profile(self) -> pulumi.Output[Optional['outputs.OSProfileResponse']]:
        """
        Specifies the operating system settings for the HANA instance.
        """
        return pulumi.get(self, "os_profile")

    @property
    @pulumi.getter(name="partnerNodeId")
    def partner_node_id(self) -> pulumi.Output[Optional[str]]:
        """
        ARM ID of another HanaInstance that will share a network with this HanaInstance
        """
        return pulumi.get(self, "partner_node_id")

    @property
    @pulumi.getter(name="powerState")
    def power_state(self) -> pulumi.Output[Optional[str]]:
        """
        Resource power state
        """
        return pulumi.get(self, "power_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        State of provisioning of the HanaInstance
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="proximityPlacementGroup")
    def proximity_placement_group(self) -> pulumi.Output[Optional[str]]:
        """
        Resource proximity placement group
        """
        return pulumi.get(self, "proximity_placement_group")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> pulumi.Output[Optional['outputs.StorageProfileResponse']]:
        """
        Specifies the storage settings for the HANA instance disks.
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

