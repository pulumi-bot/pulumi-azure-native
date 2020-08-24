# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['DataExport']


class DataExport(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_tables: Optional[pulumi.Input[bool]] = None,
                 created_date: Optional[pulumi.Input[str]] = None,
                 data_export_id: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 event_hub_name: Optional[pulumi.Input[str]] = None,
                 last_modified_date: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 table_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The top level data export resource container.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_tables: When ‘true’, all workspace's tables are exported.
        :param pulumi.Input[str] created_date: The latest data export rule modification time.
        :param pulumi.Input[str] data_export_id: The data export rule ID.
        :param pulumi.Input[bool] enable: Active when enabled.
        :param pulumi.Input[str] event_hub_name: Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
        :param pulumi.Input[str] last_modified_date: Date and time when the export was last modified.
        :param pulumi.Input[str] name: The data export rule name.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] resource_id: The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
        :param pulumi.Input[List[pulumi.Input[str]]] table_names: An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
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

            __props__['all_tables'] = all_tables
            __props__['created_date'] = created_date
            __props__['data_export_id'] = data_export_id
            __props__['enable'] = enable
            __props__['event_hub_name'] = event_hub_name
            __props__['last_modified_date'] = last_modified_date
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_id is None:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            __props__['table_names'] = table_names
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['type'] = None
        super(DataExport, __self__).__init__(
            'azurerm:operationalinsights/v20200801:DataExport',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataExport':
        """
        Get an existing DataExport resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DataExport(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allTables")
    def all_tables(self) -> Optional[bool]:
        """
        When ‘true’, all workspace's tables are exported.
        """
        return pulumi.get(self, "all_tables")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> Optional[str]:
        """
        The latest data export rule modification time.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter(name="dataExportId")
    def data_export_id(self) -> Optional[str]:
        """
        The data export rule ID.
        """
        return pulumi.get(self, "data_export_id")

    @property
    @pulumi.getter
    def enable(self) -> Optional[bool]:
        """
        Active when enabled.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="eventHubName")
    def event_hub_name(self) -> Optional[str]:
        """
        Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
        """
        return pulumi.get(self, "event_hub_name")

    @property
    @pulumi.getter(name="lastModifiedDate")
    def last_modified_date(self) -> Optional[str]:
        """
        Date and time when the export was last modified.
        """
        return pulumi.get(self, "last_modified_date")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="tableNames")
    def table_names(self) -> Optional[List[str]]:
        """
        An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
        """
        return pulumi.get(self, "table_names")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

