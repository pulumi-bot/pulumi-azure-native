# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'GetEventHubConnectionResult',
    'AwaitableGetEventHubConnectionResult',
    'get_event_hub_connection',
]

@pulumi.output_type
class GetEventHubConnectionResult:
    """
    Class representing an event hub connection.
    """
    def __init__(__self__, consumer_group=None, data_format=None, event_hub_resource_id=None, location=None, mapping_rule_name=None, name=None, table_name=None, type=None):
        if consumer_group and not isinstance(consumer_group, str):
            raise TypeError("Expected argument 'consumer_group' to be a str")
        pulumi.set(__self__, "consumer_group", consumer_group)
        if data_format and not isinstance(data_format, str):
            raise TypeError("Expected argument 'data_format' to be a str")
        pulumi.set(__self__, "data_format", data_format)
        if event_hub_resource_id and not isinstance(event_hub_resource_id, str):
            raise TypeError("Expected argument 'event_hub_resource_id' to be a str")
        pulumi.set(__self__, "event_hub_resource_id", event_hub_resource_id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if mapping_rule_name and not isinstance(mapping_rule_name, str):
            raise TypeError("Expected argument 'mapping_rule_name' to be a str")
        pulumi.set(__self__, "mapping_rule_name", mapping_rule_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if table_name and not isinstance(table_name, str):
            raise TypeError("Expected argument 'table_name' to be a str")
        pulumi.set(__self__, "table_name", table_name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="consumerGroup")
    def consumer_group(self) -> str:
        """
        The event hub consumer group.
        """
        return pulumi.get(self, "consumer_group")

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> Optional[str]:
        """
        The data format of the message. Optionally the data format can be added to each message.
        """
        return pulumi.get(self, "data_format")

    @property
    @pulumi.getter(name="eventHubResourceId")
    def event_hub_resource_id(self) -> str:
        """
        The resource ID of the event hub to be used to create a data connection.
        """
        return pulumi.get(self, "event_hub_resource_id")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mappingRuleName")
    def mapping_rule_name(self) -> Optional[str]:
        """
        The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
        """
        return pulumi.get(self, "mapping_rule_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[str]:
        """
        The table where the data should be ingested. Optionally the table information can be added to each message.
        """
        return pulumi.get(self, "table_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetEventHubConnectionResult(GetEventHubConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventHubConnectionResult(
            consumer_group=self.consumer_group,
            data_format=self.data_format,
            event_hub_resource_id=self.event_hub_resource_id,
            location=self.location,
            mapping_rule_name=self.mapping_rule_name,
            name=self.name,
            table_name=self.table_name,
            type=self.type)


def get_event_hub_connection(cluster_name: Optional[str] = None,
                             database_name: Optional[str] = None,
                             event_hub_connection_name: Optional[str] = None,
                             resource_group_name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventHubConnectionResult:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_name: The name of the Kusto cluster.
    :param str database_name: The name of the database in the Kusto cluster.
    :param str event_hub_connection_name: The name of the event hub connection.
    :param str resource_group_name: The name of the resource group containing the Kusto cluster.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['databaseName'] = database_name
    __args__['eventHubConnectionName'] = event_hub_connection_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:kusto/v20170907privatepreview:getEventHubConnection', __args__, opts=opts, typ=GetEventHubConnectionResult).value

    return AwaitableGetEventHubConnectionResult(
        consumer_group=__ret__.consumer_group,
        data_format=__ret__.data_format,
        event_hub_resource_id=__ret__.event_hub_resource_id,
        location=__ret__.location,
        mapping_rule_name=__ret__.mapping_rule_name,
        name=__ret__.name,
        table_name=__ret__.table_name,
        type=__ret__.type)
