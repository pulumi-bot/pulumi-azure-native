# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class DedicatedHostGroup(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Dedicated Host Group Properties.
      * `hosts` (`list`) - A list of references to all dedicated hosts in the dedicated host group.
        * `id` (`str`) - Resource Id

      * `platform_fault_domain_count` (`float`) - Number of fault domains that the host group can span.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    zones: pulumi.Output[list]
    """
    Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, platform_fault_domain_count=None, resource_group_name=None, tags=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Specifies information about the dedicated host group that the dedicated hosts should be assigned to. <br><br> Currently, a dedicated host can only be added to a dedicated host group at creation time. An existing dedicated host cannot be added to another dedicated host group.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the dedicated host group.
        :param pulumi.Input[float] platform_fault_domain_count: Number of fault domains that the host group can span.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[list] zones: Availability Zone to use for this host group. Only single zone is supported. The zone can be assigned only during creation. If not provided, the group supports all zones in the region. If provided, enforces each host in the group to be in the same zone.
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if platform_fault_domain_count is None:
                raise TypeError("Missing required property 'platform_fault_domain_count'")
            __props__['platform_fault_domain_count'] = platform_fault_domain_count
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['properties'] = None
            __props__['type'] = None
        super(DedicatedHostGroup, __self__).__init__(
            'azurerm:compute/v20190301:DedicatedHostGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing DedicatedHostGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DedicatedHostGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop