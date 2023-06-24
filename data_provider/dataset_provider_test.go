package data_provider_test

import (
	"testing"
	"time"

	"github.com/chwzx/DataSharingSystem/manager"

	"github.com/chwzx/DataSharingSystem/data_provider"
)

func TestRunDatasetProvider(t *testing.T) {
	go manager.RunManager(5007, 5008, "../manager/assets", "info", "../logging/log.log",
		"../key_management/keys_certificates")
	time.Sleep(1 * time.Second)

	go data_provider.RunDatasetProvider("Data_provider1", "../data_provider/datasets", "info",
		"../logging/log.log", "localhost:5008", "../key_management/keys_certificates",
		[]string{"all"})
	time.Sleep(1 * time.Second)
}
