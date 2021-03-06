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

__all__ = ['DataPool']


class DataPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 data_pool_name: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataPoolLocationArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ADP Data Pool

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the ADP account
        :param pulumi.Input[str] data_pool_name: The name of the Data Pool
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataPoolLocationArgs']]]] locations: Gets or sets the collection of locations where Data Pool resources should be created
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
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

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['data_pool_name'] = data_pool_name
            if locations is None and not opts.urn:
                raise TypeError("Missing required property 'locations'")
            __props__['locations'] = locations
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['data_pool_id'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['system_data'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:autonomousdevelopmentplatform/v20210201preview:DataPool"), pulumi.Alias(type_="azure-native:autonomousdevelopmentplatform:DataPool"), pulumi.Alias(type_="azure-nextgen:autonomousdevelopmentplatform:DataPool"), pulumi.Alias(type_="azure-native:autonomousdevelopmentplatform/v20200701preview:DataPool"), pulumi.Alias(type_="azure-nextgen:autonomousdevelopmentplatform/v20200701preview:DataPool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DataPool, __self__).__init__(
            'azure-native:autonomousdevelopmentplatform/v20210201preview:DataPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataPool':
        """
        Get an existing DataPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data_pool_id"] = None
        __props__["locations"] = None
        __props__["name"] = None
        __props__["provisioning_state"] = None
        __props__["system_data"] = None
        __props__["type"] = None
        return DataPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataPoolId")
    def data_pool_id(self) -> pulumi.Output[str]:
        """
        The Data Pool's data-plane ID
        """
        return pulumi.get(self, "data_pool_id")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Sequence['outputs.DataPoolLocationResponse']]:
        """
        Gets or sets the collection of locations where Data Pool resources should be created
        """
        return pulumi.get(self, "locations")

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
        Gets the status of the data pool at the time the operation was called
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        The system meta data relating to this resource
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

