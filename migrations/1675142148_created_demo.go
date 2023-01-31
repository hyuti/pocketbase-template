package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "hqglnt3tnjz9ph2",
			"created": "2023-01-31 05:15:48.661Z",
			"updated": "2023-01-31 05:15:48.661Z",
			"name": "demo",
			"type": "base",
			"system": false,
			"schema": [
				{
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
				}
			],
			"listRule": "",
			"viewRule": "",
			"createRule": "",
			"updateRule": "",
			"deleteRule": "",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("hqglnt3tnjz9ph2")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
