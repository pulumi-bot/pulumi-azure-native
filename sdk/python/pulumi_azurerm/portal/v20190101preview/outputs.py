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
    'DashboardLensResponse',
    'DashboardPartsResponse',
    'DashboardPartsResponsePosition',
]

@pulumi.output_type
class DashboardLensResponse(dict):
    """
    A dashboard lens.
    """
    def __init__(__self__, *,
                 order: float,
                 parts: Mapping[str, 'outputs.DashboardPartsResponse'],
                 metadata: Optional[Mapping[str, Mapping[str, Any]]] = None):
        """
        A dashboard lens.
        :param float order: The lens order.
        :param Mapping[str, 'DashboardPartsResponseArgs'] parts: The dashboard parts.
        :param Mapping[str, Mapping[str, Any]] metadata: The dashboard len's metadata.
        """
        pulumi.set(__self__, "order", order)
        pulumi.set(__self__, "parts", parts)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)

    @property
    @pulumi.getter
    def order(self) -> float:
        """
        The lens order.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter
    def parts(self) -> Mapping[str, 'outputs.DashboardPartsResponse']:
        """
        The dashboard parts.
        """
        return pulumi.get(self, "parts")

    @property
    @pulumi.getter
    def metadata(self) -> Optional[Mapping[str, Mapping[str, Any]]]:
        """
        The dashboard len's metadata.
        """
        return pulumi.get(self, "metadata")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DashboardPartsResponse(dict):
    """
    A dashboard part.
    """
    def __init__(__self__, *,
                 position: 'outputs.DashboardPartsResponsePosition',
                 metadata: Optional[Mapping[str, Mapping[str, Any]]] = None):
        """
        A dashboard part.
        :param 'DashboardPartsResponsePositionArgs' position: The dashboard's part position.
        :param Mapping[str, Mapping[str, Any]] metadata: The dashboard part's metadata.
        """
        pulumi.set(__self__, "position", position)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)

    @property
    @pulumi.getter
    def position(self) -> 'outputs.DashboardPartsResponsePosition':
        """
        The dashboard's part position.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def metadata(self) -> Optional[Mapping[str, Mapping[str, Any]]]:
        """
        The dashboard part's metadata.
        """
        return pulumi.get(self, "metadata")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DashboardPartsResponsePosition(dict):
    """
    The dashboard's part position.
    """
    def __init__(__self__, *,
                 col_span: float,
                 row_span: float,
                 x: float,
                 y: float,
                 metadata: Optional[Mapping[str, Mapping[str, Any]]] = None):
        """
        The dashboard's part position.
        :param float col_span: The dashboard's part column span.
        :param float row_span: The dashboard's part row span.
        :param float x: The dashboard's part x coordinate.
        :param float y: The dashboard's part y coordinate.
        :param Mapping[str, Mapping[str, Any]] metadata: The dashboard part's metadata.
        """
        pulumi.set(__self__, "col_span", col_span)
        pulumi.set(__self__, "row_span", row_span)
        pulumi.set(__self__, "x", x)
        pulumi.set(__self__, "y", y)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)

    @property
    @pulumi.getter(name="colSpan")
    def col_span(self) -> float:
        """
        The dashboard's part column span.
        """
        return pulumi.get(self, "col_span")

    @property
    @pulumi.getter(name="rowSpan")
    def row_span(self) -> float:
        """
        The dashboard's part row span.
        """
        return pulumi.get(self, "row_span")

    @property
    @pulumi.getter
    def x(self) -> float:
        """
        The dashboard's part x coordinate.
        """
        return pulumi.get(self, "x")

    @property
    @pulumi.getter
    def y(self) -> float:
        """
        The dashboard's part y coordinate.
        """
        return pulumi.get(self, "y")

    @property
    @pulumi.getter
    def metadata(self) -> Optional[Mapping[str, Mapping[str, Any]]]:
        """
        The dashboard part's metadata.
        """
        return pulumi.get(self, "metadata")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


