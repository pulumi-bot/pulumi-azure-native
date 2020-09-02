// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AlertsManagement.V20190505Preview.Outputs
{

    [OutputType]
    public sealed class ConditionsResponseResult
    {
        /// <summary>
        /// filter alerts by alert context (payload)
        /// </summary>
        public readonly Outputs.ConditionResponseResult? AlertContext;
        /// <summary>
        /// filter alerts by alert rule id
        /// </summary>
        public readonly Outputs.ConditionResponseResult? AlertRuleId;
        /// <summary>
        /// filter alerts by alert rule description
        /// </summary>
        public readonly Outputs.ConditionResponseResult? Description;
        /// <summary>
        /// filter alerts by monitor condition
        /// </summary>
        public readonly Outputs.ConditionResponseResult? MonitorCondition;
        /// <summary>
        /// filter alerts by monitor service
        /// </summary>
        public readonly Outputs.ConditionResponseResult? MonitorService;
        /// <summary>
        /// filter alerts by severity
        /// </summary>
        public readonly Outputs.ConditionResponseResult? Severity;
        /// <summary>
        /// filter alerts by target resource type
        /// </summary>
        public readonly Outputs.ConditionResponseResult? TargetResourceType;

        [OutputConstructor]
        private ConditionsResponseResult(
            Outputs.ConditionResponseResult? alertContext,

            Outputs.ConditionResponseResult? alertRuleId,

            Outputs.ConditionResponseResult? description,

            Outputs.ConditionResponseResult? monitorCondition,

            Outputs.ConditionResponseResult? monitorService,

            Outputs.ConditionResponseResult? severity,

            Outputs.ConditionResponseResult? targetResourceType)
        {
            AlertContext = alertContext;
            AlertRuleId = alertRuleId;
            Description = description;
            MonitorCondition = monitorCondition;
            MonitorService = monitorService;
            Severity = severity;
            TargetResourceType = targetResourceType;
        }
    }
}
