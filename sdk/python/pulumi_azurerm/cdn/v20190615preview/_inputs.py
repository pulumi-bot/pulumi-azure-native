# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'CustomRuleArgs',
    'CustomRuleListArgs',
    'DeepCreatedOriginArgs',
    'DeliveryRuleArgs',
    'DeliveryRuleActionArgs',
    'DeliveryRuleConditionArgs',
    'EndpointPropertiesUpdateParametersDeliveryPolicyArgs',
    'EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs',
    'GeoFilterArgs',
    'ManagedRuleGroupOverrideArgs',
    'ManagedRuleOverrideArgs',
    'ManagedRuleSetArgs',
    'ManagedRuleSetListArgs',
    'MatchConditionArgs',
    'RateLimitRuleArgs',
    'RateLimitRuleListArgs',
    'SkuArgs',
    'PolicySettingsArgs',
]

@pulumi.input_type
class CustomRuleArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 match_conditions: pulumi.Input[List[pulumi.Input['MatchConditionArgs']]],
                 name: pulumi.Input[str],
                 priority: pulumi.Input[float],
                 enabled_state: Optional[pulumi.Input[str]] = None):
        """
        Defines the common attributes for a custom rule that can be included in a waf policy
        :param pulumi.Input[str] action: Describes what action to be applied when rule matches
        :param pulumi.Input[List[pulumi.Input['MatchConditionArgs']]] match_conditions: List of match conditions.
        :param pulumi.Input[str] name: Defines the name of the custom rule
        :param pulumi.Input[float] priority: Defines in what order this rule be evaluated in the overall list of custom rules
        :param pulumi.Input[str] enabled_state: Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "match_conditions", match_conditions)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "priority", priority)
        if enabled_state is not None:
            pulumi.set(__self__, "enabled_state", enabled_state)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Describes what action to be applied when rule matches
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="matchConditions")
    def match_conditions(self) -> pulumi.Input[List[pulumi.Input['MatchConditionArgs']]]:
        """
        List of match conditions.
        """
        return pulumi.get(self, "match_conditions")

    @match_conditions.setter
    def match_conditions(self, value: pulumi.Input[List[pulumi.Input['MatchConditionArgs']]]):
        pulumi.set(self, "match_conditions", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Defines the name of the custom rule
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[float]:
        """
        Defines in what order this rule be evaluated in the overall list of custom rules
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[float]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="enabledState")
    def enabled_state(self) -> Optional[pulumi.Input[str]]:
        """
        Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
        """
        return pulumi.get(self, "enabled_state")

    @enabled_state.setter
    def enabled_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enabled_state", value)


@pulumi.input_type
class CustomRuleListArgs:
    def __init__(__self__, *,
                 rules: Optional[pulumi.Input[List[pulumi.Input['CustomRuleArgs']]]] = None):
        """
        Defines contents of custom rules
        :param pulumi.Input[List[pulumi.Input['CustomRuleArgs']]] rules: List of rules
        """
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[List[pulumi.Input['CustomRuleArgs']]]]:
        """
        List of rules
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[List[pulumi.Input['CustomRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class DeepCreatedOriginArgs:
    def __init__(__self__, *,
                 host_name: pulumi.Input[str],
                 name: pulumi.Input[str],
                 http_port: Optional[pulumi.Input[float]] = None,
                 https_port: Optional[pulumi.Input[float]] = None):
        """
        The main origin of CDN content which is added when creating a CDN endpoint.
        :param pulumi.Input[str] host_name: The address of the origin. It can be a domain name, IPv4 address, or IPv6 address.
        :param pulumi.Input[str] name: Origin name
        :param pulumi.Input[float] http_port: The value of the HTTP port. Must be between 1 and 65535
        :param pulumi.Input[float] https_port: The value of the HTTPS port. Must be between 1 and 65535
        """
        pulumi.set(__self__, "host_name", host_name)
        pulumi.set(__self__, "name", name)
        if http_port is not None:
            pulumi.set(__self__, "http_port", http_port)
        if https_port is not None:
            pulumi.set(__self__, "https_port", https_port)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> pulumi.Input[str]:
        """
        The address of the origin. It can be a domain name, IPv4 address, or IPv6 address.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Origin name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="httpPort")
    def http_port(self) -> Optional[pulumi.Input[float]]:
        """
        The value of the HTTP port. Must be between 1 and 65535
        """
        return pulumi.get(self, "http_port")

    @http_port.setter
    def http_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "http_port", value)

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> Optional[pulumi.Input[float]]:
        """
        The value of the HTTPS port. Must be between 1 and 65535
        """
        return pulumi.get(self, "https_port")

    @https_port.setter
    def https_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "https_port", value)


