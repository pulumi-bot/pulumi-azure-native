# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ActionResponse',
    'ScheduleResponse',
    'SourceResponse',
]

@pulumi.output_type
class ActionResponse(dict):
    """
    Action descriptor.
    """
    def __init__(__self__, *,
                 odata_type: str):
        """
        Action descriptor.
        :param str odata_type: Specifies the action. Supported values - AlertingAction, LogToMetricAction
        """
        pulumi.set(__self__, "odata_type", odata_type)

    @property
    @pulumi.getter(name="odataType")
    def odata_type(self) -> str:
        """
        Specifies the action. Supported values - AlertingAction, LogToMetricAction
        """
        return pulumi.get(self, "odata_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleResponse(dict):
    """
    Defines how often to run the search and the time interval.
    """
    def __init__(__self__, *,
                 frequency_in_minutes: float,
                 time_window_in_minutes: float):
        """
        Defines how often to run the search and the time interval.
        :param float frequency_in_minutes: frequency (in minutes) at which rule condition should be evaluated.
        :param float time_window_in_minutes: Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
        """
        pulumi.set(__self__, "frequency_in_minutes", frequency_in_minutes)
        pulumi.set(__self__, "time_window_in_minutes", time_window_in_minutes)

    @property
    @pulumi.getter(name="frequencyInMinutes")
    def frequency_in_minutes(self) -> float:
        """
        frequency (in minutes) at which rule condition should be evaluated.
        """
        return pulumi.get(self, "frequency_in_minutes")

    @property
    @pulumi.getter(name="timeWindowInMinutes")
    def time_window_in_minutes(self) -> float:
        """
        Time window for which data needs to be fetched for query (should be greater than or equal to frequencyInMinutes).
        """
        return pulumi.get(self, "time_window_in_minutes")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SourceResponse(dict):
    """
    Specifies the log search query.
    """
    def __init__(__self__, *,
                 data_source_id: str,
                 authorized_resources: Optional[List[str]] = None,
                 query: Optional[str] = None,
                 query_type: Optional[str] = None):
        """
        Specifies the log search query.
        :param str data_source_id: The resource uri over which log search query is to be run.
        :param List[str] authorized_resources: List of  Resource referred into query
        :param str query: Log search query. Required for action type - AlertingAction
        :param str query_type: Set value to 'ResultCount' .
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
    def data_source_id(self) -> str:
        """
        The resource uri over which log search query is to be run.
        """
        return pulumi.get(self, "data_source_id")

    @property
    @pulumi.getter(name="authorizedResources")
    def authorized_resources(self) -> Optional[List[str]]:
        """
        List of  Resource referred into query
        """
        return pulumi.get(self, "authorized_resources")

    @property
    @pulumi.getter
    def query(self) -> Optional[str]:
        """
        Log search query. Required for action type - AlertingAction
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="queryType")
    def query_type(self) -> Optional[str]:
        """
        Set value to 'ResultCount' .
        """
        return pulumi.get(self, "query_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


