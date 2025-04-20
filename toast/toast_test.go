package toast

import (
	"testing"

	"github.com/mondeendeguise/zestaurant/toast/schema/menus/v2"

	"embed"
	"encoding/json"
	"time"

	"github.com/Rican7/conjson"
	"github.com/Rican7/conjson/transform"
)

//go:embed mock
var mockJson embed.FS

const METADATA_FILE string = "mock/metadata.json"
const MENUS_FILE string = "mock/menus.json"

// https://stackoverflow.com/a/38596248
const TIME_LAYOUT_ISO8601 string = "2006-01-02T15:04:05-0700"

// TODO: abstract some of this away
//       create a get() function that supports dependency injection

func TestGetMetadata(t *testing.T) {
	b, err := mockJson.ReadFile(METADATA_FILE)
	if err != nil {
		t.Errorf("failed to read file `%s`", METADATA_FILE)
	}

	metadata := menus.Metadata{}
	var unmarshaler json.Unmarshaler = conjson.NewUnmarshaler(&metadata, transform.ConventionalKeys())

	err = json.Unmarshal(b, unmarshaler)
	if err != nil {
		t.Errorf("%s", err)
	}
}

// TODO: ensure we only send GET requests when `lastUpdated` is modified

func TestCheckLastUpdated(t *testing.T) {
	b, err := mockJson.ReadFile(METADATA_FILE)
	if err != nil {
		t.Errorf("failed to read file `%s`", METADATA_FILE)
	}

	metadata := menus.Metadata{}
	var unmarshaler json.Unmarshaler = conjson.NewUnmarshaler(&metadata, transform.ConventionalKeys())

	err = json.Unmarshal(b, unmarshaler)
	if err != nil {
		t.Errorf("%s", err)
	}

	lastUpdated, err := time.Parse(TIME_LAYOUT_ISO8601, metadata.LastUpdated)
	if err != nil {
		t.Errorf("%s", err)
	}

	var timestamp time.Time = time.Now()

	if !lastUpdated.Before(timestamp) {
		t.Errorf("lastUpdated has an impossible timestamp")
	}
}

func TestUnmarshalMenus(t *testing.T) {
	b, err := mockJson.ReadFile(MENUS_FILE)
	if err != nil {
		t.Errorf("failed to read file `%s`", MENUS_FILE)
	}

	restaurant := menus.Restaurant{}
	var unmarshaler json.Unmarshaler = conjson.NewUnmarshaler(&restaurant, transform.ConventionalKeys())

	err = json.Unmarshal(b, unmarshaler)
	if err != nil {
		t.Errorf("%s", err)
	}
}
