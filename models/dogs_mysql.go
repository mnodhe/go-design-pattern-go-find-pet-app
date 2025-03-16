package models

import (
	"context"
	"log"
	"time"
)

func (d *DogBreed) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	query := `select id,breed,weight_low_lbs,weight_high_lbs,
       cast(((weight_low_lbs+weight_high_lbs)/2) as unsigned) as average_weight ,
       lifespan , coalesce(details,''),
       coalesce(alternate_names,''),coalesce(geographic_origin,'')
       from dog_breeds order by breed`
	var breeds []*DogBreed
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b DogBreed
		err = rows.Scan(
			&b.ID,
			&b.Breed,
			&b.WeightLowLBs,
			&b.WeightHighLBs,
			&b.AverageWeight,
			&b.Lifespan,
			&b.Details,
			&b.AlternateName,
			&b.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &b)
	}
	return breeds, nil
}
