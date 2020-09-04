# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['IntegrationAccountMap']


class IntegrationAccountMap(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 integration_account_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 map_name: Optional[pulumi.Input[str]] = None,
                 map_type: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 parameters_schema: Optional[pulumi.Input[pulumi.InputType['IntegrationAccountMapPropertiesParametersSchemaArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The integration account map.

        ## Example Usage
        ### Create or update a map

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        integration_account_map = azurerm.logic.latest.IntegrationAccountMap("integrationAccountMap",
            content=\"\"\"<?xml version="1.0" encoding="UTF-16"?>
        <xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:msxsl="urn:schemas-microsoft-com:xslt" xmlns:var="http://schemas.microsoft.com/BizTalk/2003/var" exclude-result-prefixes="msxsl var s0 userCSharp" version="1.0" xmlns:ns0="http://BizTalk_Server_Project4.StringFunctoidsDestinationSchema" xmlns:s0="http://BizTalk_Server_Project4.StringFunctoidsSourceSchema" xmlns:userCSharp="http://schemas.microsoft.com/BizTalk/2003/userCSharp">
          <xsl:import href="http://btsfunctoids.blob.core.windows.net/functoids/functoids.xslt" />
          <xsl:output omit-xml-declaration="yes" method="xml" version="1.0" />
          <xsl:template match="/">
            <xsl:apply-templates select="/s0:Root" />
          </xsl:template>
          <xsl:template match="/s0:Root">
            <xsl:variable name="var:v1" select="userCSharp:StringFind(string(StringFindSource/text()) , &quot;SearchString&quot;)" />
            <xsl:variable name="var:v2" select="userCSharp:StringLeft(string(StringLeftSource/text()) , &quot;2&quot;)" />
            <xsl:variable name="var:v3" select="userCSharp:StringRight(string(StringRightSource/text()) , &quot;2&quot;)" />
            <xsl:variable name="var:v4" select="userCSharp:StringUpperCase(string(UppercaseSource/text()))" />
            <xsl:variable name="var:v5" select="userCSharp:StringLowerCase(string(LowercaseSource/text()))" />
            <xsl:variable name="var:v6" select="userCSharp:StringSize(string(SizeSource/text()))" />
            <xsl:variable name="var:v7" select="userCSharp:StringSubstring(string(StringExtractSource/text()) , &quot;0&quot; , &quot;2&quot;)" />
            <xsl:variable name="var:v8" select="userCSharp:StringConcat(string(StringConcatSource/text()))" />
            <xsl:variable name="var:v9" select="userCSharp:StringTrimLeft(string(StringLeftTrimSource/text()))" />
            <xsl:variable name="var:v10" select="userCSharp:StringTrimRight(string(StringRightTrimSource/text()))" />
            <ns0:Root>
              <StringFindDestination>
                <xsl:value-of select="$var:v1" />
              </StringFindDestination>
              <StringLeftDestination>
                <xsl:value-of select="$var:v2" />
              </StringLeftDestination>
              <StringRightDestination>
                <xsl:value-of select="$var:v3" />
              </StringRightDestination>
              <UppercaseDestination>
                <xsl:value-of select="$var:v4" />
              </UppercaseDestination>
              <LowercaseDestination>
                <xsl:value-of select="$var:v5" />
              </LowercaseDestination>
              <SizeDestination>
                <xsl:value-of select="$var:v6" />
              </SizeDestination>
              <StringExtractDestination>
                <xsl:value-of select="$var:v7" />
              </StringExtractDestination>
              <StringConcatDestination>
                <xsl:value-of select="$var:v8" />
              </StringConcatDestination>
              <StringLeftTrimDestination>
                <xsl:value-of select="$var:v9" />
              </StringLeftTrimDestination>
              <StringRightTrimDestination>
                <xsl:value-of select="$var:v10" />
              </StringRightTrimDestination>
            </ns0:Root>
          </xsl:template>
        </xsl:stylesheet>\"\"\",
            content_type="application/xml",
            integration_account_name="testIntegrationAccount",
            location="westus",
            map_name="testMap",
            map_type="Xslt",
            metadata={},
            resource_group_name="testResourceGroup")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: The content.
        :param pulumi.Input[str] content_type: The content type.
        :param pulumi.Input[str] integration_account_name: The integration account name.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[str] map_name: The integration account map name.
        :param pulumi.Input[str] map_type: The map type.
        :param pulumi.Input[Mapping[str, Any]] metadata: The metadata.
        :param pulumi.Input[pulumi.InputType['IntegrationAccountMapPropertiesParametersSchemaArgs']] parameters_schema: The parameters schema of integration account map.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The resource tags.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['content'] = content
            __props__['content_type'] = content_type
            if integration_account_name is None:
                raise TypeError("Missing required property 'integration_account_name'")
            __props__['integration_account_name'] = integration_account_name
            __props__['location'] = location
            if map_name is None:
                raise TypeError("Missing required property 'map_name'")
            __props__['map_name'] = map_name
            if map_type is None:
                raise TypeError("Missing required property 'map_type'")
            __props__['map_type'] = map_type
            __props__['metadata'] = metadata
            __props__['parameters_schema'] = parameters_schema
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['changed_time'] = None
            __props__['content_link'] = None
            __props__['created_time'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:logic/v20150801preview:IntegrationAccountMap"), pulumi.Alias(type_="azurerm:logic/v20160601:IntegrationAccountMap"), pulumi.Alias(type_="azurerm:logic/v20180701preview:IntegrationAccountMap"), pulumi.Alias(type_="azurerm:logic/v20190501:IntegrationAccountMap")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(IntegrationAccountMap, __self__).__init__(
            'azurerm:logic/latest:IntegrationAccountMap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IntegrationAccountMap':
        """
        Get an existing IntegrationAccountMap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IntegrationAccountMap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="changedTime")
    def changed_time(self) -> pulumi.Output[str]:
        """
        The changed time.
        """
        return pulumi.get(self, "changed_time")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[Optional[str]]:
        """
        The content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentLink")
    def content_link(self) -> pulumi.Output['outputs.ContentLinkResponse']:
        """
        The content link.
        """
        return pulumi.get(self, "content_link")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[Optional[str]]:
        """
        The content type.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        The created time.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mapType")
    def map_type(self) -> pulumi.Output[str]:
        """
        The map type.
        """
        return pulumi.get(self, "map_type")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Gets the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parametersSchema")
    def parameters_schema(self) -> pulumi.Output[Optional['outputs.IntegrationAccountMapPropertiesResponseParametersSchema']]:
        """
        The parameters schema of integration account map.
        """
        return pulumi.get(self, "parameters_schema")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Gets the resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

