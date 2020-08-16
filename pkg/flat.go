package flatfox

type Attribute struct {
}

type Category struct {
}

type CoverImage struct {
}

type Document struct {
}

type Image struct {
}

type ObjectCategory string

const (
	APARTMENT ObjectCategory = "APARTMENT"
	HOUSE     ObjectCategory = "HOUSE"
	SHARED    ObjectCategory = "SHARED"
	PARK      ObjectCategory = "PARK"
	INDUSTRY  ObjectCategory = "INDUSTRY"
	SECONDARY ObjectCategory = "SECONDARY"
	PROPERTY  ObjectCategory = "PROPERTY"
)

type OfferType string

type Organization struct {
}

type PriceDisplayType string

type Flat struct {
	Attributes         *[]Attribute      `json:"attributes"`
	Category           *Category         `json:"category"`
	City               *string          `json:"city"`
	CoverImage         *CoverImage      `json:"cover_image"`
	CoverImageUrl      *string          `json:"cover_image_url"`
	Description        *string          `json:"description"`
	DescriptionTitle   *string          `json:"description_title"`
	Documents          *[]Document       `json:"documents"`
	Floor              *int32           `json:"floor"`
	Images             []Image          `json:"images"`
	IsFurnished        *bool             `json:"is_furnished"`
	IsSellingFurniture *bool             `json:"is_selling_furniture"`
	IsInRegion         *bool             `json:"is_in_region"`
	IsLiked            *bool             `json:"is_liked"`
	IsDisliked         *bool             `json:"is_disliked"`
	IsTemporary        *bool             `json:"is_temporary"`
	Latitude           *float64         `json:"latitude"`
	LivingSpace        *int32           `json:"livingspace"`
	Longitude          *float64         `json:"longitude"`
	MovingDate         *string          `json:"moving_date"`
	MovingDateType     *string          `json:"moving_date_type"`
	NumberOfRooms      *string          `json:"number_of_rooms"`
	ObjectCategory     *ObjectCategory   `json:"object_category"`
	ObjectType         *string          `json:"object_type"`
	OfferType          *OfferType        `json:"offer_type"`
	Organization       *Organization    `json:"organization"`
	Pk                 int32            `json:"pk"`
	PriceDisplay       *int32           `json:"price_display"`
	PriceDisplayType   *PriceDisplayType `json:"price_display_type"`
	PriceUnit          *string          `json:"price_unit"`
	PublicAddress      *string          `json:"public_address"`
	PublicSubtitle     *string          `json:"public_subtitle"`
	PublicTitle        *string          `json:"public_title"`
	RefHouse           *string          `json:"ref_house"`
	RefObject          *string          `json:"ref_object"`
	RefProperty        *string          `json:"ref_property"`
	RentCharges        *int32           `json:"rent_charges"`
	RentDisplay        *int32           `json:"rent_display"`
	RentGross          *int32           `json:"rent_gross"`
	RentNet            *int32           `json:"rent_net"`
	RentTitle          *string          `json:"rent_title"`
	ShortTitle         *string          `json:"short_title"`
	ShortUrl           *string          `json:"short_url"`
	ShowExactAddress   *bool             `json:"show_exact_address"`
	Slug               *string          `json:"slug"`
	Status             *string          `json:"status"`
	Street             *string          `json:"street"`
	SurfaceLiving      *int32           `json:"surface_living"`
	ThreadUrl          *string          `json:"thread_url"`
	Title              *string          `json:"title"`
	TourUrl            *string          `json:"tour_url"`
	Url                *string          `json:"url"`
	VideoUrl           *string          `json:"video_url"`
	WebsiteUrl         *string          `json:"website_url"`
	YearBuilt          *int32           `json:"year_built"`
	Zipcode            *int32           `json:"zipcode"`
}
