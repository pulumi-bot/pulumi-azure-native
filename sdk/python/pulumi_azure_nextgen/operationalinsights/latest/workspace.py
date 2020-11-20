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

__all__ = ['Workspace']


class Workspace(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 e_tag: Optional[pulumi.Input[str]] = None,
                 force_cmk_for_query: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 public_network_access_for_ingestion: Optional[pulumi.Input[str]] = None,
                 public_network_access_for_query: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retention_in_days: Optional[pulumi.Input[int]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['WorkspaceSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workspace_capping: Optional[pulumi.Input[pulumi.InputType['WorkspaceCappingArgs']]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The top level Workspace resource container.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] e_tag: The ETag of the workspace.
        :param pulumi.Input[bool] force_cmk_for_query: Indicates whether customer managed storage is mandatory for query management.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] provisioning_state: The provisioning state of the workspace.
        :param pulumi.Input[str] public_network_access_for_ingestion: The network access type for accessing Log Analytics ingestion.
        :param pulumi.Input[str] public_network_access_for_query: The network access type for accessing Log Analytics query.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[int] retention_in_days: The workspace data retention in days, between 30 and 730.
        :param pulumi.Input[pulumi.InputType['WorkspaceSkuArgs']] sku: The SKU of the workspace.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[pulumi.InputType['WorkspaceCappingArgs']] workspace_capping: The daily volume cap for ingestion.
        :param pulumi.Input[str] workspace_name: The name of the workspace.
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

            __props__['e_tag'] = e_tag
            __props__['force_cmk_for_query'] = force_cmk_for_query
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['provisioning_state'] = provisioning_state
            __props__['public_network_access_for_ingestion'] = public_network_access_for_ingestion
            __props__['public_network_access_for_query'] = public_network_access_for_query
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['retention_in_days'] = retention_in_days
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['workspace_capping'] = workspace_capping
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['customer_id'] = None
            __props__['name'] = None
            __props__['private_link_scoped_resources'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:operationalinsights/v20151101preview:Workspace"), pulumi.Alias(type_="azure-nextgen:operationalinsights/v20200301preview:Workspace"), pulumi.Alias(type_="azure-nextgen:operationalinsights/v20200801:Workspace"), pulumi.Alias(type_="azure-nextgen:operationalinsights/v20201001:Workspace")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Workspace, __self__).__init__(
            'azure-nextgen:operationalinsights/latest:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Workspace':
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Workspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customerId")
    def customer_id(self) -> pulumi.Output[str]:
        """
        This is a read-only property. Represents the ID associated with the workspace.
        """
        return pulumi.get(self, "customer_id")

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> pulumi.Output[Optional[str]]:
        """
        The ETag of the workspace.
        """
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter(name="forceCmkForQuery")
    def force_cmk_for_query(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether customer managed storage is mandatory for query management.
        """
        return pulumi.get(self, "force_cmk_for_query")

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
    @pulumi.getter(name="privateLinkScopedResources")
    def private_link_scoped_resources(self) -> pulumi.Output[Sequence['outputs.PrivateLinkScopedResourceResponse']]:
        """
        List of linked private link scope resources.
        """
        return pulumi.get(self, "private_link_scoped_resources")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        The provisioning state of the workspace.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccessForIngestion")
    def public_network_access_for_ingestion(self) -> pulumi.Output[Optional[str]]:
        """
        The network access type for accessing Log Analytics ingestion.
        """
        return pulumi.get(self, "public_network_access_for_ingestion")

    @property
    @pulumi.getter(name="publicNetworkAccessForQuery")
    def public_network_access_for_query(self) -> pulumi.Output[Optional[str]]:
        """
        The network access type for accessing Log Analytics query.
        """
        return pulumi.get(self, "public_network_access_for_query")

    @property
    @pulumi.getter(name="retentionInDays")
    def retention_in_days(self) -> pulumi.Output[Optional[int]]:
        """
        The workspace data retention in days, between 30 and 730.
        """
        return pulumi.get(self, "retention_in_days")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.WorkspaceSkuResponse']]:
        """
        The SKU of the workspace.
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

    @property
    @pulumi.getter(name="workspaceCapping")
    def workspace_capping(self) -> pulumi.Output[Optional['outputs.WorkspaceCappingResponse']]:
        """
        The daily volume cap for ingestion.
        """
        return pulumi.get(self, "workspace_capping")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

