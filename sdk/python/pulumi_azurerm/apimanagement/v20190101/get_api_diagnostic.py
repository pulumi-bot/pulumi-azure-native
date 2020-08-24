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
    'GetApiDiagnosticResult',
    'AwaitableGetApiDiagnosticResult',
    'get_api_diagnostic',
]

@pulumi.output_type
class GetApiDiagnosticResult:
    """
    Diagnostic details.
    """
    def __init__(__self__, always_log=None, backend=None, enable_http_correlation_headers=None, frontend=None, http_correlation_protocol=None, logger_id=None, name=None, sampling=None, type=None, verbosity=None):
        if always_log and not isinstance(always_log, str):
            raise TypeError("Expected argument 'always_log' to be a str")
        pulumi.set(__self__, "always_log", always_log)
        if backend and not isinstance(backend, dict):
            raise TypeError("Expected argument 'backend' to be a dict")
        pulumi.set(__self__, "backend", backend)
        if enable_http_correlation_headers and not isinstance(enable_http_correlation_headers, bool):
            raise TypeError("Expected argument 'enable_http_correlation_headers' to be a bool")
        pulumi.set(__self__, "enable_http_correlation_headers", enable_http_correlation_headers)
        if frontend and not isinstance(frontend, dict):
            raise TypeError("Expected argument 'frontend' to be a dict")
        pulumi.set(__self__, "frontend", frontend)
        if http_correlation_protocol and not isinstance(http_correlation_protocol, str):
            raise TypeError("Expected argument 'http_correlation_protocol' to be a str")
        pulumi.set(__self__, "http_correlation_protocol", http_correlation_protocol)
        if logger_id and not isinstance(logger_id, str):
            raise TypeError("Expected argument 'logger_id' to be a str")
        pulumi.set(__self__, "logger_id", logger_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sampling and not isinstance(sampling, dict):
            raise TypeError("Expected argument 'sampling' to be a dict")
        pulumi.set(__self__, "sampling", sampling)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if verbosity and not isinstance(verbosity, str):
            raise TypeError("Expected argument 'verbosity' to be a str")
        pulumi.set(__self__, "verbosity", verbosity)

    @property
    @pulumi.getter(name="alwaysLog")
    def always_log(self) -> Optional[str]:
        """
        Specifies for what type of messages sampling settings should not apply.
        """
        return pulumi.get(self, "always_log")

    @property
    @pulumi.getter
    def backend(self) -> Optional['outputs.PipelineDiagnosticSettingsResponse']:
        """
        Diagnostic settings for incoming/outgoing HTTP messages to the Backend
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="enableHttpCorrelationHeaders")
    def enable_http_correlation_headers(self) -> Optional[bool]:
        """
        Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
        """
        return pulumi.get(self, "enable_http_correlation_headers")

    @property
    @pulumi.getter
    def frontend(self) -> Optional['outputs.PipelineDiagnosticSettingsResponse']:
        """
        Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
        """
        return pulumi.get(self, "frontend")

    @property
    @pulumi.getter(name="httpCorrelationProtocol")
    def http_correlation_protocol(self) -> Optional[str]:
        """
        Sets correlation protocol to use for Application Insights diagnostics.
        """
        return pulumi.get(self, "http_correlation_protocol")

    @property
    @pulumi.getter(name="loggerId")
    def logger_id(self) -> str:
        """
        Resource Id of a target logger.
        """
        return pulumi.get(self, "logger_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sampling(self) -> Optional['outputs.SamplingSettingsResponse']:
        """
        Sampling settings for Diagnostic.
        """
        return pulumi.get(self, "sampling")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def verbosity(self) -> Optional[str]:
        """
        The verbosity level applied to traces emitted by trace policies.
        """
        return pulumi.get(self, "verbosity")


class AwaitableGetApiDiagnosticResult(GetApiDiagnosticResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiDiagnosticResult(
            always_log=self.always_log,
            backend=self.backend,
            enable_http_correlation_headers=self.enable_http_correlation_headers,
            frontend=self.frontend,
            http_correlation_protocol=self.http_correlation_protocol,
            logger_id=self.logger_id,
            name=self.name,
            sampling=self.sampling,
            type=self.type,
            verbosity=self.verbosity)


def get_api_diagnostic(api_id: Optional[str] = None,
                       name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       service_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiDiagnosticResult:
    """
    Use this data source to access information about an existing resource.

    :param str api_id: API identifier. Must be unique in the current API Management service instance.
    :param str name: Diagnostic identifier. Must be unique in the current API Management service instance.
    :param str resource_group_name: The name of the resource group.
    :param str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:apimanagement/v20190101:getApiDiagnostic', __args__, opts=opts, typ=GetApiDiagnosticResult).value

    return AwaitableGetApiDiagnosticResult(
        always_log=__ret__.always_log,
        backend=__ret__.backend,
        enable_http_correlation_headers=__ret__.enable_http_correlation_headers,
        frontend=__ret__.frontend,
        http_correlation_protocol=__ret__.http_correlation_protocol,
        logger_id=__ret__.logger_id,
        name=__ret__.name,
        sampling=__ret__.sampling,
        type=__ret__.type,
        verbosity=__ret__.verbosity)
