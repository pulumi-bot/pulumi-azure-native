# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class WCFRelay(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Properties of WcfRelay
      * `created_at` (`str`) - The time the WCFRelay was created.
      * `is_dynamic` (`bool`) - true if the relay is dynamic; otherwise, false.
      * `listener_count` (`float`) - The number of listeners for this relay. min : 1 and max:25 supported
      * `relay_type` (`str`) - WCFRelay Type.
      * `requires_client_authorization` (`bool`) - true if client authorization is needed for this relay; otherwise, false.
      * `requires_transport_security` (`bool`) - true if transport security is needed for this relay; otherwise, false.
      * `updated_at` (`str`) - The time the namespace was updated.
      * `user_metadata` (`str`) - usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, name=None, namespace_name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Description of WcfRelays Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The relay name
        :param pulumi.Input[str] namespace_name: The Namespace Name
        :param pulumi.Input[dict] properties: Properties of WcfRelay
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.

        The **properties** object supports the following:

          * `relay_type` (`pulumi.Input[str]`) - WCFRelay Type.
          * `requires_client_authorization` (`pulumi.Input[bool]`) - true if client authorization is needed for this relay; otherwise, false.
          * `requires_transport_security` (`pulumi.Input[bool]`) - true if transport security is needed for this relay; otherwise, false.
          * `user_metadata` (`pulumi.Input[str]`) - usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(WCFRelay, __self__).__init__(
            'azurerm:relay/v20160701:WCFRelay',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WCFRelay resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WCFRelay(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
