// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501.Outputs
{

    [OutputType]
    public sealed class MatchVariableResponseResult
    {
        /// <summary>
        /// The selector of match variable.
        /// </summary>
        public readonly string? Selector;
        /// <summary>
        /// Match Variable.
        /// </summary>
        public readonly string VariableName;

        [OutputConstructor]
        private MatchVariableResponseResult(
            string? selector,

            string variableName)
        {
            Selector = selector;
            VariableName = variableName;
        }
    }
}
