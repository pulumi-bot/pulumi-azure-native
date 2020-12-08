// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20200101.Inputs
{

    public sealed class JitNetworkAccessPortRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mutually exclusive with the "allowedSourceAddressPrefixes" parameter. Should be an IP address or CIDR, for example "192.168.0.3" or "192.168.0.0/16".
        /// </summary>
        [Input("allowedSourceAddressPrefix")]
        public Input<string>? AllowedSourceAddressPrefix { get; set; }

        [Input("allowedSourceAddressPrefixes")]
        private InputList<string>? _allowedSourceAddressPrefixes;

        /// <summary>
        /// Mutually exclusive with the "allowedSourceAddressPrefix" parameter.
        /// </summary>
        public InputList<string> AllowedSourceAddressPrefixes
        {
            get => _allowedSourceAddressPrefixes ?? (_allowedSourceAddressPrefixes = new InputList<string>());
            set => _allowedSourceAddressPrefixes = value;
        }

        /// <summary>
        /// Maximum duration requests can be made for. In ISO 8601 duration format. Minimum 5 minutes, maximum 1 day
        /// </summary>
        [Input("maxRequestAccessDuration", required: true)]
        public Input<string> MaxRequestAccessDuration { get; set; } = null!;

        [Input("number", required: true)]
        public Input<int> Number { get; set; } = null!;

        [Input("protocol", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Security.V20200101.Protocol> Protocol { get; set; } = null!;

        public JitNetworkAccessPortRuleArgs()
        {
        }
    }
}
