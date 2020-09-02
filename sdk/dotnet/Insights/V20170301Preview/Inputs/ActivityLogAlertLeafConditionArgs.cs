// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20170301Preview.Inputs
{

    /// <summary>
    /// An Activity Log alert condition that is met by comparing an activity log field and value.
    /// </summary>
    public sealed class ActivityLogAlertLeafConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field value will be compared to this value (case-insensitive) to determine if the condition is met.
        /// </summary>
        [Input("equals", required: true)]
        public Input<string> Equals { get; set; } = null!;

        /// <summary>
        /// The name of the field that this condition will examine. The possible values for this field are (case-insensitive): 'resourceId', 'category', 'caller', 'level', 'operationName', 'resourceGroup', 'resourceProvider', 'status', 'subStatus', 'resourceType', or anything beginning with 'properties.'.
        /// </summary>
        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        public ActivityLogAlertLeafConditionArgs()
        {
        }
    }
}
