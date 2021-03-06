/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mocks

import (
	"crypto/tls"
	"crypto/x509"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/endpoint"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/pkg/errors"
)

// MockConfig ...
type MockConfig struct {
	tlsEnabled             bool
	mutualTLSEnabled       bool
	errorCase              bool
	customNetworkPeerCfg   []fab.NetworkPeer
	customPeerCfg          *fab.PeerConfig
	customOrdererCfg       *fab.OrdererConfig
	customRandomOrdererCfg *fab.OrdererConfig
}

// NewMockCryptoConfig ...
func NewMockCryptoConfig() core.CryptoSuiteConfig {
	return &MockConfig{}
}

// NewMockEndpointConfig ...
func NewMockEndpointConfig() fab.EndpointConfig {
	return &MockConfig{}
}

// NewMockIdentityConfig ...
func NewMockIdentityConfig() msp.IdentityConfig {
	return &MockConfig{}
}

// NewMockCryptoConfigCustomized ...
func NewMockCryptoConfigCustomized(tlsEnabled, mutualTLSEnabled, errorCase bool) core.CryptoSuiteConfig {
	return &MockConfig{tlsEnabled: tlsEnabled, mutualTLSEnabled: mutualTLSEnabled, errorCase: errorCase}
}

// NewMockEndpointConfigCustomized ...
func NewMockEndpointConfigCustomized(tlsEnabled, mutualTLSEnabled, errorCase bool) fab.EndpointConfig {
	return &MockConfig{tlsEnabled: tlsEnabled, mutualTLSEnabled: mutualTLSEnabled, errorCase: errorCase}
}

// NewMockIdentityConfigCustomized ...
func NewMockIdentityConfigCustomized(tlsEnabled, mutualTLSEnabled, errorCase bool) msp.IdentityConfig {
	return &MockConfig{tlsEnabled: tlsEnabled, mutualTLSEnabled: mutualTLSEnabled, errorCase: errorCase}
}

// Client ...
func (c *MockConfig) Client() (*msp.ClientConfig, error) {
	clientConfig := msp.ClientConfig{}

	clientConfig.CredentialStore = msp.CredentialStoreType{
		Path: "/tmp/fabsdkgo_test/store",
	}

	if c.mutualTLSEnabled {
		mutualTLSCerts := endpoint.MutualTLSConfig{

			Client: endpoint.TLSKeyPair{
				Key: endpoint.TLSConfig{
					Path: "../../../test/fixtures/config/mutual_tls/client_sdk_go-key.pem",
					Pem:  "",
				},
				Cert: endpoint.TLSConfig{
					Path: "../../../test/fixtures/config/mutual_tls/client_sdk_go.pem",
					Pem:  "",
				},
			},
		}
		clientConfig.TLSCerts = mutualTLSCerts
	}

	return &clientConfig, nil
}

// CAConfig not implemented
func (c *MockConfig) CAConfig(org string) (*msp.CAConfig, error) {
	caConfig := msp.CAConfig{
		CAName: "org1",
	}

	return &caConfig, nil
}

//CAServerCerts Read configuration option for the server certificates for given org
func (c *MockConfig) CAServerCerts(org string) ([][]byte, error) {
	return nil, nil
}

//CAClientKey Read configuration option for the fabric CA client key for given org
func (c *MockConfig) CAClientKey(org string) ([]byte, error) {
	return nil, nil
}

//CAClientCert Read configuration option for the fabric CA client cert for given org
func (c *MockConfig) CAClientCert(org string) ([]byte, error) {
	return nil, nil
}

//TimeoutOrDefault not implemented
func (c *MockConfig) TimeoutOrDefault(arg fab.TimeoutType) time.Duration {
	return time.Second * 5
}

//Timeout not implemented
func (c *MockConfig) Timeout(arg fab.TimeoutType) time.Duration {
	return time.Second * 10
}

// PeersConfig Retrieves the fabric peers from the config file provided
func (c *MockConfig) PeersConfig(org string) ([]fab.PeerConfig, error) {
	return nil, nil
}

// PeerConfig Retrieves a specific peer from the configuration by org and name
func (c *MockConfig) PeerConfig(org string, name string) (*fab.PeerConfig, error) {
	return nil, nil
}

// PeerConfigByURL retrieves PeerConfig by URL
func (c *MockConfig) PeerConfigByURL(url string) (*fab.PeerConfig, error) {
	if url == "invalid" {
		return nil, errors.New("no peer")
	}
	if c.customPeerCfg != nil {
		return c.customPeerCfg, nil
	}
	cfg := fab.PeerConfig{
		URL: "example.com",
	}
	return &cfg, nil
}

// TLSCACertPool ...
func (c *MockConfig) TLSCACertPool(cert ...*x509.Certificate) (*x509.CertPool, error) {
	if c.errorCase {
		return nil, errors.New("just to test error scenario")
	}
	return nil, nil
}

// TcertBatchSize ...
func (c *MockConfig) TcertBatchSize() int {
	return 0
}

// SecurityAlgorithm ...
func (c *MockConfig) SecurityAlgorithm() string {
	return "SHA2"
}

// SecurityLevel ...
func (c *MockConfig) SecurityLevel() int {
	return 256

}

//SecurityProviderLibPath will be set only if provider is PKCS11
func (c *MockConfig) SecurityProviderLibPath() string {
	return ""
}

