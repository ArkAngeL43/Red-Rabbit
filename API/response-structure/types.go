package API_RESPONSE_AUTOGENERATED_FINDBYT_F_TYPES_RR6_PROJ_GEN_MAIN_FILE_NUMB238_MODG3_DBG1_TYPEFILE_STRUCTS_OUTPUTSUPP_XML_JSON_YAML_CONF_INI_HTML

// ----------- WHOIS_XML type

type WHOIS_XML_API_RESPONSE_STRUCT1_MODG3 struct {
	WhoisRecord struct {
		CreatedDate string `json:"createdDate"`
		UpdatedDate string `json:"updatedDate"`
		ExpiresDate string `json:"expiresDate"`
		Registrant  struct {
			Organization string `json:"organization"`
			State        string `json:"state"`
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			RawText      string `json:"rawText"`
		} `json:"registrant"`
		AdministrativeContact struct {
			Organization string `json:"organization"`
			State        string `json:"state"`
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			RawText      string `json:"rawText"`
		} `json:"administrativeContact"`
		TechnicalContact struct {
			Organization string `json:"organization"`
			State        string `json:"state"`
			Country      string `json:"country"`
			CountryCode  string `json:"countryCode"`
			RawText      string `json:"rawText"`
		} `json:"technicalContact"`
		DomainName  string `json:"domainName"`
		NameServers struct {
			RawText   string        `json:"rawText"`
			HostNames []string      `json:"hostNames"`
			Ips       []interface{} `json:"ips"`
		} `json:"nameServers"`
		Status       string `json:"status"`
		RawText      string `json:"rawText"`
		ParseCode    int    `json:"parseCode"`
		Header       string `json:"header"`
		StrippedText string `json:"strippedText"`
		Footer       string `json:"footer"`
		Audit        struct {
			CreatedDate string `json:"createdDate"`
			UpdatedDate string `json:"updatedDate"`
		} `json:"audit"`
		RegistrarName         string `json:"registrarName"`
		RegistrarIANAID       string `json:"registrarIANAID"`
		CreatedDateNormalized string `json:"createdDateNormalized"`
		UpdatedDateNormalized string `json:"updatedDateNormalized"`
		ExpiresDateNormalized string `json:"expiresDateNormalized"`
		RegistryData          struct {
			DomainName  string `json:"domainName"`
			NameServers struct {
				RawText   string        `json:"rawText"`
				HostNames []string      `json:"hostNames"`
				Ips       []interface{} `json:"ips"`
			} `json:"nameServers"`
			Status    string `json:"status"`
			ParseCode int    `json:"parseCode"`
			Header    string `json:"header"`
			Footer    string `json:"footer"`
			Audit     struct {
				CreatedDate string `json:"createdDate"`
				UpdatedDate string `json:"updatedDate"`
			} `json:"audit"`
			RegistrarName         string `json:"registrarName"`
			RegistrarIANAID       string `json:"registrarIANAID"`
			CreatedDateNormalized string `json:"createdDateNormalized"`
			UpdatedDateNormalized string `json:"updatedDateNormalized"`
			ExpiresDateNormalized string `json:"expiresDateNormalized"`
			WhoisServer           string `json:"whoisServer"`
		} `json:"registryData"`
		ContactEmail       string `json:"contactEmail"`
		DomainNameExt      string `json:"domainNameExt"`
		EstimatedDomainAge int    `json:"estimatedDomainAge"`
	} `json:"WhoisRecord"`
}

type API_LAYER_PHONE struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}