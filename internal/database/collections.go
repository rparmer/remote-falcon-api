package database

import "go.mongodb.org/mongo-driver/mongo/options"

type Collection struct {
	Name                    string
	Indexes                 *[]string
	CreateCollectionOptions *options.CreateCollectionOptions
}

var collections = []Collection{
	{Name: "auths"},
	{Name: "playlists", Indexes: &[]string{"name"}},
	{Name: "playlistGroups", Indexes: &[]string{"name"}},
	{Name: "preferences", CreateCollectionOptions: &options.CreateCollectionOptions{Capped: boolPointer(true), SizeInBytes: int64Pointer(1024), MaxDocuments: int64Pointer(1)}},
	{Name: "sequences", Indexes: &[]string{"name"}},
	{Name: "sequenceGroups", Indexes: &[]string{"name"}},
	{Name: "jukeboxStats"},
	{Name: "votingStats"},
	{Name: "viewerStats"},
	{Name: "viewerPages", Indexes: &[]string{"name"}},
	{Name: "viewerPageTemplates", Indexes: &[]string{"name"}},
}

func boolPointer(v bool) *bool {
	b := v
	return &b
}

func int64Pointer(v int64) *int64 {
	i := v
	return &i
}
