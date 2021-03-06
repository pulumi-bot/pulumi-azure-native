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

__all__ = ['Watchlist']


class Watchlist(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 created: Optional[pulumi.Input[str]] = None,
                 created_by: Optional[pulumi.Input[pulumi.InputType['UserInfoArgs']]] = None,
                 default_duration: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 is_deleted: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 number_of_lines_to_skip: Optional[pulumi.Input[int]] = None,
                 operational_insights_resource_provider: Optional[pulumi.Input[str]] = None,
                 provider: Optional[pulumi.Input[str]] = None,
                 raw_content: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[Union[str, 'Source']]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 updated: Optional[pulumi.Input[str]] = None,
                 updated_by: Optional[pulumi.Input[pulumi.InputType['UserInfoArgs']]] = None,
                 upload_status: Optional[pulumi.Input[str]] = None,
                 watchlist_alias: Optional[pulumi.Input[str]] = None,
                 watchlist_id: Optional[pulumi.Input[str]] = None,
                 watchlist_items_count: Optional[pulumi.Input[int]] = None,
                 watchlist_type: Optional[pulumi.Input[str]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a Watchlist in Azure Security Insights.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_type: The content type of the raw content. Example : text/csv or text/tsv 
        :param pulumi.Input[str] created: The time the watchlist was created
        :param pulumi.Input[pulumi.InputType['UserInfoArgs']] created_by: Describes a user that created the watchlist
        :param pulumi.Input[str] default_duration: The default duration of a watchlist (in ISO 8601 duration format)
        :param pulumi.Input[str] description: A description of the watchlist
        :param pulumi.Input[str] display_name: The display name of the watchlist
        :param pulumi.Input[str] etag: Etag of the azure resource
        :param pulumi.Input[bool] is_deleted: A flag that indicates if the watchlist is deleted or not
        :param pulumi.Input[Sequence[pulumi.Input[str]]] labels: List of labels relevant to this watchlist
        :param pulumi.Input[int] number_of_lines_to_skip: The number of lines in a csv/tsv content to skip before the header
        :param pulumi.Input[str] operational_insights_resource_provider: The namespace of workspaces resource provider- Microsoft.OperationalInsights.
        :param pulumi.Input[str] provider: The provider of the watchlist
        :param pulumi.Input[str] raw_content: The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
        :param pulumi.Input[Union[str, 'Source']] source: The source of the watchlist
        :param pulumi.Input[str] tenant_id: The tenantId where the watchlist belongs to
        :param pulumi.Input[str] updated: The last time the watchlist was updated
        :param pulumi.Input[pulumi.InputType['UserInfoArgs']] updated_by: Describes a user that updated the watchlist
        :param pulumi.Input[str] upload_status: The status of the Watchlist upload : New, InProgress or Complete. Pls note : When a Watchlist upload status is equal to InProgress, the Watchlist cannot be deleted
        :param pulumi.Input[str] watchlist_alias: The alias of the watchlist
        :param pulumi.Input[str] watchlist_id: The id (a Guid) of the watchlist
        :param pulumi.Input[int] watchlist_items_count: The number of Watchlist Items in the Watchlist
        :param pulumi.Input[str] watchlist_type: The type of the watchlist
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

            __props__['content_type'] = content_type
            __props__['created'] = created
            __props__['created_by'] = created_by
            __props__['default_duration'] = default_duration
            __props__['description'] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['etag'] = etag
            __props__['is_deleted'] = is_deleted
            __props__['labels'] = labels
            __props__['number_of_lines_to_skip'] = number_of_lines_to_skip
            if operational_insights_resource_provider is None and not opts.urn:
                raise TypeError("Missing required property 'operational_insights_resource_provider'")
            __props__['operational_insights_resource_provider'] = operational_insights_resource_provider
            if provider is None and not opts.urn:
                raise TypeError("Missing required property 'provider'")
            __props__['provider'] = provider
            __props__['raw_content'] = raw_content
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__['source'] = source
            __props__['tenant_id'] = tenant_id
            __props__['updated'] = updated
            __props__['updated_by'] = updated_by
            __props__['upload_status'] = upload_status
            __props__['watchlist_alias'] = watchlist_alias
            __props__['watchlist_id'] = watchlist_id
            __props__['watchlist_items_count'] = watchlist_items_count
            __props__['watchlist_type'] = watchlist_type
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:securityinsights/v20190101preview:Watchlist"), pulumi.Alias(type_="azure-native:securityinsights:Watchlist"), pulumi.Alias(type_="azure-nextgen:securityinsights:Watchlist")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Watchlist, __self__).__init__(
            'azure-native:securityinsights/v20190101preview:Watchlist',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Watchlist':
        """
        Get an existing Watchlist resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["content_type"] = None
        __props__["created"] = None
        __props__["created_by"] = None
        __props__["default_duration"] = None
        __props__["description"] = None
        __props__["display_name"] = None
        __props__["etag"] = None
        __props__["is_deleted"] = None
        __props__["labels"] = None
        __props__["name"] = None
        __props__["number_of_lines_to_skip"] = None
        __props__["provider"] = None
        __props__["raw_content"] = None
        __props__["source"] = None
        __props__["tenant_id"] = None
        __props__["type"] = None
        __props__["updated"] = None
        __props__["updated_by"] = None
        __props__["upload_status"] = None
        __props__["watchlist_alias"] = None
        __props__["watchlist_id"] = None
        __props__["watchlist_items_count"] = None
        __props__["watchlist_type"] = None
        return Watchlist(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[Optional[str]]:
        """
        The content type of the raw content. Example : text/csv or text/tsv 
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[Optional[str]]:
        """
        The time the watchlist was created
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[Optional['outputs.UserInfoResponse']]:
        """
        Describes a user that created the watchlist
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="defaultDuration")
    def default_duration(self) -> pulumi.Output[Optional[str]]:
        """
        The default duration of a watchlist (in ISO 8601 duration format)
        """
        return pulumi.get(self, "default_duration")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the watchlist
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the watchlist
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
    @pulumi.getter(name="isDeleted")
    def is_deleted(self) -> pulumi.Output[Optional[bool]]:
        """
        A flag that indicates if the watchlist is deleted or not
        """
        return pulumi.get(self, "is_deleted")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of labels relevant to this watchlist
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
    @pulumi.getter(name="numberOfLinesToSkip")
    def number_of_lines_to_skip(self) -> pulumi.Output[Optional[int]]:
        """
        The number of lines in a csv/tsv content to skip before the header
        """
        return pulumi.get(self, "number_of_lines_to_skip")

    @property
    @pulumi.getter
    def provider(self) -> pulumi.Output[str]:
        """
        The provider of the watchlist
        """
        return pulumi.get(self, "provider")

    @property
    @pulumi.getter(name="rawContent")
    def raw_content(self) -> pulumi.Output[Optional[str]]:
        """
        The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
        """
        return pulumi.get(self, "raw_content")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        The source of the watchlist
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The tenantId where the watchlist belongs to
        """
        return pulumi.get(self, "tenant_id")

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
        The last time the watchlist was updated
        """
        return pulumi.get(self, "updated")

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> pulumi.Output[Optional['outputs.UserInfoResponse']]:
        """
        Describes a user that updated the watchlist
        """
        return pulumi.get(self, "updated_by")

    @property
    @pulumi.getter(name="uploadStatus")
    def upload_status(self) -> pulumi.Output[Optional[str]]:
        """
        The status of the Watchlist upload : New, InProgress or Complete. Pls note : When a Watchlist upload status is equal to InProgress, the Watchlist cannot be deleted
        """
        return pulumi.get(self, "upload_status")

    @property
    @pulumi.getter(name="watchlistAlias")
    def watchlist_alias(self) -> pulumi.Output[Optional[str]]:
        """
        The alias of the watchlist
        """
        return pulumi.get(self, "watchlist_alias")

    @property
    @pulumi.getter(name="watchlistId")
    def watchlist_id(self) -> pulumi.Output[Optional[str]]:
        """
        The id (a Guid) of the watchlist
        """
        return pulumi.get(self, "watchlist_id")

    @property
    @pulumi.getter(name="watchlistItemsCount")
    def watchlist_items_count(self) -> pulumi.Output[Optional[int]]:
        """
        The number of Watchlist Items in the Watchlist
        """
        return pulumi.get(self, "watchlist_items_count")

    @property
    @pulumi.getter(name="watchlistType")
    def watchlist_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the watchlist
        """
        return pulumi.get(self, "watchlist_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

