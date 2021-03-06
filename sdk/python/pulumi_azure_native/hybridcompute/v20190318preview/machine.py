# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['Machine']


class Machine(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_public_key: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 physical_location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Describes a hybrid machine.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_public_key: Public Key that the client provides to be used during initial resource onboarding
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the hybrid machine.
        :param pulumi.Input[str] physical_location: Resource's Physical Location
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[str] type: The identity type.
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
            __props__['location'] = location
            __props__['name'] = name
            __props__['physical_location'] = physical_location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = type
            __props__['agent_version'] = None
            __props__['display_name'] = None
            __props__['error_details'] = None
            __props__['last_status_change'] = None
            __props__['machine_fqdn'] = None
            __props__['os_name'] = None
            __props__['os_profile'] = None
            __props__['os_version'] = None
            __props__['principal_id'] = None
            __props__['provisioning_state'] = None
            __props__['status'] = None
            __props__['tenant_id'] = None
            __props__['vm_id'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:hybridcompute/v20190318preview:Machine"), pulumi.Alias(type_="azure-native:hybridcompute:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute:Machine"), pulumi.Alias(type_="azure-native:hybridcompute/latest:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/latest:Machine"), pulumi.Alias(type_="azure-native:hybridcompute/v20190802preview:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20190802preview:Machine"), pulumi.Alias(type_="azure-native:hybridcompute/v20191212:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20191212:Machine"), pulumi.Alias(type_="azure-native:hybridcompute/v20200730preview:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20200730preview:Machine"), pulumi.Alias(type_="azure-native:hybridcompute/v20200802:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20200802:Machine"), pulumi.Alias(type_="azure-native:hybridcompute/v20200815preview:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20200815preview:Machine"), pulumi.Alias(type_="azure-native:hybridcompute/v20210128preview:Machine"), pulumi.Alias(type_="azure-nextgen:hybridcompute/v20210128preview:Machine")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Machine, __self__).__init__(
            'azure-native:hybridcompute/v20190318preview:Machine',
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

        __props__["agent_version"] = None
        __props__["client_public_key"] = None
        __props__["display_name"] = None
        __props__["error_details"] = None
        __props__["last_status_change"] = None
        __props__["location"] = None
        __props__["machine_fqdn"] = None
        __props__["name"] = None
        __props__["os_name"] = None
        __props__["os_profile"] = None
        __props__["os_version"] = None
        __props__["physical_location"] = None
        __props__["principal_id"] = None
        __props__["provisioning_state"] = None
        __props__["status"] = None
        __props__["tags"] = None
        __props__["tenant_id"] = None
        __props__["type"] = None
        __props__["vm_id"] = None
        return Machine(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="errorDetails")
    def error_details(self) -> pulumi.Output[Sequence['outputs.ErrorDetailResponse']]:
        """
        Details about the error state.
        """
        return pulumi.get(self, "error_details")

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
        Resource location
        """
        return pulumi.get(self, "location")

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
        Resource name
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
    def os_profile(self) -> pulumi.Output['outputs.OSProfileResponse']:
        """
        Specifies the operating system settings for the hybrid machine.
        """
        return pulumi.get(self, "os_profile")

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> pulumi.Output[str]:
        """
        The version of Operating System running on the hybrid machine.
        """
        return pulumi.get(self, "os_version")

    @property
    @pulumi.getter(name="physicalLocation")
    def physical_location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource's Physical Location
        """
        return pulumi.get(self, "physical_location")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Output[str]:
        """
        The identity's principal id.
        """
        return pulumi.get(self, "principal_id")

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
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The identity's tenant id.
        """
        return pulumi.get(self, "tenant_id")

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
        Specifies the hybrid machine unique ID.
        """
        return pulumi.get(self, "vm_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

