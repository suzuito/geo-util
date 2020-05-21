package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	geojson "github.com/paulmach/go.geojson"
	"golang.org/x/xerrors"
)

type Area struct {
	Code     int
	Geometry *geojson.Geometry
}

func readGeoJSON(filePath string, areaMap *map[int][]Area) error {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	coll, err := geojson.UnmarshalFeatureCollection(body)
	if err != nil {
		return err
	}
	for _, feature := range coll.Features {
		g := feature.Geometry
		prop := feature.Properties
		// pref, _ := prop["N03_001"]
		// city, _ := prop["N03_003"]
		// ward, _ := prop["N03_004"]
		// 行政コード
		// https://nlftp.mlit.go.jp/ksj/gml/codelist/AdminAreaCd.html
		codeString, _ := prop["N03_007"].(string)
		code, err := strconv.Atoi(codeString)
		if err != nil {
			return err
		}
		// fmt.Println(pref, city, ward, code, feature.Type)

		areas, exists := (*areaMap)[code]
		if !exists {
			areas = []Area{}
		}
		areas = append(areas, Area{
			Code:     code,
			Geometry: g,
		})
		(*areaMap)[code] = areas
	}
	return nil
}

func putAreas(ctx context.Context, db *sql.DB, areaMap *map[int][]Area) error {
	for code, areas := range *areaMap {
		g := geojson.NewMultiPolygonGeometry()
		for _, area := range areas {
			g.MultiPolygon = append(g.MultiPolygon, area.Geometry.Polygon)
		}
		b, err := g.MarshalJSON()
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}
		_, err = db.ExecContext(
			ctx,
			fmt.Sprintf(
				"INSERT INTO geom VALUES ('%d', ST_GeomFromGeoJSON('%s', 3))",
				code,
				string(b),
			),
		)
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}
	}
	return nil
}

func getAreas(ctx context.Context, db *sql.DB, lat, lng float64, areas *[]Area) error {
	rows, err := db.Query(
		fmt.Sprintf(
			"SELECT code FROM geom WHERE ST_Contains(g, ST_GeomFromText('POINT(%f %f)', 4326))",
			lat, lng,
		),
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		code := 0
		if err := rows.Scan(&code); err != nil {
			return err
		}
		area := Area{}
		area.Code = code
		*areas = append(*areas, area)
	}
	return err
}

func main() {
	ctx := context.Background()
	areaMaps := map[int][]Area{}
	// readGeoJSON("../nl/N03-190101_13_GML.geojson")
	readGeoJSON("../nl/N03-190101_14_GML.geojson", &areaMaps)
	//
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := putAreas(ctx, db, &areaMaps); err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	areas := []Area{}
	if err := getAreas(ctx, db, 35.5931477, 139.5910198, &areas); err != nil {
		panic(err)
	}
	fmt.Println(areas)
}
