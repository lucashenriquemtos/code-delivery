package route

import "errors"
import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat float64
	Lng float64
}
type PartialRoutePosition struct {
	ID        string    `json:"routeID"`
	ClientID  string    `json:"clientID"`
	Positions []float64 `json:"position"`
	Fineshed  bool      `json:"finished"`
}

func (r *Route) LoadPositions() error {

	if r.ID == "" {
		return errors.New("route ID is not informed")
	}

	f, err := os.Open("destinations/" + r.ID + ".txt")

	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}

		lng, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}

		r.Positions = append(r.Positions, Position{Lat: lat, Lng: lng})

	}
	return nil
}
