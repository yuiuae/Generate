package handlers_test

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGenerateUUID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// mockObj := GenerateUUID.NewMockMyInterface(mockCtrl)
	// mockObj.EXPECT().SomeMethod(4, "blah")
	// pass mockObj to a real object and play with it.
}
