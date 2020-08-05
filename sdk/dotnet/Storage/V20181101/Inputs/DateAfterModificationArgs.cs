// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20181101.Inputs
{

    /// <summary>
    /// Object to define the number of days after last modification.
    /// </summary>
    public sealed class DateAfterModificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value indicating the age in days after last modification
        /// </summary>
        [Input("daysAfterModificationGreaterThan", required: true)]
        public Input<int> DaysAfterModificationGreaterThan { get; set; } = null!;

        public DateAfterModificationArgs()
        {
        }
    }
}
