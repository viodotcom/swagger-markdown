# SAPI Backend API.

Sapi Backend is a Sapi Service that implements HTTP APIs for hotel search,
powers SAPI.search() method of Sapi SDK.

Resources: https://innovativetravel.atlassian.net/wiki/spaces/RFC/pages/2744090643/

SAPI Squad

https://findhotel.slack.com/archives/C0295LCLE4S

sapi-squad@findhotel.net

## Paths

| Path | Operations |
| --- | --- |
| [/acl](#path/acl) | GET |
| [/anchor](#path/anchor) | GET |
| [/availability](#path/availability) | GET |
| [/offer/:id](#path/offer/:id) | GET |
| [/offers](#path/offers) | GET |
| [/room](#path/room) | GET |
| [/rooms-offers](#path/rooms-offers) | GET |
| [/search](#path/search) | GET POST |


## <span id="path/acl">/acl</span>

### GET

Returns ACL decision based on the input parameters.

https://app.shortcut.com/findhotel/story/66887/add-new-endpoint-to-sapi-to-expose-acl-decision

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| UserIP | false | string | (Optional) is used to store the user's IP address | <nil> |
| UserCountry | false | string | (Optional) is used to store the user's two-letter country | <nil> |
| AnonymousID | false | string | (Optional) A cookie/browser-storage-based anonymous identifier of a user | <nil> |
| UserID | false | string | (Optional) An authenticated user ID, e.g. the Google ID of a user | <nil> |
| EmailDomain | false | string | (Optional) email domain for authenticated user as a value, if email is available. Users can't have access to CUG deals | <nil> |
**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [Decision](#definitions/Decision) |


---

## <span id="path/anchor">/anchor</span>

### GET

Returns information similar to /search but only for the anchor.

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| dayDistance | false | integer | Amount of full days from now to desired check in date (works in combination with nights parameter). | 15 |
| nights | false | integer | Number of nights of stay | 3 |
| checkIn | false | string | Check in date (YYYY-MM-DD) | 2021-10-10 |
| checkOut | false | string | Check out date (YYYY-MM-DD) | 2021-10-11 |
| lat | false | number | Latitude in degrees | <nil> |
| lon | false | number | Longitude in degrees | <nil> |
|  | false | object |  | <nil> |
| BoundingBox | false | string | topLeft and bottomRight coordinates of bounding box to perform search inside it.  The format is `LatTopLeft,LonTopLeft,LatBottomRight,LonBottomRight`  The types are all float64 numbers. | 46.650828100116044,7.123046875,45.17210966999772,1.009765625 |
| HotelID | false | string | Hotel ID for hotel search. If present, takes precedence over placeId, query and geolocation. | 1371626 |
| PlaceID | false | string | Place ID for place search. If present, takes precedence over query and geolocation. | 47319 |
| Query | false | string | Free-text query | Amsterdam city |
| cugDeals | false | array | Codes of closed user group deals to retrieve offers | signed_in,offline |
| tier | false | string | User's access tier. | member |
| label | false | string | Opaque value that will be passed to RAA for tracking purposes. | <nil> |
| deviceType | false | string | The type of the requestor's device. If it isn't specified then the server determines it from User-Agent request header. If the server couldn't determine it, then value is set to desktop. | desktop |
| countryCode | false | string | The 2-char ISO 3166 country code of a requestor. If not specified then the server determines it from the client's IP address. | <nil> |
| originId | false | string | Identifier of origin where the request was originated | c3po6twr70 |
| brand | false | string | Brand of an application that uses Sapi. Required to do RAA profile selection | findhotel |
| preferredRate | false | number | Offer’s price user saw on a CA (meta) platform | 196 |
| rooms | false | string | Rooms configuration | 2 |
| searchId | false | string | A correlation id used in Analytics to identify different searches. Sapi SDK generates a new unique value per each new user search and passes it to Sapi Backend to both /search and /offers endpoints, the same value. Value is changed when a new search initiated, check documentation for Sapi SDK for details what is considered a new search.  Sapi Backend passes it to RAA when retrieving offers.  If not provided, generated as UUID. nolint:lll | <nil> |
| sortingBoost | false | string | Indicates to boost the OSO ranking of some offers, based on the criteria in the parameter. For example freeCancellation=true:100 value will multiply the oso score by 100 for offers that have free cancellation. The boost is only supported for freeCancellation at the moment. | freeCancellation=true:100 |
| Currency | false | string | 3-char ISO currency uppercase | EUR |
| Language | false | string | Language code of a visitor | en |
| ProfileID | true | string | Profile is a set of configurations for a SAPI client | <nil> |
| Variations | false | string | Comma-separated list of AB-testing variations to apply | pp000004-tags2-b,v8th43ad-saf-search-a |
| Attributes | false | array | Comma-separated attributes to retrieve | hotelEntities |
| anonymousId | false | string | Unique ID identifying users | <nil> |
| userId | false | string | User ID is an authenticated user ID, e.g. the Google ID of a user. It is used for constructing ACL context | <nil> |
| emailDomain | false | string | User email domain is for authenticated user as a value, if email is available. | <nil> |
| screenshots | false | integer | Screenshots is the number of screenshots detected by the client | <nil> |
**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [AnchorResponse](#definitions/AnchorResponse) |


---

## <span id="path/availability">/availability</span>

### GET

Returns availability details for hotels over specified period of check-in dates.

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| token | false | string | Token is a pagination token that clients send on follow-up requests if complete=false.  Sapi uses this token to mediate long-polling response. | <nil> |
| startDate | false | string | ISO date of first check in | 2021-10-10 |
| endDate | false | string | ISO date of last check in (including). Can't be further than 61 days from StartDate. | 2021-10-11 |
| nights | false | integer | Nights is length of stay | 3 |
| cugDeals | false | array | Codes of closed user group deals to retrieve from RAA. | <nil> |
| tier | false | string | User's access tier. | <nil> |
| currency | false | string | 3-char ISO 4217 currency code. | EUR |
| hotelIds | false | array | Array of hotel ids to retrieve offers for. | <nil> |
| language | false | string | Language code that will be used for translation of strings for humans. | en |
| deviceType | false | string | The type of the requester's device. Derived from parsing User-Agent. | <nil> |
| countryCode | false | string | The 2-char ISO 3166 country code of a requester. | <nil> |
| originId | false | string | Identifier of origin of request. Mapping is configured in Offers Configuration Profile. | <nil> |
| brand | false | string | Brand of an application that uses Sapi. Required to do RAA profile selection. | <nil> |
| rooms | false | string | Room configuration for offers to retrieve. | <nil> |
| searchId | false | string | A correlation id used in Analytics to identify different searches.  Sapi SDK generates a new unique value per each new user search and passes it to Sapi Backend, the same value.  Value is changed when a new search initiated, check documentation for Sapi SDK for details what is considered a new search. Sapi Backend passes it to RAA when retrieving offers. If not provided, generated as UUID. | <nil> |
| anonymousId | false | string | Unique ID identifying users. | <nil> |
| variations | false | array | Comma-separated list of AB-testing variations to apply. | sapi4eva-tags2-b,v8th43ad-saf-search-a |
**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [AvailabilityResponse](#definitions/AvailabilityResponse) |


---

## <span id="path/offer/:id">/offer/:id</span>

### GET

Returns an offer detail for the specific offer ID.

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| id | true | string | offer ID | <nil> |
**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [Offer](#definitions/Offer) |


---

## <span id="path/offers">/offers</span>

### GET

This endpoint implements synchronous offers delivery given a list of hotel ids and other itinerary parameters.
The format of request is specified in JSON schema available at [sapi-backend/offers_query_schema.json at main · FindHotel/sapi-backend](https://github.com/FindHotel/sapi-backend/blob/main/internal/offers/schema/offers_query_schema.json)

An RAA client in Go had been implemented and exposed as a synchronous /offers HTTP endpoint, it does polling on the Backend side.

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| anchorHotelId | false | string | Hotel ID of an anchor hotel, if provided then offers for the anchor hotel will be requested in a separate polling process from RAA, using different RAA profile. | 123456 |
| checkIn | false | string | ISO date of check in | 2021-10-10 |
| checkOut | false | string | ISO date of check out | 2021-10-11 |
| clientRequestId | false | string | UUID identifier of a request that client sends. Correlation id that Sapi passes to RAA for tracking purposes. If not provided, generated as UUID for every new polling, and polling iterations will reuse the same clientRequestId. | <nil> |
| cugDeals | false | array | Codes of closed user group deals to retrieve from RAA. | <nil> |
| tier | false | string | User's access tier. | <nil> |
| currency | false | string | 3-char ISO 4217 currency code. | EUR |
| getAllOffers | false | boolean | If true, then return all offers per hotel, otherwise only top offers. | <nil> |
| hotelIds | false | array | Array of hotel ids to retrieve offers for. | <nil> |
| label | false | string | Opaque value that will be passed to RAA for tracking purposes. | <nil> |
| language | false | string | Language code that will be used for translation of strings for humans. | en |
| deviceType | false | string | The type of the requester's device. Derived from parsing User-Agent. Deprecated as of 2022-02-07, so clients can stop producing it - Sapi Backend would determine it from Visitor Context. | <nil> |
| countryCode | false | string | The 2-char ISO 3166 country code of a requester. Deprecated as of 2022-02-07, so clients can stop producing it - Sapi Backend would determine it from Visitor Context. | <nil> |
| originId | false | string | TODO: rename to referrerId. See https://app.shortcut.com/findhotel/story/48084/  Identifier of origin of request. Mapping is configured in Offers Configuration Profile. | <nil> |
| brand | false | string | Brand of an application that uses Sapi. Required to do RAA profile selection. | <nil> |
| preferredRate | false | number | The rate amount that visitor saw on a CA platform. If passed, it is attached to a request for anchorHotelId. | <nil> |
| rooms | false | string | Room configuration for offers to retrieve. | <nil> |
| searchId | false | string | A correlation id used in Analytics to identify different searches. Sapi SDK generates a new unique value per each new user search and passes it to Sapi Backend to both /search and /offers endpoints, the same value. Value is changed when a new search initiated, check documentation for Sapi SDK for details what is considered a new search.  Sapi Backend passes it to RAA when retrieving offers.  If not provided, generated as UUID. nolint:lll | <nil> |
| sortingBoost | false | string | Indicates to boost the OSO ranking of some offers, based on the criteria in the parameter. For example `freeCancellation=true:100` value will multiply the oso score by 100 for offers that have free cancellation. The boost is only supported for freeCancellation at the moment. | freeCancellation=true:100 |
| anonymousId | false | string | Unique ID identifying users. | <nil> |
| variations | false | array | Comma-separated list of AB-testing variations to apply. | pp000004-tags2-b,v8th43ad-saf-search-a |
| metadata | false | string | Metadata is an url encoded values contains additional information about the request. They will be sent to RAA as part of the url parameters in the request in the form of key=value. this value should contain a list of key/value pairs, for example we can use it to pass esd and epv values. so if we want to pass esd=xyz&epv=abc to RAA the value of metadata should be esd%3Dxyz%26epv%3Dabc | metadata=esd%3Dxyz%26epv%3Dabc |
**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [OffersResponse](#definitions/OffersResponse) |


---

## <span id="path/room">/room</span>

### GET

The endpoint returns room content for the given room id and providerCode.

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| id | true | string | Room ID | <nil> |
| provider | false | string | Provider Code, uppercase. If provided, it reduces the chance of hash collisions | <nil> |
**Responses:**

| Status Code | Description |
| --- | --- |
| 404 | [ErrorResponse](#definitions/ErrorResponse) |
| 200 | [RoomContent](#definitions/RoomContent) |
| 400 | [ErrorResponse](#definitions/ErrorResponse) |


---

## <span id="path/rooms-offers">/rooms-offers</span>

### GET

Implements proxy to RAA’s /room endpoint enriched with Room Content.

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| dayDistance | false | integer | Amount of full days from now to desired check in date (works in combination with nights parameter). | 15 |
| nights | false | integer | Number of nights of stay | 3 |
| checkIn | false | string | Check in date (YYYY-MM-DD) | 2021-10-10 |
| checkOut | false | string | Check out date (YYYY-MM-DD) | 2021-10-11 |
| language | false | string | Language code that will be used for translation of strings for humans. | en |
| currency | false | string | 3-char ISO 4217 currency code. | EUR |
| clickedOfferId | false | string | SRP offer id user clicked on | <nil> |
| clickedOfferBaseRate | false | number | Base rate value of the offer which user clicked on SRP | <nil> |
| clickedOfferTaxes | false | number | Taxes value of the offer which user clicked on SRP | <nil> |
| clickedOfferHotelFees | false | number | Hotel fees value of the offer which the user clicked on SRP | <nil> |
| hotelId | true | string | FindHotel hotel id | <nil> |
| rooms | false | string | rooms configuration | <nil> |
| userCountry | false | string | The 2-char ISO 3166 country code of a requester. If not specified then the server determines it from the client's IP address. | <nil> |
| deviceType | false | string | The type of the requester's device. If it isn't specified then the server determines it from User-Agent request header. If the server couldn't determine it, then value is set to desktop. | <nil> |
| aid | true | string | an identifier of a visitor using the service, also known as anonymous id | <nil> |
| cugDeals | false | array | comma-separated list of closed-user group deals, passed as is to RAA examples: cugDeals=offline,signed_in,sensitive | <nil> |
| tier | false | string | user access tier passed as is to RAA | tier=plus |
| providerCode | false | string | Provider Code | <nil> |
| squashing | false | string | if 1 then do perform squashing of rooms based on master_id, if not provided then squashing is not performed (default) | <nil> |
| deduplication | false | string | If 0 it disables the de-duplication feature. (it is enabled by default) deduplicateOffers returns unique offers. Input is array of offers for the same room. Returned offers are not sorted. Details at https://app.shortcut.com/findhotel/story/47743/ | <nil> |
| withoutIrrelevant | false | string | If 0 it disables the withoutIrrelevant feature. (it is enabled by default) withoutIrrelevant takes offers within the same room and returns only such offers that have better conditions across dimensions. If the free cancellation offer is cheaper than or equal to non-refundable offer than just remove the non-refundable offer. This rule should be applied when there is no or equal services as well as same pay later options offered on both offers. | <nil> |
| trafficSource | false | string | Visitor's traffic source | <nil> |
| label | false | string | Visitor's traffic source in label | <nil> |
| variations | false | array | Comma-separated list of AB-testing variations to apply. | pp000004-tags2-b,v8th43ad-saf-search-a |
| preHeat | false | string | Whether we need to enable preheating feature or not | <nil> |
| userId | false | string | User ID is an authenticated user ID, e.g. the Google ID of a user. It is used for constructing ACL context. | <nil> |
| emailDomain | false | string | User email domain is for authenticated user as a value, if email is available. | <nil> |
| searchId | false | string | Unique identifier of the current search used for analytical purposes. | <nil> |
| profileId | true | string | Profile is a set of configurations for a SAPI client | <nil> |
**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [RoomsResponse](#definitions/RoomsResponse) |


---

## <span id="path/search">/search</span>

### GET

Searches for hotels with their information.

#### Passing arrays
arrays should be passed with comma-separated strings no mater if they are numbers or strings.
There is no need to put strings in the `""` or `”`.
for example
```
boundingBox=46.650828100116044,7.123046875,45.17210966999772,1.009765625
attributes=hotelEntities,offerEntities,hotelIds
```

#### Stay parameters
There are four stay parameters `dayDistance`,`nights`,`checkIn` and `checkOut`.
According to these parameters `Stay` struct will be initialized.
```go
	type Stay struct {
		CheckIn  Date
		CheckOut Date
		Nights   int
	}
```
 The logic is as follows

- If `checkIn != empty` and `checkOut != empty` => `Nights = CheckOut - CheckIn`
- If `checkIn != empty` and `nights != nil` => `CheckOut = CheckIn + Nights`
- If `checkIn == empty` and `checkOut != empty` and `nights != nil` => `CheckIn = CheokOut - Nights`
- If `nights != nil` and `dayDistance != nil` => `CheckIn = Now() + dayDistance` and then `CheokOut = CheckIn + Nights`
- If `dayDistance != nil` and `checkOut != empty` => `CheckIn = Now() + dayDistance` and then `Nights = CheckOut - CheckIn`
- Otherwise => `CheckIn = Now() + dayDistance + the_days_til_the_last_day_of_the_week` and `CheckOut = CheckIn + 1` and `Nights = 1`

#### Sample search request
```
GET http://dikcjxfwieazv.cloudfront.net/search?offset=0&profileId=findhotel-website&hotelId=1096743&checkIn=2022-03-19&checkOut=2022-03-20&rooms=2&currency=USD&language=en&variations=pp000004-tags2-a%2Cc48b82da-sapi-backend-b%2Ca41n36ps-sapi-offers-a&anonymousId=ecfedb21-ed3f-4a28-aff4-119ef0a04fa1
```

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| dayDistance | false | integer | Amount of full days from now to desired check in date (works in combination with nights parameter). | 15 |
| nights | false | integer | Number of nights of stay | 3 |
| checkIn | false | string | Check in date (YYYY-MM-DD) | 2021-10-10 |
| checkOut | false | string | Check out date (YYYY-MM-DD) | 2021-10-11 |
| lat | false | number | Latitude in degrees | <nil> |
| lon | false | number | Longitude in degrees | <nil> |
|  | false | object |  | <nil> |
| boundingBox | false | string | topLeft and bottomRight coordinates of bounding box to perform search inside it.  The format is `LatTopLeft,LonTopLeft,LatBottomRight,LonBottomRight`  The types are all float64 numbers. | 46.650828100116044,7.123046875,45.17210966999772,1.009765625 |
| hotelId | false | string | Hotel ID for hotel search. If present, takes precedence over placeId, query and geolocation. | 1371626 |
| placeId | false | string | Place ID for place search. If present, takes precedence over query and geolocation. | 47319 |
| query | false | string | Free-text query | Amsterdam city |
| offset | false | integer | The first offset results will be skipped from the returned results.  Used for pagination. | 0 |
| pagesize | false | integer | Desired page size by the client. Use pagesize=0&hotelId=<id> to return anchor hotel only. Omitted pagesize (default) means the service decides the pagesize. | 0 |
| currency | false | string | 3-char ISO currency uppercase | EUR |
| language | false | string | Language code of a visitor | en |
| profileId | true | string | Profile is a set of configurations for a SAPI client | <nil> |
| attributes | false | array | Comma-separated attributes to retrieve | hotelEntities |
| sortField | false | string | Defines the sort by criteria | popularity |
| sortOrder | false | string | Defines the sort order Note: If equals to ascending (default value), then MagicSort is not enabled and defined by the AB-test configuration or sapiOverride, if it equals to the name of configuration magic-sort-axis in AppConfig, then use provided configuration. | ascending |
| chainIds | false | array | Comma-separated chain ids whose hotels will be promoted in the hotel rankings above the rest hotels | <nil> |
| facilities | false | array | Facility ids used for facet filtering | <nil> |
| themeIds | false | array | For facet filtering by theme ids. | 4,5 |
| guestRating | false | integer | Lower bound for filter by guestRating.overall | <nil> |
| hotelName | false | string | Name of the hotel for filter by name.<language> | <nil> |
| noHostels | false | boolean | If true, then hotels with propertyType=hostel are filtered out | <nil> |
| priceMin | false | integer | Lower boundary for filter by price | <nil> |
| priceMax | false | integer | Upper boundary for filter by price | <nil> |
| propertyTypeId | false | array | Filter by property type Beware that 0 is a valid property type. | 4,5 |
| notPropertyTypeId | false | array | Negative filter by property type | 4,5 |
| starRating | false | array | For facet filtering by star rating. | 4,5 |
| clientRequestId | false | string | UUID identifier of a request that client sends. Correlation id that Sapi passes to RAA for tracking purposes. If not provided, generated as UUID for every new polling, and polling iterations will reuse the same clientRequestId. nolint:lll | <nil> |
| cugDeals | false | array | Codes of closed user group deals to retrieve offers | signed_in,offline |
| tier | false | string | User's access tier. | member |
| label | false | string | Opaque value that will be passed to RAA for tracking purposes. | <nil> |
| deviceType | false | string | The type of the requester's device. If it isn't specified then the server determines it from User-Agent request header. If the server couldn't determine it, then value is set to desktop. | desktop |
| countryCode | false | string | The 2-char ISO 3166 country code of a requester. If not specified then the server determines it from the client's IP address. | <nil> |
| originId | false | string | Identifier of origin where the request was originated | c3po6twr70 |
| brand | false | string | Brand of an application that uses Sapi. Required to do RAA profile selection | findhotel |
| preferredRate | false | number | Offer’s price user saw on a CA (meta) platform | 196 |
| rooms | false | string | Rooms configuration | 2 |
| searchId | false | string | A correlation id used in Analytics to identify different searches. Sapi SDK generates a new unique value per each new user search and passes it to Sapi Backend to both /search and /offers endpoints, the same value. Value is changed when a new search initiated, check documentation for Sapi SDK for details what is considered a new search.  Sapi Backend passes it to RAA when retrieving offers.  If not provided, generated as UUID. nolint:lll | <nil> |
| sortingBoost | false | string | Indicates to boost the OSO ranking of some offers, based on the criteria in the parameter. For example freeCancellation=true:100 value will multiply the oso score by 100 for offers that have free cancellation. The boost is only supported for freeCancellation at the moment. | freeCancellation=true:100 |
| anonymousId | false | string | Unique ID identifying users | <nil> |
| variations | false | string | Comma-separated list of AB-testing variations to apply | pp000004-tags2-b,v8th43ad-saf-search-a |
| userId | false | string | User ID is an authenticated user ID, e.g. the Google ID of a user. It is used for constructing ACL context. | <nil> |
| emailDomain | false | string | User email domain is for authenticated user as a value, if email is available. | <nil> |
| screenshots | false | integer | Screenshots is the number of screenshots detected by the client | <nil> |
**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [SearchResponse](#definitions/SearchResponse) |


---

### POST

Searches for hotels with their information.
It works as GET /search endpoint but also accepts sapiOverride parameter in request body.

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| sapiOverride | false |  |  | <nil> |
**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [SearchResponse](#definitions/SearchResponse) |


---

## Definitions

### <span id="/definitions/ContentRoomOccupancyMaxAllowed"></span>ContentRoomOccupancyMaxAllowed

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| extra_beds | integer |  |  |
| total | integer |  |  |
| adults | integer |  |  |
| children | integer |  |  |




---

### <span id="/definitions/RoomsResponse"></span>RoomsResponse

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| rooms | [RoomsWithOffers](#/definitions/RoomsWithOffers) |  |  |




---

### <span id="/definitions/Status"></span>Status

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| anchorComplete | boolean | True if it's the last message with offers in anchor polling chain |  |
| complete | boolean |  |  |
| nonAnchorComplete | boolean | True if it's the last message with offers in non anchor polling chain |  |




---

### <span id="/definitions/AvailabilityEntry"></span>AvailabilityEntry

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| cheapestRate | [Rate](#/definitions/Rate) |  |  |
| hotelID | string |  |  |
| offers | [][Offer](#/definitions/Offer) |  |  |
| rooms | object |  |  |
| searchParams | [SearchParams](#/definitions/SearchParams) |  |  |




---

### <span id="/definitions/BedType"></span>BedType

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| id | string | Arbitrary identifier unique within the room. Uniqueness between other rooms is not guaranteed. |  |
| name | string | Human description for the type of the bed. |  |




---

### <span id="/definitions/ContentRoomImageLink"></span>ContentRoomImageLink

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| href | string |  |  |
| method | string |  |  |




---

### <span id="/definitions/HotelResponse"></span>HotelResponse

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| admDivisionLevel1 | string |  |  |
| admDivisionLevel2 | string |  |  |
| objectID | string |  |  |
| guestType | [GuestType](#/definitions/GuestType) |  |  |
| reviewCount | integer |  |  |
| _geoloc | [LatLon](#/definitions/LatLon) |  |  |
| admDivisionLevel4 | string |  |  |
| facilities | []integer |  |  |
| guestRating | [GuestRating](#/definitions/GuestRating) |  |  |
| sentiments | []integer |  |  |
| tags | [Tags](#/definitions/Tags) |  |  |
| indexedDiscountModifier | string |  |  |
| lastBooked | integer |  |  |
| placeADName | [TranslatedArray](#/definitions/TranslatedArray) |  |  |
| _rankingInfo | [RankingInfo](#/definitions/RankingInfo) |  |  |
| checkOutTime | string |  |  |
| hotelName | [TranslatedString](#/definitions/TranslatedString) |  |  |
| placeDisplayName | string |  |  |
| pricing | object |  |  |
| regularPriceRange | []integer |  |  |
| starRating | integer |  |  |
| themeIds | []integer |  |  |
| displayAddress | string |  |  |
| isDeleted | boolean |  |  |
| placeDN | [TranslatedArray](#/definitions/TranslatedArray) |  |  |
| cityID | string |  |  |
| imageURIs | []string |  |  |
| parentChainID | string |  |  |
| address | [TranslatedString](#/definitions/TranslatedString) |  |  |
| chainID | string |  |  |
| checkInTime | string |  |  |
| propertyTypeId | integer |  |  |
| country | string |  |  |
| magicRankScore | integer |  |  |
| magicRanks | [MagicSortAxes](#/definitions/MagicSortAxes) |  |  |
| admDivisionLevel3 | string |  |  |
| urls | [DatelessProviderLinks](#/definitions/DatelessProviderLinks) |  |  |




---

### <span id="/definitions/OffersResponse"></span>OffersResponse

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| clientRequestId | string |  |  |
| errors | [][Error](#/definitions/Error) |  |  |
| results | [][HotelResult](#/definitions/HotelResult) |  |  |
| status | [Status](#/definitions/Status) |  |  |




---

### <span id="/definitions/AnchorResponse"></span>AnchorResponse

AnchorResponse is a response from /anchor handler.

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| anchor |  | Anchor object based on the request. - If `HotelID != ""` => it gets Anchor by Hotel (objectID: "hotel:[hotel_object_id]" objectType: "hotel") - Else if `PlaceID != ""` => it gets Anchor by Place (objectID: "place:[place_object_id]" objectType: "place") - Else if `BoundingBox != nil` => it gets Anchor by BoundingBox (objectID: "area:id" objectType: "area") - Else if `Lat != 0` and `Lon != 0` => it gets Anchor by Nearby (objectID: "point:id" objectType: "point") - Else it gets Anchor by the `Query` |  |
| anchorHotelId | string | If the SearchType is `hotel` and we have a hotel object in our Anchor, it would be the ID of that hotel |  |
| anchorType | string | AnchorType is either `hotel` or `place` |  |
| exchangeRates | object | Map of exchange rates for `EUR` and the user specified currency |  |
| hotelEntities | [HotelEntities](#/definitions/HotelEntities) |  |  |
| lov | [][Item](#/definitions/Item) |  |  |
| searchParameters | [AnchorRequest](#/definitions/AnchorRequest) |  |  |




---

### <span id="/definitions/ContentBedrooms"></span>ContentBedrooms

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| description | string | Bedroom description |  |
| name | string | Name of bedroom |  |
| bed_configurations | [][ContentBedroomsBedConfigurations](#/definitions/ContentBedroomsBedConfigurations) | How beds are configured in the bedroom |  |




---

### <span id="/definitions/TranslatedArray"></span>TranslatedArray

**Type:** map[*]->[]string







---

### <span id="/definitions/Metadata"></span>Metadata

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| cachedAt | string |  |  |
| feedID | string |  |  |
| originalAccessTier | string | In case of offer was promoted from private to public access originalAccessTier field will store original accessTier and will not be cleared by promotions |  |
| providerCampaign | string |  |  |
| providerOfferId | string |  |  |
| providerRateType | string |  |  |




---

### <span id="/definitions/Offer"></span>Offer

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| extraParams |  |  |  |
| matchedDim | [MatchedDim](#/definitions/MatchedDim) |  |  |
| matchedOfferPriceDiff | number | In case of match, contains the absolute price diff in the same currency used to return the price. The diff is positive in case of the new price is higher and negative in case of the new price is lower. |  |
| availableRooms | integer | The number of similar rooms the provider still have available. |  |
| cancellationPenalties | [][CancellationPenalty](#/definitions/CancellationPenalty) | The list of penalties applied to the cancellation of the offer. |  |
| id | string | Offer ID from Offers data model. |  |
| matchType | string ([enums](#/enums/matchType)) | Type of clicked offer matching. 'exact' means price and all terms are matched. 'by_price' means price and some of terms (but not all) are matched. 'by_terms' means all terms are matched. Terms are freeCancellation, services, room name, payLater, offerType (public or private). |  |
| providerRateType | string | The rateType in the providers terms. |  |
| services | array | List of services available for this offer. |  |
| tags | []string |  |  |
| canPayLater | boolean | Shows if the user can be charged later for the offer. |  |
| cug | []string | Access tier from Offers data model. Contains the minimum tier the user should be in order to access this offer. |  |
| isClicked | boolean | True if the offer was matched with clicked offer from the search page |  |
| providerRateId | string |  |  |
| links | [][RoomLink](#/definitions/RoomLink) | Array of one item containing a link to book the offer. |  |
| prices | [][RoomPrice](#/definitions/RoomPrice) | Array of prices containing user currency where chargeable are multiplied by number of rooms. |  |
| providerCode | string | The code of the provider that is selling the offer. |  |


## Enums

**<span id="/enums/matchType"></span>matchType:**

| matchType |
| --- |
|exact, by_price, by_terms|





---

### <span id="/definitions/Nearby"></span>Nearby

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| lat | number | Latitude in degrees |  |
| lon | number | Longitude in degrees |  |
| precision | [PrecisionRanges](#/definitions/PrecisionRanges) |  |  |




---

### <span id="/definitions/SearchResponse"></span>SearchResponse

SearchResponse is a response from /search handler.

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| facets | object | Map of Facets returned from the search result |  |
| hasMoreResults | boolean | HasMoreResults shows if there are more results exist for the given search request, and client can request more by providing offset query parameter. |  |
| resultsCount | integer | These values are needed so client calculate whether there are more results by formula hasMoreResults = resultsCount + offset <= resultsCountTotal. Deprecated as of v1.2.x - clients need to migrate to HasMoreResults attribute. |  |
| searchParameters | [SearchRequest](#/definitions/SearchRequest) |  |  |
| anchorHotelId | string | If the SearchType is `hotel` and we have a hotel object in our Anchor, it would be the ID of that hotel |  |
| anchorType | string | AnchorType is either `hotel` or `place` |  |
| hotelIds | []string | List of all hotel IDs in the search result |  |
| lov | [][Item](#/definitions/Item) |  |  |
| offerEntities | [OffersMap](#/definitions/OffersMap) |  |  |
| offset | integer | These values are needed so client calculate whether there are more results by formula hasMoreResults = resultsCount + offset <= resultsCountTotal. Deprecated as of v1.2.x - clients need to migrate to HasMoreResults attribute. |  |
| resultsCountTotal | integer | These values are needed so client calculate whether there are more results by formula hasMoreResults = resultsCount + offset <= resultsCountTotal. Deprecated as of v1.2.x - clients need to migrate to HasMoreResults attribute. |  |
| anchor |  | Anchor object based on the request. - If `HotelID != ""` => it gets Anchor by Hotel (objectID: "hotel:[hotel_object_id]" objectType: "hotel") - Else if `PlaceID != ""` => it gets Anchor by Place (objectID: "place:[place_object_id]" objectType: "place") - Else if `BoundingBox != nil` => it gets Anchor by BoundingBox (objectID: "area:id" objectType: "area") - Else if `Lat != 0` and `Lon != 0` => it gets Anchor by Nearby (objectID: "point:id" objectType: "point") - Else it gets Anchor by the `Query` |  |
| hotelEntities | [HotelEntities](#/definitions/HotelEntities) |  |  |
| hotelsHaveStaticPosition | boolean | HotelsHaveStaticPosition reflects whether the hotel position will depend on offers returned from RAA. |  |
| sortingInfo | [SortingInfo](#/definitions/SortingInfo) |  |  |
| exchangeRates | object | Map of exchange rates for `EUR` and the user specified currency |  |
| searchType | [Type](#/definitions/Type) |  |  |




---

### <span id="/definitions/Tags"></span>Tags

Tags is a slice of Tag.  Usually it is materialized for a single
hotel, however it's not checked in the code.

[][Tag](#/definitions/Tag)





---

### <span id="/definitions/TranslatedString"></span>TranslatedString

**Type:** map[*]->string





---

### <span id="/definitions/AvailabilityResponse"></span>AvailabilityResponse

AvailabilityResponse models response of GET /availability endpoint.

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| availability | object |  |  |
| status | object |  |  |




---

### <span id="/definitions/ContentBedroomsBedConfigurations"></span>ContentBedroomsBedConfigurations

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| count | integer |  |  |
| description | string |  |  |
| name | string |  |  |
| size | string |  |  |
| type | string |  |  |




---

### <span id="/definitions/GuestRating"></span>GuestRating

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| pricing | number |  |  |
| rooms | number |  |  |
| service | number |  |  |
| cleanliness | number |  |  |
| dining | number |  |  |
| facilities | number |  |  |
| location | number |  |  |
| overall | number |  |  |




---

### <span id="/definitions/RoomsWithOffers"></span>RoomsWithOffers

[][RoomWithOffers](#/definitions/RoomWithOffers)





---

### <span id="/definitions/SearchQuery"></span>SearchQuery

SearchQuery is a RAA URL search query parameters.

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| CountryCode | string |  |  |
| Label | string |  |  |
| OffersCount | integer |  |  |
| PreferredRate | number |  |  |
| Rooms | string |  |  |
| AnonymousID | string |  |  |
| Destination | []string |  |  |
| RoomLimit | integer |  |  |
| SortingBoost | string |  |  |
| TopOffersCount | integer |  |  |
| CheckIn | string |  |  |
| CheckOut | string |  |  |
| Locale | string |  |  |
| Metadata | string |  |  |
| SearchID | string |  |  |
| UserAgent | string |  |  |
| UserIP | string |  |  |
| ClientRequestID | string |  |  |
| CugDeals | []string |  |  |
| Currency | string |  |  |
| DeviceType | [DeviceType](#/definitions/DeviceType) |  |  |
| Tier | string |  |  |




---

### <span id="/definitions/Chargeable"></span>Chargeable

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| base | string | Base Rate of the offer without considering taxes and fees. |  |
| taxes | string | Contains the amount of taxes to be paid for this offer. |  |




---

### <span id="/definitions/ErrorResponse"></span>ErrorResponse

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| message | string |  |  |




---

### <span id="/definitions/RAARoom"></span>RAARoom

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| capacity | integer | capacity of the room |  |
| language | string |  |  |
| name | string |  |  |




---

### <span id="/definitions/StayRequest"></span>StayRequest

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| dayDistance | integer | Amount of full days from now to desired check in date (works in combination with nights parameter). |  |
| nights | integer | Number of nights of stay |  |
| checkIn | string | Check in date (YYYY-MM-DD) | 2021-10-10 |
| checkOut | string | Check out date (YYYY-MM-DD) | 2021-10-11 |




---

### <span id="/definitions/HotelResult"></span>HotelResult

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| rooms | object |  |  |
| anchorRate | [Rate](#/definitions/Rate) |  |  |
| availableOffersCount | integer | Total number of offers available. |  |
| discount | [Discount](#/definitions/Discount) |  |  |
| fetchedAllOffers | boolean | complete flag at the the hotel. |  |
| hasMoreOffers | boolean | HasMoreOffers is true when there are more offers than topOfferLimit available. |  |
| id | string |  |  |
| offers | [][Offer](#/definitions/Offer) |  |  |




---

### <span id="/definitions/Image"></span>Image

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| url | string |  |  |




---

### <span id="/definitions/ContentRoomAmenity"></span>ContentRoomAmenity

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| id | string |  |  |
| name | [TranslatedString](#/definitions/TranslatedString) |  |  |




---

### <span id="/definitions/Hotel"></span>Hotel

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| isDeleted | boolean |  |  |
| lastBooked | integer |  |  |
| sentiments | []integer |  |  |
| themeIds | []integer |  |  |
| admDivisionLevel2 | string |  |  |
| admDivisionLevel3 | string |  |  |
| guestType | [GuestType](#/definitions/GuestType) |  |  |
| hotelName | [TranslatedString](#/definitions/TranslatedString) |  |  |
| chainID | string |  |  |
| magicRankScore | integer |  |  |
| objectID | string |  |  |
| placeADName | [TranslatedArray](#/definitions/TranslatedArray) |  |  |
| urls | [DatelessProviderLinks](#/definitions/DatelessProviderLinks) |  |  |
| _rankingInfo | [RankingInfo](#/definitions/RankingInfo) |  |  |
| admDivisionLevel4 | string |  |  |
| checkInTime | string |  |  |
| facilities | []integer |  |  |
| imageURIs | []string |  |  |
| magicRanks | [MagicSortAxes](#/definitions/MagicSortAxes) |  |  |
| propertyTypeId | integer |  |  |
| reviewCount | integer |  |  |
| admDivisionLevel1 | string |  |  |
| country | string |  |  |
| address | [TranslatedString](#/definitions/TranslatedString) |  |  |
| checkOutTime | string |  |  |
| cityID | string |  |  |
| guestRating | [GuestRating](#/definitions/GuestRating) |  |  |
| parentChainID | string |  |  |
| placeDN | [TranslatedArray](#/definitions/TranslatedArray) |  |  |
| starRating | integer |  |  |
| _geoloc | [LatLon](#/definitions/LatLon) |  |  |
| pricing | object |  |  |
| tags | [Tags](#/definitions/Tags) |  |  |




---

### <span id="/definitions/RoomWithOffers"></span>RoomWithOffers

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| smokingOptionsAvailable | boolean |  |  |
| squashedIds | []string | List of Squashed IDs |  |
| amenities | []string | All amenities available in the room. |  |
| id | string | Identifier which is constructed by hashing of room attributes like provider code, provider hotel id and provider room id. Content's RoomDB is addressable by this identifier, i.e. knowing id it's possible to retrieve the room content. |  |
| masterId | string | Identifier for a room after room-level mapping. The room-level mapping operates on a room level and maps together multiple rooms coming from different providers, or from the same provider. |  |
| name | string | Room name in the requested language, if no requested language available, it falls back to English. |  |
| offers | [Offers](#/definitions/Offers) |  |  |
| raaName | string | Room name from RAA. |  |
| bedTypes | [][BedType](#/definitions/BedType) | Array of bed types that exist in the room. |  |
| description | string | Room description. |  |
| hasClickedOffer | boolean | True if the room contains an offer from the search page. |  |
| images | [][Image](#/definitions/Image) | List of rooms images URLs. |  |
| occupationPerRoom | integer | Maximum number of people that can stay in the room. |  |




---

### <span id="/definitions/SearchRequest"></span>SearchRequest

SearchRequest defines URL query parameters for incoming request to
search endpoint.

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| pagesize | integer | Desired page size by the client. Use pagesize=0&hotelId=<id> to return anchor hotel only. Omitted pagesize (default) means the service decides the pagesize. |  |
| query | string | Free-text query | Amsterdam city |
| variations | string | Comma-separated list of AB-testing variations to apply | pp000004-tags2-b,v8th43ad-saf-search-a |
| clientRequestId | string | UUID identifier of a request that client sends. Correlation id that Sapi passes to RAA for tracking purposes. If not provided, generated as UUID for every new polling, and polling iterations will reuse the same clientRequestId. nolint:lll |  |
| dayDistance | integer | Amount of full days from now to desired check in date (works in combination with nights parameter). |  |
| originId | string ([enums](#/enums/originId)) | Identifier of origin where the request was originated | c3po6twr70 |
| offset | integer | The first offset results will be skipped from the returned results.  Used for pagination. |  |
| precision | [PrecisionRanges](#/definitions/PrecisionRanges) |  |  |
| priceMin | integer | Lower boundary for filter by price |  |
| profileId | string | Profile is a set of configurations for a SAPI client |  |
| searchId | string | A correlation id used in Analytics to identify different searches. Sapi SDK generates a new unique value per each new user search and passes it to Sapi Backend to both /search and /offers endpoints, the same value. Value is changed when a new search initiated, check documentation for Sapi SDK for details what is considered a new search.  Sapi Backend passes it to RAA when retrieving offers.  If not provided, generated as UUID. nolint:lll |  |
| hotelName | string | Name of the hotel for filter by name.<language> |  |
| language | string ([enums](#/enums/language)) | Language code of a visitor | en |
| nights | integer | Number of nights of stay |  |
| tier | string | User's access tier. | member |
| emailDomain | string | User email domain is for authenticated user as a value, if email is available. |  |
| rooms | string | Rooms configuration | 2 |
| sortingBoost | string | Indicates to boost the OSO ranking of some offers, based on the criteria in the parameter. For example freeCancellation=true:100 value will multiply the oso score by 100 for offers that have free cancellation. The boost is only supported for freeCancellation at the moment. | freeCancellation=true:100 |
| brand | string ([enums](#/enums/brand)) | Brand of an application that uses Sapi. Required to do RAA profile selection | findhotel |
| checkIn | string | Check in date (YYYY-MM-DD) | 2021-10-10 |
| currency | string | 3-char ISO currency uppercase | EUR |
| notPropertyTypeId | []integer | Negative filter by property type | 4,5 |
| placeId | string | Place ID for place search. If present, takes precedence over query and geolocation. | 47319 |
| sortField | string ([enums](#/enums/sortField)) | Defines the sort by criteria | popularity |
| anonymousId | string | Unique ID identifying users |  |
| countryCode | string | The 2-char ISO 3166 country code of a requester. If not specified then the server determines it from the client's IP address. |  |
| facilities | []integer | Facility ids used for facet filtering |  |
| themeIds | []integer | For facet filtering by theme ids. | 4,5 |
| userId | string | User ID is an authenticated user ID, e.g. the Google ID of a user. It is used for constructing ACL context. |  |
| guestRating | integer | Lower bound for filter by guestRating.overall |  |
| noHostels | boolean | If true, then hotels with propertyType=hostel are filtered out |  |
| preferredRate | number | Offer’s price user saw on a CA (meta) platform |  |
| priceMax | integer | Upper boundary for filter by price |  |
| sortOrder | string | Defines the sort order Note: If equals to ascending (default value), then MagicSort is not enabled and defined by the AB-test configuration or sapiOverride, if it equals to the name of configuration magic-sort-axis in AppConfig, then use provided configuration. | ascending |
| boundingBox | string | topLeft and bottomRight coordinates of bounding box to perform search inside it.  The format is `LatTopLeft,LonTopLeft,LatBottomRight,LonBottomRight`  The types are all float64 numbers. | 46.650828100116044,7.123046875,45.17210966999772,1.009765625 |
| checkOut | string | Check out date (YYYY-MM-DD) | 2021-10-11 |
| deviceType | string ([enums](#/enums/deviceType)) | The type of the requester's device. If it isn't specified then the server determines it from User-Agent request header. If the server couldn't determine it, then value is set to desktop. | desktop |
| starRating | []integer | For facet filtering by star rating. | 4,5 |
| attributes | []string | Comma-separated attributes to retrieve | hotelEntities |
| lon | number | Longitude in degrees |  |
| label | string | Opaque value that will be passed to RAA for tracking purposes. |  |
| lat | number | Latitude in degrees |  |
| propertyTypeId | []integer | Filter by property type Beware that 0 is a valid property type. | 4,5 |
| screenshots | integer | Screenshots is the number of screenshots detected by the client |  |
| chainIds | []string | Comma-separated chain ids whose hotels will be promoted in the hotel rankings above the rest hotels |  |
| cugDeals | []string ([enums](#/enums/cugDeals)) | Codes of closed user group deals to retrieve offers | signed_in,offline |
| hotelId | string | Hotel ID for hotel search. If present, takes precedence over placeId, query and geolocation. | 1371626 |


## Enums

**<span id="/enums/originId"></span>originId:**

| originId |
| --- |
|c3po6twr70, r2d2m73kn8, ig88zpd1k7, bb8lf9nscr|

**<span id="/enums/language"></span>language:**

| language |
| --- |
|ar, da, de, en, es, fi, fr, he, hu, id, it, iw, ja, ko, ms, nb, nl, no, nn, pl, pt, pt-BR, ru, sv, th, tr, zh, zh-CN, zh-HK, zh-TW|

**<span id="/enums/brand"></span>brand:**

| brand |
| --- |
|findhotel, etrip, vio|

**<span id="/enums/sortField"></span>sortField:**

| sortField |
| --- |
|popularity, price, privateDeals, guestRating|

**<span id="/enums/deviceType"></span>deviceType:**

| deviceType |
| --- |
|desktop, mobile, tablet|

**<span id="/enums/cugDeals"></span>cugDeals:**

| cugDeals |
| --- |
|signed_in, offline, sensitive, prime, backup|





---

### <span id="/definitions/DeviceType"></span>DeviceType

**Type:** string



---

### <span id="/definitions/RoomContent"></span>RoomContent

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| amenities | [][ContentRoomAmenity](#/definitions/ContentRoomAmenity) | Room amenities |  |
| fht_id | string | FindHotel Room ID |  |
| images | [][ContentRoomImage](#/definitions/ContentRoomImage) | Room images |  |
| master_id | string | MasterRoomId is an identifier for a room after room-level mapping. The room-level mapping operates on a room level and maps together multiple rooms coming from different providers, or from the same provider. |  |
| room_info | [ContentRoomRoomInfo](#/definitions/ContentRoomRoomInfo) |  |  |
| description | [TranslatedString](#/definitions/TranslatedString) |  |  |
| hotel_id | string | Provider hotel id |  |
| name | [TranslatedString](#/definitions/TranslatedString) |  |  |
| occupancy | [ContentRoomOccupancy](#/definitions/ContentRoomOccupancy) |  |  |
| provider_code | string | Room Provider Code |  |
| sanitized_name | [TranslatedString](#/definitions/TranslatedString) |  |  |




---

### <span id="/definitions/ContentRoomRoomInfoCount"></span>ContentRoomRoomInfoCount

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| bedrooms | integer |  |  |
| bathrooms | integer |  |  |




---

### <span id="/definitions/Variations"></span>Variations

[]string





---

### <span id="/definitions/DatelessProviderLinks"></span>DatelessProviderLinks

[][DatelessProviderLink](#/definitions/DatelessProviderLink)





---

### <span id="/definitions/Discount"></span>Discount

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| discountProvider | string |  |  |
| hasDiscountProvider | boolean |  |  |
| hasParityProvider | boolean |  |  |
| modifier | string |  |  |




---

### <span id="/definitions/HotelEntities"></span>HotelEntities

HotelEntities is a map of Hotel Entities with tags which are relevant to the requested stay

**Type:** map[*]->[#HotelResponse](#/definitions/HotelResponse)





---

### <span id="/definitions/Offers"></span>Offers

[][Offer](#/definitions/Offer)





---

### <span id="/definitions/RankingInfo"></span>RankingInfo

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| words | integer |  |  |
| proximityDistance | integer |  |  |
| userScore | integer |  |  |
| filters | integer |  |  |
| firstMatchedWord | integer |  |  |
| geoDistance | integer |  |  |
| geoPrecision | integer |  |  |
| nbExactWords | integer |  |  |
| nbTypos | integer |  |  |




---

### <span id="/definitions/Room"></span>Room

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| description | string | Room description. |  |
| id | string | Identifier which is constructed by hashing of room attributes like provider code, provider hotel id and provider room id. Content's RoomDB is addressable by this identifier, i.e. knowing id it's possible to retrieve the room content. |  |
| images | [][Image](#/definitions/Image) | List of rooms images URLs. |  |
| masterId | string | Identifier for a room after room-level mapping. The room-level mapping operates on a room level and maps together multiple rooms coming from different providers, or from the same provider. |  |
| squashedIds | []string | List of Squashed IDs |  |
| amenities | []string | All amenities available in the room. |  |
| bedTypes | [][BedType](#/definitions/BedType) | Array of bed types that exist in the room. |  |
| name | string | Room name in the requested language, if no requested language available, it falls back to English. |  |
| occupationPerRoom | integer | Maximum number of people that can stay in the room. |  |
| raaName | string | Room name from RAA. |  |
| smokingOptionsAvailable | boolean |  |  |




---

### <span id="/definitions/CancellationPenalty"></span>CancellationPenalty

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| end | string | The datetime of when this policy ends. |  |
| start | string | The datetime of when this policy starts to be applied. |  |
| amount | number | The amount of money that is charged in specified currency in case of a cancellation in this policy. |  |
| currency | string | The currency of which the penalty will be charged if amount is specified. |  |




---

### <span id="/definitions/ContentRoomImage"></span>ContentRoomImage

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| hero_image | boolean | Whether this is the room hero(main) image |  |
| id | string | Image ID |  |
| link | [ContentRoomImageLink](#/definitions/ContentRoomImageLink) |  |  |




---

### <span id="/definitions/Error"></span>Error

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| message | string |  |  |
| params | object |  |  |
| priority | integer |  |  |
| providerCode | string |  |  |
| type | integer |  |  |




---

### <span id="/definitions/Rate"></span>Rate

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| base | number | The rate exclusive of any taxes and hotel fees. |  |
| hotelFees | number | The sum of all mandatory taxes and fees that the customer will need to pay at the hotel. For example, a resort fee. |  |
| taxes | number | Value added tax (VAT). |  |




---

### <span id="/definitions/Tag"></span>Tag

Tag is a string of a form "xYYMMDD-N" where x is one-char modifier
and YYMMDD is a date of check in in YYMMDD format, N is the number
of nights, integer value.  Tags present on a hotel level and
implement an offers indexing process.  Tags added when indexing
offers stream. Missing tag for a given stay means the were no
requests for offers.

**Type:** string



---

### <span id="/definitions/ContentRoomOccupancy"></span>ContentRoomOccupancy

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| max_allowed | [ContentRoomOccupancyMaxAllowed](#/definitions/ContentRoomOccupancyMaxAllowed) |  |  |




---

### <span id="/definitions/Decision"></span>Decision

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| source | string ([enums](#/enums/source)) | Source of the ACL decision |  |
| action | string ([enums](#/enums/action)) | Action which should be taken based on this decision |  |
| createdAt | string | Time of creating the decision |  |
| ruleID | string | ACL rule ID. if the rule comes from Algolia it's the same as ObjectID |  |


## Enums

**<span id="/enums/source"></span>source:**

| source |
| --- |
|Static, Live|

**<span id="/enums/action"></span>action:**

| action |
| --- |
|ALLOW, DENY, UNKNOWN|





---

### <span id="/definitions/LatLon"></span>LatLon

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| lat | number | Latitude in degrees |  |
| lon | number | Longitude in degrees |  |




---

### <span id="/definitions/MatchedDim"></span>MatchedDim

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| freeCancellation | boolean | True if matched and clicked offers both have or don't have free cancellation |  |
| offerType | boolean | True if matched and clicked offers both are public or private |  |
| payLater | boolean | True if matched and clicked offers have the same canPayLater value |  |
| price | boolean | True if total price diff percentage with clicked offer was <= 1% |  |
| room | boolean | True if matched and clicked offers have the same room name |  |
| services | boolean | True if matched and clicked offers have the same services |  |




---

### <span id="/definitions/PrecisionRange"></span>PrecisionRange

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| from | integer |  |  |
| value | integer |  |  |




---

### <span id="/definitions/PrecisionRanges"></span>PrecisionRanges

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| ranges | [][PrecisionRange](#/definitions/PrecisionRange) |  |  |




---

### <span id="/definitions/SortingInfo"></span>SortingInfo

SortingInfo includes information about the sorting(including SortType and whether MagicSort is activated)

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| ActivateMagicSort | boolean |  |  |
| MagicSortAxes | [MagicSortAxes](#/definitions/MagicSortAxes) |  |  |
| MagicSortError | string |  |  |
| SortingType | [sortingType](#/definitions/sortingType) |  |  |




---

### <span id="/definitions/ContentRoomRoomInfo"></span>ContentRoomRoomInfo

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| count | [ContentRoomRoomInfoCount](#/definitions/ContentRoomRoomInfoCount) |  |  |
| type | string | Room type |  |
| area | [ContentRoomRoomInfoArea](#/definitions/ContentRoomRoomInfoArea) |  |  |
| bedrooms | [][ContentBedrooms](#/definitions/ContentBedrooms) |  |  |




---

### <span id="/definitions/ContentRoomRoomInfoArea"></span>ContentRoomRoomInfoArea

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| square_feet | number |  |  |
| square_meters | number |  |  |




---

### <span id="/definitions/BreakdownFee"></span>BreakdownFee

[]**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| total | string |  |  |
| type | string |  |  |






---

### <span id="/definitions/Calendar"></span>Calendar

Calendar maps check in date (ISO-formatted as string) to availability entry.

**Type:** map[*]->[#AvailabilityEntry](#/definitions/AvailabilityEntry)





---

### <span id="/definitions/RoomPrice"></span>RoomPrice

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| rate | [OfferRate](#/definitions/OfferRate) |  |  |
| type | string | Price type. |  |
| chargeable | [Chargeable](#/definitions/Chargeable) |  |  |
| currencyCode | string | The currency applied to the rates of the offer, this currency is always the same as requested by the user. |  |
| hotelFees | [HotelFees](#/definitions/HotelFees) |  |  |




---

### <span id="/definitions/BoundingBox"></span>BoundingBox

BoundingBox represents a "rectangle" between provided coordinates of
top-left and bottom-right corners.

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| LatBottomRight | number |  |  |
| LatTopLeft | number |  |  |
| LonBottomRight | number |  |  |
| LonTopLeft | number |  |  |




---

### <span id="/definitions/GuestType"></span>GuestType

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| solo | integer |  |  |
| business | integer |  |  |
| couples | integer |  |  |
| families | integer |  |  |
| groups | integer |  |  |




---

### <span id="/definitions/DatelessProviderLink"></span>DatelessProviderLink

DatelessProviderLink holds raw urls pointing to provider websites
for a given hotel.

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| provider | string |  |  |
| url | string |  |  |




---

### <span id="/definitions/MagicSortAxes"></span>MagicSortAxes

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| Discount | integer | First calculate usual total rate minus cheapest total rate, then calculate dense rank from high to low. |  |
| Distance | integer | Dense rank by squared distance between current hotel and anchor location (anchor hotell coordinates / place coordinates / boundingBox center), (x1-x2)^2+(y1-y2)^2. For the future consider rounding. From low to high. |  |
| GeoDistance | integer | Dense rank by _rankingInfo.geoDistance values, from low to high. |  |
| HSO | integer | Dense rank by _rankingInfo.filters values, from high to low. |  |
| HasFHTOffer | integer | Dense rank by boolean flag whether exists a FHT Offer within top offers. |  |
| HasPrivateFHTOffer | integer | Dense rank by boolean flag whether exists a private FHT offer within top offers. |  |
| Location | integer | Dense rank by guestRating.Location from high to low. |  |
| Anchor | integer | Dense rank by similarity of "Who stays here?" - guestType struct where values are least squared difference. (attr1 - attr2)^2. From low to high. |  |
| Rating | integer | Dense rank by guestRating.Overall from high to low. |  |
| HasPrivateOffer | integer | Dense rank by boolean flag whether exists a private offer within top offers. |  |
| ReviewCount | integer | Dense rank by review count from high to low. |  |
| Disparity | integer | First calculate the difference between median and minimum total rate of top offers for each hotel. Then calculate dense rank for these values from high to low. |  |




---

### <span id="/definitions/Package"></span>Package

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| amenities | []string | An array of strings (enums) with the amenities of the offer. |  |
| canPayLater | boolean | Indicates if the user can be charged later for the offer. |  |




---

### <span id="/definitions/AnchorRequest"></span>AnchorRequest

AnchorRequest defines URL query parameters for incoming request to
anchor endpoint.

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| brand | string ([enums](#/enums/brand)) | Brand of an application that uses Sapi. Required to do RAA profile selection | findhotel |
| emailDomain | string | User email domain is for authenticated user as a value, if email is available. |  |
| searchId | string | A correlation id used in Analytics to identify different searches. Sapi SDK generates a new unique value per each new user search and passes it to Sapi Backend to both /search and /offers endpoints, the same value. Value is changed when a new search initiated, check documentation for Sapi SDK for details what is considered a new search.  Sapi Backend passes it to RAA when retrieving offers.  If not provided, generated as UUID. nolint:lll |  |
| checkIn | string | Check in date (YYYY-MM-DD) | 2021-10-10 |
| label | string | Opaque value that will be passed to RAA for tracking purposes. |  |
| lon | number | Longitude in degrees |  |
| tier | string | User's access tier. | member |
| HotelID | string | Hotel ID for hotel search. If present, takes precedence over placeId, query and geolocation. | 1371626 |
| PlaceID | string | Place ID for place search. If present, takes precedence over query and geolocation. | 47319 |
| anonymousId | string | Unique ID identifying users |  |
| sortingBoost | string | Indicates to boost the OSO ranking of some offers, based on the criteria in the parameter. For example freeCancellation=true:100 value will multiply the oso score by 100 for offers that have free cancellation. The boost is only supported for freeCancellation at the moment. | freeCancellation=true:100 |
| ProfileID | string | Profile is a set of configurations for a SAPI client |  |
| checkOut | string | Check out date (YYYY-MM-DD) | 2021-10-11 |
| cugDeals | []string ([enums](#/enums/cugDeals)) | Codes of closed user group deals to retrieve offers | signed_in,offline |
| lat | number | Latitude in degrees |  |
| Language | string ([enums](#/enums/Language)) | Language code of a visitor | en |
| nights | integer | Number of nights of stay |  |
| rooms | string | Rooms configuration | 2 |
| userId | string | User ID is an authenticated user ID, e.g. the Google ID of a user. It is used for constructing ACL context |  |
| BoundingBox | string | topLeft and bottomRight coordinates of bounding box to perform search inside it.  The format is `LatTopLeft,LonTopLeft,LatBottomRight,LonBottomRight`  The types are all float64 numbers. | 46.650828100116044,7.123046875,45.17210966999772,1.009765625 |
| Currency | string | 3-char ISO currency uppercase | EUR |
| Variations | string | Comma-separated list of AB-testing variations to apply | pp000004-tags2-b,v8th43ad-saf-search-a |
| precision | [PrecisionRanges](#/definitions/PrecisionRanges) |  |  |
| preferredRate | number | Offer’s price user saw on a CA (meta) platform |  |
| Attributes | []string | Comma-separated attributes to retrieve | hotelEntities |
| countryCode | string | The 2-char ISO 3166 country code of a requestor. If not specified then the server determines it from the client's IP address. |  |
| deviceType | string ([enums](#/enums/deviceType)) | The type of the requestor's device. If it isn't specified then the server determines it from User-Agent request header. If the server couldn't determine it, then value is set to desktop. | desktop |
| originId | string ([enums](#/enums/originId)) | Identifier of origin where the request was originated | c3po6twr70 |
| Query | string | Free-text query | Amsterdam city |
| dayDistance | integer | Amount of full days from now to desired check in date (works in combination with nights parameter). |  |
| screenshots | integer | Screenshots is the number of screenshots detected by the client |  |


## Enums

**<span id="/enums/cugDeals"></span>cugDeals:**

| cugDeals |
| --- |
|signed_in, offline, sensitive, prime, backup|

**<span id="/enums/Language"></span>Language:**

| Language |
| --- |
|ar, da, de, en, es, fi, fr, he, hu, id, it, iw, ja, ko, ms, nb, nl, no, nn, pl, pt, pt-BR, ru, sv, th, tr, zh, zh-CN, zh-HK, zh-TW|

**<span id="/enums/deviceType"></span>deviceType:**

| deviceType |
| --- |
|desktop, mobile, tablet|

**<span id="/enums/originId"></span>originId:**

| originId |
| --- |
|c3po6twr70, r2d2m73kn8, ig88zpd1k7, bb8lf9nscr|

**<span id="/enums/brand"></span>brand:**

| brand |
| --- |
|findhotel, etrip, vio|





---

### <span id="/definitions/ClickInfo"></span>ClickInfo

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| isClicked | boolean | True if the offer was matched with clicked offer from the search page |  |
| matchType | string ([enums](#/enums/matchType)) | Type of clicked offer matching. 'exact' means price and all terms are matched. 'by_price' means price and some of terms (but not all) are matched. 'by_terms' means all terms are matched. Terms are freeCancellation, services, room name, payLater, offerType (public or private). |  |
| matchedDim | [MatchedDim](#/definitions/MatchedDim) |  |  |
| matchedOfferPriceDiff | number | In case of match, contains the absolute price diff in the same currency used to return the price. The diff is positive in case of the new price is higher and negative in case of the new price is lower. |  |


## Enums

**<span id="/enums/matchType"></span>matchType:**

| matchType |
| --- |
|exact, by_price, by_terms|





---

### <span id="/definitions/OfferRate"></span>OfferRate

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| taxes | number | Contains the amount of taxes to be paid for this offer. |  |
| base | number | Base Rate of the offer without considering taxes and fees. |  |
| hotelFees | number | Other costs attributed to this offer. |  |




---

### <span id="/definitions/OffersMap"></span>OffersMap

OffersMap is offers map data based on each HotelID retrieved from RAA

**Type:** map[*]->[][Offer](#/definitions/Offer)







---

### <span id="/definitions/RoomLink"></span>RoomLink

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| href | string | The URL that the user should be redirected to book this offer. |  |
| method | string | HTTP method to use with href. |  |
| type | string | The type of link. |  |




---

### <span id="/definitions/SearchParams"></span>SearchParams

SearchParams is all parameters needed to send request to RAA for search endpoint.

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| Profile | string |  |  |
| Query | [SearchQuery](#/definitions/SearchQuery) |  |  |
| URL | string | Client settings |  |
| IsAnchor | boolean |  |  |




---

### <span id="/definitions/Type"></span>Type

Type can be either `hotel`, `place`, `map`, `nearby` or `query`

according to the request parameters (hotelID, PlaceID, BoundingBox, Nearby or Query)

**Type:** string



---

### <span id="/definitions/sortingType"></span>sortingType

**Type:** string



---

### <span id="/definitions/HotelFees"></span>HotelFees

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| breakdown | [][BreakdownFee](#/definitions/BreakdownFee) |  |  |
| total | string |  |  |




---

### <span id="/definitions/Item"></span>Item

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| categoryID | integer |  |  |
| id | integer |  |  |
| objectID | string |  |  |
| value | [TranslatedString](#/definitions/TranslatedString) |  |  |




---

