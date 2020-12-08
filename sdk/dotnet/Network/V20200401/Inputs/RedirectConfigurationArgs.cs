// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200401.Inputs
{

    /// <summary>
    /// Describes Redirect Route.
    /// </summary>
    public sealed class RedirectConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fragment to add to the redirect URL. Fragment is the part of the URL that comes after #. Do not include the #.
        /// </summary>
        [Input("customFragment")]
        public Input<string>? CustomFragment { get; set; }

        /// <summary>
        /// Host to redirect. Leave empty to use the incoming host as the destination host.
        /// </summary>
        [Input("customHost")]
        public Input<string>? CustomHost { get; set; }

        /// <summary>
        /// The full path to redirect. Path cannot be empty and must start with /. Leave empty to use the incoming path as destination path.
        /// </summary>
        [Input("customPath")]
        public Input<string>? CustomPath { get; set; }

        /// <summary>
        /// The set of query strings to be placed in the redirect URL. Setting this value would replace any existing query string; leave empty to preserve the incoming query string. Query string must be in &lt;key&gt;=&lt;value&gt; format. The first ? and &amp; will be added automatically so do not include them in the front, but do separate multiple query strings with &amp;.
        /// </summary>
        [Input("customQueryString")]
        public Input<string>? CustomQueryString { get; set; }

        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        /// <summary>
        /// The protocol of the destination to where the traffic is redirected
        /// </summary>
        [Input("redirectProtocol")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200401.FrontDoorRedirectProtocol>? RedirectProtocol { get; set; }

        /// <summary>
        /// The redirect type the rule will use when redirecting traffic.
        /// </summary>
        [Input("redirectType")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200401.FrontDoorRedirectType>? RedirectType { get; set; }

        public RedirectConfigurationArgs()
        {
        }
    }
}
