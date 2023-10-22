package service

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	sequenceCollection      = "sequences"
	sequenceGroupCollection = "sequenceGroups"
)

type Sequence struct {
	Id          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string              `json:"name" bson:"name"`
	DisplayName string              `json:"displayName" bson:"displayName"`
	Duration    int                 `json:"duration" bson:"duration"`
	ImageUrl    string              `json:"imageUrl" bson:"imageUrl"`
	Active      bool                `json:"active" bson:"active"`
	Psa         bool                `json:"psa" bson:"psa"`
	Type        string              `json:"type" bson:"type"`
}

type SequenceGroup struct {
	Id        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name" bson:"name"`
	Sequences *[]string           `json:"sequences" bson:"sequences"`
}

func (s *Service) CreateSequence(seq Sequence) error {
	if err := s.db.Create(sequenceCollection, &seq); err != nil {
		return err
	}
	return nil
}

func (s *Service) BulkCreateSequences(sequences []Sequence) error {
	var create []mongo.WriteModel
	for _, sequence := range sequences {
		create = append(create, mongo.NewInsertOneModel().
			SetDocument(sequence))
	}
	if err := s.db.BulkWrite(sequenceCollection, create); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateSequence(seq Sequence) error {
	count, err := s.db.Update(sequenceCollection, bson.D{{Key: "_id", Value: seq.Id}}, &seq)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (s *Service) BulkUpdateSequences(sequences []Sequence) error {
	var update []mongo.WriteModel
	for _, sequence := range sequences {
		update = append(update, mongo.NewUpdateOneModel().
			SetFilter(bson.D{{Key: "_id", Value: sequence.Id}}).
			SetUpdate(bson.D{{Key: "$set", Value: sequence}}))
	}
	if err := s.db.BulkWrite(sequenceCollection, update); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetSequence(name string) (*Sequence, error) {
	var sequence Sequence
	if err := s.db.Get(sequenceCollection, bson.D{{Key: "name", Value: name}}, &sequence); err != nil {
		return nil, err
	}
	return &sequence, nil
}

func (s *Service) ListSequences() ([]Sequence, error) {
	var sequences []Sequence
	if err := s.db.List(sequenceCollection, bson.D{}, &sequences); err != nil {
		return nil, err
	}
	return sequences, nil
}

func (s *Service) DeleteSequence(seq Sequence) error {
	if err := s.db.Delete(sequenceCollection, bson.D{{Key: "_id", Value: seq.Id}}); err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateSequenceGroup(sg SequenceGroup) error {
	if err := s.db.Create(sequenceGroupCollection, &sg); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateSequenceGroup(sg SequenceGroup) error {
	count, err := s.db.Update(sequenceGroupCollection, bson.D{{Key: "_id", Value: sg.Id}}, &sg)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (s *Service) GetSequenceGroup(name string) (*SequenceGroup, error) {
	var sg SequenceGroup
	if err := s.db.Get(sequenceGroupCollection, bson.D{{Key: "name", Value: name}}, &sg); err != nil {
		return nil, err
	}
	return &sg, nil
}

func (s *Service) ListSequenceGroups() (*[]SequenceGroup, error) {
	var sequenceGroups []SequenceGroup
	if err := s.db.List(sequenceGroupCollection, bson.D{}, &sequenceGroups); err != nil {
		return nil, err
	}
	return &sequenceGroups, nil
}

func (s *Service) DeleteSequenceGroup(sg SequenceGroup) error {
	if err := s.db.Delete(sequenceGroupCollection, bson.D{{Key: "_id", Value: sg.Id}}); err != nil {
		return err
	}
	return nil
}
