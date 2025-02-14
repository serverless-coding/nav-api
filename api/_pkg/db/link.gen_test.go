// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/serverless-coding/nav-api/api/_pkg/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Link{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Link{}) fail: %s", err)
	}
}

func Test_linkQuery(t *testing.T) {
	link := newLink(db)
	link = *link.As(link.TableName())
	_do := link.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(link.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <link> fail:", err)
		return
	}

	_, ok := link.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from link success")
	}

	err = _do.Create(&model.Link{})
	if err != nil {
		t.Error("create item in table <link> fail:", err)
	}

	err = _do.Save(&model.Link{})
	if err != nil {
		t.Error("create item in table <link> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Link{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <link> fail:", err)
	}

	_, err = _do.Select(link.ALL).Take()
	if err != nil {
		t.Error("Take() on table <link> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <link> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <link> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <link> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Link{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <link> fail:", err)
	}

	_, err = _do.Select(link.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <link> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <link> fail:", err)
	}

	_, err = _do.Select(link.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <link> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <link> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <link> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <link> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Link{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <link> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <link> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <link> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <link> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <link> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <link> fail:", err)
	}
}
