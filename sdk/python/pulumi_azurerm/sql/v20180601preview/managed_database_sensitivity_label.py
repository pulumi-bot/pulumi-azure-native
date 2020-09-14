# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['ManagedDatabaseSensitivityLabel']


class ManagedDatabaseSensitivityLabel(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 column_name: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 information_type: Optional[pulumi.Input[str]] = None,
                 information_type_id: Optional[pulumi.Input[str]] = None,
                 label_id: Optional[pulumi.Input[str]] = None,
                 label_name: Optional[pulumi.Input[str]] = None,
                 managed_instance_name: Optional[pulumi.Input[str]] = None,
                 rank: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 sensitivity_label_source: Optional[pulumi.Input[str]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A sensitivity label.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] column_name: The name of the column.
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] information_type: The information type.
        :param pulumi.Input[str] information_type_id: The information type ID.
        :param pulumi.Input[str] label_id: The label ID.
        :param pulumi.Input[str] label_name: The label name.
        :param pulumi.Input[str] managed_instance_name: The name of the managed instance.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[str] schema_name: The name of the schema.
        :param pulumi.Input[str] sensitivity_label_source: The source of the sensitivity label.
        :param pulumi.Input[str] table_name: The name of the table.
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

            if column_name is None:
                raise TypeError("Missing required property 'column_name'")
            __props__['column_name'] = column_name
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['information_type'] = information_type
            __props__['information_type_id'] = information_type_id
            __props__['label_id'] = label_id
            __props__['label_name'] = label_name
            if managed_instance_name is None:
                raise TypeError("Missing required property 'managed_instance_name'")
            __props__['managed_instance_name'] = managed_instance_name
            __props__['rank'] = rank
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if schema_name is None:
                raise TypeError("Missing required property 'schema_name'")
            __props__['schema_name'] = schema_name
            if sensitivity_label_source is None:
                raise TypeError("Missing required property 'sensitivity_label_source'")
            __props__['sensitivity_label_source'] = sensitivity_label_source
            if table_name is None:
                raise TypeError("Missing required property 'table_name'")
            __props__['table_name'] = table_name
            __props__['is_disabled'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:sql/latest:ManagedDatabaseSensitivityLabel")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ManagedDatabaseSensitivityLabel, __self__).__init__(
            'azurerm:sql/v20180601preview:ManagedDatabaseSensitivityLabel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ManagedDatabaseSensitivityLabel':
        """
        Get an existing ManagedDatabaseSensitivityLabel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ManagedDatabaseSensitivityLabel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="informationType")
    def information_type(self) -> pulumi.Output[Optional[str]]:
        """
        The information type.
        """
        return pulumi.get(self, "information_type")

    @property
    @pulumi.getter(name="informationTypeId")
    def information_type_id(self) -> pulumi.Output[Optional[str]]:
        """
        The information type ID.
        """
        return pulumi.get(self, "information_type_id")

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> pulumi.Output[bool]:
        """
        Is sensitivity recommendation disabled. Applicable for recommended sensitivity label only. Specifies whether the sensitivity recommendation on this column is disabled (dismissed) or not.
        """
        return pulumi.get(self, "is_disabled")

    @property
    @pulumi.getter(name="labelId")
    def label_id(self) -> pulumi.Output[Optional[str]]:
        """
        The label ID.
        """
        return pulumi.get(self, "label_id")

    @property
    @pulumi.getter(name="labelName")
    def label_name(self) -> pulumi.Output[Optional[str]]:
        """
        The label name.
        """
        return pulumi.get(self, "label_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rank(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "rank")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

