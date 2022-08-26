package config

import (
	"context"
	"strings"
	"testing"

	"github.com/edaniels/golog"
	"go.viam.com/test"
)

func TestStoreToCache(t *testing.T) {
	logger := golog.NewTestLogger(t)
	ctx := context.Background()
	cfg, err := FromReader(ctx, "", strings.NewReader(`{}`), logger)

	test.That(t, err, test.ShouldBeNil)

	cloud := &Cloud{
		ManagedBy:        "acme",
		SignalingAddress: "abc",
		ID:               "forCachingTest",
		Secret:           "ghi",
		FQDN:             "fqdn",
		LocalFQDN:        "localFqdn",
		TLSCertificate:   "cert",
		TLSPrivateKey:    "key",
	}
	cfg.Cloud = cloud

	// store our config to the cloud
	err = storeToCache(cfg.Cloud.ID, cfg)
	test.That(t, err, test.ShouldBeNil)

	// read config from cloud, confirm consistency
	cloudCfg, err := readFromCloud(ctx, cfg, nil, true, true, logger)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, cloudCfg, test.ShouldResemble, cfg)

	// Modify our config
	newRemote := Remote{Name: "test", Address: "foo", Prefix: true}
	cfg.Remotes = append(cfg.Remotes, newRemote)

	// read config from cloud again, confirm that the cached config differs from cfg
	cloudCfg2, err := readFromCloud(ctx, cfg, nil, true, true, logger)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, cloudCfg2, test.ShouldNotResemble, cfg)

	// store the updated config to the cloud
	err = storeToCache(cfg.Cloud.ID, cfg)
	test.That(t, err, test.ShouldBeNil)

	// read updated cloud config, confirm that it now matches our updated cfg
	cloudCfg3, err := readFromCloud(ctx, cfg, nil, true, true, logger)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, cloudCfg3, test.ShouldResemble, cfg)
}

func TestShouldCheckForCert(t *testing.T) {
	cloud1 := Cloud{
		ManagedBy:        "acme",
		SignalingAddress: "abc",
		ID:               "forCachingTest",
		Secret:           "ghi",
		FQDN:             "fqdn",
		LocalFQDN:        "localFqdn",
		TLSCertificate:   "cert",
		TLSPrivateKey:    "key",
	}
	cloud2 := cloud1
	test.That(t, shouldCheckForCert(&cloud1, &cloud2), test.ShouldBeFalse)

	cloud2.TLSCertificate = "abc"
	test.That(t, shouldCheckForCert(&cloud1, &cloud2), test.ShouldBeFalse)

	cloud2 = cloud1
	cloud2.LocationSecret = "something else"
	test.That(t, shouldCheckForCert(&cloud1, &cloud2), test.ShouldBeTrue)
}

func TestProcessConfig(t *testing.T) {
	unprocessedConfig := Config{
		ConfigFilePath: "path",
	}

	cfg, err := processConfig(&unprocessedConfig, true)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, *cfg, test.ShouldResemble, unprocessedConfig)
}
