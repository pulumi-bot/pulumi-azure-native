# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['CaseRelation']


class CaseRelation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 case_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 operational_insights_resource_provider: Optional[pulumi.Input[str]] = None,
                 relation_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_relation_node: Optional[pulumi.Input[pulumi.InputType['RelationNodeArgs']]] = None,
                 target_relation_node: Optional[pulumi.Input[pulumi.InputType['RelationNodeArgs']]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a case relation

        ## Example Usage
        ### Creates or updates a case relation.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        case_relation = azurerm.securityinsights.v20190101preview.CaseRelation("caseRelation",
            case_id="afbd324f-6c48-459c-8710-8d1e1cd03812",
            operational_insights_resource_provider="Microsoft.OperationalInsights",
            relation_name="4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014",
            resource_group_name="myRg",
            source_relation_node={
                "relationNodeId": "afbd324f-6c48-459c-8710-8d1e1cd03812",
            },
            target_relation_node={
                "relationNodeId": "2216d0e1-91e3-4902-89fd-d2df8c535096",
            },
            workspace_name="myWorkspace")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] case_id: Case ID
        :param pulumi.Input[str] etag: ETag for relation
        :param pulumi.Input[str] operational_insights_resource_provider: The namespace of workspaces resource provider- Microsoft.OperationalInsights.
        :param pulumi.Input[str] relation_name: Name of relation
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
        :param pulumi.Input[pulumi.InputType['RelationNodeArgs']] source_relation_node: Relation source node
        :param pulumi.Input[pulumi.InputType['RelationNodeArgs']] target_relation_node: Relation target node
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

            if case_id is None:
                raise TypeError("Missing required property 'case_id'")
            __props__['case_id'] = case_id
            __props__['etag'] = etag
            if operational_insights_resource_provider is None:
                raise TypeError("Missing required property 'operational_insights_resource_provider'")
            __props__['operational_insights_resource_provider'] = operational_insights_resource_provider
            if relation_name is None:
                raise TypeError("Missing required property 'relation_name'")
            __props__['relation_name'] = relation_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['source_relation_node'] = source_relation_node
            __props__['target_relation_node'] = target_relation_node
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['bookmark_id'] = None
            __props__['bookmark_name'] = None
            __props__['case_identifier'] = None
            __props__['kind'] = None
            __props__['name'] = None
            __props__['type'] = None
        super(CaseRelation, __self__).__init__(
            'azurerm:securityinsights/v20190101preview:CaseRelation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CaseRelation':
        """
        Get an existing CaseRelation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return CaseRelation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bookmarkId")
    def bookmark_id(self) -> pulumi.Output[str]:
        """
        The case related bookmark id
        """
        return pulumi.get(self, "bookmark_id")

    @property
    @pulumi.getter(name="bookmarkName")
    def bookmark_name(self) -> pulumi.Output[Optional[str]]:
        """
        The case related bookmark name
        """
        return pulumi.get(self, "bookmark_name")

    @property
    @pulumi.getter(name="caseIdentifier")
    def case_identifier(self) -> pulumi.Output[str]:
        """
        The case identifier
        """
        return pulumi.get(self, "case_identifier")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        ETag for relation
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The type of relation node
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="relationName")
    def relation_name(self) -> pulumi.Output[str]:
        """
        Name of relation
        """
        return pulumi.get(self, "relation_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

