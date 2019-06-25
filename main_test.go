package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/dev001hajipro/riot-go/common"
)

const APIKey = "RGAPI-xxx"

func TestGetAllTheLeagueEntries(t *testing.T) {
	name := "ホトトギス"
	expectedName := "ホトトギス"
	summonerInfo := SummonerDTO{}
	if err := SummonersByName(APIKey, name, &summonerInfo); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		if summonerInfo.Name != expectedName {
			t.Fail()
		}
		fmt.Printf("summonerInfo:%+v\n", summonerInfo)
	}

	s := []LeagueEntryDTO{}
	if err := GetAllTheLeagueEntries(APIKey, common.RankedSolo5x5, common.Platinum, common.DivisionI, &s); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		//fmt.Printf("GetAllTheLeagueEntries:%+v\n", s)
	}
}

func TestSummonersByName(t *testing.T) {

	// SUMMONER-V4
	name := "29tomato"
	expectedName := "29tomato"
	summonerInfo := SummonerDTO{}
	if err := SummonersByName(APIKey, name, &summonerInfo); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		if summonerInfo.Name != expectedName {
			t.Fail()
		}
		fmt.Printf("summonerInfo:%+v\n", summonerInfo)
	}

	summonerInfo2 := SummonerDTO{}
	if err := SummonersByAccountID(APIKey, summonerInfo.AccountID, &summonerInfo2); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		if summonerInfo2.Name != expectedName {
			t.Fail()
		}
		fmt.Printf("summonerInfo2:%+v\n", summonerInfo2)
	}

	summonerInfo3 := SummonerDTO{}
	if err := SummonersByPUUID(APIKey, summonerInfo.PUUID, &summonerInfo3); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		if summonerInfo3.Name != expectedName {
			t.Fail()
		}
		fmt.Printf("summonerInfo3:%+v\n", summonerInfo3)
	}

	summonerInfo4 := SummonerDTO{}
	if err := SummonersBySummonerID(APIKey, summonerInfo.ID, &summonerInfo4); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		if summonerInfo4.Name != expectedName {
			t.Fail()
		}
		fmt.Printf("summonerInfo4:%+v\n", summonerInfo4)
	}

	score := 0
	if err := ScoresBySummoner(APIKey, summonerInfo.ID, &score); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		fmt.Printf("scores:%d\n", score)
	}

	c1 := ChampionInfo{}
	if err := ChampionRotations(APIKey, &c1); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		//fmt.Printf("%+v\n", c1)
	}

	// LEAGUE-V4
	g1 := LeagueListDTO{}
	if err := GetTheChallengerLeagueForGivenQueue(APIKey, common.RankedSolo5x5, &g1); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		//fmt.Printf("%+v\n", g1)
	}

	s := []LeagueEntryDTO{}
	if err := GetLeagueEntriesBy(summonerInfo.ID, APIKey, &s); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		fmt.Printf("league entry: %+v\n", s)
	}

	grandmasters := LeagueListDTO{}
	if err := GetTheGrandmasterLeagueForGivenQueue(APIKey, common.RankedSolo5x5, &grandmasters); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		if grandmasters.Tier != "GRANDMASTER" {
			t.Fail()
		}
		//fmt.Printf("%+v\n", grandmasters)
		//println("league id:", grandmasters.LeagueID)
	}

	leagueList1 := LeagueListDTO{}
	if err := GetLeagueWithID(APIKey, grandmasters.LeagueID, &leagueList1); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		if grandmasters.Tier != "GRANDMASTER" {
			t.Fail()
		}
	}

	masters := LeagueListDTO{}
	if err := GetTheMasterLeagueForGivenQueue(APIKey, common.RankedSolo5x5, &masters); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		if masters.Tier != "MASTER" {
			t.Fail()
		}
		//fmt.Printf("%+v\n", masters)
		println("league id:", masters.LeagueID)
	}

	// LOL-STATUS-V3
	status := ShardStatus{}
	if err := GetLeagueOfLegendsStatus(APIKey, common.JP1, &status); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		fmt.Printf("%+v\n", status)
	}

	// MATCH-V4

	m1 := MatchListDTO{}
	if err := GetMatchlistForGamesPlayed(APIKey, summonerInfo.AccountID, &m1); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		//fmt.Printf("%+v\n", m1)
	}

	ca1 := []ChampionMasteryDTO{}
	if err := ChampionMasteriesBySummoner(APIKey, summonerInfo.ID, &ca1); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		//fmt.Printf("%+v\n", ca1)
	}

	cmDTO := ChampionMasteryDTO{}
	championID := 1
	if err := ChampionMasteryBySummonerAndChampionID(APIKey, summonerInfo.ID, championID, &cmDTO); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		fmt.Printf("%+v\n", cmDTO)
	}
}
