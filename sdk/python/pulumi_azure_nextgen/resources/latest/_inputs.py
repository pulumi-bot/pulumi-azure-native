# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'DebugSettingArgs',
    'DeploymentPropertiesArgs',
    'ExpressionEvaluationOptionsArgs',
    'IdentityArgs',
    'ManagedServiceIdentityArgs',
    'OnErrorDeploymentArgs',
    'ParametersLinkArgs',
    'PlanArgs',
    'SkuArgs',
    'TagsArgs',
    'TemplateLinkArgs',
    'UserAssignedIdentityArgs',
]

@pulumi.input_type
class DebugSettingArgs:
    def __init__(__self__, *,
                 detail_level: Optional[pulumi.Input[str]] = None):
        """
        The debug setting.
        :param pulumi.Input[str] detail_level: Specifies the type of information to log for debugging. The permitted values are none, requestContent, responseContent, or both requestContent and responseContent separated by a comma. The default is none. When setting this value, carefully consider the type of information you are passing in during deployment. By logging information about the request or response, you could potentially expose sensitive data that is retrieved through the deployment operations.
        """
        if detail_level is not None:
            pulumi.set(__self__, "detail_level", detail_level)

    @property
    @pulumi.getter(name="detailLevel")
    def detail_level(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of information to log for debugging. The permitted values are none, requestContent, responseContent, or both requestContent and responseContent separated by a comma. The default is none. When setting this value, carefully consider the type of information you are passing in during deployment. By logging information about the request or response, you could potentially expose sensitive data that is retrieved through the deployment operations.
        """
        return pulumi.get(self, "detail_level")

    @detail_level.setter
    def detail_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "detail_level", value)


@pulumi.input_type
class DeploymentPropertiesArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input[str],
                 debug_setting: Optional[pulumi.Input['DebugSettingArgs']] = None,
                 expression_evaluation_options: Optional[pulumi.Input['ExpressionEvaluationOptionsArgs']] = None,
                 on_error_deployment: Optional[pulumi.Input['OnErrorDeploymentArgs']] = None,
                 parameters: Optional[Any] = None,
                 parameters_link: Optional[pulumi.Input['ParametersLinkArgs']] = None,
                 template: Optional[Any] = None,
                 template_link: Optional[pulumi.Input['TemplateLinkArgs']] = None):
        """
        Deployment properties.
        :param pulumi.Input[str] mode: The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode, resources are deployed without deleting existing resources that are not included in the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included in the template are deleted. Be careful when using Complete mode as you may unintentionally delete resources.
        :param pulumi.Input['DebugSettingArgs'] debug_setting: The debug setting of the deployment.
        :param pulumi.Input['ExpressionEvaluationOptionsArgs'] expression_evaluation_options: Specifies whether template expressions are evaluated within the scope of the parent template or nested template. Only applicable to nested templates. If not specified, default value is outer.
        :param pulumi.Input['OnErrorDeploymentArgs'] on_error_deployment: The deployment on error behavior.
        :param Any parameters: Name and value pairs that define the deployment parameters for the template. You use this element when you want to provide the parameter values directly in the request rather than link to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
        :param pulumi.Input['ParametersLinkArgs'] parameters_link: The URI of parameters file. You use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
        :param Any template: The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
        :param pulumi.Input['TemplateLinkArgs'] template_link: The URI of the template. Use either the templateLink property or the template property, but not both.
        """
        pulumi.set(__self__, "mode", mode)
        if debug_setting is not None:
            pulumi.set(__self__, "debug_setting", debug_setting)
        if expression_evaluation_options is not None:
            pulumi.set(__self__, "expression_evaluation_options", expression_evaluation_options)
        if on_error_deployment is not None:
            pulumi.set(__self__, "on_error_deployment", on_error_deployment)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if parameters_link is not None:
            pulumi.set(__self__, "parameters_link", parameters_link)
        if template is not None:
            pulumi.set(__self__, "template", template)
        if template_link is not None:
            pulumi.set(__self__, "template_link", template_link)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode, resources are deployed without deleting existing resources that are not included in the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included in the template are deleted. Be careful when using Complete mode as you may unintentionally delete resources.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="debugSetting")
    def debug_setting(self) -> Optional[pulumi.Input['DebugSettingArgs']]:
        """
        The debug setting of the deployment.
        """
        return pulumi.get(self, "debug_setting")

    @debug_setting.setter
    def debug_setting(self, value: Optional[pulumi.Input['DebugSettingArgs']]):
        pulumi.set(self, "debug_setting", value)

    @property
    @pulumi.getter(name="expressionEvaluationOptions")
    def expression_evaluation_options(self) -> Optional[pulumi.Input['ExpressionEvaluationOptionsArgs']]:
        """
        Specifies whether template expressions are evaluated within the scope of the parent template or nested template. Only applicable to nested templates. If not specified, default value is outer.
        """
        return pulumi.get(self, "expression_evaluation_options")

    @expression_evaluation_options.setter
    def expression_evaluation_options(self, value: Optional[pulumi.Input['ExpressionEvaluationOptionsArgs']]):
        pulumi.set(self, "expression_evaluation_options", value)

    @property
    @pulumi.getter(name="onErrorDeployment")
    def on_error_deployment(self) -> Optional[pulumi.Input['OnErrorDeploymentArgs']]:
        """
        The deployment on error behavior.
        """
        return pulumi.get(self, "on_error_deployment")

    @on_error_deployment.setter
    def on_error_deployment(self, value: Optional[pulumi.Input['OnErrorDeploymentArgs']]):
        pulumi.set(self, "on_error_deployment", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Any]:
        """
        Name and value pairs that define the deployment parameters for the template. You use this element when you want to provide the parameter values directly in the request rather than link to an existing parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or a well formed JSON string.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[Any]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="parametersLink")
    def parameters_link(self) -> Optional[pulumi.Input['ParametersLinkArgs']]:
        """
        The URI of parameters file. You use this element to link to an existing parameters file. Use either the parametersLink property or the parameters property, but not both.
        """
        return pulumi.get(self, "parameters_link")

    @parameters_link.setter
    def parameters_link(self, value: Optional[pulumi.Input['ParametersLinkArgs']]):
        pulumi.set(self, "parameters_link", value)

    @property
    @pulumi.getter
    def template(self) -> Optional[Any]:
        """
        The template content. You use this element when you want to pass the template syntax directly in the request rather than link to an existing template. It can be a JObject or well-formed JSON string. Use either the templateLink property or the template property, but not both.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[Any]):
        pulumi.set(self, "template", value)

    @property
    @pulumi.getter(name="templateLink")
    def template_link(self) -> Optional[pulumi.Input['TemplateLinkArgs']]:
        """
        The URI of the template. Use either the templateLink property or the template property, but not both.
        """
        return pulumi.get(self, "template_link")

    @template_link.setter
    def template_link(self, value: Optional[pulumi.Input['TemplateLinkArgs']]):
        pulumi.set(self, "template_link", value)


@pulumi.input_type
class ExpressionEvaluationOptionsArgs:
    def __init__(__self__, *,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        Specifies whether template expressions are evaluated within the scope of the parent template or nested template.
        :param pulumi.Input[str] scope: The scope to be used for evaluation of parameters, variables and functions in a nested template.
        """
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        The scope to be used for evaluation of parameters, variables and functions in a nested template.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


