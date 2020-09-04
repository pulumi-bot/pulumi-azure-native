# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'RegistrationInfoArgs',
]

@pulumi.input_type
class RegistrationInfoArgs:
    def __init__(__self__, *,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 registration_token_operation: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        Represents a RegistrationInfo definition.
        :param pulumi.Input[str] expiration_time: Expiration time of registration token.
        :param pulumi.Input[str] registration_token_operation: The type of resetting the token.
        :param pulumi.Input[str] token: The registration token base64 encoded string.
        """
        if expiration_time is not None:
            pulumi.set(__self__, "expiration_time", expiration_time)
        if registration_token_operation is not None:
            pulumi.set(__self__, "registration_token_operation", registration_token_operation)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration time of registration token.
        """
        return pulumi.get(self, "expiration_time")

    @expiration_time.setter
    def expiration_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_time", value)

    @property
    @pulumi.getter(name="registrationTokenOperation")
    def registration_token_operation(self) -> Optional[pulumi.Input[str]]:
        """
        The type of resetting the token.
        """
        return pulumi.get(self, "registration_token_operation")

    @registration_token_operation.setter
    def registration_token_operation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registration_token_operation", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The registration token base64 encoded string.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


