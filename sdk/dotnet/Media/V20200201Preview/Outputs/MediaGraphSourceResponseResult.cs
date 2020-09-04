// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20200201Preview.Outputs
{

    [OutputType]
    public sealed class MediaGraphSourceResponseResult
    {
        /// <summary>
        /// Source name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The discriminator for derived types.
        /// </summary>
        public readonly string OdataType;

        [OutputConstructor]
        private MediaGraphSourceResponseResult(
            string name,

            string odataType)
        {
            Name = name;
            OdataType = odataType;
        }
    }
}
