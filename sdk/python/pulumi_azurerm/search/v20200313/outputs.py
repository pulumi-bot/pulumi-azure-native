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
    'IdentityResponse',
    'IpRuleResponse',
    'NetworkRuleSetResponse',
    'PrivateEndpointConnectionPropertiesResponse',
    'PrivateEndpointConnectionPropertiesResponsePrivateEndpoint',
    'PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState',
    'PrivateEndpointConnectionResponse',
    'QueryKeyResponseResult',
    'SkuResponse',
]

@pulumi.output_type
class IdentityResponse(dict):
    """
    Identity for the resource.
    """
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: str):
        """
        Identity for the resource.
        :param str principal_id: The principal ID of resource identity.
        :param str tenant_id: The tenant ID of resource.
        :param str type: The identity type.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal ID of resource identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The tenant ID of resource.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IpRuleResponse(dict):
    """
    The IP restriction rule of the Azure Cognitive Search service.
    """
    def __init__(__self__, *,
                 value: Optional[str] = None):
        """
        The IP restriction rule of the Azure Cognitive Search service.
        :param str value: Value corresponding to a single IPv4 address (eg., 123.1.2.3) or an IP range in CIDR format (eg., 123.1.2.3/24) to be allowed.
        """
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        Value corresponding to a single IPv4 address (eg., 123.1.2.3) or an IP range in CIDR format (eg., 123.1.2.3/24) to be allowed.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class NetworkRuleSetResponse(dict):
    """
    Network specific rules that determine how the Azure Cognitive Search service may be reached.
    """
    def __init__(__self__, *,
                 ip_rules: Optional[List['outputs.IpRuleResponse']] = None):
        """
        Network specific rules that determine how the Azure Cognitive Search service may be reached.
        :param List['IpRuleResponseArgs'] ip_rules: A list of IP restriction rules that defines the inbound network(s) with allowing access to the search service endpoint. At the meantime, all other public IP networks are blocked by the firewall. These restriction rules are applied only when the 'publicNetworkAccess' of the search service is 'enabled'; otherwise, traffic over public interface is not allowed even with any public IP rules, and private endpoint connections would be the exclusive access method.
        """
        if ip_rules is not None:
            pulumi.set(__self__, "ip_rules", ip_rules)

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> Optional[List['outputs.IpRuleResponse']]:
        """
        A list of IP restriction rules that defines the inbound network(s) with allowing access to the search service endpoint. At the meantime, all other public IP networks are blocked by the firewall. These restriction rules are applied only when the 'publicNetworkAccess' of the search service is 'enabled'; otherwise, traffic over public interface is not allowed even with any public IP rules, and private endpoint connections would be the exclusive access method.
        """
        return pulumi.get(self, "ip_rules")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointConnectionPropertiesResponse(dict):
    """
    Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
    """
    def __init__(__self__, *,
                 private_endpoint: Optional['outputs.PrivateEndpointConnectionPropertiesResponsePrivateEndpoint'] = None,
                 private_link_service_connection_state: Optional['outputs.PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState'] = None):
        """
        Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
        :param 'PrivateEndpointConnectionPropertiesResponsePrivateEndpointArgs' private_endpoint: The private endpoint resource from Microsoft.Network provider.
        :param 'PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs' private_link_service_connection_state: Describes the current state of an existing Private Link Service connection to the Azure Private Endpoint.
        """
        if private_endpoint is not None:
            pulumi.set(__self__, "private_endpoint", private_endpoint)
        if private_link_service_connection_state is not None:
            pulumi.set(__self__, "private_link_service_connection_state", private_link_service_connection_state)

    @property
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> Optional['outputs.PrivateEndpointConnectionPropertiesResponsePrivateEndpoint']:
        """
        The private endpoint resource from Microsoft.Network provider.
        """
        return pulumi.get(self, "private_endpoint")

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> Optional['outputs.PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState']:
        """
        Describes the current state of an existing Private Link Service connection to the Azure Private Endpoint.
        """
        return pulumi.get(self, "private_link_service_connection_state")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointConnectionPropertiesResponsePrivateEndpoint(dict):
    """
    The private endpoint resource from Microsoft.Network provider.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None):
        """
        The private endpoint resource from Microsoft.Network provider.
        :param str id: The resource id of the private endpoint resource from Microsoft.Network provider.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The resource id of the private endpoint resource from Microsoft.Network provider.
        """
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState(dict):
    """
    Describes the current state of an existing Private Link Service connection to the Azure Private Endpoint.
    """
    def __init__(__self__, *,
                 actions_required: Optional[str] = None,
                 description: Optional[str] = None,
                 status: Optional[str] = None):
        """
        Describes the current state of an existing Private Link Service connection to the Azure Private Endpoint.
        :param str actions_required: A description of any extra actions that may be required.
        :param str description: The description for the private link service connection state.
        :param str status: Status of the the private link service connection. Can be Pending, Approved, Rejected, or Disconnected.
        """
        if actions_required is not None:
            pulumi.set(__self__, "actions_required", actions_required)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="actionsRequired")
    def actions_required(self) -> Optional[str]:
        """
        A description of any extra actions that may be required.
        """
        return pulumi.get(self, "actions_required")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description for the private link service connection state.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Status of the the private link service connection. Can be Pending, Approved, Rejected, or Disconnected.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointConnectionResponse(dict):
    """
    Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
    """
    def __init__(__self__, *,
                 id: str,
                 name: str,
                 type: str,
                 properties: Optional['outputs.PrivateEndpointConnectionPropertiesResponse'] = None):
        """
        Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
        :param str id: The ID of the private endpoint connection. This can be used with the Azure Resource Manager to link resources together.
        :param str name: The name of the private endpoint connection.
        :param str type: The resource type.
        :param 'PrivateEndpointConnectionPropertiesResponseArgs' properties: Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the private endpoint connection. This can be used with the Azure Resource Manager to link resources together.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the private endpoint connection.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def properties(self) -> Optional['outputs.PrivateEndpointConnectionPropertiesResponse']:
        """
        Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
        """
        return pulumi.get(self, "properties")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class QueryKeyResponseResult(dict):
    """
    Describes an API key for a given Azure Cognitive Search service that has permissions for query operations only.
    """
    def __init__(__self__, *,
                 key: str,
                 name: str):
        """
        Describes an API key for a given Azure Cognitive Search service that has permissions for query operations only.
        :param str key: The value of the query API key.
        :param str name: The name of the query API key; may be empty.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The value of the query API key.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the query API key; may be empty.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class SkuResponse(dict):
    """
    Defines the SKU of an Azure Cognitive Search Service, which determines price tier and capacity limits.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        Defines the SKU of an Azure Cognitive Search Service, which determines price tier and capacity limits.
        :param str name: The SKU of the Search service. Valid values include: 'free': Shared service. 'basic': Dedicated service with up to 3 replicas. 'standard': Dedicated service with up to 12 partitions and 12 replicas. 'standard2': Similar to standard, but with more capacity per search unit. 'standard3': The largest Standard offering with up to 12 partitions and 12 replicas (or up to 3 partitions with more indexes if you also set the hostingMode property to 'highDensity'). 'storage_optimized_l1': Supports 1TB per partition, up to 12 partitions. 'storage_optimized_l2': Supports 2TB per partition, up to 12 partitions.'
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The SKU of the Search service. Valid values include: 'free': Shared service. 'basic': Dedicated service with up to 3 replicas. 'standard': Dedicated service with up to 12 partitions and 12 replicas. 'standard2': Similar to standard, but with more capacity per search unit. 'standard3': The largest Standard offering with up to 12 partitions and 12 replicas (or up to 3 partitions with more indexes if you also set the hostingMode property to 'highDensity'). 'storage_optimized_l1': Supports 1TB per partition, up to 12 partitions. 'storage_optimized_l2': Supports 2TB per partition, up to 12 partitions.'
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


