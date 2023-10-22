package service

// type Playlist struct {
// 	Id                   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name                 string              `json:"name" bson:"name"`
// 	PrettyName           string              `json:"prettyName" bson:"prettyName"`
// 	Duration             int                 `json:"duration" bson:"duration"`
// 	Visible              bool                `json:"visible" bson:"visible"`
// 	Votes                int                 `json:"votes" bson:"votes"`
// 	VoteTime             time.Time           `json:"voteTime" bson:"voteTime"`
// 	VotesTotal           int                 `json:"votesTotal" bson:"votesTotal"`
// 	Index                int                 `json:"index" bson:"index"`
// 	Order                int                 `json:"order" bson:"0rder"`
// 	ImageUrl             string              `json:"imageUrl" bson:"imageUrl"`
// 	IsActive             bool                `json:"isActive" bson:"isActive"`
// 	OwnerVoted           bool                `json:"ownerVoted" bson:"ownerVoted"`
// 	SequenceVisibleCount int                 `json:"sequenceVisibleCount" bson:"sequenceVisibleCount"`
// 	Type                 string              `json:"type" bson:"type"`
// 	GroupName            string              `json:"groupName" bson:"groupName"`
// 	Current              bool                `json:"current" bson:"current"`
// }

func (s *Service) PluginToggleViewerControl() error {
	prefs, err := s.GetPreferences()
	if err != nil {
		return err
	}
	prefs.ViewerControlEnabled = !prefs.ViewerControlEnabled
	if err := s.UpdatePreferences(*prefs); err != nil {
		return err
	}
	return nil
}

func (s *Service) PluginUpateCurrentPlaying() error {
	return nil
}

func (s *Service) PluginUpateNextPlaying() error {
	return nil
}

func (s *Service) PluginHighestVotedSequence() error {
	return nil
}

func (s *Service) PluginSyncSequences(sequences []Sequence) error {
	allSequences, err := s.ListSequences()
	var update []Sequence
	var create []Sequence
	if err != nil {
		return err
	}
	for _, seq := range allSequences {
		if !containsSequence(seq.Name, &allSequences) {
			seq.Active = false
			update = append(update, seq)
		} else {
			seq.Active = true
			create = append(create, seq)
		}
	}
	s.BulkUpdateSequences(update)
	s.BulkCreateSequences(create)
	return nil
}

func containsSequence(name string, sequences *[]Sequence) bool {
	for _, s := range *sequences {
		if s.Name == name {
			return true
		}
	}
	return false
}
