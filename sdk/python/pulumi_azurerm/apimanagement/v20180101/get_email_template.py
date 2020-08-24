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
    'GetEmailTemplateResult',
    'AwaitableGetEmailTemplateResult',
    'get_email_template',
]

@pulumi.output_type
class GetEmailTemplateResult:
    """
    Email Template details.
    """
    def __init__(__self__, body=None, description=None, is_default=None, name=None, parameters=None, subject=None, title=None, type=None):
        if body and not isinstance(body, str):
            raise TypeError("Expected argument 'body' to be a str")
        pulumi.set(__self__, "body", body)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parameters and not isinstance(parameters, list):
            raise TypeError("Expected argument 'parameters' to be a list")
        pulumi.set(__self__, "parameters", parameters)
        if subject and not isinstance(subject, str):
            raise TypeError("Expected argument 'subject' to be a str")
        pulumi.set(__self__, "subject", subject)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def body(self) -> str:
        """
        Email Template Body. This should be a valid XDocument
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the Email Template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> bool:
        """
        Whether the template is the default template provided by Api Management or has been edited.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[List['outputs.EmailTemplateParametersContractPropertiesResponse']]:
        """
        Email Template Parameter values.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def subject(self) -> str:
        """
        Subject of the Template.
        """
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter
    def title(self) -> Optional[str]:
        """
        Title of the Template.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetEmailTemplateResult(GetEmailTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEmailTemplateResult(
            body=self.body,
            description=self.description,
            is_default=self.is_default,
            name=self.name,
            parameters=self.parameters,
            subject=self.subject,
            title=self.title,
            type=self.type)


def get_email_template(name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       service_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEmailTemplateResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Email Template Name Identifier.
    :param str resource_group_name: The name of the resource group.
    :param str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:apimanagement/v20180101:getEmailTemplate', __args__, opts=opts, typ=GetEmailTemplateResult).value

    return AwaitableGetEmailTemplateResult(
        body=__ret__.body,
        description=__ret__.description,
        is_default=__ret__.is_default,
        name=__ret__.name,
        parameters=__ret__.parameters,
        subject=__ret__.subject,
        title=__ret__.title,
        type=__ret__.type)
