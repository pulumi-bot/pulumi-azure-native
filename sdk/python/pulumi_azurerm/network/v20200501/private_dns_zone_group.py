# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PrivateDnsZoneGroup(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    name: pulumi.Output[str]
    """
    Name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the private dns zone group.
      * `private_dns_zone_configs` (`list`) - A collection of private dns zone configurations of the private dns zone group.
        * `name` (`str`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the private dns zone configuration.
          * `private_dns_zone_id` (`str`) - The resource id of the private dns zone.
          * `record_sets` (`list`) - A collection of information regarding a recordSet, holding information to identify private resources.
            * `fqdn` (`str`) - Fqdn that resolves to private endpoint ip address.
            * `ip_addresses` (`list`) - The private ip address of the private endpoint.
            * `provisioning_state` (`str`) - The provisioning state of the recordset.
            * `record_set_name` (`str`) - Recordset name.
            * `record_type` (`str`) - Resource record type.
            * `ttl` (`float`) - Recordset time to live.

      * `provisioning_state` (`str`) - The provisioning state of the private dns zone group resource.
    """
    def __init__(__self__, resource_name, opts=None, id=None, name=None, private_endpoint_name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Private dns zone group resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the private dns zone group.
        :param pulumi.Input[str] private_endpoint_name: The name of the private endpoint.
        :param pulumi.Input[dict] properties: Properties of the private dns zone group.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.

        The **properties** object supports the following:

          * `private_dns_zone_configs` (`pulumi.Input[list]`) - A collection of private dns zone configurations of the private dns zone group.
            * `name` (`pulumi.Input[str]`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`) - Properties of the private dns zone configuration.
              * `private_dns_zone_id` (`pulumi.Input[str]`) - The resource id of the private dns zone.
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

            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if private_endpoint_name is None:
                raise TypeError("Missing required property 'private_endpoint_name'")
            __props__['private_endpoint_name'] = private_endpoint_name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['etag'] = None
        super(PrivateDnsZoneGroup, __self__).__init__(
            'azurerm:network/v20200501:PrivateDnsZoneGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing PrivateDnsZoneGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PrivateDnsZoneGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
