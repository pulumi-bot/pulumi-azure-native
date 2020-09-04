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
    'AdvancedFilterResponse',
    'DeadLetterDestinationResponse',
    'EventSubscriptionDestinationResponse',
    'EventSubscriptionFilterResponse',
    'InputSchemaMappingResponse',
    'RetryPolicyResponse',
]

@pulumi.output_type
class AdvancedFilterResponse(dict):
    """
    This is the base type that represents an advanced filter. To configure an advanced filter, do not directly instantiate an object of this class. Instead, instantiate an object of a derived class such as BoolEqualsAdvancedFilter, NumberInAdvancedFilter, StringEqualsAdvancedFilter etc. depending on the type of the key based on which you want to filter.
    """
    def __init__(__self__, *,
                 operator_type: str,
                 key: Optional[str] = None):
        """
        This is the base type that represents an advanced filter. To configure an advanced filter, do not directly instantiate an object of this class. Instead, instantiate an object of a derived class such as BoolEqualsAdvancedFilter, NumberInAdvancedFilter, StringEqualsAdvancedFilter etc. depending on the type of the key based on which you want to filter.
        :param str operator_type: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
        :param str key: The field/property in the event based on which you want to filter.
        """
        pulumi.set(__self__, "operator_type", operator_type)
        if key is not None:
            pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter(name="operatorType")
    def operator_type(self) -> str:
        """
        The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
        """
        return pulumi.get(self, "operator_type")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        """
        The field/property in the event based on which you want to filter.
        """
        return pulumi.get(self, "key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DeadLetterDestinationResponse(dict):
    """
    Information about the dead letter destination for an event subscription. To configure a deadletter destination, do not directly instantiate an object of this class. Instead, instantiate an object of a derived class. Currently, StorageBlobDeadLetterDestination is the only class that derives from this class.
    """
    def __init__(__self__, *,
                 endpoint_type: str):
        """
        Information about the dead letter destination for an event subscription. To configure a deadletter destination, do not directly instantiate an object of this class. Instead, instantiate an object of a derived class. Currently, StorageBlobDeadLetterDestination is the only class that derives from this class.
        :param str endpoint_type: Type of the endpoint for the dead letter destination
        """
        pulumi.set(__self__, "endpoint_type", endpoint_type)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> str:
        """
        Type of the endpoint for the dead letter destination
        """
        return pulumi.get(self, "endpoint_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EventSubscriptionDestinationResponse(dict):
    """
    Information about the destination for an event subscription
    """
    def __init__(__self__, *,
                 endpoint_type: str):
        """
        Information about the destination for an event subscription
        :param str endpoint_type: Type of the endpoint for the event subscription destination
        """
        pulumi.set(__self__, "endpoint_type", endpoint_type)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> str:
        """
        Type of the endpoint for the event subscription destination
        """
        return pulumi.get(self, "endpoint_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EventSubscriptionFilterResponse(dict):
    """
    Filter for the Event Subscription
    """
    def __init__(__self__, *,
                 advanced_filters: Optional[List['outputs.AdvancedFilterResponse']] = None,
                 included_event_types: Optional[List[str]] = None,
                 is_subject_case_sensitive: Optional[bool] = None,
                 subject_begins_with: Optional[str] = None,
                 subject_ends_with: Optional[str] = None):
        """
        Filter for the Event Subscription
        :param List['AdvancedFilterResponseArgs'] advanced_filters: An array of advanced filters that are used for filtering event subscriptions.
        :param List[str] included_event_types: A list of applicable event types that need to be part of the event subscription. If it is desired to subscribe to all default event types, set the IncludedEventTypes to null.
        :param bool is_subject_case_sensitive: Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter 
               should be compared in a case sensitive manner.
        :param str subject_begins_with: An optional string to filter events for an event subscription based on a resource path prefix.
               The format of this depends on the publisher of the events. 
               Wildcard characters are not supported in this path.
        :param str subject_ends_with: An optional string to filter events for an event subscription based on a resource path suffix.
               Wildcard characters are not supported in this path.
        """
        if advanced_filters is not None:
            pulumi.set(__self__, "advanced_filters", advanced_filters)
        if included_event_types is not None:
            pulumi.set(__self__, "included_event_types", included_event_types)
        if is_subject_case_sensitive is not None:
            pulumi.set(__self__, "is_subject_case_sensitive", is_subject_case_sensitive)
        if subject_begins_with is not None:
            pulumi.set(__self__, "subject_begins_with", subject_begins_with)
        if subject_ends_with is not None:
            pulumi.set(__self__, "subject_ends_with", subject_ends_with)

    @property
    @pulumi.getter(name="advancedFilters")
    def advanced_filters(self) -> Optional[List['outputs.AdvancedFilterResponse']]:
        """
        An array of advanced filters that are used for filtering event subscriptions.
        """
        return pulumi.get(self, "advanced_filters")

    @property
    @pulumi.getter(name="includedEventTypes")
    def included_event_types(self) -> Optional[List[str]]:
        """
        A list of applicable event types that need to be part of the event subscription. If it is desired to subscribe to all default event types, set the IncludedEventTypes to null.
        """
        return pulumi.get(self, "included_event_types")

    @property
    @pulumi.getter(name="isSubjectCaseSensitive")
    def is_subject_case_sensitive(self) -> Optional[bool]:
        """
        Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter 
        should be compared in a case sensitive manner.
        """
        return pulumi.get(self, "is_subject_case_sensitive")

    @property
    @pulumi.getter(name="subjectBeginsWith")
    def subject_begins_with(self) -> Optional[str]:
        """
        An optional string to filter events for an event subscription based on a resource path prefix.
        The format of this depends on the publisher of the events. 
        Wildcard characters are not supported in this path.
        """
        return pulumi.get(self, "subject_begins_with")

    @property
    @pulumi.getter(name="subjectEndsWith")
    def subject_ends_with(self) -> Optional[str]:
        """
        An optional string to filter events for an event subscription based on a resource path suffix.
        Wildcard characters are not supported in this path.
        """
        return pulumi.get(self, "subject_ends_with")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InputSchemaMappingResponse(dict):
    """
    By default, Event Grid expects events to be in the Event Grid event schema. Specifying an input schema mapping enables publishing to Event Grid using a custom input schema. Currently, the only supported type of InputSchemaMapping is 'JsonInputSchemaMapping'.
    """
    def __init__(__self__, *,
                 input_schema_mapping_type: Optional[str] = None):
        """
        By default, Event Grid expects events to be in the Event Grid event schema. Specifying an input schema mapping enables publishing to Event Grid using a custom input schema. Currently, the only supported type of InputSchemaMapping is 'JsonInputSchemaMapping'.
        :param str input_schema_mapping_type: Type of the custom mapping
        """
        if input_schema_mapping_type is not None:
            pulumi.set(__self__, "input_schema_mapping_type", input_schema_mapping_type)

    @property
    @pulumi.getter(name="inputSchemaMappingType")
    def input_schema_mapping_type(self) -> Optional[str]:
        """
        Type of the custom mapping
        """
        return pulumi.get(self, "input_schema_mapping_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RetryPolicyResponse(dict):
    """
    Information about the retry policy for an event subscription
    """
    def __init__(__self__, *,
                 event_time_to_live_in_minutes: Optional[float] = None,
                 max_delivery_attempts: Optional[float] = None):
        """
        Information about the retry policy for an event subscription
        :param float event_time_to_live_in_minutes: Time To Live (in minutes) for events.
        :param float max_delivery_attempts: Maximum number of delivery retry attempts for events.
        """
        if event_time_to_live_in_minutes is not None:
            pulumi.set(__self__, "event_time_to_live_in_minutes", event_time_to_live_in_minutes)
        if max_delivery_attempts is not None:
            pulumi.set(__self__, "max_delivery_attempts", max_delivery_attempts)

    @property
    @pulumi.getter(name="eventTimeToLiveInMinutes")
    def event_time_to_live_in_minutes(self) -> Optional[float]:
        """
        Time To Live (in minutes) for events.
        """
        return pulumi.get(self, "event_time_to_live_in_minutes")

    @property
    @pulumi.getter(name="maxDeliveryAttempts")
    def max_delivery_attempts(self) -> Optional[float]:
        """
        Maximum number of delivery retry attempts for events.
        """
        return pulumi.get(self, "max_delivery_attempts")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