// OrderersConfig returns a list of defined orderers
func (c *MockConfig) OrderersConfig() ([]fab.OrdererConfig, error) {
	oConfig, err := c.OrdererConfig("")

	return []fab.OrdererConfig{*oConfig}, err
}

//SetCustomNetworkPeerCfg sets custom orderer config for unit-tests
func (c *MockConfig) SetCustomNetworkPeerCfg(customNetworkPeerCfg []fab.NetworkPeer) {
	c.customNetworkPeerCfg = customNetworkPeerCfg
}

//SetCustomPeerCfg sets custom orderer config for unit-tests
func (c *MockConfig) SetCustomPeerCfg(customPeerCfg *fab.PeerConfig) {
	c.customPeerCfg = customPeerCfg
}

//SetCustomOrdererCfg sets custom orderer config for unit-tests
func (c *MockConfig) SetCustomOrdererCfg(customOrdererCfg *fab.OrdererConfig) {
	c.customOrdererCfg = customOrdererCfg
}

//SetCustomRandomOrdererCfg sets custom random orderer config for unit-tests
func (c *MockConfig) SetCustomRandomOrdererCfg(customRandomOrdererCfg *fab.OrdererConfig) {
	c.customRandomOrdererCfg = customRandomOrdererCfg
}

// OrdererConfig not implemented
func (c *MockConfig) OrdererConfig(name string) (*fab.OrdererConfig, error) {
	if name == "Invalid" {
		return nil, errors.New("no orderer")
	}
	if c.customOrdererCfg != nil {
		return c.customOrdererCfg, nil
	}
	oConfig := fab.OrdererConfig{
		URL: "example.com",
	}

	return &oConfig, nil
}

// MSPID not implemented
func (c *MockConfig) MSPID(org string) (string, error) {
	return "", nil
}

// PeerMSPID not implemented
func (c *MockConfig) PeerMSPID(name string) (string, error) {
	return "", nil
}

// KeyStorePath ...
func (c *MockConfig) KeyStorePath() string {
	return "/tmp/fabsdkgo_test"
}

// CredentialStorePath ...
func (c *MockConfig) CredentialStorePath() string {
	return "/tmp/userstore"
}

// CAKeyStorePath not implemented
func (c *MockConfig) CAKeyStorePath() string {
	return "/tmp/fabsdkgo_test"
}

// CryptoConfigPath ...
func (c *MockConfig) CryptoConfigPath() string {
	return ""
}

// NetworkConfig not implemented
func (c *MockConfig) NetworkConfig() (*fab.NetworkConfig, error) {
	return nil, nil
}

// ChannelConfig returns the channel configuration
func (c *MockConfig) ChannelConfig(name string) (*fab.ChannelNetworkConfig, error) {
	return &fab.ChannelNetworkConfig{Policies: fab.ChannelPolicies{}}, nil
}

// ChannelPeers returns the channel peers configuration
func (c *MockConfig) ChannelPeers(name string) ([]fab.ChannelPeer, error) {

	if name == "noChannelPeers" {
		return nil, nil
	}

	peerChCfg := fab.PeerChannelConfig{EndorsingPeer: true, ChaincodeQuery: true, LedgerQuery: true, EventSource: true}
	if name == "noEndpoints" {
		peerChCfg = fab.PeerChannelConfig{EndorsingPeer: false, ChaincodeQuery: false, LedgerQuery: false, EventSource: false}
	}

	mockPeer := fab.ChannelPeer{PeerChannelConfig: peerChCfg, NetworkPeer: fab.NetworkPeer{PeerConfig: fab.PeerConfig{URL: "example.com"}}}
	return []fab.ChannelPeer{mockPeer}, nil
}

// ChannelOrderers returns a list of channel orderers
func (c *MockConfig) ChannelOrderers(name string) ([]fab.OrdererConfig, error) {
	if name == "Invalid" {
		return nil, errors.New("no orderer")
	}

	oConfig, err := c.OrdererConfig("")

	return []fab.OrdererConfig{*oConfig}, err
}

// NetworkPeers returns the mock network peers configuration
func (c *MockConfig) NetworkPeers() ([]fab.NetworkPeer, error) {
	if c.customNetworkPeerCfg != nil {
		return c.customNetworkPeerCfg, nil
	}
	return nil, errors.New("no config")
}

// SecurityProvider ...
func (c *MockConfig) SecurityProvider() string {
	return "sw"
}

// SecurityProviderLabel ...
func (c *MockConfig) SecurityProviderLabel() string {
	return ""
}

//SecurityProviderPin ...
func (c *MockConfig) SecurityProviderPin() string {
	return ""
}

//SoftVerify flag
func (c *MockConfig) SoftVerify() bool {
	return false
}

// IsSecurityEnabled ...
func (c *MockConfig) IsSecurityEnabled() bool {
	return false
}

// TLSClientCerts ...
func (c *MockConfig) TLSClientCerts() ([]tls.Certificate, error) {
	return nil, nil
}

// EventServiceType returns the type of event service client to use
func (c *MockConfig) EventServiceType() fab.EventServiceType {
	return fab.DeliverEventServiceType
}

// Lookup gets the Value from config file by Key
func (c *MockConfig) Lookup(key string) (interface{}, bool) {
	if key == "invalid" {
		return nil, false
	}
	value, ok := c.Lookup(key)
	if !ok {
		return nil, false
	}
	return value, true
}
