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
    'AdvancedScheduleMonthlyOccurrenceResponse',
    'AdvancedScheduleResponse',
    'ConnectionTypeAssociationPropertyResponse',
    'ContentHashResponse',
    'ContentLinkResponse',
    'ContentSourceResponse',
    'DscConfigurationAssociationPropertyResponse',
    'DscConfigurationParameterResponse',
    'FieldDefinitionResponse',
    'KeyResponseResult',
    'ModuleErrorInfoResponse',
    'RunbookAssociationPropertyResponse',
    'RunbookDraftResponse',
    'RunbookParameterResponse',
    'ScheduleAssociationPropertyResponse',
    'SkuResponse',
]

@pulumi.output_type
class AdvancedScheduleMonthlyOccurrenceResponse(dict):
    """
    The properties of the create advanced schedule monthly occurrence.
    """
    def __init__(__self__, *,
                 day: Optional[str] = None,
                 occurrence: Optional[float] = None):
        """
        The properties of the create advanced schedule monthly occurrence.
        :param str day: Day of the occurrence. Must be one of monday, tuesday, wednesday, thursday, friday, saturday, sunday.
        :param float occurrence: Occurrence of the week within the month. Must be between 1 and 5
        """
        if day is not None:
            pulumi.set(__self__, "day", day)
        if occurrence is not None:
            pulumi.set(__self__, "occurrence", occurrence)

    @property
    @pulumi.getter
    def day(self) -> Optional[str]:
        """
        Day of the occurrence. Must be one of monday, tuesday, wednesday, thursday, friday, saturday, sunday.
        """
        return pulumi.get(self, "day")

    @property
    @pulumi.getter
    def occurrence(self) -> Optional[float]:
        """
        Occurrence of the week within the month. Must be between 1 and 5
        """
        return pulumi.get(self, "occurrence")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AdvancedScheduleResponse(dict):
    """
    The properties of the create Advanced Schedule.
    """
    def __init__(__self__, *,
                 month_days: Optional[List[float]] = None,
                 monthly_occurrences: Optional[List['outputs.AdvancedScheduleMonthlyOccurrenceResponse']] = None,
                 week_days: Optional[List[str]] = None):
        """
        The properties of the create Advanced Schedule.
        :param List[float] month_days: Days of the month that the job should execute on. Must be between 1 and 31.
        :param List['AdvancedScheduleMonthlyOccurrenceResponseArgs'] monthly_occurrences: Occurrences of days within a month.
        :param List[str] week_days: Days of the week that the job should execute on.
        """
        if month_days is not None:
            pulumi.set(__self__, "month_days", month_days)
        if monthly_occurrences is not None:
            pulumi.set(__self__, "monthly_occurrences", monthly_occurrences)
        if week_days is not None:
            pulumi.set(__self__, "week_days", week_days)

    @property
    @pulumi.getter(name="monthDays")
    def month_days(self) -> Optional[List[float]]:
        """
        Days of the month that the job should execute on. Must be between 1 and 31.
        """
        return pulumi.get(self, "month_days")

    @property
    @pulumi.getter(name="monthlyOccurrences")
    def monthly_occurrences(self) -> Optional[List['outputs.AdvancedScheduleMonthlyOccurrenceResponse']]:
        """
        Occurrences of days within a month.
        """
        return pulumi.get(self, "monthly_occurrences")

    @property
    @pulumi.getter(name="weekDays")
    def week_days(self) -> Optional[List[str]]:
        """
        Days of the week that the job should execute on.
        """
        return pulumi.get(self, "week_days")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ConnectionTypeAssociationPropertyResponse(dict):
    """
    The connection type property associated with the entity.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        The connection type property associated with the entity.
        :param str name: Gets or sets the name of the connection type.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the name of the connection type.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContentHashResponse(dict):
    """
    Definition of the runbook property type.
    """
    def __init__(__self__, *,
                 algorithm: str,
                 value: str):
        """
        Definition of the runbook property type.
        :param str algorithm: Gets or sets the content hash algorithm used to hash the content.
        :param str value: Gets or sets expected hash value of the content.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        Gets or sets the content hash algorithm used to hash the content.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Gets or sets expected hash value of the content.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContentLinkResponse(dict):
    """
    Definition of the content link.
    """
    def __init__(__self__, *,
                 content_hash: Optional['outputs.ContentHashResponse'] = None,
                 uri: Optional[str] = None,
                 version: Optional[str] = None):
        """
        Definition of the content link.
        :param 'ContentHashResponseArgs' content_hash: Gets or sets the hash.
        :param str uri: Gets or sets the uri of the runbook content.
        :param str version: Gets or sets the version of the content.
        """
        if content_hash is not None:
            pulumi.set(__self__, "content_hash", content_hash)
        if uri is not None:
            pulumi.set(__self__, "uri", uri)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="contentHash")
    def content_hash(self) -> Optional['outputs.ContentHashResponse']:
        """
        Gets or sets the hash.
        """
        return pulumi.get(self, "content_hash")

    @property
    @pulumi.getter
    def uri(self) -> Optional[str]:
        """
        Gets or sets the uri of the runbook content.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Gets or sets the version of the content.
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContentSourceResponse(dict):
    """
    Definition of the content source.
    """
    def __init__(__self__, *,
                 hash: Optional['outputs.ContentHashResponse'] = None,
                 type: Optional[str] = None,
                 value: Optional[str] = None,
                 version: Optional[str] = None):
        """
        Definition of the content source.
        :param 'ContentHashResponseArgs' hash: Gets or sets the hash.
        :param str type: Gets or sets the content source type.
        :param str value: Gets or sets the value of the content. This is based on the content source type.
        :param str version: Gets or sets the version of the content.
        """
        if hash is not None:
            pulumi.set(__self__, "hash", hash)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def hash(self) -> Optional['outputs.ContentHashResponse']:
        """
        Gets or sets the hash.
        """
        return pulumi.get(self, "hash")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Gets or sets the content source type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        Gets or sets the value of the content. This is based on the content source type.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Gets or sets the version of the content.
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DscConfigurationAssociationPropertyResponse(dict):
    """
    The Dsc configuration property associated with the entity.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        The Dsc configuration property associated with the entity.
        :param str name: Gets or sets the name of the Dsc configuration.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the name of the Dsc configuration.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DscConfigurationParameterResponse(dict):
    """
    Definition of the configuration parameter type.
    """
    def __init__(__self__, *,
                 default_value: Optional[str] = None,
                 is_mandatory: Optional[bool] = None,
                 position: Optional[float] = None,
                 type: Optional[str] = None):
        """
        Definition of the configuration parameter type.
        :param str default_value: Gets or sets the default value of parameter.
        :param bool is_mandatory: Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
        :param float position: Get or sets the position of the parameter.
        :param str type: Gets or sets the type of the parameter.
        """
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if is_mandatory is not None:
            pulumi.set(__self__, "is_mandatory", is_mandatory)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[str]:
        """
        Gets or sets the default value of parameter.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter(name="isMandatory")
    def is_mandatory(self) -> Optional[bool]:
        """
        Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
        """
        return pulumi.get(self, "is_mandatory")

    @property
    @pulumi.getter
    def position(self) -> Optional[float]:
        """
        Get or sets the position of the parameter.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Gets or sets the type of the parameter.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FieldDefinitionResponse(dict):
    """
    Definition of the connection fields.
    """
    def __init__(__self__, *,
                 type: str,
                 is_encrypted: Optional[bool] = None,
                 is_optional: Optional[bool] = None):
        """
        Definition of the connection fields.
        :param str type: Gets or sets the type of the connection field definition.
        :param bool is_encrypted: Gets or sets the isEncrypted flag of the connection field definition.
        :param bool is_optional: Gets or sets the isOptional flag of the connection field definition.
        """
        pulumi.set(__self__, "type", type)
        if is_encrypted is not None:
            pulumi.set(__self__, "is_encrypted", is_encrypted)
        if is_optional is not None:
            pulumi.set(__self__, "is_optional", is_optional)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Gets or sets the type of the connection field definition.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="isEncrypted")
    def is_encrypted(self) -> Optional[bool]:
        """
        Gets or sets the isEncrypted flag of the connection field definition.
        """
        return pulumi.get(self, "is_encrypted")

    @property
    @pulumi.getter(name="isOptional")
    def is_optional(self) -> Optional[bool]:
        """
        Gets or sets the isOptional flag of the connection field definition.
        """
        return pulumi.get(self, "is_optional")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class KeyResponseResult(dict):
    """
    Automation key which is used to register a DSC Node
    """
    def __init__(__self__, *,
                 key_name: str,
                 permissions: str,
                 value: str):
        """
        Automation key which is used to register a DSC Node
        :param str key_name: Automation key name.
        :param str permissions: Automation key permissions.
        :param str value: Value of the Automation Key used for registration.
        """
        pulumi.set(__self__, "key_name", key_name)
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> str:
        """
        Automation key name.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter
    def permissions(self) -> str:
        """
        Automation key permissions.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Value of the Automation Key used for registration.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ModuleErrorInfoResponse(dict):
    """
    Definition of the module error info type.
    """
    def __init__(__self__, *,
                 code: Optional[str] = None,
                 message: Optional[str] = None):
        """
        Definition of the module error info type.
        :param str code: Gets or sets the error code.
        :param str message: Gets or sets the error message.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if message is not None:
            pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def code(self) -> Optional[str]:
        """
        Gets or sets the error code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        Gets or sets the error message.
        """
        return pulumi.get(self, "message")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RunbookAssociationPropertyResponse(dict):
    """
    The runbook property associated with the entity.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        The runbook property associated with the entity.
        :param str name: Gets or sets the name of the runbook.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the name of the runbook.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RunbookDraftResponse(dict):
    def __init__(__self__, *,
                 creation_time: Optional[str] = None,
                 draft_content_link: Optional['outputs.ContentLinkResponse'] = None,
                 in_edit: Optional[bool] = None,
                 last_modified_time: Optional[str] = None,
                 output_types: Optional[List[str]] = None,
                 parameters: Optional[Mapping[str, 'outputs.RunbookParameterResponse']] = None):
        """
        :param str creation_time: Gets or sets the creation time of the runbook draft.
        :param 'ContentLinkResponseArgs' draft_content_link: Gets or sets the draft runbook content link.
        :param bool in_edit: Gets or sets whether runbook is in edit mode.
        :param str last_modified_time: Gets or sets the last modified time of the runbook draft.
        :param List[str] output_types: Gets or sets the runbook output types.
        :param Mapping[str, 'RunbookParameterResponseArgs'] parameters: Gets or sets the runbook draft parameters.
        """
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if draft_content_link is not None:
            pulumi.set(__self__, "draft_content_link", draft_content_link)
        if in_edit is not None:
            pulumi.set(__self__, "in_edit", in_edit)
        if last_modified_time is not None:
            pulumi.set(__self__, "last_modified_time", last_modified_time)
        if output_types is not None:
            pulumi.set(__self__, "output_types", output_types)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        """
        Gets or sets the creation time of the runbook draft.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="draftContentLink")
    def draft_content_link(self) -> Optional['outputs.ContentLinkResponse']:
        """
        Gets or sets the draft runbook content link.
        """
        return pulumi.get(self, "draft_content_link")

    @property
    @pulumi.getter(name="inEdit")
    def in_edit(self) -> Optional[bool]:
        """
        Gets or sets whether runbook is in edit mode.
        """
        return pulumi.get(self, "in_edit")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[str]:
        """
        Gets or sets the last modified time of the runbook draft.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="outputTypes")
    def output_types(self) -> Optional[List[str]]:
        """
        Gets or sets the runbook output types.
        """
        return pulumi.get(self, "output_types")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, 'outputs.RunbookParameterResponse']]:
        """
        Gets or sets the runbook draft parameters.
        """
        return pulumi.get(self, "parameters")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RunbookParameterResponse(dict):
    """
    Definition of the runbook parameter type.
    """
    def __init__(__self__, *,
                 default_value: Optional[str] = None,
                 is_mandatory: Optional[bool] = None,
                 position: Optional[float] = None,
                 type: Optional[str] = None):
        """
        Definition of the runbook parameter type.
        :param str default_value: Gets or sets the default value of parameter.
        :param bool is_mandatory: Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
        :param float position: Get or sets the position of the parameter.
        :param str type: Gets or sets the type of the parameter.
        """
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if is_mandatory is not None:
            pulumi.set(__self__, "is_mandatory", is_mandatory)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[str]:
        """
        Gets or sets the default value of parameter.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter(name="isMandatory")
    def is_mandatory(self) -> Optional[bool]:
        """
        Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
        """
        return pulumi.get(self, "is_mandatory")

    @property
    @pulumi.getter
    def position(self) -> Optional[float]:
        """
        Get or sets the position of the parameter.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Gets or sets the type of the parameter.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleAssociationPropertyResponse(dict):
    """
    The schedule property associated with the entity.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        The schedule property associated with the entity.
        :param str name: Gets or sets the name of the Schedule.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the name of the Schedule.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    The account SKU.
    """
    def __init__(__self__, *,
                 name: str,
                 capacity: Optional[float] = None,
                 family: Optional[str] = None):
        """
        The account SKU.
        :param str name: Gets or sets the SKU name of the account.
        :param float capacity: Gets or sets the SKU capacity.
        :param str family: Gets or sets the SKU family.
        """
        pulumi.set(__self__, "name", name)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if family is not None:
            pulumi.set(__self__, "family", family)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Gets or sets the SKU name of the account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def capacity(self) -> Optional[float]:
        """
        Gets or sets the SKU capacity.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def family(self) -> Optional[str]:
        """
        Gets or sets the SKU family.
        """
        return pulumi.get(self, "family")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


