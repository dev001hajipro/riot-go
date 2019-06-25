package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dev001hajipro/riot-go/common"
)

///////////////////////////////////////////////////////////////////////////////
// CHAMPION-MASTERY-V4

// ChampionMasteryDTO contains single Champion Mastery information for player and champion combination.
type ChampionMasteryDTO struct {
	ChestGranted                 bool   `json:"chestGranted"`                 // Is chest granted for this champion or not in current season.
	ChampionLevel                int32  `json:"championLevel"`                //Champion level for specified player and champion combination.
	ChampionPoints               int32  `json:"championPoints"`               //Total number of champion points for this player and champion combination - they are used to determine championLevel.
	ChampionID                   int64  `json:"championId"`                   //Champion ID for this entry.
	ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel"` //Number of points needed to achieve next level. Zero if player reached maximum champion level for this champion.
	LastPlayTime                 int64  `json:"lastPlayTime"`                 //Last time this champion was played by this player - in Unix milliseconds time format.
	TokensEarned                 int32  `json:"tokensEarned"`                 //The token earned for this champion to levelup.
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel"` //Number of points earned since current level has been achieved.
	SummonerID                   string `json:"summonerId"`                   //Summoner ID for this entry. (Encrypted)
}

// ChampionMasteriesBySummoner get all champion mastery entries sorted by number of champion points descending
func ChampionMasteriesBySummoner(apiKey string, encryptedSummonerID string, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/%s", encryptedSummonerID)
	return doRequest(apiKey, url, target)
}

// ChampionMasteryBySummonerAndChampionID get a champion mastery by player ID and champion ID.
func ChampionMasteryBySummonerAndChampionID(apiKey string, encryptedSummonerID string, championID int, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/%s/by-champion/%d", encryptedSummonerID, championID)
	return doRequest(apiKey, url, target)
}

// ScoresBySummoner get a player's total champion mastery score, which is the sum of individual champion mastery levels.
func ScoresBySummoner(apiKey string, encryptedSummonerID string, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com//lol/champion-mastery/v4/scores/by-summoner/%s", encryptedSummonerID)
	return doRequest(apiKey, url, target)
}

///////////////////////////////////////////////////////////////////////////////
// 	CHAMPION-V3

// ChampionInfo champion rotations, including free-to-play and low-level free-to-play rotations.
type ChampionInfo struct {
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	FreeChampionIds              []int `json:"freeChampionIds"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// ChampionRotations returns champion rotations.
func ChampionRotations(apiKey string, target interface{}) error {
	return doRequest(apiKey, "https://jp1.api.riotgames.com/lol/platform/v3/champion-rotations", target)
}

///////////////////////////////////////////////////////////////////////////////
// LEAGUE-V4

// LeagueListDTO get the challenger league for given queue.
type LeagueListDTO struct {
	LeagueID string     `json:"leagueId"`
	Tier     string     `json:"tier"`
	Entries  []struct { // LeagueItemDTO
		SummonerName string `json:"summonerName"`
		HotStreak    bool   `json:"hotStreak"`
		MiniSeries   struct {
			//MiniSeriesDTO
			Progress string `json:"progress"`
			Losses   int32  `json:"losses"`
			Target   int32  `json:"target"`
			Wins     int32  `json:"wins"`
		}
		Wins         int32  `json:"wins"`
		Veteran      bool   `json:"veteran"`
		Losses       int32  `json:"losses"`
		FreshBlood   bool   `json:"freshBlood"`
		Inactive     bool   `json:"inactive"`
		Rank         string `json:"rank"`
		SummonerID   string `json:"summonerId"` // Player's summonerId (Encrypted)
		LeaguePoints int32  `json:"leaguePoints"`
	}
	Queue string `json:"queue"`
	Name  string `json:"name"`
}

// GetTheChallengerLeagueForGivenQueue get the challenger league for given queue.
func GetTheChallengerLeagueForGivenQueue(apiKey string, queue common.LOLQueue, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/league/v4/challengerleagues/by-queue/%s", queue)
	return doRequest(apiKey, url, target)
}

// LeagueEntryDTO is a entry information.
type LeagueEntryDTO struct {
	QueueType    string `json:"queueType"`
	SummonerName string `json:"summonerName"`
	HotStreak    bool   `json:"hotStreak"`
	MiniSeries   struct {
		Progress string `json:"progress"`
		Losses   int32  `json:"losses"`
		Target   int32  `json:"target"`
		Wins     int32  `json:"wins"`
	}
	Wins         int32  `json:"wins"`
	Veteran      bool   `json:"veteran"`
	Losses       int32  `json:"losses"`
	Rank         string `json:"rank"`
	LeagueID     string `json:"leagueId"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	Tier         string `json:"tier"`
	SummonerID   string `json:"summonerId"` //Player's summonerId (Encrypted)
	LeaguePoints int32  `json:"leaguePoints"`
}

// GetLeagueEntriesBy get league entries in all queues for a given summoner ID.
func GetLeagueEntriesBy(encryptedSummonerID string, apiKey string, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/league/v4/entries/by-summoner/%s", encryptedSummonerID)
	return doRequest(apiKey, url, target)
}

// GetAllTheLeagueEntries get all the league entries.
func GetAllTheLeagueEntries(apiKey string, queue common.LOLQueue, tier common.Tier, division common.Division, target interface{}) error {
	// todo: append page query parameter. &page=1
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/league/v4/entries/%s/%s/%s", queue, tier, division)
	return doRequest(apiKey, url, target)
}

