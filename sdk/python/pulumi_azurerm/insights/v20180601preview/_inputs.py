# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'DataSourceArgs',
    'DataSourceConfigurationArgs',
    'EtwEventConfigurationArgs',
    'EtwProviderConfigurationArgs',
    'EventLogConfigurationArgs',
    'PerformanceCounterConfigurationArgs',
    'SinkConfigurationArgs',
]

@pulumi.input_type
class DataSourceArgs:
    def __init__(__self__, *,
                 configuration: pulumi.Input['DataSourceConfigurationArgs'],
                 kind: pulumi.Input[str],
                 sinks: pulumi.Input[List[pulumi.Input['SinkConfigurationArgs']]]):
        """
        Data source object contains configuration to collect telemetry and one or more sinks to send that telemetry data to
        :param pulumi.Input[str] kind: Datasource kind
        """
        pulumi.set(__self__, "configuration", configuration)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "sinks", sinks)

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Input['DataSourceConfigurationArgs']:
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: pulumi.Input['DataSourceConfigurationArgs']):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        """
        Datasource kind
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def sinks(self) -> pulumi.Input[List[pulumi.Input['SinkConfigurationArgs']]]:
        return pulumi.get(self, "sinks")

    @sinks.setter
    def sinks(self, value: pulumi.Input[List[pulumi.Input['SinkConfigurationArgs']]]):
        pulumi.set(self, "sinks", value)


@pulumi.input_type
class DataSourceConfigurationArgs:
    def __init__(__self__, *,
                 event_logs: Optional[pulumi.Input[List[pulumi.Input['EventLogConfigurationArgs']]]] = None,
                 perf_counters: Optional[pulumi.Input[List[pulumi.Input['PerformanceCounterConfigurationArgs']]]] = None,
                 providers: Optional[pulumi.Input[List[pulumi.Input['EtwProviderConfigurationArgs']]]] = None):
        """
        :param pulumi.Input[List[pulumi.Input['EventLogConfigurationArgs']]] event_logs: Windows event logs configuration.
        :param pulumi.Input[List[pulumi.Input['PerformanceCounterConfigurationArgs']]] perf_counters: Performance counter configuration
        :param pulumi.Input[List[pulumi.Input['EtwProviderConfigurationArgs']]] providers: ETW providers configuration
        """
        if event_logs is not None:
            pulumi.set(__self__, "event_logs", event_logs)
        if perf_counters is not None:
            pulumi.set(__self__, "perf_counters", perf_counters)
        if providers is not None:
            pulumi.set(__self__, "providers", providers)

    @property
    @pulumi.getter(name="eventLogs")
    def event_logs(self) -> Optional[pulumi.Input[List[pulumi.Input['EventLogConfigurationArgs']]]]:
        """
        Windows event logs configuration.
        """
        return pulumi.get(self, "event_logs")

    @event_logs.setter
    def event_logs(self, value: Optional[pulumi.Input[List[pulumi.Input['EventLogConfigurationArgs']]]]):
        pulumi.set(self, "event_logs", value)

    @property
    @pulumi.getter(name="perfCounters")
    def perf_counters(self) -> Optional[pulumi.Input[List[pulumi.Input['PerformanceCounterConfigurationArgs']]]]:
        """
        Performance counter configuration
        """
        return pulumi.get(self, "perf_counters")

    @perf_counters.setter
    def perf_counters(self, value: Optional[pulumi.Input[List[pulumi.Input['PerformanceCounterConfigurationArgs']]]]):
        pulumi.set(self, "perf_counters", value)

    @property
    @pulumi.getter
    def providers(self) -> Optional[pulumi.Input[List[pulumi.Input['EtwProviderConfigurationArgs']]]]:
        """
        ETW providers configuration
        """
        return pulumi.get(self, "providers")

    @providers.setter
    def providers(self, value: Optional[pulumi.Input[List[pulumi.Input['EtwProviderConfigurationArgs']]]]):
        pulumi.set(self, "providers", value)


@pulumi.input_type
class EtwEventConfigurationArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[float],
                 name: pulumi.Input[str],
                 filter: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[float]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[float]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)


@pulumi.input_type
class EtwProviderConfigurationArgs:
    def __init__(__self__, *,
                 events: pulumi.Input[List[pulumi.Input['EtwEventConfigurationArgs']]],
                 id: pulumi.Input[str]):
        pulumi.set(__self__, "events", events)
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def events(self) -> pulumi.Input[List[pulumi.Input['EtwEventConfigurationArgs']]]:
        return pulumi.get(self, "events")

    @events.setter
    def events(self, value: pulumi.Input[List[pulumi.Input['EtwEventConfigurationArgs']]]):
        pulumi.set(self, "events", value)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class EventLogConfigurationArgs:
    def __init__(__self__, *,
                 log_name: pulumi.Input[str],
                 filter: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "log_name", log_name)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter(name="logName")
    def log_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "log_name")

    @log_name.setter
    def log_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_name", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)


@pulumi.input_type
class PerformanceCounterConfigurationArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 sampling_period: pulumi.Input[str],
                 instance: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "sampling_period", sampling_period)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="samplingPeriod")
    def sampling_period(self) -> pulumi.Input[str]:
        return pulumi.get(self, "sampling_period")

    @sampling_period.setter
    def sampling_period(self, value: pulumi.Input[str]):
        pulumi.set(self, "sampling_period", value)

    @property
    @pulumi.getter
    def instance(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance", value)


@pulumi.input_type
class SinkConfigurationArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[str]):
        pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)


