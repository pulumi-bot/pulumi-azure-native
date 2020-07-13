# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ServiceApi(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Api entity contract properties.
      * `api_revision` (`str`) - Describes the Revision of the Api. If no value is provided, default revision 1 is created
      * `api_revision_description` (`str`) - Description of the Api Revision.
      * `api_version` (`str`) - Indicates the Version identifier of the API if the API is versioned
      * `api_version_description` (`str`) - Description of the Api Version.
      * `api_version_set` (`dict`) - Version set details
        * `description` (`str`) - Description of API Version Set.
        * `id` (`str`) - Identifier for existing API Version Set. Omit this value to create a new Version Set.
        * `name` (`str`) - The display Name of the API Version Set.
        * `version_header_name` (`str`) - Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
        * `version_query_name` (`str`) - Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
        * `versioning_scheme` (`str`) - An value that determines where the API Version identifer will be located in a HTTP request.

      * `api_version_set_id` (`str`) - A resource identifier for the related ApiVersionSet.
      * `authentication_settings` (`dict`) - Collection of authentication settings included into this API.
        * `o_auth2` (`dict`) - OAuth2 Authentication settings
          * `authorization_server_id` (`str`) - OAuth authorization server identifier.
          * `scope` (`str`) - operations scope.

        * `openid` (`dict`) - OpenID Connect Authentication Settings
          * `bearer_token_sending_methods` (`list`) - How to send token to the server.
          * `openid_provider_id` (`str`) - OAuth authorization server identifier.

      * `description` (`str`) - Description of the API. May include HTML formatting tags.
      * `display_name` (`str`) - API name. Must be 1 to 300 characters long.
      * `is_current` (`bool`) - Indicates if API revision is current api revision.
      * `is_online` (`bool`) - Indicates if API revision is accessible via the gateway.
      * `path` (`str`) - Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
      * `protocols` (`list`) - Describes on which protocols the operations in this API can be invoked.
      * `service_url` (`str`) - Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
      * `source_api_id` (`str`) - API identifier of the source API.
      * `subscription_key_parameter_names` (`dict`) - Protocols over which API is made available.
        * `header` (`str`) - Subscription key header name.
        * `query` (`str`) - Subscription key query string parameter name.

      * `subscription_required` (`bool`) - Specifies whether an API or Product subscription is required for accessing the API.
      * `type` (`str`) - Type of API.
    """
    type: pulumi.Output[str]
    """
    Resource type for API Management resource.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, service_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Api details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        :param pulumi.Input[dict] properties: Api entity create of update properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.

        The **properties** object supports the following:

          * `api_revision` (`pulumi.Input[str]`) - Describes the Revision of the Api. If no value is provided, default revision 1 is created
          * `api_revision_description` (`pulumi.Input[str]`) - Description of the Api Revision.
          * `api_type` (`pulumi.Input[str]`) - Type of Api to create. 
             * `http` creates a SOAP to REST API 
             * `soap` creates a SOAP pass-through API .
          * `api_version` (`pulumi.Input[str]`) - Indicates the Version identifier of the API if the API is versioned
          * `api_version_description` (`pulumi.Input[str]`) - Description of the Api Version.
          * `api_version_set` (`pulumi.Input[dict]`) - Version set details
            * `description` (`pulumi.Input[str]`) - Description of API Version Set.
            * `id` (`pulumi.Input[str]`) - Identifier for existing API Version Set. Omit this value to create a new Version Set.
            * `name` (`pulumi.Input[str]`) - The display Name of the API Version Set.
            * `version_header_name` (`pulumi.Input[str]`) - Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
            * `version_query_name` (`pulumi.Input[str]`) - Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
            * `versioning_scheme` (`pulumi.Input[str]`) - An value that determines where the API Version identifer will be located in a HTTP request.

          * `api_version_set_id` (`pulumi.Input[str]`) - A resource identifier for the related ApiVersionSet.
          * `authentication_settings` (`pulumi.Input[dict]`) - Collection of authentication settings included into this API.
            * `o_auth2` (`pulumi.Input[dict]`) - OAuth2 Authentication settings
              * `authorization_server_id` (`pulumi.Input[str]`) - OAuth authorization server identifier.
              * `scope` (`pulumi.Input[str]`) - operations scope.

            * `openid` (`pulumi.Input[dict]`) - OpenID Connect Authentication Settings
              * `bearer_token_sending_methods` (`pulumi.Input[list]`) - How to send token to the server.
              * `openid_provider_id` (`pulumi.Input[str]`) - OAuth authorization server identifier.

          * `description` (`pulumi.Input[str]`) - Description of the API. May include HTML formatting tags.
          * `display_name` (`pulumi.Input[str]`) - API name. Must be 1 to 300 characters long.
          * `format` (`pulumi.Input[str]`) - Format of the Content in which the API is getting imported.
          * `is_current` (`pulumi.Input[bool]`) - Indicates if API revision is current api revision.
          * `path` (`pulumi.Input[str]`) - Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
          * `protocols` (`pulumi.Input[list]`) - Describes on which protocols the operations in this API can be invoked.
          * `service_url` (`pulumi.Input[str]`) - Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
          * `source_api_id` (`pulumi.Input[str]`) - API identifier of the source API.
          * `subscription_key_parameter_names` (`pulumi.Input[dict]`) - Protocols over which API is made available.
            * `header` (`pulumi.Input[str]`) - Subscription key header name.
            * `query` (`pulumi.Input[str]`) - Subscription key query string parameter name.

          * `subscription_required` (`pulumi.Input[bool]`) - Specifies whether an API or Product subscription is required for accessing the API.
          * `type` (`pulumi.Input[str]`) - Type of API.
          * `value` (`pulumi.Input[str]`) - Content value when Importing an API.
          * `wsdl_selector` (`pulumi.Input[dict]`) - Criteria to limit import of WSDL to a subset of the document.
            * `wsdl_endpoint_name` (`pulumi.Input[str]`) - Name of endpoint(port) to import from WSDL
            * `wsdl_service_name` (`pulumi.Input[str]`) - Name of service to import from WSDL
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['type'] = None
        super(ServiceApi, __self__).__init__(
            'azurerm:apimanagement:ServiceApi',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ServiceApi resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ServiceApi(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
