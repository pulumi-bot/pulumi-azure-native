# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ExpressRouteCircuitAuthorization(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated
    """
    name: pulumi.Output[str]
    """
    Gets name of the resource that is unique within a resource group. This name can be used to access the resource
    """
    properties: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, authorization_key=None, authorization_use_status=None, circuit_name=None, etag=None, id=None, name=None, provisioning_state=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Authorization in a ExpressRouteCircuit resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_key: Gets or sets the authorization key
        :param pulumi.Input[str] authorization_use_status: Gets or sets AuthorizationUseStatus
        :param pulumi.Input[str] circuit_name: The name of the express route circuit.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] name: The name of the authorization.
        :param pulumi.Input[str] provisioning_state: Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
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

            __props__['authorization_key'] = authorization_key
            __props__['authorization_use_status'] = authorization_use_status
            if circuit_name is None:
                raise TypeError("Missing required property 'circuit_name'")
            __props__['circuit_name'] = circuit_name
            __props__['etag'] = etag
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['properties'] = None
        super(ExpressRouteCircuitAuthorization, __self__).__init__(
            'azurerm:network/v20160601:ExpressRouteCircuitAuthorization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ExpressRouteCircuitAuthorization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteCircuitAuthorization(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
