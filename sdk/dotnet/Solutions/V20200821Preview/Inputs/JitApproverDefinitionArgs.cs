// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Solutions.V20200821Preview.Inputs
{

    /// <summary>
    /// JIT approver definition.
    /// </summary>
    public sealed class JitApproverDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The approver display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The approver service principal Id.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The approver type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public JitApproverDefinitionArgs()
        {
        }
    }
}
