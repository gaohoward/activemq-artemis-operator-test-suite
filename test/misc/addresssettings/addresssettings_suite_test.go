package addresssettings

import (
	"github.com/artemiscloud/activemq-artemis-operator-test-suite/pkg/bdw"
	"github.com/artemiscloud/activemq-artemis-operator-test-suite/test"
	"github.com/rh-messaging/shipshape/pkg/framework"
	"testing"
)

func TestAddressSettings(t *testing.T) {
	test.PrepareNamespace(t, "addresssetting", "Address Setting test suite")
}

func TestMain(m *testing.M) {
	test.Initialize(m)
}

func setEnv(ctx1 *framework.ContextData, brokerDeployer *bdw.BrokerDeploymentWrapper) {
	brokerDeployer.WithWait(true).
		WithContext(ctx1).
		WithBrokerClient(sw.BrokerClient).
		WithCustomImage(test.Config.BrokerImageName).
		WithName(DeployName).
		WithLts(!test.Config.NeedsLatestCR).
		WithIncreasedTimeout(test.Config.TimeoutMultiplier)
}
