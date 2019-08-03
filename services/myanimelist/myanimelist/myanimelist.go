package myanimelist

import (
	log "github.com/sirupsen/logrus"

	"github.com/greenmochi/ultimate/services/myanimelist/myanimelist/data"
	"github.com/greenmochi/ultimate/services/myanimelist/myanimelist/database"
	"github.com/greenmochi/ultimate/services/myanimelist/myanimelist/fetch"
	"github.com/greenmochi/ultimate/services/myanimelist/myanimelist/store"
)

// MyAnimeList handles main application logic
type MyAnimeList struct {
	store *store.Store
	db    *database.Database
}

// New TODO
func New() *MyAnimeList {
	return &MyAnimeList{
		store: store.New(),
		db:    database.New(),
	}
}

// InitDB TODO
func (mal *MyAnimeList) InitDB(path string) {
	if err := mal.db.Open(path); err != nil {
		log.Error(err)
		return
	}
	log.Infof("Successfully initialized database: %s", path)
}

// CloseDB TODO
func (mal *MyAnimeList) CloseDB() {
	mal.db.Close()
}

// StoreAnimeList TODO
func (mal *MyAnimeList) storeUserAnimeList(user string, userAnimeList *data.UserAnimeList) error {
	mal.store.SetUserAnimeList(user, userAnimeList)
	if err := mal.db.InsertUserAnimeList(userAnimeList); err != nil {
		log.Errorf("Failed to insert %s anime list into database: %s", user, err)
		return err
	}
	log.Infof("Successfully saved user %s anime list in the database", user)
	return nil
}

// GetUserAnimeList TODO
func (mal *MyAnimeList) GetUserAnimeList(user string) (*data.UserAnimeList, error) {
	// Check if the store already has the user's anime list
	if userAnimeList := mal.store.GetUserAnimeList(user); userAnimeList != nil {
		return userAnimeList, nil
	}
	// Check if the database already has the user's anime list
	if ok := mal.db.UserAnimeListExists(user); ok {
		if userAnimeList, err := mal.db.RetrieveUserAnimeList(user); err == nil {
			mal.store.SetUserAnimeList(user, userAnimeList)
			return userAnimeList, nil
		}
	}
	// Retrieve a new user anime list
	userAnimeList, err := fetch.UserAnimeList(user)
	if err != nil {
		return nil, err
	}
	if err := mal.storeUserAnimeList(user, userAnimeList); err != nil {
		return nil, err
	}
	return userAnimeList, nil
}

func (mal *MyAnimeList) storeAnimeSearchResults(results []*data.AnimeSearchResult) {
	mal.store.SetAnimeSearchResults(results)
	log.Infof("Stored %d anime search results", len(results))
}

// SearchAnime TODO
func (mal *MyAnimeList) SearchAnime(query string) ([]*data.AnimeSearchResult, error) {
	results, err := fetch.AnimeSearchResults(query)
	if err != nil {
		return nil, err
	}
	mal.storeAnimeSearchResults(results)
	return results, nil
}

func (mal *MyAnimeList) storeAnime(anime *data.Anime) error {
	mal.store.SetAnime(anime.ID, anime)
	if err := mal.db.InsertAnime(anime); err != nil {
		log.Errorf("Failed to insert anime %s into database: %s", anime.Title, err)
		return err
	}
	log.Infof("Successfully saved anime %s to the database", anime.Title)
	return nil
}

// GetAnimeBySearchResult TODO
func (mal *MyAnimeList) GetAnimeBySearchResult(result *data.AnimeSearchResult) (*data.Anime, error) {
	if anime := mal.store.GetAnimeByTitle(result.Title); anime != nil {
		return anime, nil
	}
	if ok := mal.db.AnimeByTitleExists(result.Title); ok {
		if anime, err := mal.db.RetrieveAnimeByTitle(result.Title); err == nil {
			mal.store.SetAnime(anime.ID, anime)
			return anime, nil
		}
	}
	anime, err := fetch.AnimeBySearchResult(result)
	if err != nil {
		return nil, err
	}
	if err := mal.storeAnime(anime); err != nil {
		return nil, err
	}
	return anime, nil
}

// GetAnimeByID TODO
func (mal *MyAnimeList) GetAnimeByID(id int) (*data.Anime, error) {
	if anime := mal.store.GetAnime(id); anime != nil {
		return anime, nil
	}
	if ok := mal.db.AnimeByIDExists(id); ok {
		if anime, err := mal.db.RetrieveAnimeByID(id); err == nil {
			mal.store.SetAnime(anime.ID, anime)
			return anime, nil
		}
	}
	anime, err := fetch.AnimeByID(id)
	if err != nil {
		return nil, err
	}
	if err := mal.storeAnime(anime); err != nil {
		return nil, err
	}
	return anime, nil
}
