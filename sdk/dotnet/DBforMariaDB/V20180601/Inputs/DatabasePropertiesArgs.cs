// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBforMariaDB.V20180601.Inputs
{

    /// <summary>
    /// The properties of a database.
    /// </summary>
    public sealed class DatabasePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The charset of the database.
        /// </summary>
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        /// <summary>
        /// The collation of the database.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        public DatabasePropertiesArgs()
        {
        }
    }
}