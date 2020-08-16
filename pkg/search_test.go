package flatfox_test

import (
	"errors"
	flatfox "github.com/denysvitali/go-flatfox/pkg"
	"net/http"
	"testing"
)

func TestParseFlat(t *testing.T){

}

func TestGetPin(t *testing.T){
	ffClient := flatfox.NewClient(http.DefaultClient)


	var minRooms float32 = 5.0

	flats, err := ffClient.PinSearch(10, &flatfox.Bounds{
			North: 47.4554872,
			West:  8.5439301,
			South: 47.2027764,
			East:  8.7732697,
		},
		&flatfox.PinSearchOptions{MinRooms: &minRooms},
	)

	if err != nil {
		t.Fatal(err)
	}

	if flats == nil {
		t.Fatal(errors.New("flats is nil"))
	}

	if len(*flats) == 0 {
		t.Fatal(errors.New("flats array is empty"))
	}
}