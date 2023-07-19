package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"overflowing/src/structs"
)

/// this file is responsible for handling the stackoverflow api

/// get stackoverflow stats
func GetStats(userId string) (structs.Stats, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", getApiUrl(userId), nil)

	if err != nil {
		return structs.Stats{}, err
	}

	res, err := client.Do(req)

	if err != nil || res.StatusCode != 200 {
		return structs.Stats{}, fmt.Errorf("failed getting stats")
	}

	resBody, err := ioutil.ReadAll(res.Body);
	if err != nil {
		return structs.Stats{}, err
	}

	var stats structs.StackStats
	err = json.Unmarshal(resBody, &stats);

	if err != nil {
		return structs.Stats{}, err
	}
	// since we don't want to return empty stats, we use defer to close the body after we return the stats
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	return structs.Stats{
		ID:		 userId,
		Name:    stats.Items[0].DisplayName,
		Reputation: stats.Items[0].Reputation,
		Gold:	stats.Items[0].BadgeCounts.Gold,
		Silver:	stats.Items[0].BadgeCounts.Silver,
		Bronze:	stats.Items[0].BadgeCounts.Bronze,
		ImageUrl: stats.Items[0].ProfileImage,
	}, nil;

}

/// helper function, for parsing the url link
func getApiUrl(userId string) string {
	return fmt.Sprintf("https://api.stackexchange.com/2.3/users/%s?order=desc&sort=reputation&site=stackoverflow", userId)
}