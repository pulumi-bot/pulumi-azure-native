// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class ExpressionRootResponseResult
    {
        /// <summary>
        /// The azure resource error info.
        /// </summary>
        public readonly Outputs.AzureResourceErrorInfoResponseResult? Error;
        /// <summary>
        /// The path.
        /// </summary>
        public readonly string? Path;
        public readonly ImmutableArray<Outputs.ExpressionResponseResult> Subexpressions;
        public readonly string? Text;
        public readonly ImmutableDictionary<string, object>? Value;

        [OutputConstructor]
        private ExpressionRootResponseResult(
            Outputs.AzureResourceErrorInfoResponseResult? error,

            string? path,

            ImmutableArray<Outputs.ExpressionResponseResult> subexpressions,

            string? text,

            ImmutableDictionary<string, object>? value)
        {
            Error = error;
            Path = path;
            Subexpressions = subexpressions;
            Text = text;
            Value = value;
        }
    }
}
