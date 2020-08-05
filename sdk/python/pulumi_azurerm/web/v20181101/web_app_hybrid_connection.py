# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class WebAppHybridConnection(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource.
    """
    name: pulumi.Output[str]
    """
    Resource Name.
    """
    properties: pulumi.Output[dict]
    """
    HybridConnection resource specific properties
      * `hostname` (`str`) - The hostname of the endpoint.
      * `port` (`float`) - The port of the endpoint.
      * `relay_arm_uri` (`str`) - The ARM URI to the Service Bus relay.
      * `relay_name` (`str`) - The name of the Service Bus relay.
      * `send_key_name` (`str`) - The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
      * `send_key_value` (`str`) - The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
        normally, use the POST /listKeys API instead.
      * `service_bus_namespace` (`str`) - The name of the Service Bus namespace.
      * `service_bus_suffix` (`str`) - The suffix for the service bus endpoint. By default this is .servicebus.windows.net
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, hostname=None, kind=None, name=None, namespace_name=None, port=None, relay_arm_uri=None, resource_group_name=None, send_key_name=None, send_key_value=None, service_bus_namespace=None, service_bus_suffix=None, __props__=None, __name__=None, __opts__=None):
        """
        Hybrid Connection contract. This is used to configure a Hybrid Connection.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hostname: The hostname of the endpoint.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] name: The name of the Service Bus relay.
        :param pulumi.Input[str] namespace_name: The namespace for this hybrid connection.
        :param pulumi.Input[float] port: The port of the endpoint.
        :param pulumi.Input[str] relay_arm_uri: The ARM URI to the Service Bus relay.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[str] send_key_name: The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
        :param pulumi.Input[str] send_key_value: The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
               normally, use the POST /listKeys API instead.
        :param pulumi.Input[str] service_bus_namespace: The name of the Service Bus namespace.
        :param pulumi.Input[str] service_bus_suffix: The suffix for the service bus endpoint. By default this is .servicebus.windows.net
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

            __props__['hostname'] = hostname
            __props__['kind'] = kind
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['port'] = port
            __props__['relay_arm_uri'] = relay_arm_uri
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['send_key_name'] = send_key_name
            __props__['send_key_value'] = send_key_value
            __props__['service_bus_namespace'] = service_bus_namespace
            __props__['service_bus_suffix'] = service_bus_suffix
            __props__['properties'] = None
            __props__['type'] = None
        super(WebAppHybridConnection, __self__).__init__(
            'azurerm:web/v20181101:WebAppHybridConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WebAppHybridConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebAppHybridConnection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
