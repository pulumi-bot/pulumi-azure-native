// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CustomerInsights.V20170426.Inputs
{

    /// <summary>
    /// The complete operation.
    /// </summary>
    public sealed class ConnectorMappingCompleteOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of completion operation.
        /// </summary>
        [Input("completionOperationType")]
        public Input<Pulumi.AzureNextGen.CustomerInsights.V20170426.CompletionOperationTypes>? CompletionOperationType { get; set; }

        /// <summary>
        /// The destination folder where files will be moved to once the import is done.
        /// </summary>
        [Input("destinationFolder")]
        public Input<string>? DestinationFolder { get; set; }

        public ConnectorMappingCompleteOperationArgs()
        {
        }
    }
}
