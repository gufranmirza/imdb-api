package movieinterface

import "github.com/gufranmirza/imdb-api/database/dbmodels"

// MovieReq holds data to create a or update a movie
type MovieReq struct {
	*dbmodels.Movie
}

// MovieResp holds data movie response
type MovieResp struct {
	*dbmodels.Movie
}
