// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Batch.V20200501
{
    /// <summary>
    /// The default value is Pool. If the pool is running Windows a value of Task should be specified if stricter isolation between tasks is required. For example, if the task mutates the registry in a way which could impact other tasks, or if certificates have been specified on the pool which should not be accessible by normal tasks but should be accessible by start tasks.
    /// </summary>
    [EnumType]
    public readonly struct AutoUserScope : IEquatable<AutoUserScope>
    {
        private readonly string _value;

        private AutoUserScope(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Specifies that the service should create a new user for the task.
        /// </summary>
        public static AutoUserScope Task { get; } = new AutoUserScope("Task");
        /// <summary>
        /// Specifies that the task runs as the common auto user account which is created on every node in a pool.
        /// </summary>
        public static AutoUserScope Pool { get; } = new AutoUserScope("Pool");

        public static bool operator ==(AutoUserScope left, AutoUserScope right) => left.Equals(right);
        public static bool operator !=(AutoUserScope left, AutoUserScope right) => !left.Equals(right);

        public static explicit operator string(AutoUserScope value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AutoUserScope other && Equals(other);
        public bool Equals(AutoUserScope other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Values are:
    /// 
    ///  none - The caching mode for the disk is not enabled.
    ///  readOnly - The caching mode for the disk is read only.
    ///  readWrite - The caching mode for the disk is read and write.
    /// 
    ///  The default value for caching is none. For information about the caching options see: https://blogs.msdn.microsoft.com/windowsazurestorage/2012/06/27/exploring-windows-azure-drives-disks-and-images/.
    /// </summary>
    [EnumType]
    public readonly struct CachingType : IEquatable<CachingType>
    {
        private readonly string _value;

        private CachingType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The caching mode for the disk is not enabled.
        /// </summary>
        public static CachingType None { get; } = new CachingType("None");
        /// <summary>
        /// The caching mode for the disk is read only.
        /// </summary>
        public static CachingType ReadOnly { get; } = new CachingType("ReadOnly");
        /// <summary>
        /// The caching mode for the disk is read and write.
        /// </summary>
        public static CachingType ReadWrite { get; } = new CachingType("ReadWrite");

        public static bool operator ==(CachingType left, CachingType right) => left.Equals(right);
        public static bool operator !=(CachingType left, CachingType right) => !left.Equals(right);

        public static explicit operator string(CachingType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CachingType other && Equals(other);
        public bool Equals(CachingType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
    /// </summary>
    [EnumType]
    public readonly struct CertificateFormat : IEquatable<CertificateFormat>
    {
        private readonly string _value;

        private CertificateFormat(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The certificate is a PFX (PKCS#12) formatted certificate or certificate chain.
        /// </summary>
        public static CertificateFormat Pfx { get; } = new CertificateFormat("Pfx");
        /// <summary>
        /// The certificate is a base64-encoded X.509 certificate.
        /// </summary>
        public static CertificateFormat Cer { get; } = new CertificateFormat("Cer");

        public static bool operator ==(CertificateFormat left, CertificateFormat right) => left.Equals(right);
        public static bool operator !=(CertificateFormat left, CertificateFormat right) => !left.Equals(right);

        public static explicit operator string(CertificateFormat value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CertificateFormat other && Equals(other);
        public bool Equals(CertificateFormat other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The default value is currentUser. This property is applicable only for pools configured with Windows nodes (that is, created with cloudServiceConfiguration, or with virtualMachineConfiguration using a Windows image reference). For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
    /// </summary>
    [EnumType]
    public readonly struct CertificateStoreLocation : IEquatable<CertificateStoreLocation>
    {
        private readonly string _value;

        private CertificateStoreLocation(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Certificates should be installed to the CurrentUser certificate store.
        /// </summary>
        public static CertificateStoreLocation CurrentUser { get; } = new CertificateStoreLocation("CurrentUser");
        /// <summary>
        /// Certificates should be installed to the LocalMachine certificate store.
        /// </summary>
        public static CertificateStoreLocation LocalMachine { get; } = new CertificateStoreLocation("LocalMachine");

        public static bool operator ==(CertificateStoreLocation left, CertificateStoreLocation right) => left.Equals(right);
        public static bool operator !=(CertificateStoreLocation left, CertificateStoreLocation right) => !left.Equals(right);

        public static explicit operator string(CertificateStoreLocation value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CertificateStoreLocation other && Equals(other);
        public bool Equals(CertificateStoreLocation other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CertificateVisibility : IEquatable<CertificateVisibility>
    {
        private readonly string _value;

        private CertificateVisibility(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The certificate should be visible to the user account under which the start task is run. Note that if AutoUser Scope is Pool for both the StartTask and a Task, this certificate will be visible to the Task as well.
        /// </summary>
        public static CertificateVisibility StartTask { get; } = new CertificateVisibility("StartTask");
        /// <summary>
        /// The certificate should be visible to the user accounts under which job tasks are run.
        /// </summary>
        public static CertificateVisibility Task { get; } = new CertificateVisibility("Task");
        /// <summary>
        /// The certificate should be visible to the user accounts under which users remotely access the node.
        /// </summary>
        public static CertificateVisibility RemoteUser { get; } = new CertificateVisibility("RemoteUser");

        public static bool operator ==(CertificateVisibility left, CertificateVisibility right) => left.Equals(right);
        public static bool operator !=(CertificateVisibility left, CertificateVisibility right) => !left.Equals(right);

        public static explicit operator string(CertificateVisibility value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CertificateVisibility other && Equals(other);
        public bool Equals(CertificateVisibility other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// If omitted, the default value is Requeue.
    /// </summary>
    [EnumType]
    public readonly struct ComputeNodeDeallocationOption : IEquatable<ComputeNodeDeallocationOption>
    {
        private readonly string _value;

        private ComputeNodeDeallocationOption(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Terminate running task processes and requeue the tasks. The tasks will run again when a node is available. Remove nodes as soon as tasks have been terminated.
        /// </summary>
        public static ComputeNodeDeallocationOption Requeue { get; } = new ComputeNodeDeallocationOption("Requeue");
        /// <summary>
        /// Terminate running tasks. The tasks will be completed with failureInfo indicating that they were terminated, and will not run again. Remove nodes as soon as tasks have been terminated.
        /// </summary>
        public static ComputeNodeDeallocationOption Terminate { get; } = new ComputeNodeDeallocationOption("Terminate");
        /// <summary>
        /// Allow currently running tasks to complete. Schedule no new tasks while waiting. Remove nodes when all tasks have completed.
        /// </summary>
        public static ComputeNodeDeallocationOption TaskCompletion { get; } = new ComputeNodeDeallocationOption("TaskCompletion");
        /// <summary>
        /// Allow currently running tasks to complete, then wait for all task data retention periods to expire. Schedule no new tasks while waiting. Remove nodes when all task retention periods have expired.
        /// </summary>
        public static ComputeNodeDeallocationOption RetainedData { get; } = new ComputeNodeDeallocationOption("RetainedData");

        public static bool operator ==(ComputeNodeDeallocationOption left, ComputeNodeDeallocationOption right) => left.Equals(right);
        public static bool operator !=(ComputeNodeDeallocationOption left, ComputeNodeDeallocationOption right) => !left.Equals(right);

        public static explicit operator string(ComputeNodeDeallocationOption value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeNodeDeallocationOption other && Equals(other);
        public bool Equals(ComputeNodeDeallocationOption other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ComputeNodeFillType : IEquatable<ComputeNodeFillType>
    {
        private readonly string _value;

        private ComputeNodeFillType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Tasks should be assigned evenly across all nodes in the pool.
        /// </summary>
        public static ComputeNodeFillType Spread { get; } = new ComputeNodeFillType("Spread");
        /// <summary>
        /// As many tasks as possible (maxTasksPerNode) should be assigned to each node in the pool before any tasks are assigned to the next node in the pool.
        /// </summary>
        public static ComputeNodeFillType Pack { get; } = new ComputeNodeFillType("Pack");

        public static bool operator ==(ComputeNodeFillType left, ComputeNodeFillType right) => left.Equals(right);
        public static bool operator !=(ComputeNodeFillType left, ComputeNodeFillType right) => !left.Equals(right);

        public static explicit operator string(ComputeNodeFillType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeNodeFillType other && Equals(other);
        public bool Equals(ComputeNodeFillType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ContainerType : IEquatable<ContainerType>
    {
        private readonly string _value;

        private ContainerType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// A Docker compatible container technology will be used to launch the containers.
        /// </summary>
        public static ContainerType DockerCompatible { get; } = new ContainerType("DockerCompatible");

        public static bool operator ==(ContainerType left, ContainerType right) => left.Equals(right);
        public static bool operator !=(ContainerType left, ContainerType right) => !left.Equals(right);

        public static explicit operator string(ContainerType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContainerType other && Equals(other);
        public bool Equals(ContainerType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ContainerWorkingDirectory : IEquatable<ContainerWorkingDirectory>
    {
        private readonly string _value;

        private ContainerWorkingDirectory(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Use the standard Batch service task working directory, which will contain the Task resource files populated by Batch.
        /// </summary>
        public static ContainerWorkingDirectory TaskWorkingDirectory { get; } = new ContainerWorkingDirectory("TaskWorkingDirectory");
        /// <summary>
        /// Using container image defined working directory. Beware that this directory will not contain the resource files downloaded by Batch.
        /// </summary>
        public static ContainerWorkingDirectory ContainerImageDefault { get; } = new ContainerWorkingDirectory("ContainerImageDefault");

        public static bool operator ==(ContainerWorkingDirectory left, ContainerWorkingDirectory right) => left.Equals(right);
        public static bool operator !=(ContainerWorkingDirectory left, ContainerWorkingDirectory right) => !left.Equals(right);

        public static explicit operator string(ContainerWorkingDirectory value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ContainerWorkingDirectory other && Equals(other);
        public bool Equals(ContainerWorkingDirectory other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// If omitted, no disks on the compute nodes in the pool will be encrypted.
    /// </summary>
    [EnumType]
    public readonly struct DiskEncryptionTarget : IEquatable<DiskEncryptionTarget>
    {
        private readonly string _value;

        private DiskEncryptionTarget(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The OS Disk on the compute node is encrypted.
        /// </summary>
        public static DiskEncryptionTarget OsDisk { get; } = new DiskEncryptionTarget("OsDisk");
        /// <summary>
        /// The temporary disk on the compute node is encrypted. On Linux this encryption applies to other partitions (such as those on mounted data disks) when encryption occurs at boot time.
        /// </summary>
        public static DiskEncryptionTarget TemporaryDisk { get; } = new DiskEncryptionTarget("TemporaryDisk");

        public static bool operator ==(DiskEncryptionTarget left, DiskEncryptionTarget right) => left.Equals(right);
        public static bool operator !=(DiskEncryptionTarget left, DiskEncryptionTarget right) => !left.Equals(right);

        public static explicit operator string(DiskEncryptionTarget value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DiskEncryptionTarget other && Equals(other);
        public bool Equals(DiskEncryptionTarget other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// nonAdmin - The auto user is a standard user without elevated access. admin - The auto user is a user with elevated access and operates with full Administrator permissions. The default value is nonAdmin.
    /// </summary>
    [EnumType]
    public readonly struct ElevationLevel : IEquatable<ElevationLevel>
    {
        private readonly string _value;

        private ElevationLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The user is a standard user without elevated access.
        /// </summary>
        public static ElevationLevel NonAdmin { get; } = new ElevationLevel("NonAdmin");
        /// <summary>
        /// The user is a user with elevated access and operates with full Administrator permissions.
        /// </summary>
        public static ElevationLevel Admin { get; } = new ElevationLevel("Admin");

        public static bool operator ==(ElevationLevel left, ElevationLevel right) => left.Equals(right);
        public static bool operator !=(ElevationLevel left, ElevationLevel right) => !left.Equals(right);

        public static explicit operator string(ElevationLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ElevationLevel other && Equals(other);
        public bool Equals(ElevationLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The default value is BatchManaged
    /// </summary>
    [EnumType]
    public readonly struct IPAddressProvisioningType : IEquatable<IPAddressProvisioningType>
    {
        private readonly string _value;

        private IPAddressProvisioningType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// A public IP will be created and managed by Batch. There may be multiple public IPs depending on the size of the Pool.
        /// </summary>
        public static IPAddressProvisioningType BatchManaged { get; } = new IPAddressProvisioningType("BatchManaged");
        /// <summary>
        /// Public IPs are provided by the user and will be used to provision the Compute Nodes.
        /// </summary>
        public static IPAddressProvisioningType UserManaged { get; } = new IPAddressProvisioningType("UserManaged");
        /// <summary>
        /// No public IP Address will be created for the Compute Nodes in the Pool.
        /// </summary>
        public static IPAddressProvisioningType NoPublicIPAddresses { get; } = new IPAddressProvisioningType("NoPublicIPAddresses");

        public static bool operator ==(IPAddressProvisioningType left, IPAddressProvisioningType right) => left.Equals(right);
        public static bool operator !=(IPAddressProvisioningType left, IPAddressProvisioningType right) => !left.Equals(right);

        public static explicit operator string(IPAddressProvisioningType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IPAddressProvisioningType other && Equals(other);
        public bool Equals(IPAddressProvisioningType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct InboundEndpointProtocol : IEquatable<InboundEndpointProtocol>
    {
        private readonly string _value;

        private InboundEndpointProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Use TCP for the endpoint.
        /// </summary>
        public static InboundEndpointProtocol TCP { get; } = new InboundEndpointProtocol("TCP");
        /// <summary>
        /// Use UDP for the endpoint.
        /// </summary>
        public static InboundEndpointProtocol UDP { get; } = new InboundEndpointProtocol("UDP");

        public static bool operator ==(InboundEndpointProtocol left, InboundEndpointProtocol right) => left.Equals(right);
        public static bool operator !=(InboundEndpointProtocol left, InboundEndpointProtocol right) => !left.Equals(right);

        public static explicit operator string(InboundEndpointProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InboundEndpointProtocol other && Equals(other);
        public bool Equals(InboundEndpointProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
    /// </summary>
    [EnumType]
    public readonly struct InterNodeCommunicationState : IEquatable<InterNodeCommunicationState>
    {
        private readonly string _value;

        private InterNodeCommunicationState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Enable network communication between virtual machines.
        /// </summary>
        public static InterNodeCommunicationState Enabled { get; } = new InterNodeCommunicationState("Enabled");
        /// <summary>
        /// Disable network communication between virtual machines.
        /// </summary>
        public static InterNodeCommunicationState Disabled { get; } = new InterNodeCommunicationState("Disabled");

        public static bool operator ==(InterNodeCommunicationState left, InterNodeCommunicationState right) => left.Equals(right);
        public static bool operator !=(InterNodeCommunicationState left, InterNodeCommunicationState right) => !left.Equals(right);

        public static explicit operator string(InterNodeCommunicationState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InterNodeCommunicationState other && Equals(other);
        public bool Equals(InterNodeCommunicationState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the key source.
    /// </summary>
    [EnumType]
    public readonly struct KeySource : IEquatable<KeySource>
    {
        private readonly string _value;

        private KeySource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Batch creates and manages the encryption keys used to protect the account data.
        /// </summary>
        public static KeySource Microsoft_Batch { get; } = new KeySource("Microsoft.Batch");
        /// <summary>
        /// The encryption keys used to protect the account data are stored in an external key vault. If this is set then the Batch Account identity must be set to `SystemAssigned` and a valid Key Identifier must also be supplied under the keyVaultProperties.
        /// </summary>
        public static KeySource Microsoft_KeyVault { get; } = new KeySource("Microsoft.KeyVault");

        public static bool operator ==(KeySource left, KeySource right) => left.Equals(right);
        public static bool operator !=(KeySource left, KeySource right) => !left.Equals(right);

        public static explicit operator string(KeySource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is KeySource other && Equals(other);
        public bool Equals(KeySource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies login mode for the user. The default value for VirtualMachineConfiguration pools is interactive mode and for CloudServiceConfiguration pools is batch mode.
    /// </summary>
    [EnumType]
    public readonly struct LoginMode : IEquatable<LoginMode>
    {
        private readonly string _value;

        private LoginMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The LOGON32_LOGON_BATCH Win32 login mode. The batch login mode is recommended for long running parallel processes.
        /// </summary>
        public static LoginMode Batch { get; } = new LoginMode("Batch");
        /// <summary>
        /// The LOGON32_LOGON_INTERACTIVE Win32 login mode. Some applications require having permissions associated with the interactive login mode. If this is the case for an application used in your task, then this option is recommended.
        /// </summary>
        public static LoginMode Interactive { get; } = new LoginMode("Interactive");

        public static bool operator ==(LoginMode left, LoginMode right) => left.Equals(right);
        public static bool operator !=(LoginMode left, LoginMode right) => !left.Equals(right);

        public static explicit operator string(LoginMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LoginMode other && Equals(other);
        public bool Equals(LoginMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NetworkSecurityGroupRuleAccess : IEquatable<NetworkSecurityGroupRuleAccess>
    {
        private readonly string _value;

        private NetworkSecurityGroupRuleAccess(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Allow access.
        /// </summary>
        public static NetworkSecurityGroupRuleAccess Allow { get; } = new NetworkSecurityGroupRuleAccess("Allow");
        /// <summary>
        /// Deny access.
        /// </summary>
        public static NetworkSecurityGroupRuleAccess Deny { get; } = new NetworkSecurityGroupRuleAccess("Deny");

        public static bool operator ==(NetworkSecurityGroupRuleAccess left, NetworkSecurityGroupRuleAccess right) => left.Equals(right);
        public static bool operator !=(NetworkSecurityGroupRuleAccess left, NetworkSecurityGroupRuleAccess right) => !left.Equals(right);

        public static explicit operator string(NetworkSecurityGroupRuleAccess value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NetworkSecurityGroupRuleAccess other && Equals(other);
        public bool Equals(NetworkSecurityGroupRuleAccess other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is UserSubscription, clients must use Azure Active Directory. The default is BatchService.
    /// </summary>
    [EnumType]
    public readonly struct PoolAllocationMode : IEquatable<PoolAllocationMode>
    {
        private readonly string _value;

        private PoolAllocationMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Pools will be allocated in subscriptions owned by the Batch service.
        /// </summary>
        public static PoolAllocationMode BatchService { get; } = new PoolAllocationMode("BatchService");
        /// <summary>
        /// Pools will be allocated in a subscription owned by the user.
        /// </summary>
        public static PoolAllocationMode UserSubscription { get; } = new PoolAllocationMode("UserSubscription");

        public static bool operator ==(PoolAllocationMode left, PoolAllocationMode right) => left.Equals(right);
        public static bool operator !=(PoolAllocationMode left, PoolAllocationMode right) => !left.Equals(right);

        public static explicit operator string(PoolAllocationMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PoolAllocationMode other && Equals(other);
        public bool Equals(PoolAllocationMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// If not specified, the default value is 'enabled'.
    /// </summary>
    [EnumType]
    public readonly struct PublicNetworkAccessType : IEquatable<PublicNetworkAccessType>
    {
        private readonly string _value;

        private PublicNetworkAccessType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Enables connectivity to Azure Batch through public DNS.
        /// </summary>
        public static PublicNetworkAccessType Enabled { get; } = new PublicNetworkAccessType("Enabled");
        /// <summary>
        /// Disables public connectivity and enables private connectivity to Azure Batch Service through private endpoint resource.
        /// </summary>
        public static PublicNetworkAccessType Disabled { get; } = new PublicNetworkAccessType("Disabled");

        public static bool operator ==(PublicNetworkAccessType left, PublicNetworkAccessType right) => left.Equals(right);
        public static bool operator !=(PublicNetworkAccessType left, PublicNetworkAccessType right) => !left.Equals(right);

        public static explicit operator string(PublicNetworkAccessType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PublicNetworkAccessType other && Equals(other);
        public bool Equals(PublicNetworkAccessType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of identity used for the Batch account.
    /// </summary>
    [EnumType]
    public readonly struct ResourceIdentityType : IEquatable<ResourceIdentityType>
    {
        private readonly string _value;

        private ResourceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Batch account has a system assigned identity with it.
        /// </summary>
        public static ResourceIdentityType SystemAssigned { get; } = new ResourceIdentityType("SystemAssigned");
        /// <summary>
        /// Batch account has no identity associated with it. Setting `None` in update account will remove existing identities.
        /// </summary>
        public static ResourceIdentityType None { get; } = new ResourceIdentityType("None");

        public static bool operator ==(ResourceIdentityType left, ResourceIdentityType right) => left.Equals(right);
        public static bool operator !=(ResourceIdentityType left, ResourceIdentityType right) => !left.Equals(right);

        public static explicit operator string(ResourceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourceIdentityType other && Equals(other);
        public bool Equals(ResourceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// If omitted, the default is "Standard_LRS". Values are:
    /// 
    ///  Standard_LRS - The data disk should use standard locally redundant storage.
    ///  Premium_LRS - The data disk should use premium locally redundant storage.
    /// </summary>
    [EnumType]
    public readonly struct StorageAccountType : IEquatable<StorageAccountType>
    {
        private readonly string _value;

        private StorageAccountType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The data disk should use standard locally redundant storage.
        /// </summary>
        public static StorageAccountType Standard_LRS { get; } = new StorageAccountType("Standard_LRS");
        /// <summary>
        /// The data disk should use premium locally redundant storage.
        /// </summary>
        public static StorageAccountType Premium_LRS { get; } = new StorageAccountType("Premium_LRS");

        public static bool operator ==(StorageAccountType left, StorageAccountType right) => left.Equals(right);
        public static bool operator !=(StorageAccountType left, StorageAccountType right) => !left.Equals(right);

        public static explicit operator string(StorageAccountType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StorageAccountType other && Equals(other);
        public bool Equals(StorageAccountType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
