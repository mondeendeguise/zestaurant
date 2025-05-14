package toast

import (
	"testing"

	"github.com/mondeendeguise/zestaurant/toast"
	"github.com/mondeendeguise/zestaurant/toast/schema/menus/v2"
	"github.com/mondeendeguise/zestaurant/toast/mock"
)

func TestGetMenusSuccess(t *testing.T) {
	response, err := mock.MockResponse(mock.Menus, "200 Successful operation", 200)
	if err != nil {
		t.Error(err)
	}

	var _ menus.MenusInfo
	_, err = toast.HandleMenusResponse(response)
	if err != nil {
		t.Error(err)
	}
}
