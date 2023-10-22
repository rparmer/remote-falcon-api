package service

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	viewerPageCollection         = "viewerPages"
	viewerPageTemplateCollection = "viewerPageTemplates"
)

type ViewerPage struct {
	Id       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string              `json:"name" bson:"name"`
	Html     string              `json:"html" bson:"html"`
	Title    string              `json:"title" bson:"title"`
	IconLink string              `json:"iconLink" bson:"iconLink"`
}

type ViewerPageTemplate struct {
	Id   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string              `json:"name" bson:"name"`
	Html string              `json:"html" bson:"html"`
}

func (s *Service) CreateViewerPage(p ViewerPage) error {
	if p.Name == "" {
		return fmt.Errorf("no name provided")
	}
	if err := s.db.Create(viewerPageCollection, &p); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateViewerPage(p ViewerPage) error {
	count, err := s.db.Update(viewerPageCollection, bson.D{{Key: "_id", Value: p.Id}}, &p)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (s *Service) GetViewerPage(name string) (*ViewerPage, error) {
	var page ViewerPage
	if err := s.db.Get(viewerPageCollection, bson.D{{Key: "name", Value: name}}, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (s *Service) ListViewerPages() (*[]ViewerPage, error) {
	var pages []ViewerPage
	if err := s.db.List(viewerPageCollection, bson.D{}, &pages); err != nil {
		return nil, err
	}
	return &pages, nil
}

func (s *Service) DeleteViewerPage(p ViewerPage) error {
	if err := s.db.Delete(viewerPageCollection, bson.D{{Key: "_id", Value: p.Id}}); err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateViewerPageTemplate(p ViewerPageTemplate) error {
	if p.Name == "" {
		return fmt.Errorf("no name provided")
	}
	if err := s.db.Create(viewerPageTemplateCollection, &p); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateViewerPageTemplate(p ViewerPageTemplate) error {
	count, err := s.db.Update(viewerPageTemplateCollection, bson.D{{Key: "_id", Value: p.Id}}, &p)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no records found")
	}
	return nil
}

func (s *Service) GetViewerPageTemplate(name string) (*ViewerPageTemplate, error) {
	var template ViewerPageTemplate
	if err := s.db.Get(viewerPageTemplateCollection, bson.D{{Key: "name", Value: name}}, &template); err != nil {
		return nil, err
	}
	return &template, nil
}

func (s *Service) ListViewerPageTemplates() (*[]ViewerPageTemplate, error) {
	var templates []ViewerPageTemplate
	if err := s.db.List(viewerPageTemplateCollection, bson.D{}, &templates); err != nil {
		return nil, err
	}
	return &templates, nil
}

func (s *Service) DeleteViewerPageTemplate(p ViewerPageTemplate) error {
	if err := s.db.Delete(viewerPageTemplateCollection, bson.D{{Key: "_id", Value: p.Id}}); err != nil {
		return err
	}
	return nil
}
