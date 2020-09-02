// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180715Preview.Outputs
{

    [OutputType]
    public sealed class ODataErrorResponseResult
    {
        /// <summary>
        /// The machine-readable description of the error, such as 'InvalidRequest' or 'InternalServerError'
        /// </summary>
        public readonly string? Code;
        /// <summary>
        /// Inner errors that caused this error
        /// </summary>
        public readonly ImmutableArray<Outputs.ODataErrorResponseResult> Details;
        /// <summary>
        /// The human-readable description of the error
        /// </summary>
        public readonly string? Message;

        [OutputConstructor]
        private ODataErrorResponseResult(
            string? code,

            ImmutableArray<Outputs.ODataErrorResponseResult> details,

            string? message)
        {
            Code = code;
            Details = details;
            Message = message;
        }
    }
}
