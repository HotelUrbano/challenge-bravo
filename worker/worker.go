package worker

import (
	"github.com/labstack/gommon/log"
	"github.com/schonmann/challenge-bravo/config"
	"github.com/schonmann/challenge-bravo/keys"
	"github.com/schonmann/challenge-bravo/redis"
	"github.com/schonmann/challenge-bravo/worker/external"
	"time"
)

/**
  Worker that run's periodically as of it's configured
  update interval. It will find current rates using any
  retrieve strategy and save these on Redis right after.
*/

func StartWorker(cfg *config.AppConfig) {

	intervalTime := time.Millisecond * time.Duration(cfg.Worker.UpdateInterval)

	for {
		strategy := external.GetRetrieveQuotasStrategy()
		response, err := strategy.RetrieveRates()

		if err != nil || len(response.Quotas) == 0 {
			log.Fatalf("Error getting/parsing json from external API: %v", err)
			time.Sleep(intervalTime)
			continue
		}

		log.Infof("Keys to be set: %v", response.Quotas)

		newQuotas := make([]interface{}, 0)

		for currency, quota := range response.Quotas {
			newQuotas = append(newQuotas, keys.QuotaKey(currency), quota)
		}

		if _, err := redis.MSet(newQuotas...); err != nil {
			log.Fatalf("Error setting new rates from external API: %v", err)
		}

		time.Sleep(intervalTime)
	}
}
