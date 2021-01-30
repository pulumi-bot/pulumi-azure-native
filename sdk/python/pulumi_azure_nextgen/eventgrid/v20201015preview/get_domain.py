# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetDomainResult',
    'AwaitableGetDomainResult',
    'get_domain',
]

@pulumi.output_type
class GetDomainResult:
    """
    EventGrid Domain.
    """
    def __init__(__self__, endpoint=None, id=None, identity=None, inbound_ip_rules=None, input_schema=None, input_schema_mapping=None, location=None, metric_resource_id=None, name=None, private_endpoint_connections=None, provisioning_state=None, public_network_access=None, sku=None, tags=None, type=None):
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if inbound_ip_rules and not isinstance(inbound_ip_rules, list):
            raise TypeError("Expected argument 'inbound_ip_rules' to be a list")
        pulumi.set(__self__, "inbound_ip_rules", inbound_ip_rules)
        if input_schema and not isinstance(input_schema, str):
            raise TypeError("Expected argument 'input_schema' to be a str")
        pulumi.set(__self__, "input_schema", input_schema)
        if input_schema_mapping and not isinstance(input_schema_mapping, dict):
            raise TypeError("Expected argument 'input_schema_mapping' to be a dict")
        pulumi.set(__self__, "input_schema_mapping", input_schema_mapping)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if metric_resource_id and not isinstance(metric_resource_id, str):
            raise TypeError("Expected argument 'metric_resource_id' to be a str")
        pulumi.set(__self__, "metric_resource_id", metric_resource_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_endpoint_connections and not isinstance(private_endpoint_connections, list):
            raise TypeError("Expected argument 'private_endpoint_connections' to be a list")
        pulumi.set(__self__, "private_endpoint_connections", private_endpoint_connections)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if public_network_access and not isinstance(public_network_access, str):
            raise TypeError("Expected argument 'public_network_access' to be a str")
        pulumi.set(__self__, "public_network_access", public_network_access)
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
    @pulumi.getter
    def endpoint(self) -> str:
        """
        Endpoint for the domain.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified identifier of the resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.IdentityInfoResponse']:
        """
        Identity information for the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="inboundIpRules")
    def inbound_ip_rules(self) -> Optional[Sequence['outputs.InboundIpRuleResponse']]:
        """
        This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
        """
        return pulumi.get(self, "inbound_ip_rules")

    @property
    @pulumi.getter(name="inputSchema")
    def input_schema(self) -> Optional[str]:
        """
        This determines the format that Event Grid should expect for incoming events published to the domain.
        """
        return pulumi.get(self, "input_schema")

    @property
    @pulumi.getter(name="inputSchemaMapping")
    def input_schema_mapping(self) -> Optional['outputs.JsonInputSchemaMappingResponse']:
        """
        Information about the InputSchemaMapping which specified the info about mapping event payload.
        """
        return pulumi.get(self, "input_schema_mapping")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="metricResourceId")
    def metric_resource_id(self) -> str:
        """
        Metric resource id for the domain.
        """
        return pulumi.get(self, "metric_resource_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> Optional[Sequence['outputs.PrivateEndpointConnectionResponse']]:
        """
        List of private endpoint connections.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state of the domain.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[str]:
        """
        This determines if traffic is allowed over public network. By default it is enabled. 
        You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ResourceSkuResponse']:
        """
        The Sku pricing tier for the domain.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetDomainResult(GetDomainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainResult(
            endpoint=self.endpoint,
            id=self.id,
            identity=self.identity,
            inbound_ip_rules=self.inbound_ip_rules,
            input_schema=self.input_schema,
            input_schema_mapping=self.input_schema_mapping,
            location=self.location,
            metric_resource_id=self.metric_resource_id,
            name=self.name,
            private_endpoint_connections=self.private_endpoint_connections,
            provisioning_state=self.provisioning_state,
            public_network_access=self.public_network_access,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_domain(domain_name: Optional[str] = None,
               resource_group_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainResult:
    """
    Use this data source to access information about an existing resource.

    :param str domain_name: Name of the domain.
    :param str resource_group_name: The name of the resource group within the user's subscription.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:eventgrid/v20201015preview:getDomain', __args__, opts=opts, typ=GetDomainResult).value

    return AwaitableGetDomainResult(
        endpoint=__ret__.endpoint,
        id=__ret__.id,
        identity=__ret__.identity,
        inbound_ip_rules=__ret__.inbound_ip_rules,
        input_schema=__ret__.input_schema,
        input_schema_mapping=__ret__.input_schema_mapping,
        location=__ret__.location,
        metric_resource_id=__ret__.metric_resource_id,
        name=__ret__.name,
        private_endpoint_connections=__ret__.private_endpoint_connections,
        provisioning_state=__ret__.provisioning_state,
        public_network_access=__ret__.public_network_access,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type)
