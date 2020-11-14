# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['DisasterRecoveryConfig']


class DisasterRecoveryConfig(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 alternate_name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 partner_namespace: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Single item in List or Get Alias(Disaster Recovery configuration) operation

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The Disaster Recovery configuration name
        :param pulumi.Input[str] alternate_name: Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
        :param pulumi.Input[str] namespace_name: The namespace name
        :param pulumi.Input[str] partner_namespace: ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
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

            if alias is None:
                raise TypeError("Missing required property 'alias'")
            __props__['alias'] = alias
            __props__['alternate_name'] = alternate_name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['partner_namespace'] = partner_namespace
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['name'] = None
            __props__['pending_replication_operations_count'] = None
            __props__['provisioning_state'] = None
            __props__['role'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:servicebus/latest:DisasterRecoveryConfig"), pulumi.Alias(type_="azure-nextgen:servicebus/v20170401:DisasterRecoveryConfig")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DisasterRecoveryConfig, __self__).__init__(
            'azure-nextgen:servicebus/v20180101preview:DisasterRecoveryConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DisasterRecoveryConfig':
        """
        Get an existing DisasterRecoveryConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DisasterRecoveryConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alternateName")
    def alternate_name(self) -> pulumi.Output[Optional[str]]:
        """
        Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
        """
        return pulumi.get(self, "alternate_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partnerNamespace")
    def partner_namespace(self) -> pulumi.Output[Optional[str]]:
        """
        ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
        """
        return pulumi.get(self, "partner_namespace")

    @property
    @pulumi.getter(name="pendingReplicationOperationsCount")
    def pending_replication_operations_count(self) -> pulumi.Output[int]:
        """
        Number of entities pending to be replicated.
        """
        return pulumi.get(self, "pending_replication_operations_count")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

