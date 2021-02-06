# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ConnectedCluster']


class ConnectedCluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_public_key_certificate: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 distribution: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['ConnectedClusterIdentityArgs']]] = None,
                 infrastructure: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 provisioning_state: Optional[pulumi.Input[Union[str, 'ProvisioningState']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a connected cluster.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_public_key_certificate: Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
        :param pulumi.Input[str] cluster_name: The name of the Kubernetes cluster on which get is called.
        :param pulumi.Input[str] distribution: The Kubernetes distribution running on this connected cluster.
        :param pulumi.Input[pulumi.InputType['ConnectedClusterIdentityArgs']] identity: The identity of the connected cluster.
        :param pulumi.Input[str] infrastructure: The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[Union[str, 'ProvisioningState']] provisioning_state: Provisioning state of the connected cluster resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
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

            if agent_public_key_certificate is None and not opts.urn:
                raise TypeError("Missing required property 'agent_public_key_certificate'")
            __props__['agent_public_key_certificate'] = agent_public_key_certificate
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['distribution'] = distribution
            if identity is None and not opts.urn:
                raise TypeError("Missing required property 'identity'")
            __props__['identity'] = identity
            __props__['infrastructure'] = infrastructure
            __props__['location'] = location
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['agent_version'] = None
            __props__['connectivity_status'] = None
            __props__['kubernetes_version'] = None
            __props__['last_connectivity_time'] = None
            __props__['managed_identity_certificate_expiration_time'] = None
            __props__['name'] = None
            __props__['offering'] = None
            __props__['system_data'] = None
            __props__['total_core_count'] = None
            __props__['total_node_count'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:kubernetes/v20200101preview:ConnectedCluster"), pulumi.Alias(type_="azure-nextgen:kubernetes/v20210401preview:ConnectedCluster")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ConnectedCluster, __self__).__init__(
            'azure-nextgen:kubernetes/v20210301:ConnectedCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConnectedCluster':
        """
        Get an existing ConnectedCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ConnectedCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentPublicKeyCertificate")
    def agent_public_key_certificate(self) -> pulumi.Output[str]:
        """
        Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
        """
        return pulumi.get(self, "agent_public_key_certificate")

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> pulumi.Output[str]:
        """
        Version of the agent running on the connected cluster resource
        """
        return pulumi.get(self, "agent_version")

    @property
    @pulumi.getter(name="connectivityStatus")
    def connectivity_status(self) -> pulumi.Output[str]:
        """
        Represents the connectivity status of the connected cluster.
        """
        return pulumi.get(self, "connectivity_status")

    @property
    @pulumi.getter
    def distribution(self) -> pulumi.Output[Optional[str]]:
        """
        The Kubernetes distribution running on this connected cluster.
        """
        return pulumi.get(self, "distribution")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output['outputs.ConnectedClusterIdentityResponse']:
        """
        The identity of the connected cluster.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def infrastructure(self) -> pulumi.Output[Optional[str]]:
        """
        The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
        """
        return pulumi.get(self, "infrastructure")

    @property
    @pulumi.getter(name="kubernetesVersion")
    def kubernetes_version(self) -> pulumi.Output[str]:
        """
        The Kubernetes version of the connected cluster resource
        """
        return pulumi.get(self, "kubernetes_version")

    @property
    @pulumi.getter(name="lastConnectivityTime")
    def last_connectivity_time(self) -> pulumi.Output[str]:
        """
        Time representing the last instance when heart beat was received from the cluster
        """
        return pulumi.get(self, "last_connectivity_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedIdentityCertificateExpirationTime")
    def managed_identity_certificate_expiration_time(self) -> pulumi.Output[str]:
        """
        Expiration time of the managed identity certificate
        """
        return pulumi.get(self, "managed_identity_certificate_expiration_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def offering(self) -> pulumi.Output[str]:
        """
        Connected cluster offering
        """
        return pulumi.get(self, "offering")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        Provisioning state of the connected cluster resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Metadata pertaining to creation and last modification of the resource
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCoreCount")
    def total_core_count(self) -> pulumi.Output[int]:
        """
        Number of CPU cores present in the connected cluster resource
        """
        return pulumi.get(self, "total_core_count")

    @property
    @pulumi.getter(name="totalNodeCount")
    def total_node_count(self) -> pulumi.Output[int]:
        """
        Number of nodes present in the connected cluster resource
        """
        return pulumi.get(self, "total_node_count")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

