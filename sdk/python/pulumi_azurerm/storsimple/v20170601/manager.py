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

__all__ = ['Manager']


class Manager(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cis_intrinsic_settings: Optional[pulumi.Input[pulumi.InputType['ManagerIntrinsicSettingsArgs']]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ManagerSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The StorSimple Manager.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ManagerIntrinsicSettingsArgs']] cis_intrinsic_settings: Represents the type of StorSimple Manager.
        :param pulumi.Input[str] etag: The etag of the manager.
        :param pulumi.Input[str] location: The geo location of the resource.
        :param pulumi.Input[str] name: The manager name
        :param pulumi.Input[str] provisioning_state: Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
        :param pulumi.Input[str] resource_group_name: The resource group name
        :param pulumi.Input[pulumi.InputType['ManagerSkuArgs']] sku: Specifies the Sku.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags attached to the resource.
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

            __props__['cis_intrinsic_settings'] = cis_intrinsic_settings
            __props__['etag'] = etag
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:storsimple/v20161001:Manager")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Manager, __self__).__init__(
            'azurerm:storsimple/v20170601:Manager',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Manager':
        """
        Get an existing Manager resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Manager(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cisIntrinsicSettings")
    def cis_intrinsic_settings(self) -> Optional['outputs.ManagerIntrinsicSettingsResponse']:
        """
        Represents the type of StorSimple Manager.
        """
        return pulumi.get(self, "cis_intrinsic_settings")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        The etag of the manager.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ManagerSkuResponse']:
        """
        Specifies the Sku.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags attached to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

