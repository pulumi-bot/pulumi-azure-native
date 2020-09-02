// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppPlatform.V20190501Preview.Outputs
{

    [OutputType]
    public sealed class ConfigServerGitPropertyResponseResult
    {
        /// <summary>
        /// Public sshKey of git repository.
        /// </summary>
        public readonly string? HostKey;
        /// <summary>
        /// SshKey algorithm of git repository.
        /// </summary>
        public readonly string? HostKeyAlgorithm;
        /// <summary>
        /// Label of the repository
        /// </summary>
        public readonly string? Label;
        /// <summary>
        /// Password of git repository basic auth.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Private sshKey algorithm of git repository.
        /// </summary>
        public readonly string? PrivateKey;
        /// <summary>
        /// Repositories of git.
        /// </summary>
        public readonly ImmutableArray<Outputs.GitPatternRepositoryResponseResult> Repositories;
        /// <summary>
        /// Searching path of the repository
        /// </summary>
        public readonly ImmutableArray<string> SearchPaths;
        /// <summary>
        /// Strict host key checking or not.
        /// </summary>
        public readonly bool? StrictHostKeyChecking;
        /// <summary>
        /// URI of the repository
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// Username of git repository basic auth.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private ConfigServerGitPropertyResponseResult(
            string? hostKey,

            string? hostKeyAlgorithm,

            string? label,

            string? password,

            string? privateKey,

            ImmutableArray<Outputs.GitPatternRepositoryResponseResult> repositories,

            ImmutableArray<string> searchPaths,

            bool? strictHostKeyChecking,

            string uri,

            string? username)
        {
            HostKey = hostKey;
            HostKeyAlgorithm = hostKeyAlgorithm;
            Label = label;
            Password = password;
            PrivateKey = privateKey;
            Repositories = repositories;
            SearchPaths = searchPaths;
            StrictHostKeyChecking = strictHostKeyChecking;
            Uri = uri;
            Username = username;
        }
    }
}
