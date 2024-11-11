package entities

type ArtistDataFormat struct {
	ArtistID   	string `json:"artist_id" bson:"user_id,omitempty"`
	Artistname 	string `json:"artist_name" bson:"artist_name,omitempty"`
	Region		string `json:"region" bson:"region,omitempty"`
	Genre   	string `json:"genre" bson:"genre,omitempty"`
	Songs   	[]SongDataFormat `json:"song" bson:"song,omitempty"`
}

type SongDataFormat struct {
	SongID   	string `json:"song_id" bson:"song_id,omitempty"`
	SongName 	string `json:"song_name" bson:"song_name,omitempty"`
	Time		string `json:"time" bson:"time,omitempty"`
}