# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'AdvancedFilterArgs',
    'DeadLetterDestinationArgs',
    'EventSubscriptionDestinationArgs',
    'EventSubscriptionFilterArgs',
    'RetryPolicyArgs',
]

@pulumi.input_type
class AdvancedFilterArgs:
    def __init__(__self__, *,
                 operator_type: pulumi.Input[str],
                 key: Optional[pulumi.Input[str]] = None):
        """
        This is the base type that represents an advanced filter. To configure an advanced filter, do not directly instantiate an object of this class. Instead, instantiate an object of a derived class such as BoolEqualsAdvancedFilter, NumberInAdvancedFilter, StringEqualsAdvancedFilter etc. depending on the type of the key based on which you want to filter.
        :param pulumi.Input[str] operator_type: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
        :param pulumi.Input[str] key: The field/property in the event based on which you want to filter.
        """
        pulumi.set(__self__, "operator_type", operator_type)
        if key is not None:
            pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter(name="operatorType")
    def operator_type(self) -> pulumi.Input[str]:
        """
        The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
        """
        return pulumi.get(self, "operator_type")

    @operator_type.setter
    def operator_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "operator_type", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The field/property in the event based on which you want to filter.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)


@pulumi.input_type
class DeadLetterDestinationArgs:
    def __init__(__self__, *,
                 endpoint_type: pulumi.Input[str]):
        """
        Information about the dead letter destination for an event subscription. To configure a deadletter destination, do not directly instantiate an object of this class. Instead, instantiate an object of a derived class. Currently, StorageBlobDeadLetterDestination is the only class that derives from this class.
        :param pulumi.Input[str] endpoint_type: Type of the endpoint for the dead letter destination
        """
        pulumi.set(__self__, "endpoint_type", endpoint_type)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> pulumi.Input[str]:
        """
        Type of the endpoint for the dead letter destination
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_type", value)


@pulumi.input_type
class EventSubscriptionDestinationArgs:
    def __init__(__self__, *,
                 endpoint_type: pulumi.Input[str]):
        """
        Information about the destination for an event subscription
        :param pulumi.Input[str] endpoint_type: Type of the endpoint for the event subscription destination
        """
        pulumi.set(__self__, "endpoint_type", endpoint_type)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> pulumi.Input[str]:
        """
        Type of the endpoint for the event subscription destination
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_type", value)


@pulumi.input_type
class EventSubscriptionFilterArgs:
    def __init__(__self__, *,
                 advanced_filters: Optional[pulumi.Input[List[pulumi.Input['AdvancedFilterArgs']]]] = None,
                 included_event_types: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 is_subject_case_sensitive: Optional[pulumi.Input[bool]] = None,
                 subject_begins_with: Optional[pulumi.Input[str]] = None,
                 subject_ends_with: Optional[pulumi.Input[str]] = None):
        """
        Filter for the Event Subscription.
        :param pulumi.Input[List[pulumi.Input['AdvancedFilterArgs']]] advanced_filters: An array of advanced filters that are used for filtering event subscriptions.
        :param pulumi.Input[List[pulumi.Input[str]]] included_event_types: A list of applicable event types that need to be part of the event subscription. If it is desired to subscribe to all default event types, set the IncludedEventTypes to null.
        :param pulumi.Input[bool] is_subject_case_sensitive: Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter 
               should be compared in a case sensitive manner.
        :param pulumi.Input[str] subject_begins_with: An optional string to filter events for an event subscription based on a resource path prefix.
               The format of this depends on the publisher of the events. 
               Wildcard characters are not supported in this path.
        :param pulumi.Input[str] subject_ends_with: An optional string to filter events for an event subscription based on a resource path suffix.
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
    def advanced_filters(self) -> Optional[pulumi.Input[List[pulumi.Input['AdvancedFilterArgs']]]]:
        """
        An array of advanced filters that are used for filtering event subscriptions.
        """
        return pulumi.get(self, "advanced_filters")

    @advanced_filters.setter
    def advanced_filters(self, value: Optional[pulumi.Input[List[pulumi.Input['AdvancedFilterArgs']]]]):
        pulumi.set(self, "advanced_filters", value)

    @property
    @pulumi.getter(name="includedEventTypes")
    def included_event_types(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        A list of applicable event types that need to be part of the event subscription. If it is desired to subscribe to all default event types, set the IncludedEventTypes to null.
        """
        return pulumi.get(self, "included_event_types")

    @included_event_types.setter
    def included_event_types(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "included_event_types", value)

    @property
    @pulumi.getter(name="isSubjectCaseSensitive")
    def is_subject_case_sensitive(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter 
        should be compared in a case sensitive manner.
        """
        return pulumi.get(self, "is_subject_case_sensitive")

    @is_subject_case_sensitive.setter
    def is_subject_case_sensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_subject_case_sensitive", value)

    @property
    @pulumi.getter(name="subjectBeginsWith")
    def subject_begins_with(self) -> Optional[pulumi.Input[str]]:
        """
        An optional string to filter events for an event subscription based on a resource path prefix.
        The format of this depends on the publisher of the events. 
        Wildcard characters are not supported in this path.
        """
        return pulumi.get(self, "subject_begins_with")

    @subject_begins_with.setter
    def subject_begins_with(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject_begins_with", value)

    @property
    @pulumi.getter(name="subjectEndsWith")
    def subject_ends_with(self) -> Optional[pulumi.Input[str]]:
        """
        An optional string to filter events for an event subscription based on a resource path suffix.
        Wildcard characters are not supported in this path.
        """
        return pulumi.get(self, "subject_ends_with")

    @subject_ends_with.setter
    def subject_ends_with(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject_ends_with", value)


@pulumi.input_type
class RetryPolicyArgs:
    def __init__(__self__, *,
                 event_time_to_live_in_minutes: Optional[pulumi.Input[float]] = None,
                 max_delivery_attempts: Optional[pulumi.Input[float]] = None):
        """
        Information about the retry policy for an event subscription.
        :param pulumi.Input[float] event_time_to_live_in_minutes: Time To Live (in minutes) for events.
        :param pulumi.Input[float] max_delivery_attempts: Maximum number of delivery retry attempts for events.
        """
        if event_time_to_live_in_minutes is not None:
            pulumi.set(__self__, "event_time_to_live_in_minutes", event_time_to_live_in_minutes)
        if max_delivery_attempts is not None:
            pulumi.set(__self__, "max_delivery_attempts", max_delivery_attempts)

    @property
    @pulumi.getter(name="eventTimeToLiveInMinutes")
    def event_time_to_live_in_minutes(self) -> Optional[pulumi.Input[float]]:
        """
        Time To Live (in minutes) for events.
        """
        return pulumi.get(self, "event_time_to_live_in_minutes")

    @event_time_to_live_in_minutes.setter
    def event_time_to_live_in_minutes(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "event_time_to_live_in_minutes", value)

    @property
    @pulumi.getter(name="maxDeliveryAttempts")
    def max_delivery_attempts(self) -> Optional[pulumi.Input[float]]:
        """
        Maximum number of delivery retry attempts for events.
        """
        return pulumi.get(self, "max_delivery_attempts")

    @max_delivery_attempts.setter
    def max_delivery_attempts(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_delivery_attempts", value)


