# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ExpressRoutePort(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    identity: pulumi.Output[dict]
    """
    The identity of ExpressRoutePort, if configured.
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
    ExpressRoutePort properties.
      * `allocation_date` (`str`) - Date of the physical port allocation to be used in Letter of Authorization.
      * `bandwidth_in_gbps` (`float`) - Bandwidth of procured ports in Gbps.
      * `circuits` (`list`) - Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.
        * `id` (`str`) - Resource ID.

      * `encapsulation` (`str`) - Encapsulation method on physical ports.
      * `ether_type` (`str`) - Ether type of the physical port.
      * `links` (`list`) - The set of physical links of the ExpressRoutePort resource.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Name of child port resource that is unique among child port resources of the parent.
        * `properties` (`dict`) - ExpressRouteLink properties.
          * `admin_state` (`str`) - Administrative state of the physical port.
          * `connector_type` (`str`) - Physical fiber port type.
          * `interface_name` (`str`) - Name of Azure router interface.
          * `mac_sec_config` (`dict`) - MacSec configuration.
            * `cak_secret_identifier` (`str`) - Keyvault Secret Identifier URL containing Mac security CAK key.
            * `cipher` (`str`) - Mac security cipher.
            * `ckn_secret_identifier` (`str`) - Keyvault Secret Identifier URL containing Mac security CKN key.

          * `patch_panel_id` (`str`) - Mapping between physical port to patch panel port.
          * `provisioning_state` (`str`) - The provisioning state of the express route link resource.
          * `rack_id` (`str`) - Mapping of physical patch panel to rack.
          * `router_name` (`str`) - Name of Azure router associated with physical port.

      * `mtu` (`str`) - Maximum transmission unit of the physical port pair(s).
      * `peering_location` (`str`) - The name of the peering location that the ExpressRoutePort is mapped to physically.
      * `provisioned_bandwidth_in_gbps` (`float`) - Aggregate Gbps of associated circuit bandwidths.
      * `provisioning_state` (`str`) - The provisioning state of the express route port resource.
      * `resource_guid` (`str`) - The resource GUID property of the express route port resource.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, id=None, identity=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        ExpressRoutePort resource definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[dict] identity: The identity of ExpressRoutePort, if configured.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the ExpressRoutePort resource.
        :param pulumi.Input[dict] properties: ExpressRoutePort properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual machine.
          * `user_assigned_identities` (`pulumi.Input[dict]`) - The list of user identities associated with resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.

        The **properties** object supports the following:

          * `bandwidth_in_gbps` (`pulumi.Input[float]`) - Bandwidth of procured ports in Gbps.
          * `encapsulation` (`pulumi.Input[str]`) - Encapsulation method on physical ports.
          * `links` (`pulumi.Input[list]`) - The set of physical links of the ExpressRoutePort resource.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - Name of child port resource that is unique among child port resources of the parent.
            * `properties` (`pulumi.Input[dict]`) - ExpressRouteLink properties.
              * `admin_state` (`pulumi.Input[str]`) - Administrative state of the physical port.
              * `mac_sec_config` (`pulumi.Input[dict]`) - MacSec configuration.
                * `cak_secret_identifier` (`pulumi.Input[str]`) - Keyvault Secret Identifier URL containing Mac security CAK key.
                * `cipher` (`pulumi.Input[str]`) - Mac security cipher.
                * `ckn_secret_identifier` (`pulumi.Input[str]`) - Keyvault Secret Identifier URL containing Mac security CKN key.

          * `peering_location` (`pulumi.Input[str]`) - The name of the peering location that the ExpressRoutePort is mapped to physically.
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
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['type'] = None
        super(ExpressRoutePort, __self__).__init__(
            'azurerm:network/v20191201:ExpressRoutePort',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ExpressRoutePort resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRoutePort(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
