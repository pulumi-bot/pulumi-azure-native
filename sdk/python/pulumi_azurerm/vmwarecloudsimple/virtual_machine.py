# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class VirtualMachine(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Azure region
    """
    name: pulumi.Output[str]
    """
    {virtualMachineName}
    """
    properties: pulumi.Output[dict]
    """
    Virtual machine properties
      * `amount_of_ram` (`float`) - The amount of memory
      * `controllers` (`list`) - The list of Virtual Disks' Controllers
        * `id` (`str`) - Controller's id
        * `name` (`str`) - The display name of Controller
        * `sub_type` (`str`) - dik controller subtype (VMWARE_PARAVIRTUAL, BUS_PARALLEL, LSI_PARALLEL, LSI_SAS)
        * `type` (`str`) - disk controller type (SCSI)

      * `customization` (`dict`) - Virtual machine properties
        * `dns_servers` (`list`) - List of dns servers to use
        * `host_name` (`str`) - Virtual Machine hostname
        * `password` (`str`) - Password for login
        * `policy_id` (`str`) - id of customization policy
        * `username` (`str`) - Username for login

      * `disks` (`list`) - The list of Virtual Disks
        * `controller_id` (`str`) - Disk's Controller id
        * `independence_mode` (`str`) - Disk's independence mode type
        * `total_size` (`float`) - Disk's total size
        * `virtual_disk_id` (`str`) - Disk's id
        * `virtual_disk_name` (`str`) - Disk's display name

      * `dnsname` (`str`) - The DNS name of Virtual Machine in VCenter
      * `expose_to_guest_vm` (`bool`) - Expose Guest OS or not
      * `folder` (`str`) - The path to virtual machine folder in VCenter
      * `guest_os` (`str`) - The name of Guest OS
      * `guest_os_type` (`str`) - The Guest OS type
      * `nics` (`list`) - The list of Virtual NICs
        * `customization` (`dict`) - guest OS customization for nic
          * `allocation` (`str`) - IP address allocation method
          * `dns_servers` (`list`) - List of dns servers to use
          * `gateway` (`list`) - Gateway addresses assigned to nic
          * `ip_address` (`str`) - Static ip address for nic
          * `mask` (`str`) - Network mask for nic
          * `primary_wins_server` (`str`) - primary WINS server for Windows
          * `secondary_wins_server` (`str`) - secondary WINS server for Windows

        * `ip_addresses` (`list`) - NIC ip address
        * `mac_address` (`str`) - NIC MAC address
        * `network` (`dict`) - Virtual Network
          * `assignable` (`bool`) - can be used in vm creation/deletion
          * `id` (`str`) - virtual network id (privateCloudId:vsphereId)
          * `location` (`str`) - Azure region
          * `name` (`str`) - {VirtualNetworkName}
          * `properties` (`dict`) - Virtual Network properties
            * `private_cloud_id` (`str`) - The Private Cloud id

          * `type` (`str`) - {resourceProviderNamespace}/{resourceType}

        * `nic_type` (`str`) - NIC type
        * `power_on_boot` (`bool`) - Is NIC powered on/off on boot
        * `virtual_nic_id` (`str`) - NIC id
        * `virtual_nic_name` (`str`) - NIC name

      * `number_of_cores` (`float`) - The number of CPU cores
      * `password` (`str`) - Password for login. Deprecated - use customization property
      * `private_cloud_id` (`str`) - Private Cloud Id
      * `provisioning_state` (`str`) - The provisioning status of the resource
      * `public_ip` (`str`) - The public ip of Virtual Machine
      * `resource_pool` (`dict`) - Virtual Machines Resource Pool
        * `id` (`str`) - resource pool id (privateCloudId:vsphereId)
        * `location` (`str`) - Azure region
        * `name` (`str`) - {ResourcePoolName}
        * `private_cloud_id` (`str`) - The Private Cloud Id
        * `properties` (`dict`) - Resource pool properties
          * `full_name` (`str`) - Hierarchical resource pool name

        * `type` (`str`) - {resourceProviderNamespace}/{resourceType}

      * `status` (`str`) - The status of Virtual machine
      * `template_id` (`str`) - Virtual Machine Template Id
      * `username` (`str`) - Username for login. Deprecated - use customization property
      * `v_sphere_networks` (`list`) - The list of Virtual VSphere Networks
      * `vm_id` (`str`) - The internal id of Virtual Machine in VCenter
      * `vmwaretools` (`str`) - VMware tools version
    """
    tags: pulumi.Output[dict]
    """
    The list of tags
    """
    type: pulumi.Output[str]
    """
    {resourceProviderNamespace}/{resourceType}
    """
    def __init__(__self__, resource_name, opts=None, referer=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Virtual machine model

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] referer: referer url
        :param pulumi.Input[str] location: Azure region
        :param pulumi.Input[str] name: virtual machine name
        :param pulumi.Input[dict] properties: Virtual machine properties
        :param pulumi.Input[str] resource_group_name: The name of the resource group
        :param pulumi.Input[dict] tags: The list of tags

        The **properties** object supports the following:

          * `amount_of_ram` (`pulumi.Input[float]`) - The amount of memory
          * `customization` (`pulumi.Input[dict]`) - Virtual machine properties
            * `dns_servers` (`pulumi.Input[list]`) - List of dns servers to use
            * `host_name` (`pulumi.Input[str]`) - Virtual Machine hostname
            * `password` (`pulumi.Input[str]`) - Password for login
            * `policy_id` (`pulumi.Input[str]`) - id of customization policy
            * `username` (`pulumi.Input[str]`) - Username for login

          * `disks` (`pulumi.Input[list]`) - The list of Virtual Disks
            * `controller_id` (`pulumi.Input[str]`) - Disk's Controller id
            * `independence_mode` (`pulumi.Input[str]`) - Disk's independence mode type
            * `total_size` (`pulumi.Input[float]`) - Disk's total size
            * `virtual_disk_id` (`pulumi.Input[str]`) - Disk's id

          * `expose_to_guest_vm` (`pulumi.Input[bool]`) - Expose Guest OS or not
          * `nics` (`pulumi.Input[list]`) - The list of Virtual NICs
            * `customization` (`pulumi.Input[dict]`) - guest OS customization for nic
              * `allocation` (`pulumi.Input[str]`) - IP address allocation method
              * `dns_servers` (`pulumi.Input[list]`) - List of dns servers to use
              * `gateway` (`pulumi.Input[list]`) - Gateway addresses assigned to nic
              * `ip_address` (`pulumi.Input[str]`) - Static ip address for nic
              * `mask` (`pulumi.Input[str]`) - Network mask for nic
              * `primary_wins_server` (`pulumi.Input[str]`) - primary WINS server for Windows
              * `secondary_wins_server` (`pulumi.Input[str]`) - secondary WINS server for Windows

            * `ip_addresses` (`pulumi.Input[list]`) - NIC ip address
            * `mac_address` (`pulumi.Input[str]`) - NIC MAC address
            * `network` (`pulumi.Input[dict]`) - Virtual Network
              * `id` (`pulumi.Input[str]`) - virtual network id (privateCloudId:vsphereId)
              * `properties` (`pulumi.Input[dict]`) - Virtual Network properties

            * `nic_type` (`pulumi.Input[str]`) - NIC type
            * `power_on_boot` (`pulumi.Input[bool]`) - Is NIC powered on/off on boot
            * `virtual_nic_id` (`pulumi.Input[str]`) - NIC id

          * `number_of_cores` (`pulumi.Input[float]`) - The number of CPU cores
          * `password` (`pulumi.Input[str]`) - Password for login. Deprecated - use customization property
          * `private_cloud_id` (`pulumi.Input[str]`) - Private Cloud Id
          * `resource_pool` (`pulumi.Input[dict]`) - Virtual Machines Resource Pool
            * `id` (`pulumi.Input[str]`) - resource pool id (privateCloudId:vsphereId)
            * `properties` (`pulumi.Input[dict]`) - Resource pool properties

          * `template_id` (`pulumi.Input[str]`) - Virtual Machine Template Id
          * `username` (`pulumi.Input[str]`) - Username for login. Deprecated - use customization property
          * `v_sphere_networks` (`pulumi.Input[list]`) - The list of Virtual VSphere Networks
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if referer is None:
                raise TypeError("Missing required property 'referer'")
            __props__['referer'] = referer
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
            __props__['tags'] = tags
            __props__['type'] = None
        super(VirtualMachine, __self__).__init__(
            'azurerm:vmwarecloudsimple:VirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualMachine(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
