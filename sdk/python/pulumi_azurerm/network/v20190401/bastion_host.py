# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class BastionHost(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Gets a unique read-only string that changes whenever the resource is updated.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Represents the bastion host resource.
      * `dns_name` (`str`) - FQDN for the endpoint on which bastion host is accessible.
      * `ip_configurations` (`list`) - IP configuration of the Bastion Host resource.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Represents the ip configuration associated with the resource.
          * `private_ip_allocation_method` (`str`) - Private IP allocation method.
          * `provisioning_state` (`str`) - The provisioning state of the resource.
          * `public_ip_address` (`dict`) - Reference of the PublicIP resource.
            * `id` (`str`) - Resource ID.

          * `subnet` (`dict`) - Reference of the subnet resource.

        * `type` (`str`) - Ip configuration type.

      * `provisioning_state` (`str`) - The provisioning state of the resource.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, dns_name=None, id=None, ip_configurations=None, location=None, name=None, provisioning_state=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Bastion Host resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dns_name: FQDN for the endpoint on which bastion host is accessible.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[list] ip_configurations: IP configuration of the Bastion Host resource.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the Bastion Host.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags.

        The **ip_configurations** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `private_ip_allocation_method` (`pulumi.Input[str]`) - Private IP allocation method.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource.
          * `public_ip_address` (`pulumi.Input[dict]`) - Reference of the PublicIP resource.
            * `id` (`pulumi.Input[str]`) - Resource ID.

          * `subnet` (`pulumi.Input[dict]`) - Reference of the subnet resource.
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

            __props__['dns_name'] = dns_name
            __props__['id'] = id
            __props__['ip_configurations'] = ip_configurations
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(BastionHost, __self__).__init__(
            'azurerm:network/v20190401:BastionHost',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing BastionHost resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return BastionHost(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
