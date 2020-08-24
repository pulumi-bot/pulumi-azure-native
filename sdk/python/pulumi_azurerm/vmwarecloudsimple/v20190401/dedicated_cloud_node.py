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

__all__ = ['DedicatedCloudNode']


class DedicatedCloudNode(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone_id: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nodes_count: Optional[pulumi.Input[float]] = None,
                 placement_group_id: Optional[pulumi.Input[str]] = None,
                 purchase_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['SkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Dedicated cloud node model

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone_id: Availability Zone id, e.g. "az1"
        :param pulumi.Input[str] id: SKU's id
        :param pulumi.Input[str] location: Azure region
        :param pulumi.Input[str] name: dedicated cloud node name
        :param pulumi.Input[float] nodes_count: count of nodes to create
        :param pulumi.Input[str] placement_group_id: Placement Group id, e.g. "n1"
        :param pulumi.Input[str] purchase_id: purchase id
        :param pulumi.Input[str] resource_group_name: The name of the resource group
        :param pulumi.Input[pulumi.InputType['SkuArgs']] sku: Dedicated Cloud Nodes SKU
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Dedicated Cloud Nodes tags
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

            if availability_zone_id is None:
                raise TypeError("Missing required property 'availability_zone_id'")
            __props__['availability_zone_id'] = availability_zone_id
            if id is None:
                raise TypeError("Missing required property 'id'")
            __props__['id'] = id
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if nodes_count is None:
                raise TypeError("Missing required property 'nodes_count'")
            __props__['nodes_count'] = nodes_count
            if placement_group_id is None:
                raise TypeError("Missing required property 'placement_group_id'")
            __props__['placement_group_id'] = placement_group_id
            if purchase_id is None:
                raise TypeError("Missing required property 'purchase_id'")
            __props__['purchase_id'] = purchase_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['availability_zone_name'] = None
            __props__['cloud_rack_name'] = None
            __props__['created'] = None
            __props__['placement_group_name'] = None
            __props__['private_cloud_id'] = None
            __props__['private_cloud_name'] = None
            __props__['provisioning_state'] = None
            __props__['status'] = None
            __props__['type'] = None
            __props__['vmware_cluster_name'] = None
        super(DedicatedCloudNode, __self__).__init__(
            'azurerm:vmwarecloudsimple/v20190401:DedicatedCloudNode',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DedicatedCloudNode':
        """
        Get an existing DedicatedCloudNode resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DedicatedCloudNode(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZoneId")
    def availability_zone_id(self) -> str:
        """
        Availability Zone id, e.g. "az1"
        """
        return pulumi.get(self, "availability_zone_id")

    @property
    @pulumi.getter(name="availabilityZoneName")
    def availability_zone_name(self) -> str:
        """
        Availability Zone name, e.g. "Availability Zone 1"
        """
        return pulumi.get(self, "availability_zone_name")

    @property
    @pulumi.getter(name="cloudRackName")
    def cloud_rack_name(self) -> str:
        """
        VMWare Cloud Rack Name
        """
        return pulumi.get(self, "cloud_rack_name")

    @property
    @pulumi.getter
    def created(self) -> Mapping[str, Any]:
        """
        date time the resource was created
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Azure region
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        SKU's name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodesCount")
    def nodes_count(self) -> float:
        """
        count of nodes to create
        """
        return pulumi.get(self, "nodes_count")

    @property
    @pulumi.getter(name="placementGroupId")
    def placement_group_id(self) -> str:
        """
        Placement Group id, e.g. "n1"
        """
        return pulumi.get(self, "placement_group_id")

    @property
    @pulumi.getter(name="placementGroupName")
    def placement_group_name(self) -> str:
        """
        Placement Name, e.g. "Placement Group 1"
        """
        return pulumi.get(self, "placement_group_name")

    @property
    @pulumi.getter(name="privateCloudId")
    def private_cloud_id(self) -> str:
        """
        Private Cloud Id
        """
        return pulumi.get(self, "private_cloud_id")

    @property
    @pulumi.getter(name="privateCloudName")
    def private_cloud_name(self) -> str:
        """
        Resource Pool Name
        """
        return pulumi.get(self, "private_cloud_name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning status of the resource
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="purchaseId")
    def purchase_id(self) -> str:
        """
        purchase id
        """
        return pulumi.get(self, "purchase_id")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        Dedicated Cloud Nodes SKU
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Node status, indicates is private cloud set up on this node or not
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Dedicated Cloud Nodes tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        {resourceProviderNamespace}/{resourceType}
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmwareClusterName")
    def vmware_cluster_name(self) -> str:
        """
        VMWare Cluster Name
        """
        return pulumi.get(self, "vmware_cluster_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

