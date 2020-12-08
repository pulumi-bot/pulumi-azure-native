// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.V20181101
{
    /// <summary>
    /// A hostname binding object.
    /// </summary>
    public partial class WebAppHostNameBindingSlot : Pulumi.CustomResource
    {
        /// <summary>
        /// Azure resource name.
        /// </summary>
        [Output("azureResourceName")]
        public Output<string?> AzureResourceName { get; private set; } = null!;

        /// <summary>
        /// Azure resource type.
        /// </summary>
        [Output("azureResourceType")]
        public Output<string?> AzureResourceType { get; private set; } = null!;

        /// <summary>
        /// Custom DNS record type.
        /// </summary>
        [Output("customHostNameDnsRecordType")]
        public Output<string?> CustomHostNameDnsRecordType { get; private set; } = null!;

        /// <summary>
        /// Fully qualified ARM domain resource URI.
        /// </summary>
        [Output("domainId")]
        public Output<string?> DomainId { get; private set; } = null!;

        /// <summary>
        /// Hostname type.
        /// </summary>
        [Output("hostNameType")]
        public Output<string?> HostNameType { get; private set; } = null!;

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// App Service app name.
        /// </summary>
        [Output("siteName")]
        public Output<string?> SiteName { get; private set; } = null!;

        /// <summary>
        /// SSL type
        /// </summary>
        [Output("sslState")]
        public Output<string?> SslState { get; private set; } = null!;

        /// <summary>
        /// SSL certificate thumbprint
        /// </summary>
        [Output("thumbprint")]
        public Output<string?> Thumbprint { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Virtual IP address assigned to the hostname if IP based SSL is enabled.
        /// </summary>
        [Output("virtualIP")]
        public Output<string> VirtualIP { get; private set; } = null!;


        /// <summary>
        /// Create a WebAppHostNameBindingSlot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebAppHostNameBindingSlot(string name, WebAppHostNameBindingSlotArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:web/v20181101:WebAppHostNameBindingSlot", name, args ?? new WebAppHostNameBindingSlotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebAppHostNameBindingSlot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:web/v20181101:WebAppHostNameBindingSlot", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:web/latest:WebAppHostNameBindingSlot"},
                    new Pulumi.Alias { Type = "azure-nextgen:web/v20150801:WebAppHostNameBindingSlot"},
                    new Pulumi.Alias { Type = "azure-nextgen:web/v20160801:WebAppHostNameBindingSlot"},
                    new Pulumi.Alias { Type = "azure-nextgen:web/v20180201:WebAppHostNameBindingSlot"},
                    new Pulumi.Alias { Type = "azure-nextgen:web/v20190801:WebAppHostNameBindingSlot"},
                    new Pulumi.Alias { Type = "azure-nextgen:web/v20200601:WebAppHostNameBindingSlot"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WebAppHostNameBindingSlot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebAppHostNameBindingSlot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebAppHostNameBindingSlot(name, id, options);
        }
    }

    public sealed class WebAppHostNameBindingSlotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure resource name.
        /// </summary>
        [Input("azureResourceName")]
        public Input<string>? AzureResourceName { get; set; }

        /// <summary>
        /// Azure resource type.
        /// </summary>
        [Input("azureResourceType")]
        public Input<Pulumi.AzureNextGen.Web.V20181101.AzureResourceType>? AzureResourceType { get; set; }

        /// <summary>
        /// Custom DNS record type.
        /// </summary>
        [Input("customHostNameDnsRecordType")]
        public Input<Pulumi.AzureNextGen.Web.V20181101.CustomHostNameDnsRecordType>? CustomHostNameDnsRecordType { get; set; }

        /// <summary>
        /// Fully qualified ARM domain resource URI.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// Hostname in the hostname binding.
        /// </summary>
        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        /// <summary>
        /// Hostname type.
        /// </summary>
        [Input("hostNameType")]
        public Input<Pulumi.AzureNextGen.Web.V20181101.HostNameType>? HostNameType { get; set; }

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// App Service app name.
        /// </summary>
        [Input("siteName")]
        public Input<string>? SiteName { get; set; }

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will create a binding for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public Input<string> Slot { get; set; } = null!;

        /// <summary>
        /// SSL type
        /// </summary>
        [Input("sslState")]
        public Input<Pulumi.AzureNextGen.Web.V20181101.SslState>? SslState { get; set; }

        /// <summary>
        /// SSL certificate thumbprint
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        public WebAppHostNameBindingSlotArgs()
        {
        }
    }
}