@pulumi.input_type
class IdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Identity for the resource.
        :param pulumi.Input[str] type: The identity type.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ManagedServiceIdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None,
                 user_assigned_identities: Optional[pulumi.Input[Mapping[str, pulumi.Input['UserAssignedIdentityArgs']]]] = None):
        """
        Managed identity generic object.
        :param pulumi.Input[str] type: Type of the managed identity.
        :param pulumi.Input[Mapping[str, pulumi.Input['UserAssignedIdentityArgs']]] user_assigned_identities: The list of user-assigned managed identities associated with the resource. Key is the Azure resource Id of the managed identity.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_assigned_identities is not None:
            pulumi.set(__self__, "user_assigned_identities", user_assigned_identities)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the managed identity.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userAssignedIdentities")
    def user_assigned_identities(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['UserAssignedIdentityArgs']]]]:
        """
        The list of user-assigned managed identities associated with the resource. Key is the Azure resource Id of the managed identity.
        """
        return pulumi.get(self, "user_assigned_identities")

    @user_assigned_identities.setter
    def user_assigned_identities(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['UserAssignedIdentityArgs']]]]):
        pulumi.set(self, "user_assigned_identities", value)


@pulumi.input_type
class OnErrorDeploymentArgs:
    def __init__(__self__, *,
                 deployment_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Deployment on error behavior.
        :param pulumi.Input[str] deployment_name: The deployment to be used on error case.
        :param pulumi.Input[str] type: The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
        """
        if deployment_name is not None:
            pulumi.set(__self__, "deployment_name", deployment_name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="deploymentName")
    def deployment_name(self) -> Optional[pulumi.Input[str]]:
        """
        The deployment to be used on error case.
        """
        return pulumi.get(self, "deployment_name")

    @deployment_name.setter
    def deployment_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deployment_name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ParametersLinkArgs:
    def __init__(__self__, *,
                 uri: pulumi.Input[str],
                 content_version: Optional[pulumi.Input[str]] = None):
        """
        Entity representing the reference to the deployment parameters.
        :param pulumi.Input[str] uri: The URI of the parameters file.
        :param pulumi.Input[str] content_version: If included, must match the ContentVersion in the template.
        """
        pulumi.set(__self__, "uri", uri)
        if content_version is not None:
            pulumi.set(__self__, "content_version", content_version)

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Input[str]:
        """
        The URI of the parameters file.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "uri", value)

    @property
    @pulumi.getter(name="contentVersion")
    def content_version(self) -> Optional[pulumi.Input[str]]:
        """
        If included, must match the ContentVersion in the template.
        """
        return pulumi.get(self, "content_version")

    @content_version.setter
    def content_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_version", value)


@pulumi.input_type
class PlanArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 product: Optional[pulumi.Input[str]] = None,
                 promotion_code: Optional[pulumi.Input[str]] = None,
                 publisher: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Plan for the resource.
        :param pulumi.Input[str] name: The plan ID.
        :param pulumi.Input[str] product: The offer ID.
        :param pulumi.Input[str] promotion_code: The promotion code.
        :param pulumi.Input[str] publisher: The publisher ID.
        :param pulumi.Input[str] version: The plan's version.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if product is not None:
            pulumi.set(__self__, "product", product)
        if promotion_code is not None:
            pulumi.set(__self__, "promotion_code", promotion_code)
        if publisher is not None:
            pulumi.set(__self__, "publisher", publisher)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The plan ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def product(self) -> Optional[pulumi.Input[str]]:
        """
        The offer ID.
        """
        return pulumi.get(self, "product")

    @product.setter
    def product(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product", value)

    @property
    @pulumi.getter(name="promotionCode")
    def promotion_code(self) -> Optional[pulumi.Input[str]]:
        """
        The promotion code.
        """
        return pulumi.get(self, "promotion_code")

    @promotion_code.setter
    def promotion_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "promotion_code", value)

    @property
    @pulumi.getter
    def publisher(self) -> Optional[pulumi.Input[str]]:
        """
        The publisher ID.
        """
        return pulumi.get(self, "publisher")

    @publisher.setter
    def publisher(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "publisher", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The plan's version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 capacity: Optional[pulumi.Input[int]] = None,
                 family: Optional[pulumi.Input[str]] = None,
                 model: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 tier: Optional[pulumi.Input[str]] = None):
        """
        SKU for the resource.
        :param pulumi.Input[int] capacity: The SKU capacity.
        :param pulumi.Input[str] family: The SKU family.
        :param pulumi.Input[str] model: The SKU model.
        :param pulumi.Input[str] name: The SKU name.
        :param pulumi.Input[str] size: The SKU size.
        :param pulumi.Input[str] tier: The SKU tier.
        """
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if family is not None:
            pulumi.set(__self__, "family", family)
        if model is not None:
            pulumi.set(__self__, "model", model)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The SKU capacity.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def family(self) -> Optional[pulumi.Input[str]]:
        """
        The SKU family.
        """
        return pulumi.get(self, "family")

    @family.setter
    def family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "family", value)

    @property
    @pulumi.getter
    def model(self) -> Optional[pulumi.Input[str]]:
        """
        The SKU model.
        """
        return pulumi.get(self, "model")

    @model.setter
    def model(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "model", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The SKU name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[str]]:
        """
        The SKU size.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The SKU tier.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)


