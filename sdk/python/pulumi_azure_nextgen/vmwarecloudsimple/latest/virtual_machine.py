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
                 amount_of_ram: Optional[pulumi.Input[int]] = None,
                 customization: Optional[pulumi.Input[pulumi.InputType['GuestOSCustomizationArgs']]] = None,
                 disks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualDiskArgs']]]]] = None,
                 expose_to_guest_vm: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 nics: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNicArgs']]]]] = None,
                 number_of_cores: Optional[pulumi.Input[int]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 private_cloud_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_pool: Optional[pulumi.Input[pulumi.InputType['ResourcePoolArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 v_sphere_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 virtual_machine_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Virtual machine model
        Latest API Version: 2019-04-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] amount_of_ram: The amount of memory
        :param pulumi.Input[pulumi.InputType['GuestOSCustomizationArgs']] customization: Virtual machine properties
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualDiskArgs']]]] disks: The list of Virtual Disks
        :param pulumi.Input[bool] expose_to_guest_vm: Expose Guest OS or not
        :param pulumi.Input[str] location: Azure region
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualNicArgs']]]] nics: The list of Virtual NICs
        :param pulumi.Input[int] number_of_cores: The number of CPU cores
        :param pulumi.Input[str] password: Password for login. Deprecated - use customization property
        :param pulumi.Input[str] private_cloud_id: Private Cloud Id
        :param pulumi.Input[str] resource_group_name: The name of the resource group
        :param pulumi.Input[pulumi.InputType['ResourcePoolArgs']] resource_pool: Virtual Machines Resource Pool
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The list of tags
        :param pulumi.Input[str] template_id: Virtual Machine Template Id
        :param pulumi.Input[str] username: Username for login. Deprecated - use customization property
        :param pulumi.Input[Sequence[pulumi.Input[str]]] v_sphere_networks: The list of Virtual VSphere Networks
        :param pulumi.Input[str] virtual_machine_name: virtual machine name
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

            if amount_of_ram is None and not opts.urn:
                raise TypeError("Missing required property 'amount_of_ram'")
            __props__['amount_of_ram'] = amount_of_ram
            __props__['customization'] = customization
            __props__['disks'] = disks
            __props__['expose_to_guest_vm'] = expose_to_guest_vm
            __props__['location'] = location
            __props__['nics'] = nics
            if number_of_cores is None and not opts.urn:
                raise TypeError("Missing required property 'number_of_cores'")
            __props__['number_of_cores'] = number_of_cores
            __props__['password'] = password
            if private_cloud_id is None and not opts.urn:
                raise TypeError("Missing required property 'private_cloud_id'")
            __props__['private_cloud_id'] = private_cloud_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_pool'] = resource_pool
            __props__['tags'] = tags
            __props__['template_id'] = template_id
            __props__['username'] = username
            __props__['v_sphere_networks'] = v_sphere_networks
            if virtual_machine_name is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_machine_name'")
            __props__['virtual_machine_name'] = virtual_machine_name
            __props__['controllers'] = None
            __props__['dnsname'] = None
            __props__['folder'] = None
            __props__['guest_os'] = None
            __props__['guest_os_type'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['public_ip'] = None
            __props__['status'] = None
            __props__['type'] = None
            __props__['vm_id'] = None
            __props__['vmwaretools'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:vmwarecloudsimple/v20190401:VirtualMachine")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualMachine, __self__).__init__(
            'azure-nextgen:vmwarecloudsimple/latest:VirtualMachine',
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
    @pulumi.getter(name="amountOfRam")
    def amount_of_ram(self) -> pulumi.Output[int]:
        """
        The amount of memory
        """
        return pulumi.get(self, "amount_of_ram")

    @property
    @pulumi.getter
    def controllers(self) -> pulumi.Output[Sequence['outputs.VirtualDiskControllerResponse']]:
        """
        The list of Virtual Disks' Controllers
        """
        return pulumi.get(self, "controllers")

    @property
    @pulumi.getter
    def customization(self) -> pulumi.Output[Optional['outputs.GuestOSCustomizationResponse']]:
        """
        Virtual machine properties
        """
        return pulumi.get(self, "customization")

    @property
    @pulumi.getter
    def disks(self) -> pulumi.Output[Optional[Sequence['outputs.VirtualDiskResponse']]]:
        """
        The list of Virtual Disks
        """
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter
    def dnsname(self) -> pulumi.Output[str]:
        """
        The DNS name of Virtual Machine in VCenter
        """
        return pulumi.get(self, "dnsname")

    @property
    @pulumi.getter(name="exposeToGuestVM")
    def expose_to_guest_vm(self) -> pulumi.Output[Optional[bool]]:
        """
        Expose Guest OS or not
        """
        return pulumi.get(self, "expose_to_guest_vm")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[str]:
        """
        The path to virtual machine folder in VCenter
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="guestOS")
    def guest_os(self) -> pulumi.Output[str]:
        """
        The name of Guest OS
        """
        return pulumi.get(self, "guest_os")

    @property
    @pulumi.getter(name="guestOSType")
    def guest_os_type(self) -> pulumi.Output[str]:
        """
        The Guest OS type
        """
        return pulumi.get(self, "guest_os_type")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Azure region
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        {virtualMachineName}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nics(self) -> pulumi.Output[Optional[Sequence['outputs.VirtualNicResponse']]]:
        """
        The list of Virtual NICs
        """
        return pulumi.get(self, "nics")

    @property
    @pulumi.getter(name="numberOfCores")
    def number_of_cores(self) -> pulumi.Output[int]:
        """
        The number of CPU cores
        """
        return pulumi.get(self, "number_of_cores")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Password for login. Deprecated - use customization property
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="privateCloudId")
    def private_cloud_id(self) -> pulumi.Output[str]:
        """
        Private Cloud Id
        """
        return pulumi.get(self, "private_cloud_id")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning status of the resource
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicIP")
    def public_ip(self) -> pulumi.Output[str]:
        """
        The public ip of Virtual Machine
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="resourcePool")
    def resource_pool(self) -> pulumi.Output[Optional['outputs.ResourcePoolResponse']]:
        """
        Virtual Machines Resource Pool
        """
        return pulumi.get(self, "resource_pool")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of Virtual machine
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The list of tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[Optional[str]]:
        """
        Virtual Machine Template Id
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        {resourceProviderNamespace}/{resourceType}
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        Username for login. Deprecated - use customization property
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="vSphereNetworks")
    def v_sphere_networks(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of Virtual VSphere Networks
        """
        return pulumi.get(self, "v_sphere_networks")

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> pulumi.Output[str]:
        """
        The internal id of Virtual Machine in VCenter
        """
        return pulumi.get(self, "vm_id")

    @property
    @pulumi.getter
    def vmwaretools(self) -> pulumi.Output[str]:
        """
        VMware tools version
        """
        return pulumi.get(self, "vmwaretools")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

