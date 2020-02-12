package quark_lt

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

func CheckQuarkNode(host string, port int) {
	resp, err := http.Get("http://"+host+":" + strconv.Itoa(port))
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
}
