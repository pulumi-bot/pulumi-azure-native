# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class NetworkVirtualAppliance(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    identity: pulumi.Output[dict]
    """
    The service principal that has read access to cloud-init and config blob.
      * `principal_id` (`str`) - The principal id of the system assigned identity. This property will only be provided for a system assigned identity.
      * `tenant_id` (`str`) - The tenant id of the system assigned identity. This property will only be provided for a system assigned identity.
      * `type` (`str`) - The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual machine.
      * `user_assigned_identities` (`dict`) - The list of user identities associated with resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
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
    Properties of the Network Virtual Appliance.
      * `boot_strap_configuration_blob` (`list`) - BootStrapConfigurationBlob storage URLs.
      * `cloud_init_configuration_blob` (`list`) - CloudInitConfigurationBlob storage URLs.
      * `provisioning_state` (`str`) - The provisioning state of the resource.
      * `virtual_appliance_asn` (`float`) - VirtualAppliance ASN.
      * `virtual_appliance_nics` (`list`) - List of Virtual Appliance Network Interfaces.
        * `name` (`str`) - NIC name.
        * `private_ip_address` (`str`) - Private IP address.
        * `public_ip_address` (`str`) - Public IP address.

      * `virtual_hub` (`dict`) - The Virtual Hub where Network Virtual Appliance is being deployed.
        * `id` (`str`) - Resource ID.
    """
    sku: pulumi.Output[dict]
    """
    Network Virtual Appliance SKU.
      * `bundled_scale_unit` (`str`) - Virtual Appliance Scale Unit.
      * `market_place_version` (`str`) - Virtual Appliance Version.
      * `vendor` (`str`) - Virtual Appliance Vendor.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, id=None, identity=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        NetworkVirtualAppliance Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[dict] identity: The service principal that has read access to cloud-init and config blob.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of Network Virtual Appliance.
        :param pulumi.Input[dict] properties: Properties of the Network Virtual Appliance.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: Network Virtual Appliance SKU.
        :param pulumi.Input[dict] tags: Resource tags.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual machine.
          * `user_assigned_identities` (`pulumi.Input[dict]`) - The list of user identities associated with resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.

        The **properties** object supports the following:

          * `boot_strap_configuration_blob` (`pulumi.Input[list]`) - BootStrapConfigurationBlob storage URLs.
          * `cloud_init_configuration_blob` (`pulumi.Input[list]`) - CloudInitConfigurationBlob storage URLs.
          * `virtual_appliance_asn` (`pulumi.Input[float]`) - VirtualAppliance ASN.
          * `virtual_hub` (`pulumi.Input[dict]`) - The Virtual Hub where Network Virtual Appliance is being deployed.
            * `id` (`pulumi.Input[str]`) - Resource ID.

        The **sku** object supports the following:

          * `bundled_scale_unit` (`pulumi.Input[str]`) - Virtual Appliance Scale Unit.
          * `market_place_version` (`pulumi.Input[str]`) - Virtual Appliance Version.
          * `vendor` (`pulumi.Input[str]`) - Virtual Appliance Vendor.
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
            __props__['identity'] = identity
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
            __props__['etag'] = None
            __props__['type'] = None
        super(NetworkVirtualAppliance, __self__).__init__(
            'azurerm:network/v20200301:NetworkVirtualAppliance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing NetworkVirtualAppliance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return NetworkVirtualAppliance(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
