# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class RulesEngine(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the Rules Engine Configuration.
      * `resource_state` (`str`) - Resource status.
      * `rules` (`list`) - A list of rules that define a particular Rules Engine Configuration.
        * `action` (`dict`) - Actions to perform on the request and response if all of the match conditions are met.
          * `request_header_actions` (`list`) - A list of header actions to apply from the request from AFD to the origin.
            * `header_action_type` (`str`) - Which type of manipulation to apply to the header.
            * `header_name` (`str`) - The name of the header this action will apply to.
            * `value` (`str`) - The value to update the given header name with. This value is not used if the actionType is Delete.

          * `response_header_actions` (`list`) - A list of header actions to apply from the response from AFD to the client.

        * `match_conditions` (`list`) - A list of match conditions that must meet in order for the actions of this rule to run. Having no match conditions means the actions will always run.
          * `negate_condition` (`bool`) - Describes if this is negate condition or not
          * `rules_engine_match_value` (`list`) - Match values to match against. The operator will apply to each value in here with OR semantics. If any of them match the variable with the given operator this match condition is considered a match.
          * `rules_engine_match_variable` (`str`) - Match Variable
          * `rules_engine_operator` (`str`) - Describes operator to apply to the match condition.
          * `selector` (`str`) - Name of selector in RequestHeader or RequestBody to be matched
          * `transforms` (`list`) - List of transforms

        * `match_processing_behavior` (`str`) - If this rule is a match should the rules engine continue running the remaining rules or stop. If not present, defaults to Continue.
        * `name` (`str`) - A name to refer to this specific rule.
        * `priority` (`float`) - A priority assigned to this rule. 
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, front_door_name=None, name=None, resource_group_name=None, resource_state=None, rules=None, __props__=None, __name__=None, __opts__=None):
        """
        A rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] front_door_name: Name of the Front Door which is globally unique.
        :param pulumi.Input[str] name: Name of the Rules Engine which is unique within the Front Door.
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[str] resource_state: Resource status.
        :param pulumi.Input[list] rules: A list of rules that define a particular Rules Engine Configuration.

        The **rules** object supports the following:

          * `action` (`pulumi.Input[dict]`) - Actions to perform on the request and response if all of the match conditions are met.
            * `request_header_actions` (`pulumi.Input[list]`) - A list of header actions to apply from the request from AFD to the origin.
              * `header_action_type` (`pulumi.Input[str]`) - Which type of manipulation to apply to the header.
              * `header_name` (`pulumi.Input[str]`) - The name of the header this action will apply to.
              * `value` (`pulumi.Input[str]`) - The value to update the given header name with. This value is not used if the actionType is Delete.

            * `response_header_actions` (`pulumi.Input[list]`) - A list of header actions to apply from the response from AFD to the client.

          * `match_conditions` (`pulumi.Input[list]`) - A list of match conditions that must meet in order for the actions of this rule to run. Having no match conditions means the actions will always run.
            * `negate_condition` (`pulumi.Input[bool]`) - Describes if this is negate condition or not
            * `rules_engine_match_value` (`pulumi.Input[list]`) - Match values to match against. The operator will apply to each value in here with OR semantics. If any of them match the variable with the given operator this match condition is considered a match.
            * `rules_engine_match_variable` (`pulumi.Input[str]`) - Match Variable
            * `rules_engine_operator` (`pulumi.Input[str]`) - Describes operator to apply to the match condition.
            * `selector` (`pulumi.Input[str]`) - Name of selector in RequestHeader or RequestBody to be matched
            * `transforms` (`pulumi.Input[list]`) - List of transforms

          * `match_processing_behavior` (`pulumi.Input[str]`) - If this rule is a match should the rules engine continue running the remaining rules or stop. If not present, defaults to Continue.
          * `name` (`pulumi.Input[str]`) - A name to refer to this specific rule.
          * `priority` (`pulumi.Input[float]`) - A priority assigned to this rule. 
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

            if front_door_name is None:
                raise TypeError("Missing required property 'front_door_name'")
            __props__['front_door_name'] = front_door_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_state'] = resource_state
            __props__['rules'] = rules
            __props__['properties'] = None
            __props__['type'] = None
        super(RulesEngine, __self__).__init__(
            'azurerm:network/v20200501:RulesEngine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing RulesEngine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RulesEngine(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop