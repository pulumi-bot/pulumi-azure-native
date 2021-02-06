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

__all__ = ['FileServer']


class FileServer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_disks: Optional[pulumi.Input[pulumi.InputType['DataDisksArgs']]] = None,
                 file_server_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 ssh_configuration: Optional[pulumi.Input[pulumi.InputType['SshConfigurationArgs']]] = None,
                 subnet: Optional[pulumi.Input[pulumi.InputType['ResourceIdArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vm_size: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Contains information about the File Server.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DataDisksArgs']] data_disks: Settings for the data disk which would be created for the File Server.
        :param pulumi.Input[str] file_server_name: The name of the file server within the specified resource group. File server names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
        :param pulumi.Input[str] location: The region in which to create the File Server.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[pulumi.InputType['SshConfigurationArgs']] ssh_configuration: SSH configuration settings for the VM
        :param pulumi.Input[pulumi.InputType['ResourceIdArgs']] subnet: Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The user specified tags associated with the File Server.
        :param pulumi.Input[str] vm_size: For information about available VM sizes for fileservers from the Virtual Machines Marketplace, see Sizes for Virtual Machines (Linux).
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

            if data_disks is None and not opts.urn:
                raise TypeError("Missing required property 'data_disks'")
            __props__['data_disks'] = data_disks
            if file_server_name is None and not opts.urn:
                raise TypeError("Missing required property 'file_server_name'")
            __props__['file_server_name'] = file_server_name
            __props__['location'] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if ssh_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'ssh_configuration'")
            __props__['ssh_configuration'] = ssh_configuration
            __props__['subnet'] = subnet
            __props__['tags'] = tags
            if vm_size is None and not opts.urn:
                raise TypeError("Missing required property 'vm_size'")
            __props__['vm_size'] = vm_size
            __props__['creation_time'] = None
            __props__['mount_settings'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['provisioning_state_transition_time'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:batchai/v20180301:FileServer")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(FileServer, __self__).__init__(
            'azure-nextgen:batchai/v20170901preview:FileServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FileServer':
        """
        Get an existing FileServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return FileServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="dataDisks")
    def data_disks(self) -> pulumi.Output[Optional['outputs.DataDisksResponse']]:
        """
        Settings for the data disk which would be created for the File Server.
        """
        return pulumi.get(self, "data_disks")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the resource
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mountSettings")
    def mount_settings(self) -> pulumi.Output['outputs.MountSettingsResponse']:
        """
        Details of the File Server.
        """
        return pulumi.get(self, "mount_settings")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Possible values: creating - The File Server is getting created. updating - The File Server creation has been accepted and it is getting updated. deleting - The user has requested that the File Server be deleted, and it is in the process of being deleted. failed - The File Server creation has failed with the specified errorCode. Details about the error code are specified in the message field. succeeded - The File Server creation has succeeded.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="provisioningStateTransitionTime")
    def provisioning_state_transition_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "provisioning_state_transition_time")

    @property
    @pulumi.getter(name="sshConfiguration")
    def ssh_configuration(self) -> pulumi.Output[Optional['outputs.SshConfigurationResponse']]:
        """
        SSH configuration settings for the VM
        """
        return pulumi.get(self, "ssh_configuration")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[Optional['outputs.ResourceIdResponse']]:
        """
        Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The tags of the resource
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> pulumi.Output[Optional[str]]:
        """
        For information about available VM sizes for File Server from the Virtual Machines Marketplace, see Sizes for Virtual Machines (Linux).
        """
        return pulumi.get(self, "vm_size")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

