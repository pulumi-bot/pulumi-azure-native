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

__all__ = ['Bookmark']


class Bookmark(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bookmark_id: Optional[pulumi.Input[str]] = None,
                 created: Optional[pulumi.Input[str]] = None,
                 created_by: Optional[pulumi.Input[pulumi.InputType['UserInfoArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 incident_info: Optional[pulumi.Input[pulumi.InputType['IncidentInfoArgs']]] = None,
                 labels: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 query_result: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 updated: Optional[pulumi.Input[str]] = None,
                 updated_by: Optional[pulumi.Input[pulumi.InputType['UserInfoArgs']]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a bookmark in Azure Security Insights.

        ## Example Usage
        ### Creates or updates a bookmark.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        bookmark = azurerm.operationalinsights.latest.Bookmark("bookmark",
            bookmark_id="73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            created="2019-01-01T13:15:30Z",
            created_by={
                "objectId": "2046feea-040d-4a46-9e2b-91c2941bfa70",
            },
            display_name="My bookmark",
            etag="\"0300bf09-0000-0000-0000-5c37296e0000\"",
            labels=[
                "Tag1",
                "Tag2",
            ],
            notes="Found a suspicious activity",
            query="SecurityEvent | where TimeGenerated > ago(1d) and TimeGenerated < ago(2d)",
            query_result="Security Event query result",
            resource_group_name="myRg",
            updated="2019-01-01T13:15:30Z",
            updated_by={
                "objectId": "2046feea-040d-4a46-9e2b-91c2941bfa70",
            },
            workspace_name="myWorkspace")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bookmark_id: Bookmark ID
        :param pulumi.Input[str] created: The time the bookmark was created
        :param pulumi.Input[pulumi.InputType['UserInfoArgs']] created_by: Describes a user that created the bookmark
        :param pulumi.Input[str] display_name: The display name of the bookmark
        :param pulumi.Input[str] etag: Etag of the azure resource
        :param pulumi.Input[pulumi.InputType['IncidentInfoArgs']] incident_info: Describes an incident that relates to bookmark
        :param pulumi.Input[List[pulumi.Input[str]]] labels: List of labels relevant to this bookmark
        :param pulumi.Input[str] notes: The notes of the bookmark
        :param pulumi.Input[str] query: The query of the bookmark.
        :param pulumi.Input[str] query_result: The query result of the bookmark.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
        :param pulumi.Input[str] updated: The last time the bookmark was updated
        :param pulumi.Input[pulumi.InputType['UserInfoArgs']] updated_by: Describes a user that updated the bookmark
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

            if bookmark_id is None:
                raise TypeError("Missing required property 'bookmark_id'")
            __props__['bookmark_id'] = bookmark_id
            __props__['created'] = created
            __props__['created_by'] = created_by
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['etag'] = etag
            __props__['incident_info'] = incident_info
            __props__['labels'] = labels
            __props__['notes'] = notes
            if query is None:
                raise TypeError("Missing required property 'query'")
            __props__['query'] = query
            __props__['query_result'] = query_result
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['updated'] = updated
            __props__['updated_by'] = updated_by
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:operationalinsights/v20200101:Bookmark")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Bookmark, __self__).__init__(
            'azurerm:operationalinsights/latest:Bookmark',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Bookmark':
        """
        Get an existing Bookmark resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Bookmark(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[Optional[str]]:
        """
        The time the bookmark was created
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[Optional['outputs.UserInfoResponse']]:
        """
        Describes a user that created the bookmark
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the bookmark
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        Etag of the azure resource
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="incidentInfo")
    def incident_info(self) -> pulumi.Output[Optional['outputs.IncidentInfoResponse']]:
        """
        Describes an incident that relates to bookmark
        """
        return pulumi.get(self, "incident_info")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of labels relevant to this bookmark
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        """
        The notes of the bookmark
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def query(self) -> pulumi.Output[str]:
        """
        The query of the bookmark.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="queryResult")
    def query_result(self) -> pulumi.Output[Optional[str]]:
        """
        The query result of the bookmark.
        """
        return pulumi.get(self, "query_result")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def updated(self) -> pulumi.Output[Optional[str]]:
        """
        The last time the bookmark was updated
        """
        return pulumi.get(self, "updated")

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> pulumi.Output[Optional['outputs.UserInfoResponse']]:
        """
        Describes a user that updated the bookmark
        """
        return pulumi.get(self, "updated_by")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

