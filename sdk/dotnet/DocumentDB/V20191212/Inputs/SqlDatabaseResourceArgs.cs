// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20191212.Inputs
{

    /// <summary>
    /// Cosmos DB SQL database resource object
    /// </summary>
    public sealed class SqlDatabaseResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Cosmos DB SQL database
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public SqlDatabaseResourceArgs()
        {
        }
    }
}