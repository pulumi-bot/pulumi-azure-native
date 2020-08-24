# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ActionArgs',
    'ScheduleArgs',
    'SourceArgs',
]

@pulumi.input_type
class ActionArgs:
    def __init__(__self__, *,
                 odata_type: pulumi.Input[str]):
        """
        Action descriptor.
        :param pulumi.Input[str] odata_type: Specifies the action. Supported values - AlertingAction, LogToMetricAction
        """
        pulumi.set(__self__, "odata_type", odata_type)

    @property
    @pulumi.getter(name="odataType")
    def odata_type(self) -> pulumi.Input[str]:
        """
        Specifies the action. Supported values - AlertingAction, LogToMetricAction
        """
        return pulumi.get(self, "odata_type")

    @odata_type.setter
    def odata_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "odata_type", value)


@pulumi.input_type
class ScheduleArgs:
    def __init__(__self__, *,
                 frequency_in_minutes: pulumi.Input[float],
                 time_window_in_minutes: pulumi.Input[float]):
        """
        Defines how often to run the search and the time interval.
        :param pulumi.Input[float] frequency_in_minutes: frequency (in minutes) at which rule condition should be evaluated.
        :param pulumi.Input[float] time_window_in_minutes: Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
        """
        pulumi.set(__self__, "frequency_in_minutes", frequency_in_minutes)
        pulumi.set(__self__, "time_window_in_minutes", time_window_in_minutes)

    @property
    @pulumi.getter(name="frequencyInMinutes")
    def frequency_in_minutes(self) -> pulumi.Input[float]:
        """
        frequency (in minutes) at which rule condition should be evaluated.
        """
        return pulumi.get(self, "frequency_in_minutes")

    @frequency_in_minutes.setter
    def frequency_in_minutes(self, value: pulumi.Input[float]):
        pulumi.set(self, "frequency_in_minutes", value)

    @property
    @pulumi.getter(name="timeWindowInMinutes")
    def time_window_in_minutes(self) -> pulumi.Input[float]:
        """
        Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
        """
        return pulumi.get(self, "time_window_in_minutes")

    @time_window_in_minutes.setter
    def time_window_in_minutes(self, value: pulumi.Input[float]):
        pulumi.set(self, "time_window_in_minutes", value)


@pulumi.input_type
class SourceArgs:
    def __init__(__self__, *,
                 data_source_id: pulumi.Input[str],
                 authorized_resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 query_type: Optional[pulumi.Input[str]] = None):
        """
        Specifies the log search query.
        :param pulumi.Input[str] data_source_id: The resource uri over which log search query is to be run.
        :param pulumi.Input[List[pulumi.Input[str]]] authorized_resources: List of  Resource referred into query
        :param pulumi.Input[str] query: Log search query. Required for action type - AlertingAction
        :param pulumi.Input[str] query_type: Set value to 'ResultCount' .
        """
        pulumi.set(__self__, "data_source_id", data_source_id)
        if authorized_resources is not None:
            pulumi.set(__self__, "authorized_resources", authorized_resources)
        if query is not None:
            pulumi.set(__self__, "query", query)
        if query_type is not None:
            pulumi.set(__self__, "query_type", query_type)

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> pulumi.Input[str]:
        """
        The resource uri over which log search query is to be run.
        """
        return pulumi.get(self, "data_source_id")

    @data_source_id.setter
    def data_source_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_source_id", value)

    @property
    @pulumi.getter(name="authorizedResources")
    def authorized_resources(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        List of  Resource referred into query
        """
        return pulumi.get(self, "authorized_resources")

    @authorized_resources.setter
    def authorized_resources(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_resources", value)

    @property
    @pulumi.getter
    def query(self) -> Optional[pulumi.Input[str]]:
        """
        Log search query. Required for action type - AlertingAction
        """
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter(name="queryType")
    def query_type(self) -> Optional[pulumi.Input[str]]:
        """
        Set value to 'ResultCount' .
        """
        return pulumi.get(self, "query_type")

    @query_type.setter
    def query_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_type", value)


