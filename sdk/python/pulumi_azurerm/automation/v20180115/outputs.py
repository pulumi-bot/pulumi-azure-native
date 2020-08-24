# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'DscConfigurationAssociationPropertyResponse',
]

@pulumi.output_type
class DscConfigurationAssociationPropertyResponse(dict):
    """
    The Dsc configuration property associated with the entity.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        The Dsc configuration property associated with the entity.
        :param str name: Gets or sets the name of the Dsc configuration.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the name of the Dsc configuration.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


