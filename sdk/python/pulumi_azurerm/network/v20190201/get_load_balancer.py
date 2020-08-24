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
    'GetLoadBalancerResult',
    'AwaitableGetLoadBalancerResult',
    'get_load_balancer',
]

@pulumi.output_type
class GetLoadBalancerResult:
    """
    LoadBalancer resource
    """
    def __init__(__self__, backend_address_pools=None, etag=None, frontend_ip_configurations=None, inbound_nat_pools=None, inbound_nat_rules=None, load_balancing_rules=None, location=None, name=None, outbound_rules=None, probes=None, provisioning_state=None, resource_guid=None, sku=None, tags=None, type=None):
        if backend_address_pools and not isinstance(backend_address_pools, list):
            raise TypeError("Expected argument 'backend_address_pools' to be a list")
        pulumi.set(__self__, "backend_address_pools", backend_address_pools)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if frontend_ip_configurations and not isinstance(frontend_ip_configurations, list):
            raise TypeError("Expected argument 'frontend_ip_configurations' to be a list")
        pulumi.set(__self__, "frontend_ip_configurations", frontend_ip_configurations)
        if inbound_nat_pools and not isinstance(inbound_nat_pools, list):
            raise TypeError("Expected argument 'inbound_nat_pools' to be a list")
        pulumi.set(__self__, "inbound_nat_pools", inbound_nat_pools)
        if inbound_nat_rules and not isinstance(inbound_nat_rules, list):
            raise TypeError("Expected argument 'inbound_nat_rules' to be a list")
        pulumi.set(__self__, "inbound_nat_rules", inbound_nat_rules)
        if load_balancing_rules and not isinstance(load_balancing_rules, list):
            raise TypeError("Expected argument 'load_balancing_rules' to be a list")
        pulumi.set(__self__, "load_balancing_rules", load_balancing_rules)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if outbound_rules and not isinstance(outbound_rules, list):
            raise TypeError("Expected argument 'outbound_rules' to be a list")
        pulumi.set(__self__, "outbound_rules", outbound_rules)
        if probes and not isinstance(probes, list):
            raise TypeError("Expected argument 'probes' to be a list")
        pulumi.set(__self__, "probes", probes)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_guid and not isinstance(resource_guid, str):
            raise TypeError("Expected argument 'resource_guid' to be a str")
        pulumi.set(__self__, "resource_guid", resource_guid)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="backendAddressPools")
    def backend_address_pools(self) -> Optional[List['outputs.BackendAddressPoolResponse']]:
        """
        Collection of backend address pools used by a load balancer
        """
        return pulumi.get(self, "backend_address_pools")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="frontendIPConfigurations")
    def frontend_ip_configurations(self) -> Optional[List['outputs.FrontendIPConfigurationResponse']]:
        """
        Object representing the frontend IPs to be used for the load balancer
        """
        return pulumi.get(self, "frontend_ip_configurations")

    @property
    @pulumi.getter(name="inboundNatPools")
    def inbound_nat_pools(self) -> Optional[List['outputs.InboundNatPoolResponse']]:
        """
        Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
        """
        return pulumi.get(self, "inbound_nat_pools")

    @property
    @pulumi.getter(name="inboundNatRules")
    def inbound_nat_rules(self) -> Optional[List['outputs.InboundNatRuleResponse']]:
        """
        Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
        """
        return pulumi.get(self, "inbound_nat_rules")

    @property
    @pulumi.getter(name="loadBalancingRules")
    def load_balancing_rules(self) -> Optional[List['outputs.LoadBalancingRuleResponse']]:
        """
        Object collection representing the load balancing rules Gets the provisioning 
        """
        return pulumi.get(self, "load_balancing_rules")

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
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Optional[List['outputs.OutboundRuleResponse']]:
        """
        The outbound rules.
        """
        return pulumi.get(self, "outbound_rules")

    @property
    @pulumi.getter
    def probes(self) -> Optional[List['outputs.ProbeResponse']]:
        """
        Collection of probe objects used in the load balancer
        """
        return pulumi.get(self, "probes")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> Optional[str]:
        """
        The resource GUID property of the load balancer resource.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.LoadBalancerSkuResponse']:
        """
        The load balancer SKU.
        """
        return pulumi.get(self, "sku")

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


class AwaitableGetLoadBalancerResult(GetLoadBalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancerResult(
            backend_address_pools=self.backend_address_pools,
            etag=self.etag,
            frontend_ip_configurations=self.frontend_ip_configurations,
            inbound_nat_pools=self.inbound_nat_pools,
            inbound_nat_rules=self.inbound_nat_rules,
            load_balancing_rules=self.load_balancing_rules,
            location=self.location,
            name=self.name,
            outbound_rules=self.outbound_rules,
            probes=self.probes,
            provisioning_state=self.provisioning_state,
            resource_guid=self.resource_guid,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_load_balancer(expand: Optional[str] = None,
                      name: Optional[str] = None,
                      resource_group_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancerResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: Expands referenced resources.
    :param str name: The name of the load balancer.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20190201:getLoadBalancer', __args__, opts=opts, typ=GetLoadBalancerResult).value

    return AwaitableGetLoadBalancerResult(
        backend_address_pools=__ret__.backend_address_pools,
        etag=__ret__.etag,
        frontend_ip_configurations=__ret__.frontend_ip_configurations,
        inbound_nat_pools=__ret__.inbound_nat_pools,
        inbound_nat_rules=__ret__.inbound_nat_rules,
        load_balancing_rules=__ret__.load_balancing_rules,
        location=__ret__.location,
        name=__ret__.name,
        outbound_rules=__ret__.outbound_rules,
        probes=__ret__.probes,
        provisioning_state=__ret__.provisioning_state,
        resource_guid=__ret__.resource_guid,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type)
