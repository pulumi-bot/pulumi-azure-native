# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Machine']


class Machine(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_public_key: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['MachineIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 location_data: Optional[pulumi.Input[pulumi.InputType['LocationDataArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Describes a hybrid machine.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_public_key: Public Key that the client provides to be used during initial resource onboarding
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[pulumi.InputType['LocationDataArgs']] location_data: Metadata pertaining to the geographic location of the resource.
        :param pulumi.Input[str] name: The name of the hybrid machine.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] vm_id: Specifies the hybrid machine unique ID.
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

            __props__['client_public_key'] = client_public_key
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['location_data'] = location_data
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['vm_id'] = vm_id
            __props__['ad_fqdn'] = None
            __props__['agent_version'] = None
            __props__['display_name'] = None
            __props__['dns_fqdn'] = None
            __props__['domain_name'] = None
            __props__['error_details'] = None
            __props__['extensions'] = None
            __props__['last_status_change'] = None
            __props__['machine_fqdn'] = None
            __props__['os_name'] = None
            __props__['os_profile'] = None
            __props__['os_sku'] = None
            __props__['os_version'] = None
            __props__['provisioning_state'] = None
            __props__['status'] = None
            __props__['type'] = None
            __props__['vm_uuid'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:hybridcompute/latest:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20190318preview:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20190802preview:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20191212:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20200730preview:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20200815preview:Machine")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Machine, __self__).__init__(
            'azure-nextgen:hybridcompute/v20200802:Machine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Machine':
        """
        Get an existing Machine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Machine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adFqdn")
    def ad_fqdn(self) -> pulumi.Output[str]:
        """
        Specifies the AD fully qualified display name.
        """
        return pulumi.get(self, "ad_fqdn")

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> pulumi.Output[str]:
        """
        The hybrid machine agent full version.
        """
        return pulumi.get(self, "agent_version")

    @property
    @pulumi.getter(name="clientPublicKey")
    def client_public_key(self) -> pulumi.Output[Optional[str]]:
        """
        Public Key that the client provides to be used during initial resource onboarding
        """
        return pulumi.get(self, "client_public_key")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Specifies the hybrid machine display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="dnsFqdn")
    def dns_fqdn(self) -> pulumi.Output[str]:
        """
        Specifies the DNS fully qualified display name.
        """
        return pulumi.get(self, "dns_fqdn")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        Specifies the Windows domain name.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="errorDetails")
    def error_details(self) -> pulumi.Output[Sequence['outputs.ErrorDetailResponse']]:
        """
        Details about the error state.
        """
        return pulumi.get(self, "error_details")

    @property
    @pulumi.getter
    def extensions(self) -> pulumi.Output[Sequence['outputs.MachineExtensionInstanceViewResponse']]:
        """
        Machine Extensions information
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.MachineResponseIdentity']]:
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="lastStatusChange")
    def last_status_change(self) -> pulumi.Output[str]:
        """
        The time of the last status change.
        """
        return pulumi.get(self, "last_status_change")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="locationData")
    def location_data(self) -> pulumi.Output[Optional['outputs.LocationDataResponse']]:
        """
        Metadata pertaining to the geographic location of the resource.
        """
        return pulumi.get(self, "location_data")

    @property
    @pulumi.getter(name="machineFqdn")
    def machine_fqdn(self) -> pulumi.Output[str]:
        """
        Specifies the hybrid machine FQDN.
        """
        return pulumi.get(self, "machine_fqdn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osName")
    def os_name(self) -> pulumi.Output[str]:
        """
        The Operating System running on the hybrid machine.
        """
        return pulumi.get(self, "os_name")

    @property
    @pulumi.getter(name="osProfile")
    def os_profile(self) -> pulumi.Output[Optional['outputs.MachinePropertiesResponseOsProfile']]:
        """
        Specifies the operating system settings for the hybrid machine.
        """
        return pulumi.get(self, "os_profile")

    @property
    @pulumi.getter(name="osSku")
    def os_sku(self) -> pulumi.Output[str]:
        """
        Specifies the Operating System product SKU.
        """
        return pulumi.get(self, "os_sku")

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> pulumi.Output[str]:
        """
        The version of Operating System running on the hybrid machine.
        """
        return pulumi.get(self, "os_version")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the hybrid machine agent.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the hybrid machine unique ID.
        """
        return pulumi.get(self, "vm_id")

    @property
    @pulumi.getter(name="vmUuid")
    def vm_uuid(self) -> pulumi.Output[str]:
        """
        Specifies the Arc Machine's unique SMBIOS ID
        """
        return pulumi.get(self, "vm_uuid")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