@pulumi.input_type
class TagsArgs:
    def __init__(__self__, *,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        A dictionary of name and value pairs.
        """
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class TemplateLinkArgs:
    def __init__(__self__, *,
                 content_version: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 relative_path: Optional[pulumi.Input[str]] = None,
                 uri: Optional[pulumi.Input[str]] = None):
        """
        Entity representing the reference to the template.
        :param pulumi.Input[str] content_version: If included, must match the ContentVersion in the template.
        :param pulumi.Input[str] id: The resource id of a Template Spec. Use either the id or uri property, but not both.
        :param pulumi.Input[str] relative_path: Applicable only if this template link references a Template Spec. This relativePath property can optionally be used to reference a Template Spec artifact by path.
        :param pulumi.Input[str] uri: The URI of the template to deploy. Use either the uri or id property, but not both.
        """
        if content_version is not None:
            pulumi.set(__self__, "content_version", content_version)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if relative_path is not None:
            pulumi.set(__self__, "relative_path", relative_path)
        if uri is not None:
            pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter(name="contentVersion")
    def content_version(self) -> Optional[pulumi.Input[str]]:
        """
        If included, must match the ContentVersion in the template.
        """
        return pulumi.get(self, "content_version")

    @content_version.setter
    def content_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_version", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource id of a Template Spec. Use either the id or uri property, but not both.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="relativePath")
    def relative_path(self) -> Optional[pulumi.Input[str]]:
        """
        Applicable only if this template link references a Template Spec. This relativePath property can optionally be used to reference a Template Spec artifact by path.
        """
        return pulumi.get(self, "relative_path")

    @relative_path.setter
    def relative_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "relative_path", value)

    @property
    @pulumi.getter
    def uri(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the template to deploy. Use either the uri or id property, but not both.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uri", value)


@pulumi.input_type
class UserAssignedIdentityArgs:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None):
        """
        User-assigned managed identity.
        :param pulumi.Input[str] client_id: Client App Id associated with this identity.
        :param pulumi.Input[str] principal_id: Azure Active Directory principal ID associated with this identity.
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        Client App Id associated with this identity.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        Azure Active Directory principal ID associated with this identity.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)


