package x509

func init() {
	certDirectories = append(certDirectories,
		"/system/etc/security/cacerts",    // Android system roots
		"/data/misc/keychain/certs-added", // User trusted CA folder
	)
}
