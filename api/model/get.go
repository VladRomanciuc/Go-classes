package model

import (
	"log"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)


func createClass(className string) {

	res, err := db.Query(
		f.If(
			f.Exists(f.Class(className)),
			true,
			f.CreateClass(f.Obj{"name": className})))

	if err != nil {
		panic(err)
	}

	if res != f.BooleanV(true) {
		log.Printf("Created Class: %s\n %s", className, res)
	} else {
		log.Printf("Class: %s, Already Exists\n %s", className, res)
	}
}


/*
func DBGetAll(client *f.FaunaClient, token string) (refs []f.RefV, err error) {
	value, err := client.Query(
		f.Paginate(
			f.MatchTerm(
				f.Index("entries_with_token"),
				token,
			),
		),
	)
	if err != nil {
		return nil, err
	}

	value.At(f.ObjKey("data")).Get(&refs)
	return refs, nil
}

// DBGetFromRefs - get all elements
func DBGetFrom(client *f.FaunaClient, refs []f.RefV) (entries []Entry, err error) {
	request := Mapper(refs, func(ref f.RefV) interface{} {
		return f.Get(ref)
	})
	value, err := client.Query(f.Arr(request))

	if err != nil {
		return nil, err
	}

	var elements f.ArrayV
	value.Get(&elements)

	results := make([]Entry, len(elements))
	for index, element := range elements {
		var object f.ObjectV
		element.At(f.ObjKey("data")).Get(&object)
		var entry Entry
		object.Get(&entry)
		results[index] = entry
	}

	return results, nil
}

func (entry Entry) DBGetEntry(client *f.FaunaClient) (value f.Value, err error) {
	return client.Query(
		f.Get(
			f.MatchTerm(
				f.Index("entry_with_token_user_item"),
				f.Arr{entry.Token, entry.UserID, entry.ItemID},
			),
		),
	)
}
*/