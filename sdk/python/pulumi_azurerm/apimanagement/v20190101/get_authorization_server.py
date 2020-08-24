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
    'GetAuthorizationServerResult',
    'AwaitableGetAuthorizationServerResult',
    'get_authorization_server',
]

@pulumi.output_type
class GetAuthorizationServerResult:
    """
    External OAuth authorization server settings.
    """
    def __init__(__self__, authorization_endpoint=None, authorization_methods=None, bearer_token_sending_methods=None, client_authentication_method=None, client_id=None, client_registration_endpoint=None, client_secret=None, default_scope=None, description=None, display_name=None, grant_types=None, name=None, resource_owner_password=None, resource_owner_username=None, support_state=None, token_body_parameters=None, token_endpoint=None, type=None):
        if authorization_endpoint and not isinstance(authorization_endpoint, str):
            raise TypeError("Expected argument 'authorization_endpoint' to be a str")
        pulumi.set(__self__, "authorization_endpoint", authorization_endpoint)
        if authorization_methods and not isinstance(authorization_methods, list):
            raise TypeError("Expected argument 'authorization_methods' to be a list")
        pulumi.set(__self__, "authorization_methods", authorization_methods)
        if bearer_token_sending_methods and not isinstance(bearer_token_sending_methods, list):
            raise TypeError("Expected argument 'bearer_token_sending_methods' to be a list")
        pulumi.set(__self__, "bearer_token_sending_methods", bearer_token_sending_methods)
        if client_authentication_method and not isinstance(client_authentication_method, list):
            raise TypeError("Expected argument 'client_authentication_method' to be a list")
        pulumi.set(__self__, "client_authentication_method", client_authentication_method)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_registration_endpoint and not isinstance(client_registration_endpoint, str):
            raise TypeError("Expected argument 'client_registration_endpoint' to be a str")
        pulumi.set(__self__, "client_registration_endpoint", client_registration_endpoint)
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if default_scope and not isinstance(default_scope, str):
            raise TypeError("Expected argument 'default_scope' to be a str")
        pulumi.set(__self__, "default_scope", default_scope)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if grant_types and not isinstance(grant_types, list):
            raise TypeError("Expected argument 'grant_types' to be a list")
        pulumi.set(__self__, "grant_types", grant_types)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_owner_password and not isinstance(resource_owner_password, str):
            raise TypeError("Expected argument 'resource_owner_password' to be a str")
        pulumi.set(__self__, "resource_owner_password", resource_owner_password)
        if resource_owner_username and not isinstance(resource_owner_username, str):
            raise TypeError("Expected argument 'resource_owner_username' to be a str")
        pulumi.set(__self__, "resource_owner_username", resource_owner_username)
        if support_state and not isinstance(support_state, bool):
            raise TypeError("Expected argument 'support_state' to be a bool")
        pulumi.set(__self__, "support_state", support_state)
        if token_body_parameters and not isinstance(token_body_parameters, list):
            raise TypeError("Expected argument 'token_body_parameters' to be a list")
        pulumi.set(__self__, "token_body_parameters", token_body_parameters)
        if token_endpoint and not isinstance(token_endpoint, str):
            raise TypeError("Expected argument 'token_endpoint' to be a str")
        pulumi.set(__self__, "token_endpoint", token_endpoint)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> str:
        """
        OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
        """
        return pulumi.get(self, "authorization_endpoint")

    @property
    @pulumi.getter(name="authorizationMethods")
    def authorization_methods(self) -> Optional[List[str]]:
        """
        HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
        """
        return pulumi.get(self, "authorization_methods")

    @property
    @pulumi.getter(name="bearerTokenSendingMethods")
    def bearer_token_sending_methods(self) -> Optional[List[str]]:
        """
        Specifies the mechanism by which access token is passed to the API. 
        """
        return pulumi.get(self, "bearer_token_sending_methods")

    @property
    @pulumi.getter(name="clientAuthenticationMethod")
    def client_authentication_method(self) -> Optional[List[str]]:
        """
        Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
        """
        return pulumi.get(self, "client_authentication_method")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        Client or app id registered with this authorization server.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientRegistrationEndpoint")
    def client_registration_endpoint(self) -> str:
        """
        Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
        """
        return pulumi.get(self, "client_registration_endpoint")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[str]:
        """
        Client or app secret registered with this authorization server.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="defaultScope")
    def default_scope(self) -> Optional[str]:
        """
        Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
        """
        return pulumi.get(self, "default_scope")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the authorization server. Can contain HTML formatting tags.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        User-friendly authorization server name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="grantTypes")
    def grant_types(self) -> List[str]:
        """
        Form of an authorization grant, which the client uses to request the access token.
        """
        return pulumi.get(self, "grant_types")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceOwnerPassword")
    def resource_owner_password(self) -> Optional[str]:
        """
        Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
        """
        return pulumi.get(self, "resource_owner_password")

    @property
    @pulumi.getter(name="resourceOwnerUsername")
    def resource_owner_username(self) -> Optional[str]:
        """
        Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
        """
        return pulumi.get(self, "resource_owner_username")

    @property
    @pulumi.getter(name="supportState")
    def support_state(self) -> Optional[bool]:
        """
        If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
        """
        return pulumi.get(self, "support_state")

    @property
    @pulumi.getter(name="tokenBodyParameters")
    def token_body_parameters(self) -> Optional[List['outputs.TokenBodyParameterContractResponse']]:
        """
        Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
        """
        return pulumi.get(self, "token_body_parameters")

    @property
    @pulumi.getter(name="tokenEndpoint")
    def token_endpoint(self) -> Optional[str]:
        """
        OAuth token endpoint. Contains absolute URI to entity being referenced.
        """
        return pulumi.get(self, "token_endpoint")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetAuthorizationServerResult(GetAuthorizationServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthorizationServerResult(
            authorization_endpoint=self.authorization_endpoint,
            authorization_methods=self.authorization_methods,
            bearer_token_sending_methods=self.bearer_token_sending_methods,
            client_authentication_method=self.client_authentication_method,
            client_id=self.client_id,
            client_registration_endpoint=self.client_registration_endpoint,
            client_secret=self.client_secret,
            default_scope=self.default_scope,
            description=self.description,
            display_name=self.display_name,
            grant_types=self.grant_types,
            name=self.name,
            resource_owner_password=self.resource_owner_password,
            resource_owner_username=self.resource_owner_username,
            support_state=self.support_state,
            token_body_parameters=self.token_body_parameters,
            token_endpoint=self.token_endpoint,
            type=self.type)


def get_authorization_server(name: Optional[str] = None,
                             resource_group_name: Optional[str] = None,
                             service_name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthorizationServerResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Identifier of the authorization server.
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
    __ret__ = pulumi.runtime.invoke('azurerm:apimanagement/v20190101:getAuthorizationServer', __args__, opts=opts, typ=GetAuthorizationServerResult).value

    return AwaitableGetAuthorizationServerResult(
        authorization_endpoint=__ret__.authorization_endpoint,
        authorization_methods=__ret__.authorization_methods,
        bearer_token_sending_methods=__ret__.bearer_token_sending_methods,
        client_authentication_method=__ret__.client_authentication_method,
        client_id=__ret__.client_id,
        client_registration_endpoint=__ret__.client_registration_endpoint,
        client_secret=__ret__.client_secret,
        default_scope=__ret__.default_scope,
        description=__ret__.description,
        display_name=__ret__.display_name,
        grant_types=__ret__.grant_types,
        name=__ret__.name,
        resource_owner_password=__ret__.resource_owner_password,
        resource_owner_username=__ret__.resource_owner_username,
        support_state=__ret__.support_state,
        token_body_parameters=__ret__.token_body_parameters,
        token_endpoint=__ret__.token_endpoint,
        type=__ret__.type)
