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

__all__ = ['VirtualNetworkFunction']


class VirtualNetworkFunction(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_application_parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vendor_name: Optional[pulumi.Input[str]] = None,
                 virtual_network_function_name: Optional[pulumi.Input[str]] = None,
                 virtual_network_function_user_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualNetworkFunctionUserConfigurationArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Hybrid network virtual network function resource response.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] device: The reference to the hybrid network device.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[Mapping[str, Any]] managed_application_parameters: The parameters for the managed application.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] sku_name: The sku name for the hybrid network virtual network function.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] vendor_name: The vendor name for the hybrid network virtual network function.
        :param pulumi.Input[str] virtual_network_function_name: Resource name for the hybrid network virtual network function resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualNetworkFunctionUserConfigurationArgs']]]] virtual_network_function_user_configurations: The virtual network function configurations from the user.
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

            __props__['device'] = device
            __props__['etag'] = etag
            __props__['location'] = location
            __props__['managed_application_parameters'] = managed_application_parameters
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['vendor_name'] = vendor_name
            if virtual_network_function_name is None:
                raise TypeError("Missing required property 'virtual_network_function_name'")
            __props__['virtual_network_function_name'] = virtual_network_function_name
            __props__['virtual_network_function_user_configurations'] = virtual_network_function_user_configurations
            __props__['managed_application'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['service_key'] = None
            __props__['sku_type'] = None
            __props__['type'] = None
            __props__['vendor_provisioning_state'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:hybridnetwork/latest:VirtualNetworkFunction")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualNetworkFunction, __self__).__init__(
            'azurerm:hybridnetwork/v20200101preview:VirtualNetworkFunction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualNetworkFunction':
        """
        Get an existing VirtualNetworkFunction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualNetworkFunction(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        The reference to the hybrid network device.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedApplication")
    def managed_application(self) -> pulumi.Output['outputs.SubResourceResponse']:
        """
        The resource URI of the managed application.
        """
        return pulumi.get(self, "managed_application")

    @property
    @pulumi.getter(name="managedApplicationParameters")
    def managed_application_parameters(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The parameters for the managed application.
        """
        return pulumi.get(self, "managed_application_parameters")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the hybrid network virtual network function resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serviceKey")
    def service_key(self) -> pulumi.Output[str]:
        """
        The service key for the virtual network function resource.
        """
        return pulumi.get(self, "service_key")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Output[Optional[str]]:
        """
        The sku name for the hybrid network virtual network function.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter(name="skuType")
    def sku_type(self) -> pulumi.Output[str]:
        """
        The sku type for the hybrid network virtual network function.
        """
        return pulumi.get(self, "sku_type")

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
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vendorName")
    def vendor_name(self) -> pulumi.Output[Optional[str]]:
        """
        The vendor name for the hybrid network virtual network function.
        """
        return pulumi.get(self, "vendor_name")

    @property
    @pulumi.getter(name="vendorProvisioningState")
    def vendor_provisioning_state(self) -> pulumi.Output[str]:
        """
        The vendor provisioning state for the virtual network function resource.
        """
        return pulumi.get(self, "vendor_provisioning_state")

    @property
    @pulumi.getter(name="virtualNetworkFunctionUserConfigurations")
    def virtual_network_function_user_configurations(self) -> pulumi.Output[Optional[List['outputs.VirtualNetworkFunctionUserConfigurationResponse']]]:
        """
        The virtual network function configurations from the user.
        """
        return pulumi.get(self, "virtual_network_function_user_configurations")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

