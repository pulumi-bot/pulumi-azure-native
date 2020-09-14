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

__all__ = ['HostPool']


class HostPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_rdp_property: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 host_pool_name: Optional[pulumi.Input[str]] = None,
                 host_pool_type: Optional[pulumi.Input[str]] = None,
                 load_balancer_type: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 max_session_limit: Optional[pulumi.Input[float]] = None,
                 personal_desktop_assignment_type: Optional[pulumi.Input[str]] = None,
                 preferred_app_group_type: Optional[pulumi.Input[str]] = None,
                 registration_info: Optional[pulumi.Input[pulumi.InputType['RegistrationInfoArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 ring: Optional[pulumi.Input[float]] = None,
                 sso_context: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 validation_environment: Optional[pulumi.Input[bool]] = None,
                 vm_template: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a HostPool definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] custom_rdp_property: Custom rdp property of HostPool.
        :param pulumi.Input[str] description: Description of HostPool.
        :param pulumi.Input[str] friendly_name: Friendly name of HostPool.
        :param pulumi.Input[str] host_pool_name: The name of the host pool within the specified resource group
        :param pulumi.Input[str] host_pool_type: HostPool type for desktop.
        :param pulumi.Input[str] load_balancer_type: The type of the load balancer.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[float] max_session_limit: The max session limit of HostPool.
        :param pulumi.Input[str] personal_desktop_assignment_type: PersonalDesktopAssignment type for HostPool.
        :param pulumi.Input[str] preferred_app_group_type: The type of preferred application group type, default to Desktop Application Group
        :param pulumi.Input[pulumi.InputType['RegistrationInfoArgs']] registration_info: The registration info of HostPool.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[float] ring: The ring number of HostPool.
        :param pulumi.Input[str] sso_context: Path to keyvault containing ssoContext secret.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[bool] validation_environment: Is validation environment.
        :param pulumi.Input[str] vm_template: VM template for sessionhosts configuration within hostpool.
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

            __props__['custom_rdp_property'] = custom_rdp_property
            __props__['description'] = description
            __props__['friendly_name'] = friendly_name
            if host_pool_name is None:
                raise TypeError("Missing required property 'host_pool_name'")
            __props__['host_pool_name'] = host_pool_name
            if host_pool_type is None:
                raise TypeError("Missing required property 'host_pool_type'")
            __props__['host_pool_type'] = host_pool_type
            if load_balancer_type is None:
                raise TypeError("Missing required property 'load_balancer_type'")
            __props__['load_balancer_type'] = load_balancer_type
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['max_session_limit'] = max_session_limit
            __props__['personal_desktop_assignment_type'] = personal_desktop_assignment_type
            if preferred_app_group_type is None:
                raise TypeError("Missing required property 'preferred_app_group_type'")
            __props__['preferred_app_group_type'] = preferred_app_group_type
            __props__['registration_info'] = registration_info
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['ring'] = ring
            __props__['sso_context'] = sso_context
            __props__['tags'] = tags
            __props__['validation_environment'] = validation_environment
            __props__['vm_template'] = vm_template
            __props__['application_group_references'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:desktopvirtualization/latest:HostPool"), pulumi.Alias(type_="azurerm:desktopvirtualization/v20190123preview:HostPool"), pulumi.Alias(type_="azurerm:desktopvirtualization/v20191210preview:HostPool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(HostPool, __self__).__init__(
            'azurerm:desktopvirtualization/v20190924preview:HostPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HostPool':
        """
        Get an existing HostPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return HostPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationGroupReferences")
    def application_group_references(self) -> pulumi.Output[List[str]]:
        """
        List of applicationGroup links.
        """
        return pulumi.get(self, "application_group_references")

    @property
    @pulumi.getter(name="customRdpProperty")
    def custom_rdp_property(self) -> pulumi.Output[Optional[str]]:
        """
        Custom rdp property of HostPool.
        """
        return pulumi.get(self, "custom_rdp_property")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of HostPool.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[Optional[str]]:
        """
        Friendly name of HostPool.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="hostPoolType")
    def host_pool_type(self) -> pulumi.Output[str]:
        """
        HostPool type for desktop.
        """
        return pulumi.get(self, "host_pool_type")

    @property
    @pulumi.getter(name="loadBalancerType")
    def load_balancer_type(self) -> pulumi.Output[str]:
        """
        The type of the load balancer.
        """
        return pulumi.get(self, "load_balancer_type")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxSessionLimit")
    def max_session_limit(self) -> pulumi.Output[Optional[float]]:
        """
        The max session limit of HostPool.
        """
        return pulumi.get(self, "max_session_limit")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="personalDesktopAssignmentType")
    def personal_desktop_assignment_type(self) -> pulumi.Output[Optional[str]]:
        """
        PersonalDesktopAssignment type for HostPool.
        """
        return pulumi.get(self, "personal_desktop_assignment_type")

    @property
    @pulumi.getter(name="preferredAppGroupType")
    def preferred_app_group_type(self) -> pulumi.Output[str]:
        """
        The type of preferred application group type, default to Desktop Application Group
        """
        return pulumi.get(self, "preferred_app_group_type")

    @property
    @pulumi.getter(name="registrationInfo")
    def registration_info(self) -> pulumi.Output[Optional['outputs.RegistrationInfoResponse']]:
        """
        The registration info of HostPool.
        """
        return pulumi.get(self, "registration_info")

    @property
    @pulumi.getter
    def ring(self) -> pulumi.Output[Optional[float]]:
        """
        The ring number of HostPool.
        """
        return pulumi.get(self, "ring")

    @property
    @pulumi.getter(name="ssoContext")
    def sso_context(self) -> pulumi.Output[Optional[str]]:
        """
        Path to keyvault containing ssoContext secret.
        """
        return pulumi.get(self, "sso_context")

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
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="validationEnvironment")
    def validation_environment(self) -> pulumi.Output[Optional[bool]]:
        """
        Is validation environment.
        """
        return pulumi.get(self, "validation_environment")

    @property
    @pulumi.getter(name="vmTemplate")
    def vm_template(self) -> pulumi.Output[Optional[str]]:
        """
        VM template for sessionhosts configuration within hostpool.
        """
        return pulumi.get(self, "vm_template")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

