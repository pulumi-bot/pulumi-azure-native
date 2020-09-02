# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'DatabaseVulnerabilityAssessmentRuleBaselineItemResponse',
    'JobScheduleResponse',
    'JobStepActionResponse',
    'JobStepExecutionOptionsResponse',
    'JobStepOutputResponse',
    'JobTargetResponse',
    'SkuResponse',
    'VulnerabilityAssessmentRecurringScansPropertiesResponse',
]

@pulumi.output_type
class DatabaseVulnerabilityAssessmentRuleBaselineItemResponse(dict):
    """
    Properties for an Azure SQL Database Vulnerability Assessment rule baseline's result.
    """
    def __init__(__self__, *,
                 result: List[str]):
        """
        Properties for an Azure SQL Database Vulnerability Assessment rule baseline's result.
        :param List[str] result: The rule baseline result
        """
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> List[str]:
        """
        The rule baseline result
        """
        return pulumi.get(self, "result")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobScheduleResponse(dict):
    """
    Scheduling properties of a job.
    """
    def __init__(__self__, *,
                 enabled: Optional[bool] = None,
                 end_time: Optional[str] = None,
                 interval: Optional[str] = None,
                 start_time: Optional[str] = None,
                 type: Optional[str] = None):
        """
        Scheduling properties of a job.
        :param bool enabled: Whether or not the schedule is enabled.
        :param str end_time: Schedule end time.
        :param str interval: Value of the schedule's recurring interval, if the schedule type is recurring. ISO8601 duration format.
        :param str start_time: Schedule start time.
        :param str type: Schedule interval type
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Whether or not the schedule is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[str]:
        """
        Schedule end time.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def interval(self) -> Optional[str]:
        """
        Value of the schedule's recurring interval, if the schedule type is recurring. ISO8601 duration format.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[str]:
        """
        Schedule start time.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Schedule interval type
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobStepActionResponse(dict):
    """
    The action to be executed by a job step.
    """
    def __init__(__self__, *,
                 value: str,
                 source: Optional[str] = None,
                 type: Optional[str] = None):
        """
        The action to be executed by a job step.
        :param str value: The action value, for example the text of the T-SQL script to execute.
        :param str source: The source of the action to execute.
        :param str type: Type of action being executed by the job step.
        """
        pulumi.set(__self__, "value", value)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The action value, for example the text of the T-SQL script to execute.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def source(self) -> Optional[str]:
        """
        The source of the action to execute.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Type of action being executed by the job step.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobStepExecutionOptionsResponse(dict):
    """
    The execution options of a job step.
    """
    def __init__(__self__, *,
                 initial_retry_interval_seconds: Optional[float] = None,
                 maximum_retry_interval_seconds: Optional[float] = None,
                 retry_attempts: Optional[float] = None,
                 retry_interval_backoff_multiplier: Optional[float] = None,
                 timeout_seconds: Optional[float] = None):
        """
        The execution options of a job step.
        :param float initial_retry_interval_seconds: Initial delay between retries for job step execution.
        :param float maximum_retry_interval_seconds: The maximum amount of time to wait between retries for job step execution.
        :param float retry_attempts: Maximum number of times the job step will be reattempted if the first attempt fails.
        :param float retry_interval_backoff_multiplier: The backoff multiplier for the time between retries.
        :param float timeout_seconds: Execution timeout for the job step.
        """
        if initial_retry_interval_seconds is not None:
            pulumi.set(__self__, "initial_retry_interval_seconds", initial_retry_interval_seconds)
        if maximum_retry_interval_seconds is not None:
            pulumi.set(__self__, "maximum_retry_interval_seconds", maximum_retry_interval_seconds)
        if retry_attempts is not None:
            pulumi.set(__self__, "retry_attempts", retry_attempts)
        if retry_interval_backoff_multiplier is not None:
            pulumi.set(__self__, "retry_interval_backoff_multiplier", retry_interval_backoff_multiplier)
        if timeout_seconds is not None:
            pulumi.set(__self__, "timeout_seconds", timeout_seconds)

    @property
    @pulumi.getter(name="initialRetryIntervalSeconds")
    def initial_retry_interval_seconds(self) -> Optional[float]:
        """
        Initial delay between retries for job step execution.
        """
        return pulumi.get(self, "initial_retry_interval_seconds")

    @property
    @pulumi.getter(name="maximumRetryIntervalSeconds")
    def maximum_retry_interval_seconds(self) -> Optional[float]:
        """
        The maximum amount of time to wait between retries for job step execution.
        """
        return pulumi.get(self, "maximum_retry_interval_seconds")

    @property
    @pulumi.getter(name="retryAttempts")
    def retry_attempts(self) -> Optional[float]:
        """
        Maximum number of times the job step will be reattempted if the first attempt fails.
        """
        return pulumi.get(self, "retry_attempts")

    @property
    @pulumi.getter(name="retryIntervalBackoffMultiplier")
    def retry_interval_backoff_multiplier(self) -> Optional[float]:
        """
        The backoff multiplier for the time between retries.
        """
        return pulumi.get(self, "retry_interval_backoff_multiplier")

    @property
    @pulumi.getter(name="timeoutSeconds")
    def timeout_seconds(self) -> Optional[float]:
        """
        Execution timeout for the job step.
        """
        return pulumi.get(self, "timeout_seconds")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobStepOutputResponse(dict):
    """
    The output configuration of a job step.
    """
    def __init__(__self__, *,
                 credential: str,
                 database_name: str,
                 server_name: str,
                 table_name: str,
                 resource_group_name: Optional[str] = None,
                 schema_name: Optional[str] = None,
                 subscription_id: Optional[str] = None,
                 type: Optional[str] = None):
        """
        The output configuration of a job step.
        :param str credential: The resource ID of the credential to use to connect to the output destination.
        :param str database_name: The output destination database.
        :param str server_name: The output destination server name.
        :param str table_name: The output destination table.
        :param str resource_group_name: The output destination resource group.
        :param str schema_name: The output destination schema.
        :param str subscription_id: The output destination subscription id.
        :param str type: The output destination type.
        """
        pulumi.set(__self__, "credential", credential)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "server_name", server_name)
        pulumi.set(__self__, "table_name", table_name)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if schema_name is not None:
            pulumi.set(__self__, "schema_name", schema_name)
        if subscription_id is not None:
            pulumi.set(__self__, "subscription_id", subscription_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def credential(self) -> str:
        """
        The resource ID of the credential to use to connect to the output destination.
        """
        return pulumi.get(self, "credential")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        """
        The output destination database.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> str:
        """
        The output destination server name.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> str:
        """
        The output destination table.
        """
        return pulumi.get(self, "table_name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[str]:
        """
        The output destination resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="schemaName")
    def schema_name(self) -> Optional[str]:
        """
        The output destination schema.
        """
        return pulumi.get(self, "schema_name")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> Optional[str]:
        """
        The output destination subscription id.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The output destination type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobTargetResponse(dict):
    """
    A job target, for example a specific database or a container of databases that is evaluated during job execution.
    """
    def __init__(__self__, *,
                 type: str,
                 database_name: Optional[str] = None,
                 elastic_pool_name: Optional[str] = None,
                 membership_type: Optional[str] = None,
                 refresh_credential: Optional[str] = None,
                 server_name: Optional[str] = None,
                 shard_map_name: Optional[str] = None):
        """
        A job target, for example a specific database or a container of databases that is evaluated during job execution.
        :param str type: The target type.
        :param str database_name: The target database name.
        :param str elastic_pool_name: The target elastic pool name.
        :param str membership_type: Whether the target is included or excluded from the group.
        :param str refresh_credential: The resource ID of the credential that is used during job execution to connect to the target and determine the list of databases inside the target.
        :param str server_name: The target server name.
        :param str shard_map_name: The target shard map.
        """
        pulumi.set(__self__, "type", type)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if elastic_pool_name is not None:
            pulumi.set(__self__, "elastic_pool_name", elastic_pool_name)
        if membership_type is not None:
            pulumi.set(__self__, "membership_type", membership_type)
        if refresh_credential is not None:
            pulumi.set(__self__, "refresh_credential", refresh_credential)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if shard_map_name is not None:
            pulumi.set(__self__, "shard_map_name", shard_map_name)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The target type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[str]:
        """
        The target database name.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="elasticPoolName")
    def elastic_pool_name(self) -> Optional[str]:
        """
        The target elastic pool name.
        """
        return pulumi.get(self, "elastic_pool_name")

    @property
    @pulumi.getter(name="membershipType")
    def membership_type(self) -> Optional[str]:
        """
        Whether the target is included or excluded from the group.
        """
        return pulumi.get(self, "membership_type")

    @property
    @pulumi.getter(name="refreshCredential")
    def refresh_credential(self) -> Optional[str]:
        """
        The resource ID of the credential that is used during job execution to connect to the target and determine the list of databases inside the target.
        """
        return pulumi.get(self, "refresh_credential")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[str]:
        """
        The target server name.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="shardMapName")
    def shard_map_name(self) -> Optional[str]:
        """
        The target shard map.
        """
        return pulumi.get(self, "shard_map_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    An ARM Resource SKU.
    """
    def __init__(__self__, *,
                 name: str,
                 capacity: Optional[float] = None,
                 family: Optional[str] = None,
                 size: Optional[str] = None,
                 tier: Optional[str] = None):
        """
        An ARM Resource SKU.
        :param str name: The name of the SKU, typically, a letter + Number code, e.g. P3.
        :param float capacity: Capacity of the particular SKU.
        :param str family: If the service has different generations of hardware, for the same SKU, then that can be captured here.
        :param str size: Size of the particular SKU
        :param str tier: The tier or edition of the particular SKU, e.g. Basic, Premium.
        """
        pulumi.set(__self__, "name", name)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if family is not None:
            pulumi.set(__self__, "family", family)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the SKU, typically, a letter + Number code, e.g. P3.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def capacity(self) -> Optional[float]:
        """
        Capacity of the particular SKU.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def family(self) -> Optional[str]:
        """
        If the service has different generations of hardware, for the same SKU, then that can be captured here.
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter
    def size(self) -> Optional[str]:
        """
        Size of the particular SKU
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def tier(self) -> Optional[str]:
        """
        The tier or edition of the particular SKU, e.g. Basic, Premium.
        """
        return pulumi.get(self, "tier")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VulnerabilityAssessmentRecurringScansPropertiesResponse(dict):
    """
    Properties of a Vulnerability Assessment recurring scans.
    """
    def __init__(__self__, *,
                 email_subscription_admins: Optional[bool] = None,
                 emails: Optional[List[str]] = None,
                 is_enabled: Optional[bool] = None):
        """
        Properties of a Vulnerability Assessment recurring scans.
        :param bool email_subscription_admins: Specifies that the schedule scan notification will be is sent to the subscription administrators.
        :param List[str] emails: Specifies an array of e-mail addresses to which the scan notification is sent.
        :param bool is_enabled: Recurring scans state.
        """
        if email_subscription_admins is not None:
            pulumi.set(__self__, "email_subscription_admins", email_subscription_admins)
        if emails is not None:
            pulumi.set(__self__, "emails", emails)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)

    @property
    @pulumi.getter(name="emailSubscriptionAdmins")
    def email_subscription_admins(self) -> Optional[bool]:
        """
        Specifies that the schedule scan notification will be is sent to the subscription administrators.
        """
        return pulumi.get(self, "email_subscription_admins")

    @property
    @pulumi.getter
    def emails(self) -> Optional[List[str]]:
        """
        Specifies an array of e-mail addresses to which the scan notification is sent.
        """
        return pulumi.get(self, "emails")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[bool]:
        """
        Recurring scans state.
        """
        return pulumi.get(self, "is_enabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


