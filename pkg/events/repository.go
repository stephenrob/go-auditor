package events

import "github.com/go-pg/pg/v10"

type Repository interface {
	Create(event *AuditEvent) error
	GetByID(id int) (*AuditEvent, error)
}

type DBRepository struct {
	DB *pg.DB
}

func (d *DBRepository) GetByID(id int) (*AuditEvent, error) {
	e := &AuditEvent{
		ID: id,
	}
	err := d.DB.Model(e).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (d *DBRepository) Create(event *AuditEvent) error {
	m := event
	_, err := d.DB.Model(m).Insert()

	if err != nil {
		return err
	}

	return nil
}

