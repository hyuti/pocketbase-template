package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("hqglnt3tnjz9ph2")
		if err != nil {
			return err
		}

		// update
		edit_period_message := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "soek9wtg",
			"name": "period_message",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": 1,
				"max": null,
				"pattern": ""
			}
		}`), edit_period_message)
		collection.Schema.AddField(edit_period_message)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("hqglnt3tnjz9ph2")
		if err != nil {
			return err
		}

		// update
		edit_period_message := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "soek9wtg",
			"name": "period_message",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": 1,
				"max": null,
				"pattern": ""
			}
		}`), edit_period_message)
		collection.Schema.AddField(edit_period_message)

		return dao.SaveCollection(collection)
	})
}
