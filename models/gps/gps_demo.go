//go:build demo

package gps

import (
	"embed"
	"time"
	"net/http"
	"strings"

	"github.com/merliot/dean"
	"github.com/merliot/hub/models/common"
)

type place struct {
	lat  float64
	long float64
}

var places = [...]place{
	{42.379822, -71.064941},
	{33.943203, -118.247772},
	{40.902771, -73.133850},
	{40.827454, -73.699722},
	{40.719517, -73.852211},
	{25.782721, -80.140556},
	{33.988270, -118.472023},
	{41.027779, -77.949165},
	{34.617779, -117.833611},
	{31.030001, -96.510002},
	{34.490002, -89.000000},
	{29.765934, -95.416328},
	{45.494972, -122.655609},
	{34.263302, -118.302711},
	{34.150879, -118.551651},
	{34.227444, -118.381073},
	{34.141323, -118.387833},
	{34.201115, -118.536049},
	{34.236206, -118.485085},
	{34.237923, -118.530197},
	{34.187042, -118.381256},
	{34.272259, -118.468880},
	{34.151749, -118.521431},
	{41.739685, -87.554420},
	{34.667904, -77.512085},
	{36.135910, -95.944733},
	{43.174095, -70.609909},
	{25.830500, -80.180374},
	{30.448336, -91.128960},
	{34.043926, -118.242432},
	{39.783741, -104.758385},
	{33.447041, -82.691544},
	{35.884766, -78.625053},
	{41.908802, -87.679596},
	{47.546257, -122.611740},
	{40.758556, -73.765434},
	{37.947632, -122.525261},
	{37.371067, -121.821060},
	{40.645531, -74.012383},
	{39.039829, -77.055260},
	{33.501804, -81.965118},
	{25.777643, -80.237709},
	{32.741947, -117.239571},
	{32.781025, -96.735657},
	{37.335480, -121.893028},
	{43.615231, -116.289207},
	{33.939728, -118.352882},
	{42.405594, -83.096870},
	{40.640232, -73.906059},
	{29.938885, -95.399193},
	{38.900497, -77.007507},
	{40.755684, -73.883072},
	{33.597538, -112.271828},
	{33.481136, -112.078232},
	{42.341179, -83.035378},
	{34.204529, -119.170029},
	{32.898235, -96.955223},
	{41.851215, -87.634422},
	{41.850510, -87.669006},
	{41.846996, -87.705315},
	{32.810379, -96.635460},
	{34.078159, -118.260559},
	{47.639282, -122.103020},
	{33.805309, -84.395973},
	{39.998089, -75.134109},
	{39.963692, -75.139946},
	{39.961025, -75.191750},
	{40.696011, -73.993286},
	{45.512794, -122.679565},
	{33.016113, -96.679688},
	{29.406347, -98.656769},
	{33.598892, -112.033020},
	{36.197330, -86.798996},
	{37.538994, -121.984276},
	{34.189857, -118.451355},
	{33.533482, -112.107254},
	{32.802353, -117.241676},
	{37.759121, -122.389542},
	{40.047050, -105.272148},
	{38.756649, -75.236603},
	{38.580917, -90.244598},
	{28.571934, -81.235870},
	{47.299606, -122.507942},
	{39.110298, -94.581078},
	{35.002865, -89.997658},
	{40.108448, -74.046249},
	{41.482601, -71.421448},
	{34.070152, -118.349747},
	{32.812614, -96.838730},
	{47.640541, -122.399452},
	{32.736259, -96.864586},
	{39.752174, -86.139793},
	{38.839993, -104.782753},
	{40.732689, -73.784866},
}

//go:embed *
var fs embed.FS

type targetStruct struct {
}

func (g *Gps) targetNew() {
}

func (g *Gps) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch strings.TrimPrefix(r.URL.Path, "/") {
	case "state":
		common.ShowState(g.templates, w, g)
	default:
		g.Common.API(g.templates, w, r)
	}
}

func (g *Gps) run(i *dean.Injector) {
	var msg dean.Msg
	var update = Update{Path: "update"}
	var next int

	for {
		p := places[next]
		next = (next + 1) % len(places)
		update.Lat, update.Long = p.lat, p.long
		i.Inject(msg.Marshal(update))
		time.Sleep(time.Minute)
	}
}
