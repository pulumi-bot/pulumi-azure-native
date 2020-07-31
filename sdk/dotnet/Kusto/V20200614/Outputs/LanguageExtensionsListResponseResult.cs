// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20200614.Outputs
{

    [OutputType]
    public sealed class LanguageExtensionsListResponseResult
    {
        /// <summary>
        /// The list of language extensions.
        /// </summary>
        public readonly ImmutableArray<Outputs.LanguageExtensionResponseResult> Value;

        [OutputConstructor]
        private LanguageExtensionsListResponseResult(ImmutableArray<Outputs.LanguageExtensionResponseResult> value)
        {
            Value = value;
        }
    }
}