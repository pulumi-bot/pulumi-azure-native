// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Migrate.Latest.Outputs
{

    [OutputType]
    public sealed class MoveCollectionPropertiesResponseErrors
    {
        /// <summary>
        /// The move resource error body.
        /// </summary>
        public readonly Outputs.MoveResourceErrorBodyResponse? Properties;

        [OutputConstructor]
        private MoveCollectionPropertiesResponseErrors(Outputs.MoveResourceErrorBodyResponse? properties)
        {
            Properties = properties;
        }
    }
}
