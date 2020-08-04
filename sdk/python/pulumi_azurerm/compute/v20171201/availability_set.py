# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class AvailabilitySet(pulumi.CustomResource):
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
    The instance view of a resource.
      * `platform_fault_domain_count` (`float`) - Fault Domain count.
      * `platform_update_domain_count` (`float`) - Update Domain count.
      * `statuses` (`list`) - The resource status information.
        * `code` (`str`) - The status code.
        * `display_status` (`str`) - The short localizable label for the status.
        * `level` (`str`) - The level code.
        * `message` (`str`) - The detailed status message, including for alerts and error messages.
        * `time` (`str`) - The time of the status.

      * `virtual_machines` (`list`) - A list of references to all virtual machines in the availability set.
        * `id` (`str`) - Resource Id
    """
    sku: pulumi.Output[dict]
    """
    Sku of the availability set
      * `capacity` (`float`) - Specifies the number of virtual machines in the scale set.
      * `name` (`str`) - The sku name.
      * `tier` (`str`) - Specifies the tier of virtual machines in a scale set.<br /><br /> Possible Values:<br /><br /> **Standard**<br /><br /> **Basic**
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the availability set.
        :param pulumi.Input[dict] properties: The instance view of a resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: Sku of the availability set
        :param pulumi.Input[dict] tags: Resource tags

        The **properties** object supports the following:

          * `platform_fault_domain_count` (`pulumi.Input[float]`) - Fault Domain count.
          * `platform_update_domain_count` (`pulumi.Input[float]`) - Update Domain count.
          * `virtual_machines` (`pulumi.Input[list]`) - A list of references to all virtual machines in the availability set.
            * `id` (`pulumi.Input[str]`) - Resource Id

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - Specifies the number of virtual machines in the scale set.
          * `name` (`pulumi.Input[str]`) - The sku name.
          * `tier` (`pulumi.Input[str]`) - Specifies the tier of virtual machines in a scale set.<br /><br /> Possible Values:<br /><br /> **Standard**<br /><br /> **Basic**
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
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['type'] = None
        super(AvailabilitySet, __self__).__init__(
            'azurerm:compute/v20171201:AvailabilitySet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AvailabilitySet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AvailabilitySet(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
