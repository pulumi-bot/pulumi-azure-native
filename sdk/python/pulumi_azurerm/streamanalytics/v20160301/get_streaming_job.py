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
    'GetStreamingJobResult',
    'AwaitableGetStreamingJobResult',
    'get_streaming_job',
]

@pulumi.output_type
class GetStreamingJobResult:
    """
    A streaming job object, containing all information associated with the named streaming job.
    """
    def __init__(__self__, compatibility_level=None, created_date=None, data_locale=None, etag=None, events_late_arrival_max_delay_in_seconds=None, events_out_of_order_max_delay_in_seconds=None, events_out_of_order_policy=None, functions=None, inputs=None, job_id=None, job_state=None, last_output_event_time=None, location=None, name=None, output_error_policy=None, output_start_mode=None, output_start_time=None, outputs=None, provisioning_state=None, sku=None, tags=None, transformation=None, type=None):
        if compatibility_level and not isinstance(compatibility_level, str):
            raise TypeError("Expected argument 'compatibility_level' to be a str")
        pulumi.set(__self__, "compatibility_level", compatibility_level)
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if data_locale and not isinstance(data_locale, str):
            raise TypeError("Expected argument 'data_locale' to be a str")
        pulumi.set(__self__, "data_locale", data_locale)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if events_late_arrival_max_delay_in_seconds and not isinstance(events_late_arrival_max_delay_in_seconds, float):
            raise TypeError("Expected argument 'events_late_arrival_max_delay_in_seconds' to be a float")
        pulumi.set(__self__, "events_late_arrival_max_delay_in_seconds", events_late_arrival_max_delay_in_seconds)
        if events_out_of_order_max_delay_in_seconds and not isinstance(events_out_of_order_max_delay_in_seconds, float):
            raise TypeError("Expected argument 'events_out_of_order_max_delay_in_seconds' to be a float")
        pulumi.set(__self__, "events_out_of_order_max_delay_in_seconds", events_out_of_order_max_delay_in_seconds)
        if events_out_of_order_policy and not isinstance(events_out_of_order_policy, str):
            raise TypeError("Expected argument 'events_out_of_order_policy' to be a str")
        pulumi.set(__self__, "events_out_of_order_policy", events_out_of_order_policy)
        if functions and not isinstance(functions, list):
            raise TypeError("Expected argument 'functions' to be a list")
        pulumi.set(__self__, "functions", functions)
        if inputs and not isinstance(inputs, list):
            raise TypeError("Expected argument 'inputs' to be a list")
        pulumi.set(__self__, "inputs", inputs)
        if job_id and not isinstance(job_id, str):
            raise TypeError("Expected argument 'job_id' to be a str")
        pulumi.set(__self__, "job_id", job_id)
        if job_state and not isinstance(job_state, str):
            raise TypeError("Expected argument 'job_state' to be a str")
        pulumi.set(__self__, "job_state", job_state)
        if last_output_event_time and not isinstance(last_output_event_time, str):
            raise TypeError("Expected argument 'last_output_event_time' to be a str")
        pulumi.set(__self__, "last_output_event_time", last_output_event_time)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if output_error_policy and not isinstance(output_error_policy, str):
            raise TypeError("Expected argument 'output_error_policy' to be a str")
        pulumi.set(__self__, "output_error_policy", output_error_policy)
        if output_start_mode and not isinstance(output_start_mode, str):
            raise TypeError("Expected argument 'output_start_mode' to be a str")
        pulumi.set(__self__, "output_start_mode", output_start_mode)
        if output_start_time and not isinstance(output_start_time, str):
            raise TypeError("Expected argument 'output_start_time' to be a str")
        pulumi.set(__self__, "output_start_time", output_start_time)
        if outputs and not isinstance(outputs, list):
            raise TypeError("Expected argument 'outputs' to be a list")
        pulumi.set(__self__, "outputs", outputs)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if transformation and not isinstance(transformation, dict):
            raise TypeError("Expected argument 'transformation' to be a dict")
        pulumi.set(__self__, "transformation", transformation)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="compatibilityLevel")
    def compatibility_level(self) -> Optional[str]:
        """
        Controls certain runtime behaviors of the streaming job.
        """
        return pulumi.get(self, "compatibility_level")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter(name="dataLocale")
    def data_locale(self) -> Optional[str]:
        """
        The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
        """
        return pulumi.get(self, "data_locale")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="eventsLateArrivalMaxDelayInSeconds")
    def events_late_arrival_max_delay_in_seconds(self) -> Optional[float]:
        """
        The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
        """
        return pulumi.get(self, "events_late_arrival_max_delay_in_seconds")

    @property
    @pulumi.getter(name="eventsOutOfOrderMaxDelayInSeconds")
    def events_out_of_order_max_delay_in_seconds(self) -> Optional[float]:
        """
        The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
        """
        return pulumi.get(self, "events_out_of_order_max_delay_in_seconds")

    @property
    @pulumi.getter(name="eventsOutOfOrderPolicy")
    def events_out_of_order_policy(self) -> Optional[str]:
        """
        Indicates the policy to apply to events that arrive out of order in the input event stream.
        """
        return pulumi.get(self, "events_out_of_order_policy")

    @property
    @pulumi.getter
    def functions(self) -> Optional[List['outputs.FunctionResponse']]:
        """
        A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
        """
        return pulumi.get(self, "functions")

    @property
    @pulumi.getter
    def inputs(self) -> Optional[List['outputs.InputResponse']]:
        """
        A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
        """
        return pulumi.get(self, "inputs")

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> str:
        """
        A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
        """
        return pulumi.get(self, "job_id")

    @property
    @pulumi.getter(name="jobState")
    def job_state(self) -> str:
        """
        Describes the state of the streaming job.
        """
        return pulumi.get(self, "job_state")

    @property
    @pulumi.getter(name="lastOutputEventTime")
    def last_output_event_time(self) -> str:
        """
        Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
        """
        return pulumi.get(self, "last_output_event_time")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location. Required on PUT (CreateOrReplace) requests.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputErrorPolicy")
    def output_error_policy(self) -> Optional[str]:
        """
        Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
        """
        return pulumi.get(self, "output_error_policy")

    @property
    @pulumi.getter(name="outputStartMode")
    def output_start_mode(self) -> Optional[str]:
        """
        This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
        """
        return pulumi.get(self, "output_start_mode")

    @property
    @pulumi.getter(name="outputStartTime")
    def output_start_time(self) -> Optional[str]:
        """
        Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
        """
        return pulumi.get(self, "output_start_time")

    @property
    @pulumi.getter
    def outputs(self) -> Optional[List['outputs.OutputResponse']]:
        """
        A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Describes the provisioning status of the streaming job.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def transformation(self) -> Optional['outputs.TransformationResponse']:
        """
        Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
        """
        return pulumi.get(self, "transformation")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetStreamingJobResult(GetStreamingJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStreamingJobResult(
            compatibility_level=self.compatibility_level,
            created_date=self.created_date,
            data_locale=self.data_locale,
            etag=self.etag,
            events_late_arrival_max_delay_in_seconds=self.events_late_arrival_max_delay_in_seconds,
            events_out_of_order_max_delay_in_seconds=self.events_out_of_order_max_delay_in_seconds,
            events_out_of_order_policy=self.events_out_of_order_policy,
            functions=self.functions,
            inputs=self.inputs,
            job_id=self.job_id,
            job_state=self.job_state,
            last_output_event_time=self.last_output_event_time,
            location=self.location,
            name=self.name,
            output_error_policy=self.output_error_policy,
            output_start_mode=self.output_start_mode,
            output_start_time=self.output_start_time,
            outputs=self.outputs,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            tags=self.tags,
            transformation=self.transformation,
            type=self.type)


def get_streaming_job(expand: Optional[str] = None,
                      name: Optional[str] = None,
                      resource_group_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStreamingJobResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: The $expand OData query parameter. This is a comma-separated list of additional streaming job properties to include in the response, beyond the default set returned when this parameter is absent. The default set is all streaming job properties other than 'inputs', 'transformation', 'outputs', and 'functions'.
    :param str name: The name of the streaming job.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:streamanalytics/v20160301:getStreamingJob', __args__, opts=opts, typ=GetStreamingJobResult).value

    return AwaitableGetStreamingJobResult(
        compatibility_level=__ret__.compatibility_level,
        created_date=__ret__.created_date,
        data_locale=__ret__.data_locale,
        etag=__ret__.etag,
        events_late_arrival_max_delay_in_seconds=__ret__.events_late_arrival_max_delay_in_seconds,
        events_out_of_order_max_delay_in_seconds=__ret__.events_out_of_order_max_delay_in_seconds,
        events_out_of_order_policy=__ret__.events_out_of_order_policy,
        functions=__ret__.functions,
        inputs=__ret__.inputs,
        job_id=__ret__.job_id,
        job_state=__ret__.job_state,
        last_output_event_time=__ret__.last_output_event_time,
        location=__ret__.location,
        name=__ret__.name,
        output_error_policy=__ret__.output_error_policy,
        output_start_mode=__ret__.output_start_mode,
        output_start_time=__ret__.output_start_time,
        outputs=__ret__.outputs,
        provisioning_state=__ret__.provisioning_state,
        sku=__ret__.sku,
        tags=__ret__.tags,
        transformation=__ret__.transformation,
        type=__ret__.type)
