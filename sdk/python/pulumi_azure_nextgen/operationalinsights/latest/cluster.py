# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Cluster']


class Cluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['IdentityArgs']]] = None,
                 is_availability_zones_enabled: Optional[pulumi.Input[bool]] = None,
                 is_double_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 key_vault_properties: Optional[pulumi.Input[pulumi.InputType['KeyVaultPropertiesArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ClusterSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The top level Log Analytics cluster resource container.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_type: Configures whether billing will be only on the cluster or each workspace will be billed by its proportional use. This does not change the overall billing, only how it will be distributed. Default value is 'Cluster'
        :param pulumi.Input[str] cluster_name: The name of the Log Analytics cluster.
        :param pulumi.Input[pulumi.InputType['IdentityArgs']] identity: The identity of the resource.
        :param pulumi.Input[bool] is_availability_zones_enabled: Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
        :param pulumi.Input[bool] is_double_encryption_enabled: Configures whether cluster will use double encryption. This Property can not be modified after cluster creation. Default value is 'true'
        :param pulumi.Input[pulumi.InputType['KeyVaultPropertiesArgs']] key_vault_properties: The associated key properties.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[pulumi.InputType['ClusterSkuArgs']] sku: The sku properties.
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

            __props__['billing_type'] = billing_type
            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['identity'] = identity
            __props__['is_availability_zones_enabled'] = is_availability_zones_enabled
            __props__['is_double_encryption_enabled'] = is_double_encryption_enabled
            __props__['key_vault_properties'] = key_vault_properties
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['cluster_id'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:operationalinsights/v20190801preview:Cluster"), pulumi.Alias(type_="azure-nextgen:operationalinsights/v20200301preview:Cluster"), pulumi.Alias(type_="azure-nextgen:operationalinsights/v20200801:Cluster"), pulumi.Alias(type_="azure-nextgen:operationalinsights/v20201001:Cluster")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Cluster, __self__).__init__(
            'azure-nextgen:operationalinsights/latest:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> pulumi.Output[Optional[str]]:
        """
        Configures whether billing will be only on the cluster or each workspace will be billed by its proportional use. This does not change the overall billing, only how it will be distributed. Default value is 'Cluster'
        """
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The ID associated with the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.IdentityResponse']]:
        """
        The identity of the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="isAvailabilityZonesEnabled")
    def is_availability_zones_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
        """
        return pulumi.get(self, "is_availability_zones_enabled")

    @property
    @pulumi.getter(name="isDoubleEncryptionEnabled")
    def is_double_encryption_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Configures whether cluster will use double encryption. This Property can not be modified after cluster creation. Default value is 'true'
        """
        return pulumi.get(self, "is_double_encryption_enabled")

    @property
    @pulumi.getter(name="keyVaultProperties")
    def key_vault_properties(self) -> pulumi.Output[Optional['outputs.KeyVaultPropertiesResponse']]:
        """
        The associated key properties.
        """
        return pulumi.get(self, "key_vault_properties")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the cluster.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.ClusterSkuResponse']]:
        """
        The sku properties.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

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

