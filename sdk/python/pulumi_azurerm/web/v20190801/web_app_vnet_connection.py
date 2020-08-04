# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class WebAppVnetConnection(pulumi.CustomResource):
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
    VnetInfo resource specific properties
      * `cert_blob` (`str`) - A certificate file (.cer) blob containing the public key of the private key used to authenticate a 
        Point-To-Site VPN connection.
      * `cert_thumbprint` (`str`) - The client certificate thumbprint.
      * `dns_servers` (`str`) - DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
      * `is_swift` (`bool`) - Flag that is used to denote if this is VNET injection
      * `resync_required` (`bool`) - <code>true</code> if a resync is required; otherwise, <code>false</code>.
      * `routes` (`list`) - The routes that this Virtual Network connection uses.
        * `id` (`str`) - Resource Id.
        * `kind` (`str`) - Kind of resource.
        * `name` (`str`) - Resource Name.
        * `properties` (`dict`) - VnetRoute resource specific properties
          * `end_address` (`str`) - The ending address for this route. If the start address is specified in CIDR notation, this must be omitted.
          * `route_type` (`str`) - The type of route this is:
            DEFAULT - By default, every app has routes to the local address ranges specified by RFC1918
            INHERITED - Routes inherited from the real Virtual Network routes
            STATIC - Static route set on the app only
            
            These values will be used for syncing an app's routes with those from a Virtual Network.
          * `start_address` (`str`) - The starting address for this route. This may also include a CIDR notation, in which case the end address must not be specified.

        * `type` (`str`) - Resource type.

      * `vnet_resource_id` (`str`) - The Virtual Network's resource ID.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, kind=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Virtual Network information contract.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] name: Name of an existing Virtual Network.
        :param pulumi.Input[dict] properties: VnetInfo resource specific properties
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.

        The **properties** object supports the following:

          * `cert_blob` (`pulumi.Input[str]`) - A certificate file (.cer) blob containing the public key of the private key used to authenticate a 
            Point-To-Site VPN connection.
          * `dns_servers` (`pulumi.Input[str]`) - DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
          * `is_swift` (`pulumi.Input[bool]`) - Flag that is used to denote if this is VNET injection
          * `vnet_resource_id` (`pulumi.Input[str]`) - The Virtual Network's resource ID.
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

            __props__['kind'] = kind
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(WebAppVnetConnection, __self__).__init__(
            'azurerm:web/v20190801:WebAppVnetConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WebAppVnetConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebAppVnetConnection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
