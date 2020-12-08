// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.AVS.V20200320.Inputs
{

    /// <summary>
    /// vCenter Single Sign On Identity Source
    /// </summary>
    public sealed class IdentitySourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain's NetBIOS name
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// The base distinguished name for groups
        /// </summary>
        [Input("baseGroupDN")]
        public Input<string>? BaseGroupDN { get; set; }

        /// <summary>
        /// The base distinguished name for users
        /// </summary>
        [Input("baseUserDN")]
        public Input<string>? BaseUserDN { get; set; }

        /// <summary>
        /// The domain's dns name
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The name of the identity source
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password of the Active Directory user with a minimum of read-only access to Base DN for users and groups.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Primary server URL
        /// </summary>
        [Input("primaryServer")]
        public Input<string>? PrimaryServer { get; set; }

        /// <summary>
        /// Secondary server URL
        /// </summary>
        [Input("secondaryServer")]
        public Input<string>? SecondaryServer { get; set; }

        /// <summary>
        /// Protect LDAP communication using SSL certificate (LDAPS)
        /// </summary>
        [Input("ssl")]
        public InputUnion<string, Pulumi.AzureNextGen.AVS.V20200320.SslEnum>? Ssl { get; set; }

        /// <summary>
        /// The ID of an Active Directory user with a minimum of read-only access to Base DN for users and group
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public IdentitySourceArgs()
        {
        }
    }
}
