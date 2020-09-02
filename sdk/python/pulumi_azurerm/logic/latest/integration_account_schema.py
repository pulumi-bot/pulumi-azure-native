# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['IntegrationAccountSchema']


class IntegrationAccountSchema(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 document_name: Optional[pulumi.Input[str]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 integration_account_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 schema_name: Optional[pulumi.Input[str]] = None,
                 schema_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The integration account schema.

        ## Create or update schema

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        integration_account_schema = azurerm.logic.latest.IntegrationAccountSchema("integrationAccountSchema",
            content=\"\"\"<?xml version="1.0" encoding="utf-16"?>
        <xs:schema xmlns:b="http://schemas.microsoft.com/BizTalk/2003" xmlns="http://Inbound_EDI.OrderFile" targetNamespace="http://Inbound_EDI.OrderFile" xmlns:xs="http://www.w3.org/2001/XMLSchema">
          <xs:annotation>
            <xs:appinfo>
              <b:schemaInfo default_pad_char=" " count_positions_by_byte="false" parser_optimization="speed" lookahead_depth="3" suppress_empty_nodes="false" generate_empty_nodes="true" allow_early_termination="false" early_terminate_optional_fields="false" allow_message_breakup_of_infix_root="false" compile_parse_tables="false" standard="Flat File" root_reference="OrderFile" />
              <schemaEditorExtension:schemaInfo namespaceAlias="b" extensionClass="Microsoft.BizTalk.FlatFileExtension.FlatFileExtension" standardName="Flat File" xmlns:schemaEditorExtension="http://schemas.microsoft.com/BizTalk/2003/SchemaEditorExtensions" />
            </xs:appinfo>
          </xs:annotation>
          <xs:element name="OrderFile">
            <xs:annotation>
              <xs:appinfo>
                <b:recordInfo structure="delimited" preserve_delimiter_for_empty_data="true" suppress_trailing_delimiters="false" sequence_number="1" />
              </xs:appinfo>
            </xs:annotation>
            <xs:complexType>
              <xs:sequence>
                <xs:annotation>
                  <xs:appinfo>
                    <b:groupInfo sequence_number="0" />
                  </xs:appinfo>
                </xs:annotation>
                <xs:element name="Order">
                  <xs:annotation>
                    <xs:appinfo>
                      <b:recordInfo sequence_number="1" structure="delimited" preserve_delimiter_for_empty_data="true" suppress_trailing_delimiters="false" child_delimiter_type="hex" child_delimiter="0x0D 0x0A" child_order="infix" />
                    </xs:appinfo>
                  </xs:annotation>
                  <xs:complexType>
                    <xs:sequence>
                      <xs:annotation>
                        <xs:appinfo>
                          <b:groupInfo sequence_number="0" />
                        </xs:appinfo>
                      </xs:annotation>
                      <xs:element name="Header">
                        <xs:annotation>
                          <xs:appinfo>
                            <b:recordInfo sequence_number="1" structure="delimited" preserve_delimiter_for_empty_data="true" suppress_trailing_delimiters="false" child_delimiter_type="char" child_delimiter="|" child_order="infix" tag_name="HDR|" />
                          </xs:appinfo>
                        </xs:annotation>
                        <xs:complexType>
                          <xs:sequence>
                            <xs:annotation>
                              <xs:appinfo>
                                <b:groupInfo sequence_number="0" />
                              </xs:appinfo>
                            </xs:annotation>
                            <xs:element name="PODate" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="1" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="PONumber" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo justification="left" sequence_number="2" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="CustomerID" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="3" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="CustomerContactName" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="4" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="CustomerContactPhone" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="5" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                          </xs:sequence>
                        </xs:complexType>
                      </xs:element>
                      <xs:element minOccurs="1" maxOccurs="unbounded" name="LineItems">
                        <xs:annotation>
                          <xs:appinfo>
                            <b:recordInfo sequence_number="2" structure="delimited" preserve_delimiter_for_empty_data="true" suppress_trailing_delimiters="false" child_delimiter_type="char" child_delimiter="|" child_order="infix" tag_name="DTL|" />
                          </xs:appinfo>
                        </xs:annotation>
                        <xs:complexType>
                          <xs:sequence>
                            <xs:annotation>
                              <xs:appinfo>
                                <b:groupInfo sequence_number="0" />
                              </xs:appinfo>
                            </xs:annotation>
                            <xs:element name="PONumber" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="1" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="ItemOrdered" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="2" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="Quantity" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="3" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="UOM" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="4" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="Price" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="5" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="ExtendedPrice" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="6" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                            <xs:element name="Description" type="xs:string">
                              <xs:annotation>
                                <xs:appinfo>
                                  <b:fieldInfo sequence_number="7" justification="left" />
                                </xs:appinfo>
                              </xs:annotation>
                            </xs:element>
                          </xs:sequence>
                        </xs:complexType>
                      </xs:element>
                    </xs:sequence>
                  </xs:complexType>
                </xs:element>
              </xs:sequence>
            </xs:complexType>
          </xs:element>
        </xs:schema>\"\"\",
            content_type="application/xml",
            integration_account_name="testIntegrationAccount",
            location="westus",
            metadata={},
            resource_group_name="testResourceGroup",
            schema_name="testSchema",
            schema_type="Xml",
            tags={
                "integrationAccountSchemaName": "IntegrationAccountSchema8120",
            })

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: The content.
        :param pulumi.Input[str] content_type: The content type.
        :param pulumi.Input[str] document_name: The document name.
        :param pulumi.Input[str] file_name: The file name.
        :param pulumi.Input[str] integration_account_name: The integration account name.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[Mapping[str, Any]] metadata: The metadata.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[str] schema_name: The integration account schema name.
        :param pulumi.Input[str] schema_type: The schema type.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The resource tags.
        :param pulumi.Input[str] target_namespace: The target namespace of the schema.
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
            __props__['document_name'] = document_name
            __props__['file_name'] = file_name
            if integration_account_name is None:
                raise TypeError("Missing required property 'integration_account_name'")
            __props__['integration_account_name'] = integration_account_name
            __props__['location'] = location
            __props__['metadata'] = metadata
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if schema_name is None:
                raise TypeError("Missing required property 'schema_name'")
            __props__['schema_name'] = schema_name
            if schema_type is None:
                raise TypeError("Missing required property 'schema_type'")
            __props__['schema_type'] = schema_type
            __props__['tags'] = tags
            __props__['target_namespace'] = target_namespace
            __props__['changed_time'] = None
            __props__['content_link'] = None
            __props__['created_time'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:logic/v20160601:IntegrationAccountSchema"), pulumi.Alias(type_="azurerm:logic/v20190501:IntegrationAccountSchema")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(IntegrationAccountSchema, __self__).__init__(
            'azurerm:logic/latest:IntegrationAccountSchema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IntegrationAccountSchema':
        """
        Get an existing IntegrationAccountSchema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IntegrationAccountSchema(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="documentName")
    def document_name(self) -> pulumi.Output[Optional[str]]:
        """
        The document name.
        """
        return pulumi.get(self, "document_name")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> pulumi.Output[Optional[str]]:
        """
        The file name.
        """
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The resource location.
        """
        return pulumi.get(self, "location")

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
    @pulumi.getter(name="schemaType")
    def schema_type(self) -> pulumi.Output[str]:
        """
        The schema type.
        """
        return pulumi.get(self, "schema_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetNamespace")
    def target_namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The target namespace of the schema.
        """
        return pulumi.get(self, "target_namespace")

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
