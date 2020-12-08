// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20190101Preview.Inputs
{

    /// <summary>
    /// The result of the assessment
    /// </summary>
    public sealed class AssessmentStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Programmatic code for the cause of the assessment status
        /// </summary>
        [Input("cause")]
        public Input<string>? Cause { get; set; }

        /// <summary>
        /// Programmatic code for the status of the assessment
        /// </summary>
        [Input("code", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Security.V20190101Preview.AssessmentStatusCode> Code { get; set; } = null!;

        /// <summary>
        /// Human readable description of the assessment status
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        public AssessmentStatusArgs()
        {
        }
    }
}
