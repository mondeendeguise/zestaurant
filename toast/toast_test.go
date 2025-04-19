package toast

import (
	"testing"

	"github.com/mondeendeguise/zestaurant/toast/schema/menus/v2"

	"embed"
	"encoding/json"

	"github.com/Rican7/conjson"
	"github.com/Rican7/conjson/transform"
)

//go:embed mock
var mockJson embed.FS

const menusFile string = "mock/menus.json"

// TODO: handle Metadata.LastUpdated, Restaurant.LastUpdated as timestamps
//       check that we only send GET requests when we're out of date

// TODO: abstract some of this away
//       create a get() function that supports dependency injection
func TestUnmarshalMenus(t *testing.T) {
	b, err := mockJson.ReadFile(menusFile)
	if err != nil {
		t.Errorf("failed to read file `%s`\n", menusFile)
	}

	restaurant := menus.Restaurant{}
	var unmarshaler json.Unmarshaler = conjson.NewUnmarshaler(&restaurant, transform.ConventionalKeys())

	err = json.Unmarshal(b, unmarshaler)
	if err != nil {
		t.Errorf("%s\n", err)
	}
}
