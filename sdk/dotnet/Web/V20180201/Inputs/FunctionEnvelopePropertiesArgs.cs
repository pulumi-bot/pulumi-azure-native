// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201.Inputs
{

    /// <summary>
    /// FunctionEnvelope resource specific properties
    /// </summary>
    public sealed class FunctionEnvelopePropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<object>? _config;

        /// <summary>
        /// Config information.
        /// </summary>
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        /// <summary>
        /// Config URI.
        /// </summary>
        [Input("config_href")]
        public Input<string>? Config_href { get; set; }

        [Input("files")]
        private InputMap<string>? _files;

        /// <summary>
        /// File list.
        /// </summary>
        public InputMap<string> Files
        {
            get => _files ?? (_files = new InputMap<string>());
            set => _files = value;
        }

        /// <summary>
        /// Function App ID.
        /// </summary>
        [Input("function_app_id")]
        public Input<string>? Function_app_id { get; set; }

        /// <summary>
        /// Function URI.
        /// </summary>
        [Input("href")]
        public Input<string>? Href { get; set; }

        /// <summary>
        /// The invocation URL
        /// </summary>
        [Input("invoke_url_template")]
        public Input<string>? Invoke_url_template { get; set; }

        /// <summary>
        /// Value indicating whether the function is disabled
        /// </summary>
        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// The function language
        /// </summary>
        [Input("language")]
        public Input<string>? Language { get; set; }

        /// <summary>
        /// Script URI.
        /// </summary>
        [Input("script_href")]
        public Input<string>? Script_href { get; set; }

        /// <summary>
        /// Script root path URI.
        /// </summary>
        [Input("script_root_path_href")]
        public Input<string>? Script_root_path_href { get; set; }

        /// <summary>
        /// Secrets file URI.
        /// </summary>
        [Input("secrets_file_href")]
        public Input<string>? Secrets_file_href { get; set; }

        /// <summary>
        /// Test data used when testing via the Azure Portal.
        /// </summary>
        [Input("test_data")]
        public Input<string>? Test_data { get; set; }

        /// <summary>
        /// Test data URI.
        /// </summary>
        [Input("test_data_href")]
        public Input<string>? Test_data_href { get; set; }

        public FunctionEnvelopePropertiesArgs()
        {
        }
    }
}