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
    'GetStreamingPolicyResult',
    'AwaitableGetStreamingPolicyResult',
    'get_streaming_policy',
]

@pulumi.output_type
class GetStreamingPolicyResult:
    """
    A Streaming Policy resource
    """
    def __init__(__self__, common_encryption_cbcs=None, common_encryption_cenc=None, created=None, default_content_key_policy_name=None, envelope_encryption=None, name=None, no_encryption=None, type=None):
        if common_encryption_cbcs and not isinstance(common_encryption_cbcs, dict):
            raise TypeError("Expected argument 'common_encryption_cbcs' to be a dict")
        pulumi.set(__self__, "common_encryption_cbcs", common_encryption_cbcs)
        if common_encryption_cenc and not isinstance(common_encryption_cenc, dict):
            raise TypeError("Expected argument 'common_encryption_cenc' to be a dict")
        pulumi.set(__self__, "common_encryption_cenc", common_encryption_cenc)
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if default_content_key_policy_name and not isinstance(default_content_key_policy_name, str):
            raise TypeError("Expected argument 'default_content_key_policy_name' to be a str")
        pulumi.set(__self__, "default_content_key_policy_name", default_content_key_policy_name)
        if envelope_encryption and not isinstance(envelope_encryption, dict):
            raise TypeError("Expected argument 'envelope_encryption' to be a dict")
        pulumi.set(__self__, "envelope_encryption", envelope_encryption)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if no_encryption and not isinstance(no_encryption, dict):
            raise TypeError("Expected argument 'no_encryption' to be a dict")
        pulumi.set(__self__, "no_encryption", no_encryption)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="commonEncryptionCbcs")
    def common_encryption_cbcs(self) -> Optional['outputs.CommonEncryptionCbcsResponse']:
        """
        Configuration of CommonEncryptionCbcs
        """
        return pulumi.get(self, "common_encryption_cbcs")

    @property
    @pulumi.getter(name="commonEncryptionCenc")
    def common_encryption_cenc(self) -> Optional['outputs.CommonEncryptionCencResponse']:
        """
        Configuration of CommonEncryptionCenc
        """
        return pulumi.get(self, "common_encryption_cenc")

    @property
    @pulumi.getter
    def created(self) -> str:
        """
        Creation time of Streaming Policy
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="defaultContentKeyPolicyName")
    def default_content_key_policy_name(self) -> Optional[str]:
        """
        Default ContentKey used by current Streaming Policy
        """
        return pulumi.get(self, "default_content_key_policy_name")

    @property
    @pulumi.getter(name="envelopeEncryption")
    def envelope_encryption(self) -> Optional['outputs.EnvelopeEncryptionResponse']:
        """
        Configuration of EnvelopeEncryption
        """
        return pulumi.get(self, "envelope_encryption")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noEncryption")
    def no_encryption(self) -> Optional['outputs.NoEncryptionResponse']:
        """
        Configurations of NoEncryption
        """
        return pulumi.get(self, "no_encryption")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetStreamingPolicyResult(GetStreamingPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStreamingPolicyResult(
            common_encryption_cbcs=self.common_encryption_cbcs,
            common_encryption_cenc=self.common_encryption_cenc,
            created=self.created,
            default_content_key_policy_name=self.default_content_key_policy_name,
            envelope_encryption=self.envelope_encryption,
            name=self.name,
            no_encryption=self.no_encryption,
            type=self.type)


def get_streaming_policy(account_name: Optional[str] = None,
                         resource_group_name: Optional[str] = None,
                         streaming_policy_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStreamingPolicyResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: The Media Services account name.
    :param str resource_group_name: The name of the resource group within the Azure subscription.
    :param str streaming_policy_name: The Streaming Policy name.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['streamingPolicyName'] = streaming_policy_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:media/v20180601preview:getStreamingPolicy', __args__, opts=opts, typ=GetStreamingPolicyResult).value

    return AwaitableGetStreamingPolicyResult(
        common_encryption_cbcs=__ret__.common_encryption_cbcs,
        common_encryption_cenc=__ret__.common_encryption_cenc,
        created=__ret__.created,
        default_content_key_policy_name=__ret__.default_content_key_policy_name,
        envelope_encryption=__ret__.envelope_encryption,
        name=__ret__.name,
        no_encryption=__ret__.no_encryption,
        type=__ret__.type)
