# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'DataSourceConfigurationResponse',
    'DataSourceResponse',
    'EtwEventConfigurationResponse',
    'EtwProviderConfigurationResponse',
    'EventLogConfigurationResponse',
    'PerformanceCounterConfigurationResponse',
    'SinkConfigurationResponse',
]

@pulumi.output_type
class DataSourceConfigurationResponse(dict):
    def __init__(__self__, *,
                 event_logs: Optional[List['outputs.EventLogConfigurationResponse']] = None,
                 perf_counters: Optional[List['outputs.PerformanceCounterConfigurationResponse']] = None,
                 providers: Optional[List['outputs.EtwProviderConfigurationResponse']] = None):
        """
        :param List['EventLogConfigurationResponseArgs'] event_logs: Windows event logs configuration.
        :param List['PerformanceCounterConfigurationResponseArgs'] perf_counters: Performance counter configuration
        :param List['EtwProviderConfigurationResponseArgs'] providers: ETW providers configuration
        """
        if event_logs is not None:
            pulumi.set(__self__, "event_logs", event_logs)
        if perf_counters is not None:
            pulumi.set(__self__, "perf_counters", perf_counters)
        if providers is not None:
            pulumi.set(__self__, "providers", providers)

    @property
    @pulumi.getter(name="eventLogs")
    def event_logs(self) -> Optional[List['outputs.EventLogConfigurationResponse']]:
        """
        Windows event logs configuration.
        """
        return pulumi.get(self, "event_logs")

    @property
    @pulumi.getter(name="perfCounters")
    def perf_counters(self) -> Optional[List['outputs.PerformanceCounterConfigurationResponse']]:
        """
        Performance counter configuration
        """
        return pulumi.get(self, "perf_counters")

    @property
    @pulumi.getter
    def providers(self) -> Optional[List['outputs.EtwProviderConfigurationResponse']]:
        """
        ETW providers configuration
        """
        return pulumi.get(self, "providers")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DataSourceResponse(dict):
    """
    Data source object contains configuration to collect telemetry and one or more sinks to send that telemetry data to
    """
    def __init__(__self__, *,
                 configuration: 'outputs.DataSourceConfigurationResponse',
                 kind: str,
                 sinks: List['outputs.SinkConfigurationResponse']):
        """
        Data source object contains configuration to collect telemetry and one or more sinks to send that telemetry data to
        :param str kind: Datasource kind
        """
        pulumi.set(__self__, "configuration", configuration)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "sinks", sinks)

    @property
    @pulumi.getter
    def configuration(self) -> 'outputs.DataSourceConfigurationResponse':
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Datasource kind
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def sinks(self) -> List['outputs.SinkConfigurationResponse']:
        return pulumi.get(self, "sinks")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EtwEventConfigurationResponse(dict):
    def __init__(__self__, *,
                 id: float,
                 name: str,
                 filter: Optional[str] = None):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def id(self) -> float:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EtwProviderConfigurationResponse(dict):
    def __init__(__self__, *,
                 events: List['outputs.EtwEventConfigurationResponse'],
                 id: str):
        pulumi.set(__self__, "events", events)
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def events(self) -> List['outputs.EtwEventConfigurationResponse']:
        return pulumi.get(self, "events")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EventLogConfigurationResponse(dict):
    def __init__(__self__, *,
                 log_name: str,
                 filter: Optional[str] = None):
        pulumi.set(__self__, "log_name", log_name)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter(name="logName")
    def log_name(self) -> str:
        return pulumi.get(self, "log_name")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PerformanceCounterConfigurationResponse(dict):
    def __init__(__self__, *,
                 name: str,
                 sampling_period: str,
                 instance: Optional[str] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "sampling_period", sampling_period)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="samplingPeriod")
    def sampling_period(self) -> str:
        return pulumi.get(self, "sampling_period")

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        return pulumi.get(self, "instance")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SinkConfigurationResponse(dict):
    def __init__(__self__, *,
                 kind: str):
        pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


