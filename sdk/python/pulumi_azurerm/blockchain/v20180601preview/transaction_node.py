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

__all__ = ['TransactionNode']


class TransactionNode(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blockchain_member_name: Optional[pulumi.Input[str]] = None,
                 firewall_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallRuleArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 transaction_node_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Payload of the transaction node which is the request/response of the resource provider.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] blockchain_member_name: Blockchain member name.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallRuleArgs']]]] firewall_rules: Gets or sets the firewall rules.
        :param pulumi.Input[str] location: Gets or sets the transaction node location.
        :param pulumi.Input[str] password: Sets the transaction node dns endpoint basic auth password.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[str] transaction_node_name: Transaction node name.
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

            if blockchain_member_name is None:
                raise TypeError("Missing required property 'blockchain_member_name'")
            __props__['blockchain_member_name'] = blockchain_member_name
            __props__['firewall_rules'] = firewall_rules
            __props__['location'] = location
            __props__['password'] = password
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if transaction_node_name is None:
                raise TypeError("Missing required property 'transaction_node_name'")
            __props__['transaction_node_name'] = transaction_node_name
            __props__['dns'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['public_key'] = None
            __props__['type'] = None
            __props__['user_name'] = None
        super(TransactionNode, __self__).__init__(
            'azurerm:blockchain/v20180601preview:TransactionNode',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TransactionNode':
        """
        Get an existing TransactionNode resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return TransactionNode(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dns(self) -> pulumi.Output[str]:
        """
        Gets or sets the transaction node dns endpoint.
        """
        return pulumi.get(self, "dns")

    @property
    @pulumi.getter(name="firewallRules")
    def firewall_rules(self) -> pulumi.Output[Optional[List['outputs.FirewallRuleResponse']]]:
        """
        Gets or sets the firewall rules.
        """
        return pulumi.get(self, "firewall_rules")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the transaction node location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Sets the transaction node dns endpoint basic auth password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Gets or sets the blockchain member provision state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        Gets or sets the transaction node public key.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the service - e.g. "Microsoft.Blockchain"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        Gets or sets the transaction node dns endpoint basic auth user name.
        """
        return pulumi.get(self, "user_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

