// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Resources.V20201001.Inputs
{

    /// <summary>
    /// Specifies whether template expressions are evaluated within the scope of the parent template or nested template.
    /// </summary>
    public sealed class ExpressionEvaluationOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The scope to be used for evaluation of parameters, variables and functions in a nested template.
        /// </summary>
        [Input("scope")]
        public InputUnion<string, Pulumi.AzureNextGen.Resources.V20201001.ExpressionEvaluationOptionsScopeType>? Scope { get; set; }

        public ExpressionEvaluationOptionsArgs()
        {
        }
    }
}
