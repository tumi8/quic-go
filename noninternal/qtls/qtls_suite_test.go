package qtls

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func TestQTLS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "qtls Suite")
}

var mockCtrl *gomock.Controller

var _ = BeforeEach(func() {
	mockCtrl = gomock.NewController(GinkgoT())
})

var _ = AfterEach(func() {
	mockCtrl.Finish()
})
