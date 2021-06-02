package utils

type (
	Albums struct {
		Artist string
		Track  string
	}
	AlbumsJson struct {
		Pagination struct {
			Page    int `json:"page"`
			Pages   int `json:"pages"`
			PerPage int `json:"per_page"`
			Items   int `json:"items"`
			Urls    struct {
			} `json:"urls"`
		} `json:"pagination"`
		Results []struct {
			Country     string        `json:"country"`
			Genre       []string      `json:"genre"`
			Format      []string      `json:"format"`
			Style       []interface{} `json:"style"`
			ID          int           `json:"id"`
			Label       []string      `json:"label"`
			Type        string        `json:"type"`
			Barcode     []string      `json:"barcode"`
			MasterID    int           `json:"master_id"`
			MasterURL   interface{}   `json:"master_url"`
			URI         string        `json:"uri"`
			Catno       string        `json:"catno"`
			Title       string        `json:"title"`
			Thumb       string        `json:"thumb"`
			CoverImage  string        `json:"cover_image"`
			ResourceURL string        `json:"resource_url"`
			Community   struct {
				Want int `json:"want"`
				Have int `json:"have"`
			} `json:"community"`
			FormatQuantity int `json:"format_quantity"`
			Formats        []struct {
				Name         string   `json:"name"`
				Qty          string   `json:"qty"`
				Descriptions []string `json:"descriptions"`
			} `json:"formats"`
			Year string `json:"year,omitempty"`
		} `json:"results"`
	}
	TrackJson struct {
		ID          int    `json:"id"`
		Status      string `json:"status"`
		Year        int    `json:"year"`
		ResourceURL string `json:"resource_url"`
		URI         string `json:"uri"`
		Artists     []struct {
			Name        string `json:"name"`
			Anv         string `json:"anv"`
			Join        string `json:"join"`
			Role        string `json:"role"`
			Tracks      string `json:"tracks"`
			ID          int    `json:"id"`
			ResourceURL string `json:"resource_url"`
		} `json:"artists"`
		ArtistsSort string `json:"artists_sort"`
		Labels      []struct {
			Name           string `json:"name"`
			Catno          string `json:"catno"`
			EntityType     string `json:"entity_type"`
			EntityTypeName string `json:"entity_type_name"`
			ID             int    `json:"id"`
			ResourceURL    string `json:"resource_url"`
		} `json:"labels"`
		Series    []interface{} `json:"series"`
		Companies []interface{} `json:"companies"`
		Formats   []struct {
			Name         string   `json:"name"`
			Qty          string   `json:"qty"`
			Descriptions []string `json:"descriptions"`
		} `json:"formats"`
		DataQuality string `json:"data_quality"`
		Community   struct {
			Have   int `json:"have"`
			Want   int `json:"want"`
			Rating struct {
				Count   int     `json:"count"`
				Average float64 `json:"average"`
			} `json:"rating"`
			Submitter struct {
				Username    string `json:"username"`
				ResourceURL string `json:"resource_url"`
			} `json:"submitter"`
			Contributors []struct {
				Username    string `json:"username"`
				ResourceURL string `json:"resource_url"`
			} `json:"contributors"`
			DataQuality string `json:"data_quality"`
			Status      string `json:"status"`
		} `json:"community"`
		FormatQuantity    int           `json:"format_quantity"`
		DateAdded         string        `json:"date_added"`
		DateChanged       string        `json:"date_changed"`
		NumForSale        int           `json:"num_for_sale"`
		LowestPrice       interface{}   `json:"lowest_price"`
		Title             string        `json:"title"`
		Country           string        `json:"country"`
		Released          string        `json:"released"`
		ReleasedFormatted string        `json:"released_formatted"`
		Identifiers       []interface{} `json:"identifiers"`
		Genres            []string      `json:"genres"`
		Tracklist         []struct {
			Position string `json:"position"`
			Type     string `json:"type_"`
			Title    string `json:"title"`
			Duration string `json:"duration"`
		} `json:"tracklist"`
		Extraartists []interface{} `json:"extraartists"`
		Images       []struct {
			Type        string `json:"type"`
			URI         string `json:"uri"`
			ResourceURL string `json:"resource_url"`
			URI150      string `json:"uri150"`
			Width       int    `json:"width"`
			Height      int    `json:"height"`
		} `json:"images"`
		Thumb           string `json:"thumb"`
		EstimatedWeight int    `json:"estimated_weight"`
		BlockedFromSale bool   `json:"blocked_from_sale"`
	}
	Results struct {
		Thumb   string       `json:"cover"`
		Artist  string       `json:"artist_name"`
		Country string       `json:"country"`
		Year    string       `json:"year"`
		Genres  string       `json:"genres"`
		Tracks  []TrackAlbum `json:"tracks"`
	}
	TrackAlbum struct {
		TrackName string `json:"track"`
	}
)
