package ridesumo

import (
	"image"
	"time"
	"middleware"
	"appengine/datastore"
	"log"
	// "sessionauth"
)

const (
	DatastoreNameComp string = "comp"
)

type Comp struct {
	Name string
	Description string

	// may be unbounded?
	Starts, Ends time.Time

	BonusSegments Segments

	Prize   Prize
	Entries EntrySet

	Metadata string
}

// ids of segments and activities
type Segments []int64
type Activities []int64


func GetAllComps() ([]Comp, error) {
	q := datastore.NewQuery(DatastoreNameComp)

    var comps []Comp
    _, err := q.GetAll(middleware.AppEngineContext, &comps)
    if err != nil {
    	return nil, err
    }

	return comps, nil
}

func (c *Comp) Join(u *User) (*Entry, error) {

	return nil, nil
}

func (c *Comp) Put() error {
	// var err error
	// if u.uniqueId == 0 {
	// 	key := datastore.NewIncompleteKey(c, DatastoreNameComp, nil)
	// 	key, err = datastore.Put(c, key, u)
	// 	u.uniqueId = key.IntID()
	// } else {
	// 	// key doesn't change, right??
	// 	_, err = datastore.Put(c, datastore.NewKey(c, DatastoreNameUser, "", u.uniqueId, nil), u)
	// }

	// return err
	return nil
}

type Prize struct {
	Title, Description string

	// not sure if the interface is enough
	// also not sure what the DataStore will make of this
	Photo image.Image
}

// implement PropertyLoadSaver to save the iamge
func (p *Prize) Load(c <-chan datastore.Property) error {
	log.Println("Loading Prize from datastore")
	return nil
}

func (p *Prize) Save(c chan<- datastore.Property) error {
	log.Println("Saving Prize to datastore")
	return nil
}

type Entry struct {
	Entrant      User
	CurrentScore int64
	LastUpdated  time.Time
	ScoredActitives Activities
}


type EntrySet []Entry
func (set EntrySet) Len() int {
	return len(set)
}
func (set EntrySet) Less(i, j int) bool {
    return set[i].CurrentScore < set[j].CurrentScore
}
func (set EntrySet) Swap(i, j int) {
    set[i], set[j] = set[j], set[i]
}

type Scorer interface {
	Score(e1 Entry) int64
}

type SpecialBlendScorer struct {

}

func NewSpecialBlendScorer(metadata string) (*SpecialBlendScorer, error) {

	return nil, nil
}
func (sb *SpecialBlendScorer) Score(e Entry) int64 {
	return 0
}
