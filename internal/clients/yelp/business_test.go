package yelp

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalBusiness(t *testing.T) {
	var biz Business
	err := json.Unmarshal([]byte(peterLugerBusiness), &biz)
	require.NoError(t, err)
	assert.NotEmpty(t, biz)

	assert.Equal(t, "https://s3-media2.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/o.jpg", biz.ImageURL)
	assert.Equal(t, "Steakhouses", biz.Categories[0].Title)
	assert.Equal(t, "178 Broadway", biz.Location.Address1)
	assert.Equal(t, "Brooklyn, NY 11211", biz.Location.DisplayAddress[1])
}

const peterLugerBusiness = `{
    "id": "4yPqqJDJOQX69gC66YUDkA",
    "alias": "peter-luger-brooklyn-2",
    "name": "Peter Luger",
    "image_url": "https://s3-media2.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/o.jpg",
    "is_claimed": true,
    "is_closed": false,
    "url": "https://www.yelp.com/biz/peter-luger-brooklyn-2?adjust_creative=4sK9X2Iq_bY8O95k_OGFeg&utm_campaign=yelp_api_v3&utm_medium=api_v3_business_lookup&utm_source=4sK9X2Iq_bY8O95k_OGFeg",
    "phone": "+17183877400",
    "display_phone": "(718) 387-7400",
    "review_count": 5437,
    "categories": [
        {
            "alias": "steak",
            "title": "Steakhouses"
        }
    ],
    "rating": 4,
    "location": {
        "address1": "178 Broadway",
        "address2": "",
        "address3": "",
        "city": "Brooklyn",
        "zip_code": "11211",
        "country": "US",
        "state": "NY",
        "display_address": [
            "178 Broadway",
            "Brooklyn, NY 11211"
        ],
        "cross_streets": "Driggs Ave & 6th St"
    },
    "coordinates": {
        "latitude": 40.709945,
        "longitude": -73.962478
    },
    "photos": [
        "https://s3-media2.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/o.jpg",
        "https://s3-media3.fl.yelpcdn.com/bphoto/jOoRBpCU9_2iS8z306CGOQ/o.jpg",
        "https://s3-media4.fl.yelpcdn.com/bphoto/KRpKd1ZdcYSPRDhlb12Vjw/o.jpg"
    ],
    "price": "$$$$",
    "hours": [
        {
            "open": [
                {
                    "is_overnight": false,
                    "start": "1200",
                    "end": "2130",
                    "day": 0
                },
                {
                    "is_overnight": false,
                    "start": "1200",
                    "end": "2130",
                    "day": 1
                },
                {
                    "is_overnight": false,
                    "start": "1200",
                    "end": "2130",
                    "day": 2
                },
                {
                    "is_overnight": false,
                    "start": "1200",
                    "end": "2130",
                    "day": 3
                },
                {
                    "is_overnight": false,
                    "start": "1200",
                    "end": "2230",
                    "day": 4
                },
                {
                    "is_overnight": false,
                    "start": "1200",
                    "end": "2230",
                    "day": 5
                },
                {
                    "is_overnight": false,
                    "start": "1300",
                    "end": "2130",
                    "day": 6
                }
            ],
            "hours_type": "REGULAR",
            "is_open_now": false
        }
    ],
    "transactions": []
}`
