# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['Key']


class Key(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 is_active_cmk: Optional[pulumi.Input[bool]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_vault_url: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A workspace key

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] is_active_cmk: Used to activate the workspace after a customer managed key is provided.
        :param pulumi.Input[str] key_name: The name of the workspace key
        :param pulumi.Input[str] key_vault_url: The Key Vault Url of the workspace key.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] workspace_name: The name of the workspace
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

            __props__['is_active_cmk'] = is_active_cmk
            __props__['key_name'] = key_name
            __props__['key_vault_url'] = key_vault_url
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:synapse/v20190601preview:Key"), pulumi.Alias(type_="azure-native:synapse:Key"), pulumi.Alias(type_="azure-nextgen:synapse:Key"), pulumi.Alias(type_="azure-native:synapse/latest:Key"), pulumi.Alias(type_="azure-nextgen:synapse/latest:Key"), pulumi.Alias(type_="azure-native:synapse/v20201201:Key"), pulumi.Alias(type_="azure-nextgen:synapse/v20201201:Key"), pulumi.Alias(type_="azure-native:synapse/v20210301:Key"), pulumi.Alias(type_="azure-nextgen:synapse/v20210301:Key")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Key, __self__).__init__(
            'azure-native:synapse/v20190601preview:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["is_active_cmk"] = None
        __props__["key_vault_url"] = None
        __props__["name"] = None
        __props__["type"] = None
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="isActiveCMK")
    def is_active_cmk(self) -> pulumi.Output[Optional[bool]]:
        """
        Used to activate the workspace after a customer managed key is provided.
        """
        return pulumi.get(self, "is_active_cmk")

    @property
    @pulumi.getter(name="keyVaultUrl")
    def key_vault_url(self) -> pulumi.Output[Optional[str]]:
        """
        The Key Vault Url of the workspace key.
        """
        return pulumi.get(self, "key_vault_url")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

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

