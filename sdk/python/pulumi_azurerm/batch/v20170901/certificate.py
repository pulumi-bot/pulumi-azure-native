# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['Certificate']


class Certificate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 thumbprint: Optional[pulumi.Input[str]] = None,
                 thumbprint_algorithm: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Contains information about a certificate.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the Batch account.
        :param pulumi.Input[str] data: The maximum size is 10KB.
        :param pulumi.Input[str] format: The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
        :param pulumi.Input[str] name: The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
        :param pulumi.Input[str] password: This is required if the certificate format is pfx and must be omitted if the certificate format is cer.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the Batch account.
        :param pulumi.Input[str] thumbprint: This must match the thumbprint from the name.
        :param pulumi.Input[str] thumbprint_algorithm: This must match the first portion of the certificate name. Currently required to be 'SHA1'.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if data is None:
                raise TypeError("Missing required property 'data'")
            __props__['data'] = data
            __props__['format'] = format
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['password'] = password
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['thumbprint'] = thumbprint
            __props__['thumbprint_algorithm'] = thumbprint_algorithm
            __props__['delete_certificate_error'] = None
            __props__['etag'] = None
            __props__['previous_provisioning_state'] = None
            __props__['previous_provisioning_state_transition_time'] = None
            __props__['provisioning_state'] = None
            __props__['provisioning_state_transition_time'] = None
            __props__['public_data'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:batch/v20181201:Certificate"), pulumi.Alias(type_="azurerm:batch/v20190401:Certificate"), pulumi.Alias(type_="azurerm:batch/v20190801:Certificate"), pulumi.Alias(type_="azurerm:batch/v20200301:Certificate"), pulumi.Alias(type_="azurerm:batch/v20200501:Certificate")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Certificate, __self__).__init__(
            'azurerm:batch/v20170901:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Certificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deleteCertificateError")
    def delete_certificate_error(self) -> 'outputs.DeleteCertificateErrorResponse':
        """
        This is only returned when the certificate provisioningState is 'Failed'.
        """
        return pulumi.get(self, "delete_certificate_error")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        The ETag of the resource, used for concurrency statements.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def format(self) -> Optional[str]:
        """
        The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="previousProvisioningState")
    def previous_provisioning_state(self) -> str:
        """
        The previous provisioned state of the resource
        """
        return pulumi.get(self, "previous_provisioning_state")

    @property
    @pulumi.getter(name="previousProvisioningStateTransitionTime")
    def previous_provisioning_state_transition_time(self) -> str:
        return pulumi.get(self, "previous_provisioning_state_transition_time")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Values are:

         Succeeded - The certificate is available for use in pools.
         Deleting - The user has requested that the certificate be deleted, but the delete operation has not yet completed. You may not reference the certificate when creating or updating pools.
         Failed - The user requested that the certificate be deleted, but there are pools that still have references to the certificate, or it is still installed on one or more compute nodes. (The latter can occur if the certificate has been removed from the pool, but the node has not yet restarted. Nodes refresh their certificates only when they restart.) You may use the cancel certificate delete operation to cancel the delete, or the delete certificate operation to retry the delete.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="provisioningStateTransitionTime")
    def provisioning_state_transition_time(self) -> str:
        return pulumi.get(self, "provisioning_state_transition_time")

    @property
    @pulumi.getter(name="publicData")
    def public_data(self) -> str:
        """
        The public key of the certificate.
        """
        return pulumi.get(self, "public_data")

    @property
    @pulumi.getter
    def thumbprint(self) -> Optional[str]:
        """
        This must match the thumbprint from the name.
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter(name="thumbprintAlgorithm")
    def thumbprint_algorithm(self) -> Optional[str]:
        """
        This must match the first portion of the certificate name. Currently required to be 'SHA1'.
        """
        return pulumi.get(self, "thumbprint_algorithm")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

