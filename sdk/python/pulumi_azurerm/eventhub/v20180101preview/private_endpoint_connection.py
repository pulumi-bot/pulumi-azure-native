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

__all__ = ['PrivateEndpointConnection']


class PrivateEndpointConnection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 private_endpoint: Optional[pulumi.Input[pulumi.InputType['PrivateEndpointArgs']]] = None,
                 private_endpoint_connection_name: Optional[pulumi.Input[str]] = None,
                 private_link_service_connection_state: Optional[pulumi.Input[pulumi.InputType['ConnectionStateArgs']]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Properties of the PrivateEndpointConnection.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] namespace_name: The Namespace name
        :param pulumi.Input[pulumi.InputType['PrivateEndpointArgs']] private_endpoint: The Private Endpoint resource for this Connection.
        :param pulumi.Input[str] private_endpoint_connection_name: The PrivateEndpointConnection name
        :param pulumi.Input[pulumi.InputType['ConnectionStateArgs']] private_link_service_connection_state: Details about the state of the connection.
        :param pulumi.Input[str] provisioning_state: Provisioning state of the Private Endpoint Connection.
        :param pulumi.Input[str] resource_group_name: Name of the resource group within the azure subscription.
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

            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['private_endpoint'] = private_endpoint
            if private_endpoint_connection_name is None:
                raise TypeError("Missing required property 'private_endpoint_connection_name'")
            __props__['private_endpoint_connection_name'] = private_endpoint_connection_name
            __props__['private_link_service_connection_state'] = private_link_service_connection_state
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:eventhub/latest:PrivateEndpointConnection")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PrivateEndpointConnection, __self__).__init__(
            'azurerm:eventhub/v20180101preview:PrivateEndpointConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateEndpointConnection':
        """
        Get an existing PrivateEndpointConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PrivateEndpointConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> pulumi.Output[Optional['outputs.PrivateEndpointResponse']]:
        """
        The Private Endpoint resource for this Connection.
        """
        return pulumi.get(self, "private_endpoint")

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> pulumi.Output[Optional['outputs.ConnectionStateResponse']]:
        """
        Details about the state of the connection.
        """
        return pulumi.get(self, "private_link_service_connection_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        Provisioning state of the Private Endpoint Connection.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

