# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ExpressRouteCircuit']


class ExpressRouteCircuit(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_classic_operations: Optional[pulumi.Input[bool]] = None,
                 authorizations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ExpressRouteCircuitAuthorizationArgs']]]]] = None,
                 bandwidth_in_gbps: Optional[pulumi.Input[float]] = None,
                 circuit_provisioning_state: Optional[pulumi.Input[str]] = None,
                 express_route_port: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 gateway_manager_etag: Optional[pulumi.Input[str]] = None,
                 global_reach_enabled: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peerings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_key: Optional[pulumi.Input[str]] = None,
                 service_provider_notes: Optional[pulumi.Input[str]] = None,
                 service_provider_properties: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitServiceProviderPropertiesArgs']]] = None,
                 service_provider_provisioning_state: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ExpressRouteCircuit resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_classic_operations: Allow classic operations.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ExpressRouteCircuitAuthorizationArgs']]]] authorizations: The list of authorizations.
        :param pulumi.Input[float] bandwidth_in_gbps: The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
        :param pulumi.Input[str] circuit_provisioning_state: The CircuitProvisioningState state of the resource.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] express_route_port: The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
        :param pulumi.Input[str] gateway_manager_etag: The GatewayManager Etag.
        :param pulumi.Input[bool] global_reach_enabled: Flag denoting global reach status.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the circuit.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringArgs']]]] peerings: The list of peerings.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_key: The ServiceKey.
        :param pulumi.Input[str] service_provider_notes: The ServiceProviderNotes.
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitServiceProviderPropertiesArgs']] service_provider_properties: The ServiceProviderProperties.
        :param pulumi.Input[str] service_provider_provisioning_state: The ServiceProviderProvisioningState state of the resource.
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitSkuArgs']] sku: The SKU.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
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

            __props__['allow_classic_operations'] = allow_classic_operations
            __props__['authorizations'] = authorizations
            __props__['bandwidth_in_gbps'] = bandwidth_in_gbps
            __props__['circuit_provisioning_state'] = circuit_provisioning_state
            __props__['express_route_port'] = express_route_port
            __props__['gateway_manager_etag'] = gateway_manager_etag
            __props__['global_reach_enabled'] = global_reach_enabled
            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['peerings'] = peerings
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['service_key'] = service_key
            __props__['service_provider_notes'] = service_provider_notes
            __props__['service_provider_properties'] = service_provider_properties
            __props__['service_provider_provisioning_state'] = service_provider_provisioning_state
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['provisioning_state'] = None
            __props__['stag'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20150615:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20160330:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20160601:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20160901:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20161201:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20170301:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20170601:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20170801:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20170901:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20171001:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20171101:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20180101:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20180201:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20180401:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20180601:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20180701:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20180801:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20181001:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20181101:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20181201:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20190201:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20190401:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20190601:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20190701:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20190801:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20190901:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20191101:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20200301:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20200401:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20200501:ExpressRouteCircuit"), pulumi.Alias(type_="azurerm:network/v20200601:ExpressRouteCircuit")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ExpressRouteCircuit, __self__).__init__(
            'azurerm:network/v20191201:ExpressRouteCircuit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExpressRouteCircuit':
        """
        Get an existing ExpressRouteCircuit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteCircuit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowClassicOperations")
    def allow_classic_operations(self) -> Optional[bool]:
        """
        Allow classic operations.
        """
        return pulumi.get(self, "allow_classic_operations")

    @property
    @pulumi.getter
    def authorizations(self) -> Optional[List['outputs.ExpressRouteCircuitAuthorizationResponse']]:
        """
        The list of authorizations.
        """
        return pulumi.get(self, "authorizations")

    @property
    @pulumi.getter(name="bandwidthInGbps")
    def bandwidth_in_gbps(self) -> Optional[float]:
        """
        The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
        """
        return pulumi.get(self, "bandwidth_in_gbps")

    @property
    @pulumi.getter(name="circuitProvisioningState")
    def circuit_provisioning_state(self) -> Optional[str]:
        """
        The CircuitProvisioningState state of the resource.
        """
        return pulumi.get(self, "circuit_provisioning_state")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expressRoutePort")
    def express_route_port(self) -> Optional['outputs.SubResourceResponse']:
        """
        The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
        """
        return pulumi.get(self, "express_route_port")

    @property
    @pulumi.getter(name="gatewayManagerEtag")
    def gateway_manager_etag(self) -> Optional[str]:
        """
        The GatewayManager Etag.
        """
        return pulumi.get(self, "gateway_manager_etag")

    @property
    @pulumi.getter(name="globalReachEnabled")
    def global_reach_enabled(self) -> Optional[bool]:
        """
        Flag denoting global reach status.
        """
        return pulumi.get(self, "global_reach_enabled")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def peerings(self) -> Optional[List['outputs.ExpressRouteCircuitPeeringResponse']]:
        """
        The list of peerings.
        """
        return pulumi.get(self, "peerings")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the express route circuit resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serviceKey")
    def service_key(self) -> Optional[str]:
        """
        The ServiceKey.
        """
        return pulumi.get(self, "service_key")

    @property
    @pulumi.getter(name="serviceProviderNotes")
    def service_provider_notes(self) -> Optional[str]:
        """
        The ServiceProviderNotes.
        """
        return pulumi.get(self, "service_provider_notes")

    @property
    @pulumi.getter(name="serviceProviderProperties")
    def service_provider_properties(self) -> Optional['outputs.ExpressRouteCircuitServiceProviderPropertiesResponse']:
        """
        The ServiceProviderProperties.
        """
        return pulumi.get(self, "service_provider_properties")

    @property
    @pulumi.getter(name="serviceProviderProvisioningState")
    def service_provider_provisioning_state(self) -> Optional[str]:
        """
        The ServiceProviderProvisioningState state of the resource.
        """
        return pulumi.get(self, "service_provider_provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ExpressRouteCircuitSkuResponse']:
        """
        The SKU.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def stag(self) -> float:
        """
        The identifier of the circuit traffic. Outer tag for QinQ encapsulation.
        """
        return pulumi.get(self, "stag")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

