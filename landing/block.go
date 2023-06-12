package landing

type Block struct {
	Id            string
	Type          string
	Type_for_from string
	Title         string
	Entities      []Block_entity
	Description   string
	Data          map[Personal_playlists_data]Play_contexts_data
}
