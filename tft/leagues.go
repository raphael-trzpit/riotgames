package tft

import (
	"context"

	"github.com/pkg/errors"
	"github.com/raphael-trzpit/riotgames/api"
)

const (
	endpointLeagueChallenger = "/tft/league/v1/challenger"
)

type Leagues struct {
	Client *api.Client
}

func (l *Leagues) GetChallenger(ctx context.Context) (LeagueRankings, error) {
	rankings := LeagueRankings{}
	if err := l.Client.Get(ctx,endpointLeagueChallenger, &rankings); err != nil {
		return rankings, errors.Wrap(err, "cannot get challenger")
	}

	return rankings, nil
}

type LeagueRankings struct {
	ID        string           `json:"leagueId"`
	Summoners []LeagueSummoner `json:"entries"`
	Tier      string           `json:"tier"`
	Name      string           `json:"name"`
	Queue     string           `json:"queue"`
}

type LeagueSummoner struct {
	ID         string                   `json:"summonerId"`
	Name       string                   `json:"summonerName"`
	Rank       string                   `json:"rank"`
	Points     int                      `json:"leaguePoints"`
	Wins       int                      `json:"wins"`
	Losses     int                      `json:"losses"`
	FreshBlood bool                     `json:"freshBlood"`
	HotStreak  bool                     `json:"hotStreak"`
	Veteran    bool                     `json:"veteran"`
	Inactive   bool                     `json:"inactive"`
	LastSeries LeagueSummonerLastSeries `json:"miniSeries"`
}

type LeagueSummonerLastSeries struct {
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     string `json:"wins"`
	Losses   string `json:"losses"`
}
