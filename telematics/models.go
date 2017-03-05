package telematics

import route4me "github.com/route4me/route4me-go-sdk"

type VendorSize string

const (
	VendorSizeGlobal   VendorSize = "global"
	VendorSizeRegional VendorSize = "regional"
	VendorSizeLocal    VendorSize = "local"
)

type VendorQuery struct {
	Size       VendorSize    `http:"size"`
	Integrated route4me.Bool `http:"is_integrated"`
	Feature    string        `http:"feature"`
	Country    string        `http:"country"`
	Search     string        `http:"search"`
	Page       uint          `http:"page"`
	PerPage    uint          `http:"per_page"`
}

type Vendor struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Title        string     `json:"title"`
	Slug         string     `json:"slug"`
	Description  string     `json:"description"`
	LogoURL      string     `json:"logo_url"`
	WebsiteURL   string     `json:"website_url"`
	APIDocsURL   string     `json:"api_docs_url"`
	IsIntegrated string     `json:"is_integrated"`
	Size         VendorSize `json:"size"`
	Features     []Feature  `json:"features"`
	Countries    []Country  `json:"countries"`
}

type Country struct {
	ID          string `json:"id"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}

type Feature struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	FeatureGroup string `json:"feature_group"`
}
