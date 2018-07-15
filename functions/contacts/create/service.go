package main

import (
	"errors"
	"log"
)

// contactService - handles manipulation of contact
type contactService struct {
	*contact
}

func (cs *contactService) create() (*contact, error) {
	err := cs.validate()
	if err != nil {
		log.Println(err)
		return cs.contact, err
	}

	_, err = contactDAO.insert(cs.contact)
	if err != nil {
		log.Println(err)
	}
	return cs.contact, err
}

func (cs *contactService) validate() error {
	if cs.contact.Email == "" {
		return errors.New("Email cannot be blank")
	}

	if cs.contact.Message == "" {
		return errors.New("Message cannot be blank")
	}

	return nil
}
