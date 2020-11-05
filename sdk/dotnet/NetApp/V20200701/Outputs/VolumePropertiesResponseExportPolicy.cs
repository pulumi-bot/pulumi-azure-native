// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.NetApp.V20200701.Outputs
{

    [OutputType]
    public sealed class VolumePropertiesResponseExportPolicy
    {
        /// <summary>
        /// Export policy rule
        /// </summary>
        public readonly ImmutableArray<Outputs.ExportPolicyRuleResponse> Rules;

        [OutputConstructor]
        private VolumePropertiesResponseExportPolicy(ImmutableArray<Outputs.ExportPolicyRuleResponse> rules)
        {
            Rules = rules;
        }
    }
}
