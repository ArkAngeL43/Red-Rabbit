package SUPER_Apiinfo_Shodan_Branch_one_main_type_rr6_autogenerated_file_editm_m1_tSUPER_CSUPER_DataI_TSHODAN_L1

import (
	"encoding/json"
	"fmt"
	"log"
	SUPER_TYPES "main/modules/go-main/xml/types"
	"net/http"
)

type API_Client_SHODAN struct {
	APIK string
}

var SHODAN_URL = "https://api.shodan.io"

func New(Shodan_Key string) *API_Client_SHODAN {
	return &API_Client_SHODAN{
		APIK: Shodan_Key,
	}
}

func (s *API_Client_SHODAN) APIInfo() (*SUPER_TYPES.API_Information_SHODAN, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", SHODAN_URL, s.APIK))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret SUPER_TYPES.API_Information_SHODAN
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *API_Client_SHODAN) HostSearch(q string) (*SUPER_TYPES.HostSearch, error) {
	res, err := http.Get(
		fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", SHODAN_URL, s.APIK, q),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret SUPER_TYPES.HostSearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func Run(apikey string, searchterm string) {
	s := New(apikey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.Shodan_QueryCredits,
		info.Shodan_ScanCredits)

	hostSearch, err := s.HostSearch(searchterm)
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
