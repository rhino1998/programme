package main

import (
	"fmt"
	"github.com/rhino1998/programme/backend/lib/model/routes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"googlemaps.github.io/maps"
	"log"
	"net"
)

type RoutesServer struct {
}

func (s *RoutesServer) CalcTravelTime(ctx context.Context, l *routes.Locations) (*routes.Trips, error) {
	c, err := maps.NewClient(maps.WithAPIKey("GOOGLE_API_KEY"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	origins := make([]string, len(l.Locations))
	destinations := make([]string, len(origins))
	for i := range origins {
		origins[i] = fmt.Sprintf("%f, %f", l.Locations[i].Coords.Latitude, l.Locations[i].Coords.Longitude)
		destinations[i] = origins[i]
	}

	var travelMode string
	switch l.Method {
	case 1:
		travelMode = "ModeWalking"
	case 2:
		travelMode = "ModeBicycling"
	default:
		travelMode = "ModeDriving"
	}

	r := &maps.DistanceMatrixRequest{
		Origins:      origins,
		Destinations: destinations,
		Mode:         maps.Mode(travelMode),
	}

	// routes is of type *DistanceMatrixResponse
	routesMatrix, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	rows := routesMatrix.Rows //[]maps.DistanceMatrixElementsRow
	trips := &routes.Trips{}
	for rowNum := range rows {
		for colNum := range rows[rowNum].Elements {
			if rowNum != colNum {
				trips.Trips = append(trips.Trips, TripFromMatrix(l, rows, rowNum, colNum))
			}
		}
	}

	return trips, nil
}

func TripFromMatrix(l *routes.Locations, rows []maps.DistanceMatrixElementsRow, rowNum int, colNum int) *routes.Trip {
	return &routes.Trip{
		Start: &routes.Location{
			Name: l.Locations[rowNum].Name,
			Coords: &routes.Coords{
				Latitude:  l.Locations[rowNum].Coords.Latitude,
				Longitude: l.Locations[rowNum].Coords.Longitude,
			},
		},
		End: &routes.Location{
			Name: l.Locations[colNum].Name,
			Coords: &routes.Coords{
				Latitude:  l.Locations[colNum].Coords.Latitude,
				Longitude: l.Locations[colNum].Coords.Longitude,
			},
		},
		Method:   l.Method,
		Duration: int64(rows[rowNum].Elements[colNum].Duration),
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", "8080"))

	if err != nil {
		fmt.Errorf("Error listening %v", err)
	}

	grpcServer := grpc.NewServer()
	routes.RegisterRoutesAPIServer(grpcServer, &RoutesServer{})
	grpcServer.Serve(lis)
}