@pulumi.input_type
class DeliveryRuleArgs:
    def __init__(__self__, *,
                 actions: pulumi.Input[List[pulumi.Input['DeliveryRuleActionArgs']]],
                 order: pulumi.Input[float],
                 conditions: Optional[pulumi.Input[List[pulumi.Input['DeliveryRuleConditionArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        A rule that specifies a set of actions and conditions
        :param pulumi.Input[List[pulumi.Input['DeliveryRuleActionArgs']]] actions: A list of actions that are executed when all the conditions of a rule are satisfied.
        :param pulumi.Input[float] order: The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
        :param pulumi.Input[List[pulumi.Input['DeliveryRuleConditionArgs']]] conditions: A list of conditions that must be matched for the actions to be executed
        :param pulumi.Input[str] name: Name of the rule
        """
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "order", order)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Input[List[pulumi.Input['DeliveryRuleActionArgs']]]:
        """
        A list of actions that are executed when all the conditions of a rule are satisfied.
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: pulumi.Input[List[pulumi.Input['DeliveryRuleActionArgs']]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def order(self) -> pulumi.Input[float]:
        """
        The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: pulumi.Input[float]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input[List[pulumi.Input['DeliveryRuleConditionArgs']]]]:
        """
        A list of conditions that must be matched for the actions to be executed
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input[List[pulumi.Input['DeliveryRuleConditionArgs']]]]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the rule
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DeliveryRuleActionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        """
        An action for the delivery rule.
        :param pulumi.Input[str] name: The name of the action for the delivery rule.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the action for the delivery rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DeliveryRuleConditionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        """
        A condition for the delivery rule.
        :param pulumi.Input[str] name: The name of the condition for the delivery rule.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the condition for the delivery rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class EndpointPropertiesUpdateParametersDeliveryPolicyArgs:
    def __init__(__self__, *,
                 rules: pulumi.Input[List[pulumi.Input['DeliveryRuleArgs']]],
                 description: Optional[pulumi.Input[str]] = None):
        """
        A policy that specifies the delivery rules to be used for an endpoint.
        :param pulumi.Input[List[pulumi.Input['DeliveryRuleArgs']]] rules: A list of the delivery rules.
        :param pulumi.Input[str] description: User-friendly description of the policy.
        """
        pulumi.set(__self__, "rules", rules)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[List[pulumi.Input['DeliveryRuleArgs']]]:
        """
        A list of the delivery rules.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[List[pulumi.Input['DeliveryRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User-friendly description of the policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None):
        """
        Defines the Web Application Firewall policy for the endpoint (if applicable)
        :param pulumi.Input[str] id: Resource ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class GeoFilterArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 country_codes: pulumi.Input[List[pulumi.Input[str]]],
                 relative_path: pulumi.Input[str]):
        """
        Rules defining user's geo access within a CDN endpoint.
        :param pulumi.Input[str] action: Action of the geo filter, i.e. allow or block access.
        :param pulumi.Input[List[pulumi.Input[str]]] country_codes: Two letter country codes defining user country access in a geo filter, e.g. AU, MX, US.
        :param pulumi.Input[str] relative_path: Relative path applicable to geo filter. (e.g. '/mypictures', '/mypicture/kitty.jpg', and etc.)
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "country_codes", country_codes)
        pulumi.set(__self__, "relative_path", relative_path)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Action of the geo filter, i.e. allow or block access.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="countryCodes")
    def country_codes(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        Two letter country codes defining user country access in a geo filter, e.g. AU, MX, US.
        """
        return pulumi.get(self, "country_codes")

    @country_codes.setter
    def country_codes(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "country_codes", value)

    @property
    @pulumi.getter(name="relativePath")
    def relative_path(self) -> pulumi.Input[str]:
        """
        Relative path applicable to geo filter. (e.g. '/mypictures', '/mypicture/kitty.jpg', and etc.)
        """
        return pulumi.get(self, "relative_path")

    @relative_path.setter
    def relative_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "relative_path", value)


@pulumi.input_type
class ManagedRuleGroupOverrideArgs:
    def __init__(__self__, *,
                 rule_group_name: pulumi.Input[str],
                 rules: Optional[pulumi.Input[List[pulumi.Input['ManagedRuleOverrideArgs']]]] = None):
        """
        Defines a managed rule group override setting.
        :param pulumi.Input[str] rule_group_name: Describes the managed rule group within the rule set to override
        :param pulumi.Input[List[pulumi.Input['ManagedRuleOverrideArgs']]] rules: List of rules that will be disabled. If none specified, all rules in the group will be disabled.
        """
        pulumi.set(__self__, "rule_group_name", rule_group_name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="ruleGroupName")
    def rule_group_name(self) -> pulumi.Input[str]:
        """
        Describes the managed rule group within the rule set to override
        """
        return pulumi.get(self, "rule_group_name")

    @rule_group_name.setter
    def rule_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_group_name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[List[pulumi.Input['ManagedRuleOverrideArgs']]]]:
        """
        List of rules that will be disabled. If none specified, all rules in the group will be disabled.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[List[pulumi.Input['ManagedRuleOverrideArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class ManagedRuleOverrideArgs:
    def __init__(__self__, *,
                 rule_id: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 enabled_state: Optional[pulumi.Input[str]] = None):
        """
        Defines a managed rule group override setting.
        :param pulumi.Input[str] rule_id: Identifier for the managed rule.
        :param pulumi.Input[str] action: Describes the override action to be applied when rule matches.
        :param pulumi.Input[str] enabled_state: Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
        """
        pulumi.set(__self__, "rule_id", rule_id)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if enabled_state is not None:
            pulumi.set(__self__, "enabled_state", enabled_state)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Input[str]:
        """
        Identifier for the managed rule.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Describes the override action to be applied when rule matches.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="enabledState")
    def enabled_state(self) -> Optional[pulumi.Input[str]]:
        """
        Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
        """
        return pulumi.get(self, "enabled_state")

    @enabled_state.setter
    def enabled_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enabled_state", value)


@pulumi.input_type
class ManagedRuleSetArgs:
    def __init__(__self__, *,
                 rule_set_type: pulumi.Input[str],
                 rule_set_version: pulumi.Input[str],
                 anomaly_score: Optional[pulumi.Input[float]] = None,
                 rule_group_overrides: Optional[pulumi.Input[List[pulumi.Input['ManagedRuleGroupOverrideArgs']]]] = None):
        """
        Defines a managed rule set.
        :param pulumi.Input[str] rule_set_type: Defines the rule set type to use.
        :param pulumi.Input[str] rule_set_version: Defines the version of the rule set to use.
        :param pulumi.Input[float] anomaly_score: Verizon only : If the rule set supports anomaly detection mode, this describes the threshold for blocking requests.
        :param pulumi.Input[List[pulumi.Input['ManagedRuleGroupOverrideArgs']]] rule_group_overrides: Defines the rule overrides to apply to the rule set.
        """
        pulumi.set(__self__, "rule_set_type", rule_set_type)
        pulumi.set(__self__, "rule_set_version", rule_set_version)
        if anomaly_score is not None:
            pulumi.set(__self__, "anomaly_score", anomaly_score)
        if rule_group_overrides is not None:
            pulumi.set(__self__, "rule_group_overrides", rule_group_overrides)

    @property
    @pulumi.getter(name="ruleSetType")
    def rule_set_type(self) -> pulumi.Input[str]:
        """
        Defines the rule set type to use.
        """
        return pulumi.get(self, "rule_set_type")

    @rule_set_type.setter
    def rule_set_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_set_type", value)

    @property
    @pulumi.getter(name="ruleSetVersion")
    def rule_set_version(self) -> pulumi.Input[str]:
        """
        Defines the version of the rule set to use.
        """
        return pulumi.get(self, "rule_set_version")

    @rule_set_version.setter
    def rule_set_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_set_version", value)

    @property
    @pulumi.getter(name="anomalyScore")
    def anomaly_score(self) -> Optional[pulumi.Input[float]]:
        """
        Verizon only : If the rule set supports anomaly detection mode, this describes the threshold for blocking requests.
        """
        return pulumi.get(self, "anomaly_score")

    @anomaly_score.setter
    def anomaly_score(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "anomaly_score", value)

    @property
    @pulumi.getter(name="ruleGroupOverrides")
    def rule_group_overrides(self) -> Optional[pulumi.Input[List[pulumi.Input['ManagedRuleGroupOverrideArgs']]]]:
        """
        Defines the rule overrides to apply to the rule set.
        """
        return pulumi.get(self, "rule_group_overrides")

    @rule_group_overrides.setter
    def rule_group_overrides(self, value: Optional[pulumi.Input[List[pulumi.Input['ManagedRuleGroupOverrideArgs']]]]):
        pulumi.set(self, "rule_group_overrides", value)


@pulumi.input_type
class ManagedRuleSetListArgs:
    def __init__(__self__, *,
                 managed_rule_sets: Optional[pulumi.Input[List[pulumi.Input['ManagedRuleSetArgs']]]] = None):
        """
        Defines the list of managed rule sets for the policy.
        :param pulumi.Input[List[pulumi.Input['ManagedRuleSetArgs']]] managed_rule_sets: List of rule sets.
        """
        if managed_rule_sets is not None:
            pulumi.set(__self__, "managed_rule_sets", managed_rule_sets)

    @property
    @pulumi.getter(name="managedRuleSets")
    def managed_rule_sets(self) -> Optional[pulumi.Input[List[pulumi.Input['ManagedRuleSetArgs']]]]:
        """
        List of rule sets.
        """
        return pulumi.get(self, "managed_rule_sets")

    @managed_rule_sets.setter
    def managed_rule_sets(self, value: Optional[pulumi.Input[List[pulumi.Input['ManagedRuleSetArgs']]]]):
        pulumi.set(self, "managed_rule_sets", value)


@pulumi.input_type
class MatchConditionArgs:
    def __init__(__self__, *,
                 match_value: pulumi.Input[List[pulumi.Input[str]]],
                 match_variable: pulumi.Input[str],
                 operator: pulumi.Input[str],
                 negate_condition: Optional[pulumi.Input[bool]] = None,
                 selector: Optional[pulumi.Input[str]] = None,
                 transforms: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        Define match conditions
        :param pulumi.Input[List[pulumi.Input[str]]] match_value: List of possible match values.
        :param pulumi.Input[str] match_variable: Match variable to compare against.
        :param pulumi.Input[str] operator: Describes operator to be matched
        :param pulumi.Input[bool] negate_condition: Describes if the result of this condition should be negated.
        :param pulumi.Input[str] selector: Selector can used to match a specific key for QueryString, Cookies, RequestHeader or PostArgs.
        :param pulumi.Input[List[pulumi.Input[str]]] transforms: List of transforms.
        """
        pulumi.set(__self__, "match_value", match_value)
        pulumi.set(__self__, "match_variable", match_variable)
        pulumi.set(__self__, "operator", operator)
        if negate_condition is not None:
            pulumi.set(__self__, "negate_condition", negate_condition)
        if selector is not None:
            pulumi.set(__self__, "selector", selector)
        if transforms is not None:
            pulumi.set(__self__, "transforms", transforms)

    @property
    @pulumi.getter(name="matchValue")
    def match_value(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        List of possible match values.
        """
        return pulumi.get(self, "match_value")

    @match_value.setter
    def match_value(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "match_value", value)

    @property
    @pulumi.getter(name="matchVariable")
    def match_variable(self) -> pulumi.Input[str]:
        """
        Match variable to compare against.
        """
        return pulumi.get(self, "match_variable")

    @match_variable.setter
    def match_variable(self, value: pulumi.Input[str]):
        pulumi.set(self, "match_variable", value)

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Input[str]:
        """
        Describes operator to be matched
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: pulumi.Input[str]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter(name="negateCondition")
    def negate_condition(self) -> Optional[pulumi.Input[bool]]:
        """
        Describes if the result of this condition should be negated.
        """
        return pulumi.get(self, "negate_condition")

    @negate_condition.setter
    def negate_condition(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "negate_condition", value)

    @property
    @pulumi.getter
    def selector(self) -> Optional[pulumi.Input[str]]:
        """
        Selector can used to match a specific key for QueryString, Cookies, RequestHeader or PostArgs.
        """
        return pulumi.get(self, "selector")

    @selector.setter
    def selector(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "selector", value)

    @property
    @pulumi.getter
    def transforms(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        List of transforms.
        """
        return pulumi.get(self, "transforms")

    @transforms.setter
    def transforms(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "transforms", value)


@pulumi.input_type
class RateLimitRuleArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 match_conditions: pulumi.Input[List[pulumi.Input['MatchConditionArgs']]],
                 name: pulumi.Input[str],
                 priority: pulumi.Input[float],
                 rate_limit_duration_in_minutes: pulumi.Input[float],
                 rate_limit_threshold: pulumi.Input[float],
                 enabled_state: Optional[pulumi.Input[str]] = None):
        """
        Defines a rate limiting rule that can be included in a waf policy
        :param pulumi.Input[str] action: Describes what action to be applied when rule matches
        :param pulumi.Input[List[pulumi.Input['MatchConditionArgs']]] match_conditions: List of match conditions.
        :param pulumi.Input[str] name: Defines the name of the custom rule
        :param pulumi.Input[float] priority: Defines in what order this rule be evaluated in the overall list of custom rules
        :param pulumi.Input[float] rate_limit_duration_in_minutes: Defines rate limit duration. Default is 1 minute.
        :param pulumi.Input[float] rate_limit_threshold: Defines rate limit threshold.
        :param pulumi.Input[str] enabled_state: Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "match_conditions", match_conditions)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "rate_limit_duration_in_minutes", rate_limit_duration_in_minutes)
        pulumi.set(__self__, "rate_limit_threshold", rate_limit_threshold)
        if enabled_state is not None:
            pulumi.set(__self__, "enabled_state", enabled_state)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Describes what action to be applied when rule matches
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="matchConditions")
    def match_conditions(self) -> pulumi.Input[List[pulumi.Input['MatchConditionArgs']]]:
        """
        List of match conditions.
        """
        return pulumi.get(self, "match_conditions")

    @match_conditions.setter
    def match_conditions(self, value: pulumi.Input[List[pulumi.Input['MatchConditionArgs']]]):
        pulumi.set(self, "match_conditions", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Defines the name of the custom rule
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[float]:
        """
        Defines in what order this rule be evaluated in the overall list of custom rules
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[float]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="rateLimitDurationInMinutes")
    def rate_limit_duration_in_minutes(self) -> pulumi.Input[float]:
        """
        Defines rate limit duration. Default is 1 minute.
        """
        return pulumi.get(self, "rate_limit_duration_in_minutes")

    @rate_limit_duration_in_minutes.setter
    def rate_limit_duration_in_minutes(self, value: pulumi.Input[float]):
        pulumi.set(self, "rate_limit_duration_in_minutes", value)

    @property
    @pulumi.getter(name="rateLimitThreshold")
    def rate_limit_threshold(self) -> pulumi.Input[float]:
        """
        Defines rate limit threshold.
        """
        return pulumi.get(self, "rate_limit_threshold")

    @rate_limit_threshold.setter
    def rate_limit_threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "rate_limit_threshold", value)

    @property
    @pulumi.getter(name="enabledState")
    def enabled_state(self) -> Optional[pulumi.Input[str]]:
        """
        Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
        """
        return pulumi.get(self, "enabled_state")

    @enabled_state.setter
    def enabled_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enabled_state", value)


