package toast

import (
	"github.com/mondeendeguise/zestaurant/toast/schema/menus/v2"

	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"log"
	"os"
	"time"

	// "github.com/Rican7/conjson"
	// "github.com/Rican7/conjson/transform"
)

// https://stackoverflow.com/a/38596248
const TIME_LAYOUT_ISO8601 string = "2006-01-02T15:04:05-0700"

// TODO: better error handling.
//       im lost between `log.Fatal()` and `return menusInfo, err`
func HandleMenusResponse(response *http.Response) (menusInfo menus.MenusInfo, err error) {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Llongfile)

	defer response.Body.Close()

	var data []byte
	if response.StatusCode == http.StatusOK {
		data, err = io.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			return menusInfo, err
		}
	}

	// var unmarshaler json.Unmarshaler = conjson.NewUnmarshaler(&menusInfo, transform.ConventionalKeys())
	err = json.Unmarshal(data, &menusInfo)
	if err != nil {
		log.Println(err)
		return menusInfo, err
	}

	return menusInfo, err
}

type Client interface {
	GetMenus() (menusInfo menus.MenusInfo, err error)
}

type ToastHTTPClient struct {
	URL string
	Token string
	RestaurantID string
	LastUpdated time.Time
}

func (client *ToastHTTPClient) GetMenus() (menusInfo menus.MenusInfo, err error) {
	httpClient := http.Client{}

	req, err := http.NewRequest("GET", client.URL, nil)
	if err != nil {
		return menusInfo, err
	}

	req.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", client.Token)},
		"Toast-Restaurant-External-ID": {client.RestaurantID},
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return menusInfo, err
	}

	return HandleMenusResponse(response)
}