// GetTheGrandmasterLeagueForGivenQueue get the grandmaster league for given queue.
func GetTheGrandmasterLeagueForGivenQueue(apiKey string, queue common.LOLQueue, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/league/v4/grandmasterleagues/by-queue/%s", queue)
	return doRequest(apiKey, url, target)
}

// GetLeagueWithID get league with given ID, including inactive entries.
func GetLeagueWithID(apiKey string, leagueID string, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/league/v4/leagues/%s", leagueID)
	return doRequest(apiKey, url, target)
}

// GetTheMasterLeagueForGivenQueue get the master league for given queue.
func GetTheMasterLeagueForGivenQueue(apiKey string, queue common.LOLQueue, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/league/v4/masterleagues/by-queue/%s", queue)
	return doRequest(apiKey, url, target)
}

///////////////////////////////////////////////////////////////////////////////
// LOL-STATUS-V3

// ShardStatus is the information of active or not and so on.
type ShardStatus struct {
	Name      string `json:"name"`
	RegionTag string `json:"region_tag"`
	Hostname  string `json:"hostname"`
	Services  []struct {
		Status    string `json:"status"`
		Incidents []struct {
			//Incident
			Active    bool   `json:"active"`
			CreatedAt string `json:"created_at"`
			ID        int64  `json:"id"`
			Updates   []struct {
				// Message
				Severity     string `json:"severity"`
				Author       string `json:"author"`
				CreatedAt    string `json:"created_at"`
				Translations []struct {
					// Translation
					Locale    string `json:"locale"`
					Content   string `json:"content"`
					UpdatedAt string `json:"updated_at"`
				}
				UpdatedAt string `json:"updated_at"`
				Content   string `json:"content"`
				ID        string `json:"id"`
			}
		}
		Name string `json:"name"`
		Slug string `json:"slug"`
	}
	Slug    string   `json:"slug"`
	Locales []string `json:"locales"`
}

// GetLeagueOfLegendsStatus get League of Legends status for the given shard.
func GetLeagueOfLegendsStatus(apiKey string, region common.Region, target interface{}) error {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/status/v3/shard-data", region)
	return doRequest(apiKey, url, target)
}

///////////////////////////////////////////////////////////////////////////////
// MATCH-V4

// GetMatchByMatchID get match by match ID.
func GetMatchByMatchID(apiKey string, matchID string, target interface{}) error {
	// TODO: complex dto.
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/match/v4/matches/%s", matchID)
	return doRequest(apiKey, url, target)
}

// MatchListDTO is matchlist for games played.
type MatchListDTO struct {
	Matches []struct { // MatchReferenceDto
		Lane       string `json:"startIndex"`
		GameID     int64  `json:"gameId"`
		Champion   int32  `json:"champion"`
		PlatformID string `json:"platformId"`
		Season     int32  `json:"season"`
		Queue      int32  `json:"queue"`
		Role       string `json:"role"`
		Timestamp  int64  `json:"timestamp"`
	}
	StartIndex int32 `json:"startIndex"`
	EndIndex   int32 `json:"endIndex"`
}

// GetMatchlistForGamesPlayed that Get matchlist for games played on given account ID and platform ID and filtered using given filter parameters, if an
func GetMatchlistForGamesPlayed(apiKey string, encryptedAccountID string, target interface{}) error {
	// todo: implement more fileter parameters.
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/match/v4/matchlists/by-account/%s", encryptedAccountID)
	return doRequest(apiKey, url, target)
}

///////////////////////////////////////////////////////////////////////////////
// SUMMONER-V4

// SummonerDTO represents a summoner
// a field's name of public access need Uppercase in golang.
type SummonerDTO struct {
	ProfileIconID int    `json:"profileIconId"` //ID of the summoner icon associated with the summoner.
	Name          string `json:"name"`          //Summoner name.
	PUUID         string `json:"puuid"`         //Encrypted PUUID. Exact length of 78 characters.
	SummonerLevel int64  `json:"summonerLevel"` //Summoner level associated with the summoner.
	RevisionDate  int64  `json:"revisionDate"`  //Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: profile icon change, playing the tutorial or advanced tutorial, finishing a game, summoner name change
	ID            string `json:"id"`            //Encrypted summoner ID. Max length 63 characters.
	AccountID     string `json:"accountId"`     //Encrypted account ID. Max length 56 characters.
}

// SummonersByAccountID get a summoner by account ID.
func SummonersByAccountID(apiKey string, encryptedAccountID string, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/summoner/v4/summoners/by-account/%s", encryptedAccountID)
	return doRequest(apiKey, url, target)
}

// SummonersByName get a summoner by name.
func SummonersByName(apiKey string, name string, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", name)
	return doRequest(apiKey, url, target)
}

// SummonersByPUUID get a summoner by PUUID.
func SummonersByPUUID(apiKey string, PUUID string, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s", PUUID)
	return doRequest(apiKey, url, target)
}

// SummonersBySummonerID get a summoner by PUUID.
// Consistently looking up summoner ids that don't exist will result in a blacklist.
func SummonersBySummonerID(apiKey string, encryptedSummonerID string, target interface{}) error {
	url := fmt.Sprintf("https://jp1.api.riotgames.com/lol/summoner/v4/summoners/%s", encryptedSummonerID)
	return doRequest(apiKey, url, target)
}

func doRequest(apiKey string, url string, target interface{}) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Riot-Token", apiKey)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func main() {
	const APIKey = "RGAPI-4e048f9a-fc5c-4bc2-8161-163edc2885c9"

	s1 := SummonerDTO{}
	if err := SummonersByName(APIKey, "29tomato", &s1); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v\n", s1)
	}
}
