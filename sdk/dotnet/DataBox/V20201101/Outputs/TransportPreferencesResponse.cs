// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBox.V20201101.Outputs
{

    [OutputType]
    public sealed class TransportPreferencesResponse
    {
        /// <summary>
        /// Indicates Shipment Logistics type that the customer preferred.
        /// </summary>
        public readonly string PreferredShipmentType;

        [OutputConstructor]
        private TransportPreferencesResponse(string preferredShipmentType)
        {
            PreferredShipmentType = preferredShipmentType;
        }
    }
}