@pulumi.input_type
class RateLimitRuleListArgs:
    def __init__(__self__, *,
                 rules: Optional[pulumi.Input[List[pulumi.Input['RateLimitRuleArgs']]]] = None):
        """
        Defines contents of rate limit rules
        :param pulumi.Input[List[pulumi.Input['RateLimitRuleArgs']]] rules: List of rules
        """
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[List[pulumi.Input['RateLimitRuleArgs']]]]:
        """
        List of rules
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[List[pulumi.Input['RateLimitRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
        :param pulumi.Input[str] name: Name of the pricing tier.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the pricing tier.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PolicySettingsArgs:
    def __init__(__self__, *,
                 default_custom_block_response_body: Optional[pulumi.Input[str]] = None,
                 default_custom_block_response_status_code: Optional[pulumi.Input[float]] = None,
                 default_redirect_url: Optional[pulumi.Input[str]] = None,
                 enabled_state: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None):
        """
        Defines contents of a web application firewall global configuration
        :param pulumi.Input[str] default_custom_block_response_body: If the action type is block, customer can override the response body. The body must be specified in base64 encoding.
        :param pulumi.Input[float] default_custom_block_response_status_code: If the action type is block, this field defines the default customer overridable http response status code.
        :param pulumi.Input[str] default_redirect_url: If action type is redirect, this field represents the default redirect URL for the client.
        :param pulumi.Input[str] enabled_state: describes if the policy is in enabled state or disabled state
        :param pulumi.Input[str] mode: Describes if it is in detection mode or prevention mode at policy level.
        """
        if default_custom_block_response_body is not None:
            pulumi.set(__self__, "default_custom_block_response_body", default_custom_block_response_body)
        if default_custom_block_response_status_code is not None:
            pulumi.set(__self__, "default_custom_block_response_status_code", default_custom_block_response_status_code)
        if default_redirect_url is not None:
            pulumi.set(__self__, "default_redirect_url", default_redirect_url)
        if enabled_state is not None:
            pulumi.set(__self__, "enabled_state", enabled_state)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter(name="defaultCustomBlockResponseBody")
    def default_custom_block_response_body(self) -> Optional[pulumi.Input[str]]:
        """
        If the action type is block, customer can override the response body. The body must be specified in base64 encoding.
        """
        return pulumi.get(self, "default_custom_block_response_body")

    @default_custom_block_response_body.setter
    def default_custom_block_response_body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_custom_block_response_body", value)

    @property
    @pulumi.getter(name="defaultCustomBlockResponseStatusCode")
    def default_custom_block_response_status_code(self) -> Optional[pulumi.Input[float]]:
        """
        If the action type is block, this field defines the default customer overridable http response status code.
        """
        return pulumi.get(self, "default_custom_block_response_status_code")

    @default_custom_block_response_status_code.setter
    def default_custom_block_response_status_code(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "default_custom_block_response_status_code", value)

    @property
    @pulumi.getter(name="defaultRedirectUrl")
    def default_redirect_url(self) -> Optional[pulumi.Input[str]]:
        """
        If action type is redirect, this field represents the default redirect URL for the client.
        """
        return pulumi.get(self, "default_redirect_url")

    @default_redirect_url.setter
    def default_redirect_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_redirect_url", value)

    @property
    @pulumi.getter(name="enabledState")
    def enabled_state(self) -> Optional[pulumi.Input[str]]:
        """
        describes if the policy is in enabled state or disabled state
        """
        return pulumi.get(self, "enabled_state")

    @enabled_state.setter
    def enabled_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enabled_state", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Describes if it is in detection mode or prevention mode at policy level.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)


