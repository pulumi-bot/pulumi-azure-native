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
    'AssignmentInfoResponse',
    'AssignmentReportResourceComplianceReasonResponse',
    'AssignmentReportResourceResponse',
    'AssignmentReportResponse',
    'ConfigurationInfoResponse',
    'ConfigurationParameterResponse',
    'ConfigurationSettingResponse',
    'ErrorDetailResponse',
    'GuestConfigurationAssignmentPropertiesResponse',
    'GuestConfigurationNavigationResponse',
    'MachineExtensionInstanceViewResponse',
    'MachineExtensionInstanceViewResponseStatus',
    'MachineExtensionPropertiesResponseInstanceView',
    'MachinePropertiesResponseOsProfile',
    'MachineResponseIdentity',
    'VMInfoResponse',
    'LocationDataResponse',
]

@pulumi.output_type
class AssignmentInfoResponse(dict):
    """
    Information about the guest configuration assignment.
    """
    def __init__(__self__, *,
                 name: str,
                 configuration: Optional['outputs.ConfigurationInfoResponse'] = None):
        """
        Information about the guest configuration assignment.
        :param str name: Name of the guest configuration assignment.
        :param 'ConfigurationInfoResponseArgs' configuration: Information about the configuration.
        """
        pulumi.set(__self__, "name", name)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the guest configuration assignment.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def configuration(self) -> Optional['outputs.ConfigurationInfoResponse']:
        """
        Information about the configuration.
        """
        return pulumi.get(self, "configuration")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AssignmentReportResourceComplianceReasonResponse(dict):
    """
    Reason and code for the compliance of the guest configuration assignment resource.
    """
    def __init__(__self__, *,
                 code: str,
                 phrase: str):
        """
        Reason and code for the compliance of the guest configuration assignment resource.
        :param str code: Code for the compliance of the guest configuration assignment resource.
        :param str phrase: Reason for the compliance of the guest configuration assignment resource.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "phrase", phrase)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Code for the compliance of the guest configuration assignment resource.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def phrase(self) -> str:
        """
        Reason for the compliance of the guest configuration assignment resource.
        """
        return pulumi.get(self, "phrase")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AssignmentReportResourceResponse(dict):
    """
    The guest configuration assignment resource.
    """
    def __init__(__self__, *,
                 compliance_status: str,
                 properties: Mapping[str, Any],
                 resource_id: str,
                 reasons: Optional[List['outputs.AssignmentReportResourceComplianceReasonResponse']] = None):
        """
        The guest configuration assignment resource.
        :param str compliance_status: A value indicating compliance status of the machine for the assigned guest configuration.
        :param Mapping[str, Any] properties: Properties of a guest configuration assignment resource.
        :param str resource_id: Name of the guest configuration assignment resource setting.
        :param List['AssignmentReportResourceComplianceReasonResponseArgs'] reasons: Compliance reason and reason code for a resource.
        """
        pulumi.set(__self__, "compliance_status", compliance_status)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "resource_id", resource_id)
        if reasons is not None:
            pulumi.set(__self__, "reasons", reasons)

    @property
    @pulumi.getter(name="complianceStatus")
    def compliance_status(self) -> str:
        """
        A value indicating compliance status of the machine for the assigned guest configuration.
        """
        return pulumi.get(self, "compliance_status")

    @property
    @pulumi.getter
    def properties(self) -> Mapping[str, Any]:
        """
        Properties of a guest configuration assignment resource.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        Name of the guest configuration assignment resource setting.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def reasons(self) -> Optional[List['outputs.AssignmentReportResourceComplianceReasonResponse']]:
        """
        Compliance reason and reason code for a resource.
        """
        return pulumi.get(self, "reasons")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AssignmentReportResponse(dict):
    def __init__(__self__, *,
                 compliance_status: str,
                 end_time: str,
                 id: str,
                 operation_type: str,
                 report_id: str,
                 start_time: str,
                 assignment: Optional['outputs.AssignmentInfoResponse'] = None,
                 resources: Optional[List['outputs.AssignmentReportResourceResponse']] = None,
                 vm: Optional['outputs.VMInfoResponse'] = None):
        """
        :param str compliance_status: A value indicating compliance status of the machine for the assigned guest configuration.
        :param str end_time: End date and time of the guest configuration assignment compliance status check.
        :param str id: ARM resource id of the report for the guest configuration assignment.
        :param str operation_type: Type of report, Consistency or Initial
        :param str report_id: GUID that identifies the guest configuration assignment report under a subscription, resource group.
        :param str start_time: Start date and time of the guest configuration assignment compliance status check.
        :param 'AssignmentInfoResponseArgs' assignment: Configuration details of the guest configuration assignment.
        :param List['AssignmentReportResourceResponseArgs'] resources: The list of resources for which guest configuration assignment compliance is checked.
        :param 'VMInfoResponseArgs' vm: Information about the VM.
        """
        pulumi.set(__self__, "compliance_status", compliance_status)
        pulumi.set(__self__, "end_time", end_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "operation_type", operation_type)
        pulumi.set(__self__, "report_id", report_id)
        pulumi.set(__self__, "start_time", start_time)
        if assignment is not None:
            pulumi.set(__self__, "assignment", assignment)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if vm is not None:
            pulumi.set(__self__, "vm", vm)

    @property
    @pulumi.getter(name="complianceStatus")
    def compliance_status(self) -> str:
        """
        A value indicating compliance status of the machine for the assigned guest configuration.
        """
        return pulumi.get(self, "compliance_status")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        End date and time of the guest configuration assignment compliance status check.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ARM resource id of the report for the guest configuration assignment.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="operationType")
    def operation_type(self) -> str:
        """
        Type of report, Consistency or Initial
        """
        return pulumi.get(self, "operation_type")

    @property
    @pulumi.getter(name="reportId")
    def report_id(self) -> str:
        """
        GUID that identifies the guest configuration assignment report under a subscription, resource group.
        """
        return pulumi.get(self, "report_id")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        Start date and time of the guest configuration assignment compliance status check.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def assignment(self) -> Optional['outputs.AssignmentInfoResponse']:
        """
        Configuration details of the guest configuration assignment.
        """
        return pulumi.get(self, "assignment")

    @property
    @pulumi.getter
    def resources(self) -> Optional[List['outputs.AssignmentReportResourceResponse']]:
        """
        The list of resources for which guest configuration assignment compliance is checked.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def vm(self) -> Optional['outputs.VMInfoResponse']:
        """
        Information about the VM.
        """
        return pulumi.get(self, "vm")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ConfigurationInfoResponse(dict):
    """
    Information about the configuration.
    """
    def __init__(__self__, *,
                 name: str,
                 version: str):
        """
        Information about the configuration.
        :param str name: Name of the configuration.
        :param str version: Version of the configuration.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Version of the configuration.
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ConfigurationParameterResponse(dict):
    """
    Represents a configuration parameter.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 value: Optional[str] = None):
        """
        Represents a configuration parameter.
        :param str name: Name of the configuration parameter.
        :param str value: Value of the configuration parameter.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the configuration parameter.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        Value of the configuration parameter.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ConfigurationSettingResponse(dict):
    """
    Configuration setting of LCM (Local Configuration Manager).
    """
    def __init__(__self__, *,
                 action_after_reboot: Optional[str] = None,
                 allow_module_overwrite: Optional[str] = None,
                 configuration_mode: Optional[str] = None,
                 configuration_mode_frequency_mins: Optional[float] = None,
                 reboot_if_needed: Optional[str] = None,
                 refresh_frequency_mins: Optional[float] = None):
        """
        Configuration setting of LCM (Local Configuration Manager).
        :param str action_after_reboot: Specifies what happens after a reboot during the application of a configuration. The possible values are ContinueConfiguration and StopConfiguration
        :param str allow_module_overwrite: If true - new configurations downloaded from the pull service are allowed to overwrite the old ones on the target node. Otherwise, false
        :param str configuration_mode: Specifies how the LCM(Local Configuration Manager) actually applies the configuration to the target nodes. Possible values are ApplyOnly, ApplyAndMonitor, and ApplyAndAutoCorrect.
        :param float configuration_mode_frequency_mins: How often, in minutes, the current configuration is checked and applied. This property is ignored if the ConfigurationMode property is set to ApplyOnly. The default value is 15.
        :param str reboot_if_needed: Set this to true to automatically reboot the node after a configuration that requires reboot is applied. Otherwise, you will have to manually reboot the node for any configuration that requires it. The default value is false. To use this setting when a reboot condition is enacted by something other than DSC (such as Windows Installer), combine this setting with the xPendingReboot module.
        :param float refresh_frequency_mins: The time interval, in minutes, at which the LCM checks a pull service to get updated configurations. This value is ignored if the LCM is not configured in pull mode. The default value is 30.
        """
        if action_after_reboot is not None:
            pulumi.set(__self__, "action_after_reboot", action_after_reboot)
        if allow_module_overwrite is not None:
            pulumi.set(__self__, "allow_module_overwrite", allow_module_overwrite)
        if configuration_mode is not None:
            pulumi.set(__self__, "configuration_mode", configuration_mode)
        if configuration_mode_frequency_mins is not None:
            pulumi.set(__self__, "configuration_mode_frequency_mins", configuration_mode_frequency_mins)
        if reboot_if_needed is not None:
            pulumi.set(__self__, "reboot_if_needed", reboot_if_needed)
        if refresh_frequency_mins is not None:
            pulumi.set(__self__, "refresh_frequency_mins", refresh_frequency_mins)

    @property
    @pulumi.getter(name="actionAfterReboot")
    def action_after_reboot(self) -> Optional[str]:
        """
        Specifies what happens after a reboot during the application of a configuration. The possible values are ContinueConfiguration and StopConfiguration
        """
        return pulumi.get(self, "action_after_reboot")

    @property
    @pulumi.getter(name="allowModuleOverwrite")
    def allow_module_overwrite(self) -> Optional[str]:
        """
        If true - new configurations downloaded from the pull service are allowed to overwrite the old ones on the target node. Otherwise, false
        """
        return pulumi.get(self, "allow_module_overwrite")

    @property
    @pulumi.getter(name="configurationMode")
    def configuration_mode(self) -> Optional[str]:
        """
        Specifies how the LCM(Local Configuration Manager) actually applies the configuration to the target nodes. Possible values are ApplyOnly, ApplyAndMonitor, and ApplyAndAutoCorrect.
        """
        return pulumi.get(self, "configuration_mode")

    @property
    @pulumi.getter(name="configurationModeFrequencyMins")
    def configuration_mode_frequency_mins(self) -> Optional[float]:
        """
        How often, in minutes, the current configuration is checked and applied. This property is ignored if the ConfigurationMode property is set to ApplyOnly. The default value is 15.
        """
        return pulumi.get(self, "configuration_mode_frequency_mins")

    @property
    @pulumi.getter(name="rebootIfNeeded")
    def reboot_if_needed(self) -> Optional[str]:
        """
        Set this to true to automatically reboot the node after a configuration that requires reboot is applied. Otherwise, you will have to manually reboot the node for any configuration that requires it. The default value is false. To use this setting when a reboot condition is enacted by something other than DSC (such as Windows Installer), combine this setting with the xPendingReboot module.
        """
        return pulumi.get(self, "reboot_if_needed")

    @property
    @pulumi.getter(name="refreshFrequencyMins")
    def refresh_frequency_mins(self) -> Optional[float]:
        """
        The time interval, in minutes, at which the LCM checks a pull service to get updated configurations. This value is ignored if the LCM is not configured in pull mode. The default value is 30.
        """
        return pulumi.get(self, "refresh_frequency_mins")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ErrorDetailResponse(dict):
    def __init__(__self__, *,
                 code: str,
                 message: str,
                 details: Optional[List['outputs.ErrorDetailResponse']] = None,
                 target: Optional[str] = None):
        """
        :param str code: The error's code.
        :param str message: A human readable error message.
        :param List['ErrorDetailResponseArgs'] details: Additional error details.
        :param str target: Indicates which property in the request is responsible for the error.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "message", message)
        if details is not None:
            pulumi.set(__self__, "details", details)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        The error's code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        A human readable error message.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def details(self) -> Optional[List['outputs.ErrorDetailResponse']]:
        """
        Additional error details.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        """
        Indicates which property in the request is responsible for the error.
        """
        return pulumi.get(self, "target")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GuestConfigurationAssignmentPropertiesResponse(dict):
    """
    Guest configuration assignment properties.
    """
    def __init__(__self__, *,
                 assignment_hash: str,
                 compliance_status: str,
                 last_compliance_status_checked: str,
                 latest_report_id: str,
                 provisioning_state: str,
                 target_resource_id: str,
                 context: Optional[str] = None,
                 guest_configuration: Optional['outputs.GuestConfigurationNavigationResponse'] = None,
                 latest_assignment_report: Optional['outputs.AssignmentReportResponse'] = None):
        """
        Guest configuration assignment properties.
        :param str assignment_hash: Combined hash of the configuration package and parameters.
        :param str compliance_status: A value indicating compliance status of the machine for the assigned guest configuration.
        :param str last_compliance_status_checked: Date and time when last compliance status was checked.
        :param str latest_report_id: Id of the latest report for the guest configuration assignment. 
        :param str provisioning_state: The provisioning state, which only appears in the response.
        :param str target_resource_id: VM resource Id.
        :param str context: The source which initiated the guest configuration assignment. Ex: Azure Policy
        :param 'GuestConfigurationNavigationResponseArgs' guest_configuration: The guest configuration to assign.
        :param 'AssignmentReportResponseArgs' latest_assignment_report: Last reported guest configuration assignment report.
        """
        pulumi.set(__self__, "assignment_hash", assignment_hash)
        pulumi.set(__self__, "compliance_status", compliance_status)
        pulumi.set(__self__, "last_compliance_status_checked", last_compliance_status_checked)
        pulumi.set(__self__, "latest_report_id", latest_report_id)
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        pulumi.set(__self__, "target_resource_id", target_resource_id)
        if context is not None:
            pulumi.set(__self__, "context", context)
        if guest_configuration is not None:
            pulumi.set(__self__, "guest_configuration", guest_configuration)
        if latest_assignment_report is not None:
            pulumi.set(__self__, "latest_assignment_report", latest_assignment_report)

    @property
    @pulumi.getter(name="assignmentHash")
    def assignment_hash(self) -> str:
        """
        Combined hash of the configuration package and parameters.
        """
        return pulumi.get(self, "assignment_hash")

    @property
    @pulumi.getter(name="complianceStatus")
    def compliance_status(self) -> str:
        """
        A value indicating compliance status of the machine for the assigned guest configuration.
        """
        return pulumi.get(self, "compliance_status")

    @property
    @pulumi.getter(name="lastComplianceStatusChecked")
    def last_compliance_status_checked(self) -> str:
        """
        Date and time when last compliance status was checked.
        """
        return pulumi.get(self, "last_compliance_status_checked")

    @property
    @pulumi.getter(name="latestReportId")
    def latest_report_id(self) -> str:
        """
        Id of the latest report for the guest configuration assignment. 
        """
        return pulumi.get(self, "latest_report_id")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> str:
        """
        VM resource Id.
        """
        return pulumi.get(self, "target_resource_id")

    @property
    @pulumi.getter
    def context(self) -> Optional[str]:
        """
        The source which initiated the guest configuration assignment. Ex: Azure Policy
        """
        return pulumi.get(self, "context")

    @property
    @pulumi.getter(name="guestConfiguration")
    def guest_configuration(self) -> Optional['outputs.GuestConfigurationNavigationResponse']:
        """
        The guest configuration to assign.
        """
        return pulumi.get(self, "guest_configuration")

    @property
    @pulumi.getter(name="latestAssignmentReport")
    def latest_assignment_report(self) -> Optional['outputs.AssignmentReportResponse']:
        """
        Last reported guest configuration assignment report.
        """
        return pulumi.get(self, "latest_assignment_report")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GuestConfigurationNavigationResponse(dict):
    """
    Guest configuration is an artifact that encapsulates DSC configuration and its dependencies. The artifact is a zip file containing DSC configuration (as MOF) and dependent resources and other dependencies like modules.
    """
    def __init__(__self__, *,
                 content_hash: str,
                 content_uri: str,
                 configuration_parameter: Optional[List['outputs.ConfigurationParameterResponse']] = None,
                 configuration_setting: Optional['outputs.ConfigurationSettingResponse'] = None,
                 kind: Optional[str] = None,
                 name: Optional[str] = None,
                 version: Optional[str] = None):
        """
        Guest configuration is an artifact that encapsulates DSC configuration and its dependencies. The artifact is a zip file containing DSC configuration (as MOF) and dependent resources and other dependencies like modules.
        :param str content_hash: Combined hash of the guest configuration package and configuration parameters.
        :param str content_uri: Uri of the storage where guest configuration package is uploaded.
        :param List['ConfigurationParameterResponseArgs'] configuration_parameter: The configuration parameters for the guest configuration.
        :param 'ConfigurationSettingResponseArgs' configuration_setting: The configuration setting for the guest configuration.
        :param str kind: Kind of the guest configuration. For example:DSC
        :param str name: Name of the guest configuration.
        :param str version: Version of the guest configuration.
        """
        pulumi.set(__self__, "content_hash", content_hash)
        pulumi.set(__self__, "content_uri", content_uri)
        if configuration_parameter is not None:
            pulumi.set(__self__, "configuration_parameter", configuration_parameter)
        if configuration_setting is not None:
            pulumi.set(__self__, "configuration_setting", configuration_setting)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="contentHash")
    def content_hash(self) -> str:
        """
        Combined hash of the guest configuration package and configuration parameters.
        """
        return pulumi.get(self, "content_hash")

    @property
    @pulumi.getter(name="contentUri")
    def content_uri(self) -> str:
        """
        Uri of the storage where guest configuration package is uploaded.
        """
        return pulumi.get(self, "content_uri")

    @property
    @pulumi.getter(name="configurationParameter")
    def configuration_parameter(self) -> Optional[List['outputs.ConfigurationParameterResponse']]:
        """
        The configuration parameters for the guest configuration.
        """
        return pulumi.get(self, "configuration_parameter")

    @property
    @pulumi.getter(name="configurationSetting")
    def configuration_setting(self) -> Optional['outputs.ConfigurationSettingResponse']:
        """
        The configuration setting for the guest configuration.
        """
        return pulumi.get(self, "configuration_setting")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of the guest configuration. For example:DSC
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the guest configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Version of the guest configuration.
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MachineExtensionInstanceViewResponse(dict):
    """
    Describes the Machine Extension Instance View.
    """
    def __init__(__self__, *,
                 name: str,
                 type: str,
                 type_handler_version: str,
                 status: Optional['outputs.MachineExtensionInstanceViewResponseStatus'] = None):
        """
        Describes the Machine Extension Instance View.
        :param str name: The machine extension name.
        :param str type: Specifies the type of the extension; an example is "CustomScriptExtension".
        :param str type_handler_version: Specifies the version of the script handler.
        :param 'MachineExtensionInstanceViewResponseStatusArgs' status: Instance view status.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "type_handler_version", type_handler_version)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The machine extension name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the type of the extension; an example is "CustomScriptExtension".
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="typeHandlerVersion")
    def type_handler_version(self) -> str:
        """
        Specifies the version of the script handler.
        """
        return pulumi.get(self, "type_handler_version")

    @property
    @pulumi.getter
    def status(self) -> Optional['outputs.MachineExtensionInstanceViewResponseStatus']:
        """
        Instance view status.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MachineExtensionInstanceViewResponseStatus(dict):
    """
    Instance view status.
    """
    def __init__(__self__, *,
                 code: str,
                 display_status: str,
                 level: str,
                 message: str,
                 time: str):
        """
        Instance view status.
        :param str code: The status code.
        :param str display_status: The short localizable label for the status.
        :param str level: The level code.
        :param str message: The detailed status message, including for alerts and error messages.
        :param str time: The time of the status.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "display_status", display_status)
        pulumi.set(__self__, "level", level)
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "time", time)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        The status code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter(name="displayStatus")
    def display_status(self) -> str:
        """
        The short localizable label for the status.
        """
        return pulumi.get(self, "display_status")

    @property
    @pulumi.getter
    def level(self) -> str:
        """
        The level code.
        """
        return pulumi.get(self, "level")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        The detailed status message, including for alerts and error messages.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def time(self) -> str:
        """
        The time of the status.
        """
        return pulumi.get(self, "time")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MachineExtensionPropertiesResponseInstanceView(dict):
    """
    The machine extension instance view.
    """
    def __init__(__self__, *,
                 name: str,
                 type: str,
                 type_handler_version: str,
                 status: Optional['outputs.MachineExtensionInstanceViewResponseStatus'] = None):
        """
        The machine extension instance view.
        :param str name: The machine extension name.
        :param str type: Specifies the type of the extension; an example is "CustomScriptExtension".
        :param str type_handler_version: Specifies the version of the script handler.
        :param 'MachineExtensionInstanceViewResponseStatusArgs' status: Instance view status.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "type_handler_version", type_handler_version)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The machine extension name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the type of the extension; an example is "CustomScriptExtension".
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="typeHandlerVersion")
    def type_handler_version(self) -> str:
        """
        Specifies the version of the script handler.
        """
        return pulumi.get(self, "type_handler_version")

    @property
    @pulumi.getter
    def status(self) -> Optional['outputs.MachineExtensionInstanceViewResponseStatus']:
        """
        Instance view status.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MachinePropertiesResponseOsProfile(dict):
    """
    Specifies the operating system settings for the hybrid machine.
    """
    def __init__(__self__, *,
                 computer_name: str):
        """
        Specifies the operating system settings for the hybrid machine.
        :param str computer_name: Specifies the host OS name of the hybrid machine.
        """
        pulumi.set(__self__, "computer_name", computer_name)

    @property
    @pulumi.getter(name="computerName")
    def computer_name(self) -> str:
        """
        Specifies the host OS name of the hybrid machine.
        """
        return pulumi.get(self, "computer_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MachineResponseIdentity(dict):
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: Optional[str] = None):
        """
        :param str principal_id: The identity's principal id.
        :param str tenant_id: The identity's tenant id.
        :param str type: The identity type.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The identity's principal id.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The identity's tenant id.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VMInfoResponse(dict):
    """
    Information about the VM.
    """
    def __init__(__self__, *,
                 id: str,
                 uuid: str):
        """
        Information about the VM.
        :param str id: Azure resource Id of the VM.
        :param str uuid: UUID(Universally Unique Identifier) of the VM.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Azure resource Id of the VM.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        UUID(Universally Unique Identifier) of the VM.
        """
        return pulumi.get(self, "uuid")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class LocationDataResponse(dict):
    """
    Metadata pertaining to the geographic location of the resource.
    """
    def __init__(__self__, *,
                 name: str,
                 city: Optional[str] = None,
                 country_or_region: Optional[str] = None,
                 district: Optional[str] = None):
        """
        Metadata pertaining to the geographic location of the resource.
        :param str name: A canonical name for the geographic or physical location.
        :param str city: The city or locality where the resource is located.
        :param str country_or_region: The country or region where the resource is located
        :param str district: The district, state, or province where the resource is located.
        """
        pulumi.set(__self__, "name", name)
        if city is not None:
            pulumi.set(__self__, "city", city)
        if country_or_region is not None:
            pulumi.set(__self__, "country_or_region", country_or_region)
        if district is not None:
            pulumi.set(__self__, "district", district)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A canonical name for the geographic or physical location.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def city(self) -> Optional[str]:
        """
        The city or locality where the resource is located.
        """
        return pulumi.get(self, "city")

    @property
    @pulumi.getter(name="countryOrRegion")
    def country_or_region(self) -> Optional[str]:
        """
        The country or region where the resource is located
        """
        return pulumi.get(self, "country_or_region")

    @property
    @pulumi.getter
    def district(self) -> Optional[str]:
        """
        The district, state, or province where the resource is located.
        """
        return pulumi.get(self, "district")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


