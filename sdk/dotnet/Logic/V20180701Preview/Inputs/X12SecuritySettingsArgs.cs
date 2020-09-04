// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview.Inputs
{

    /// <summary>
    /// The X12 agreement security settings.
    /// </summary>
    public sealed class X12SecuritySettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization qualifier.
        /// </summary>
        [Input("authorizationQualifier", required: true)]
        public Input<string> AuthorizationQualifier { get; set; } = null!;

        /// <summary>
        /// The authorization value.
        /// </summary>
        [Input("authorizationValue")]
        public Input<string>? AuthorizationValue { get; set; }

        /// <summary>
        /// The password value.
        /// </summary>
        [Input("passwordValue")]
        public Input<string>? PasswordValue { get; set; }

        /// <summary>
        /// The security qualifier.
        /// </summary>
        [Input("securityQualifier", required: true)]
        public Input<string> SecurityQualifier { get; set; } = null!;

        public X12SecuritySettingsArgs()
        {
        }
    }
}
