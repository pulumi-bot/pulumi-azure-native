// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801Preview.Inputs
{

    /// <summary>
    /// Message envelope that contains the common Azure resource manager properties and the resource provider specific content
    /// </summary>
    public sealed class ResponseMessageEnvelopeApiEntityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource Id. Typically id is populated only for responses to GET requests. Caller is responsible for passing in this
        ///             value for GET requests only.
        ///             For example: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupId}/providers/Microsoft.Web/sites/{sitename}
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Geo region resource belongs to e.g. SouthCentralUS, SouthEastAsia
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Azure resource manager plan
        /// </summary>
        [Input("plan")]
        public Input<Inputs.ArmPlanArgs>? Plan { get; set; }

        /// <summary>
        /// Resource specific properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ApiEntityArgs>? Properties { get; set; }

        /// <summary>
        /// Sku description of the resource
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuDescriptionArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags associated with resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of resource e.g Microsoft.Web/sites
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ResponseMessageEnvelopeApiEntityArgs()
        {
        }
    }
}
