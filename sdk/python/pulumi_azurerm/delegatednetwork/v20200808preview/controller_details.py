# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['ControllerDetails']


class ControllerDetails(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 controller_type: Optional[pulumi.Input[str]] = None,
                 kubernetes_properties: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['KubernetesPropertiesArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents an instance of a DNC controller.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] controller_type: Type of Delegated controller.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['KubernetesPropertiesArgs']]]] kubernetes_properties: properties of kubernetes clusters
        :param pulumi.Input[str] resource_group_name: The name of the Azure Resource group of which a given DelegatedNetwork resource is part. This name must be at least 1 character in length, and no more than 90.
        :param pulumi.Input[str] resource_name_: The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
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

            __props__['controller_type'] = controller_type
            __props__['kubernetes_properties'] = kubernetes_properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['dnc_app_id'] = None
            __props__['dnc_endpoint'] = None
            __props__['location'] = None
            __props__['name'] = None
            __props__['resource_guid'] = None
            __props__['state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:delegatednetwork/latest:ControllerDetails")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ControllerDetails, __self__).__init__(
            'azurerm:delegatednetwork/v20200808preview:ControllerDetails',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ControllerDetails':
        """
        Get an existing ControllerDetails resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ControllerDetails(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dncAppID")
    def dnc_app_id(self) -> pulumi.Output[Optional[str]]:
        """
        Get controller AAD ID.
        """
        return pulumi.get(self, "dnc_app_id")

    @property
    @pulumi.getter(name="dncEndpoint")
    def dnc_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        Dnc Endpoint url.
        """
        return pulumi.get(self, "dnc_endpoint")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Location of the DNC controller resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the DNC controller resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets resource GUID property of the controller resource.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of dnc controller resource.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the DNC controller  resource.(Microsoft.DelegatedNetwork/controller)
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

