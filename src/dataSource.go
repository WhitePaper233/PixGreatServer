package src

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type Index struct {
	ID         string
	PixGreatId []string
}

type Metadata struct {
	Title    string `json:"title"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	IllustId string `json:"illust_id"`
	URL      string `json:"url"`
}

var index []Index
var MetadataMap map[string]Metadata

func init() {
	// decode index json
	indexPtr, err := os.Open("./data/index/#風景2.json")
	if err != nil {
		log.Panic("An error occurred while opening index file")
		panic(err)
	}
	defer indexPtr.Close()

	decoder := json.NewDecoder(indexPtr)
	err = decoder.Decode(&index)
	if err != nil {
		log.Panic("An error occurred while decoding index file")
		panic(err)
	}

	// setup random seed
	rand.Seed(time.Now().Unix())
}

func LoadMetadata(indexList []Index) (map[string]Metadata, error) {
	MetadataMap := make(map[string]Metadata)

	var metadata Metadata
	for _, index := range indexList {
		for _, pixGreatID := range index.PixGreatId {
			path := fmt.Sprintf("./data/metadata/%s.json", pixGreatID)
			metadataContentPtr, err := os.Open(path)
			if err != nil {
				log.Errorf("An error occurred while opening metadata file: %s", err.Error())
			}
			defer metadataContentPtr.Close()

			decoder := json.NewDecoder(metadataContentPtr)
			err = decoder.Decode(&metadata)
			if err != nil {
				log.Panic("An error occurred while decode metadata file")
				panic(err)
			}
			MetadataMap[pixGreatID] = metadata
		}
	}
	return MetadataMap, nil
}

func GetIndexList() []Index {
	return index
}

func GetRandomIndex() Index {
	ret := index[rand.Intn(len(index))]
	return ret
}

func GetRandomIDString() string {
	ret := GetRandomIndex().ID
	return ret
}

func GetRandomID() (int64, error) {
	ret, err := strconv.ParseInt(GetRandomIDString(), 10, 64)
	return ret, err
}

func GetRandomPixGreatIDString() string {
	randomIndex := GetRandomIndex()
	ret := randomIndex.PixGreatId[rand.Intn(len(randomIndex.PixGreatId))]
	return ret
}

func GetRandomPixGreatID() (int64, error) {
	ret, err := strconv.ParseInt(GetRandomPixGreatIDString(), 10, 64)
	return ret, err
}

func GetRandomMetadata() (ret Metadata, err error) {
	pixGreatId := GetRandomPixGreatIDString()
	path := fmt.Sprintf("./data/metadata/%s.json", pixGreatId)
	metadataContentPtr, err := os.Open(path)
	if err != nil {
		log.Errorf("An error occurred while opening metadata file: %s", err.Error())
	}
	defer metadataContentPtr.Close()

	decoder := json.NewDecoder(metadataContentPtr)
	err = decoder.Decode(&ret)
	if err != nil {
		log.Panic("An error occurred while decode metadata file")
		panic(err)
	}
	return
}

func GetRandomMetadataFromMem(metadataMap map[string]Metadata) Metadata {
	pixGreatId := GetRandomPixGreatIDString()
	ret := metadataMap[pixGreatId]
	return ret
}
