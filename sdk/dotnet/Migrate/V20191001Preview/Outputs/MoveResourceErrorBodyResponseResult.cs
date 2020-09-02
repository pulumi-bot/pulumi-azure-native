// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate.V20191001Preview.Outputs
{

    [OutputType]
    public sealed class MoveResourceErrorBodyResponseResult
    {
        /// <summary>
        /// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// A list of additional details about the error.
        /// </summary>
        public readonly ImmutableArray<Outputs.MoveResourceErrorBodyResponseResult> Details;
        /// <summary>
        /// A message describing the error, intended to be suitable for display in a user interface.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// The target of the particular error. For example, the name of the property in error.
        /// </summary>
        public readonly string Target;

        [OutputConstructor]
        private MoveResourceErrorBodyResponseResult(
            string code,

            ImmutableArray<Outputs.MoveResourceErrorBodyResponseResult> details,

            string message,

            string target)
        {
            Code = code;
            Details = details;
            Message = message;
            Target = target;
        }
    }
}
