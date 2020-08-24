# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetExpressRouteCircuitResult',
    'AwaitableGetExpressRouteCircuitResult',
    'get_express_route_circuit',
]

@pulumi.output_type
class GetExpressRouteCircuitResult:
    """
    ExpressRouteCircuit resource.
    """
    def __init__(__self__, allow_classic_operations=None, authorizations=None, bandwidth_in_gbps=None, circuit_provisioning_state=None, etag=None, express_route_port=None, gateway_manager_etag=None, global_reach_enabled=None, location=None, name=None, peerings=None, provisioning_state=None, service_key=None, service_provider_notes=None, service_provider_properties=None, service_provider_provisioning_state=None, sku=None, stag=None, tags=None, type=None):
        if allow_classic_operations and not isinstance(allow_classic_operations, bool):
            raise TypeError("Expected argument 'allow_classic_operations' to be a bool")
        pulumi.set(__self__, "allow_classic_operations", allow_classic_operations)
        if authorizations and not isinstance(authorizations, list):
            raise TypeError("Expected argument 'authorizations' to be a list")
        pulumi.set(__self__, "authorizations", authorizations)
        if bandwidth_in_gbps and not isinstance(bandwidth_in_gbps, float):
            raise TypeError("Expected argument 'bandwidth_in_gbps' to be a float")
        pulumi.set(__self__, "bandwidth_in_gbps", bandwidth_in_gbps)
        if circuit_provisioning_state and not isinstance(circuit_provisioning_state, str):
            raise TypeError("Expected argument 'circuit_provisioning_state' to be a str")
        pulumi.set(__self__, "circuit_provisioning_state", circuit_provisioning_state)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if express_route_port and not isinstance(express_route_port, dict):
            raise TypeError("Expected argument 'express_route_port' to be a dict")
        pulumi.set(__self__, "express_route_port", express_route_port)
        if gateway_manager_etag and not isinstance(gateway_manager_etag, str):
            raise TypeError("Expected argument 'gateway_manager_etag' to be a str")
        pulumi.set(__self__, "gateway_manager_etag", gateway_manager_etag)
        if global_reach_enabled and not isinstance(global_reach_enabled, bool):
            raise TypeError("Expected argument 'global_reach_enabled' to be a bool")
        pulumi.set(__self__, "global_reach_enabled", global_reach_enabled)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if peerings and not isinstance(peerings, list):
            raise TypeError("Expected argument 'peerings' to be a list")
        pulumi.set(__self__, "peerings", peerings)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if service_key and not isinstance(service_key, str):
            raise TypeError("Expected argument 'service_key' to be a str")
        pulumi.set(__self__, "service_key", service_key)
        if service_provider_notes and not isinstance(service_provider_notes, str):
            raise TypeError("Expected argument 'service_provider_notes' to be a str")
        pulumi.set(__self__, "service_provider_notes", service_provider_notes)
        if service_provider_properties and not isinstance(service_provider_properties, dict):
            raise TypeError("Expected argument 'service_provider_properties' to be a dict")
        pulumi.set(__self__, "service_provider_properties", service_provider_properties)
        if service_provider_provisioning_state and not isinstance(service_provider_provisioning_state, str):
            raise TypeError("Expected argument 'service_provider_provisioning_state' to be a str")
        pulumi.set(__self__, "service_provider_provisioning_state", service_provider_provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if stag and not isinstance(stag, float):
            raise TypeError("Expected argument 'stag' to be a float")
        pulumi.set(__self__, "stag", stag)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

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
        Flag denoting Global reach status.
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
    def provisioning_state(self) -> Optional[str]:
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


class AwaitableGetExpressRouteCircuitResult(GetExpressRouteCircuitResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExpressRouteCircuitResult(
            allow_classic_operations=self.allow_classic_operations,
            authorizations=self.authorizations,
            bandwidth_in_gbps=self.bandwidth_in_gbps,
            circuit_provisioning_state=self.circuit_provisioning_state,
            etag=self.etag,
            express_route_port=self.express_route_port,
            gateway_manager_etag=self.gateway_manager_etag,
            global_reach_enabled=self.global_reach_enabled,
            location=self.location,
            name=self.name,
            peerings=self.peerings,
            provisioning_state=self.provisioning_state,
            service_key=self.service_key,
            service_provider_notes=self.service_provider_notes,
            service_provider_properties=self.service_provider_properties,
            service_provider_provisioning_state=self.service_provider_provisioning_state,
            sku=self.sku,
            stag=self.stag,
            tags=self.tags,
            type=self.type)


def get_express_route_circuit(name: Optional[str] = None,
                              resource_group_name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExpressRouteCircuitResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of express route circuit.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20190801:getExpressRouteCircuit', __args__, opts=opts, typ=GetExpressRouteCircuitResult).value

    return AwaitableGetExpressRouteCircuitResult(
        allow_classic_operations=__ret__.allow_classic_operations,
        authorizations=__ret__.authorizations,
        bandwidth_in_gbps=__ret__.bandwidth_in_gbps,
        circuit_provisioning_state=__ret__.circuit_provisioning_state,
        etag=__ret__.etag,
        express_route_port=__ret__.express_route_port,
        gateway_manager_etag=__ret__.gateway_manager_etag,
        global_reach_enabled=__ret__.global_reach_enabled,
        location=__ret__.location,
        name=__ret__.name,
        peerings=__ret__.peerings,
        provisioning_state=__ret__.provisioning_state,
        service_key=__ret__.service_key,
        service_provider_notes=__ret__.service_provider_notes,
        service_provider_properties=__ret__.service_provider_properties,
        service_provider_provisioning_state=__ret__.service_provider_provisioning_state,
        sku=__ret__.sku,
        stag=__ret__.stag,
        tags=__ret__.tags,
        type=__ret__.type)
