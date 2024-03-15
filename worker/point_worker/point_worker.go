package point_worker

import (
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/rpc/airdrop"
)

type PointWorker struct {
	db                *database.DB
	airDropRpcService airdrop.AirDropRpcService
	tasks             tasks.Group
}

func NewPointWorker(db *database.DB, airDropRpcService airdrop.AirDropRpcService) (*PointWorker, error) {
	pointWorker := PointWorker{
		db:                db,
		airDropRpcService: airDropRpcService,
		tasks: tasks.Group{HandleCrit: func(err error) {
			log.Error("critical error in point worker processor: %w", err)
		}},
	}
	return &pointWorker, nil
}

func (pw *PointWorker) Start() error {
	tickerScan := time.NewTicker(5 * time.Second)
	pw.tasks.Go(func() error {
		for range tickerScan.C {
			err := pw.pointsByBridge()
			if err != nil {
				log.Error(" shutting down pointsByBridge ", "err", err)
				continue
			}
		}
		return nil
	})
	pw.tasks.Go(func() error {
		for range tickerScan.C {
			err := pw.pointsByStaking()
			if err != nil {
				log.Error(" shutting down pointsByBridge ", "err", err)
				continue
			}
		}
		return nil
	})
	return nil
}

func (pw *PointWorker) pointsByBridge() error {
	log.Info(" start pointsByBridge ")
	records := pw.db.BridgeRecord.GetNotPointsRecords()
	for i := 0; i < len(records); i++ {
		bridgeRecord := records[i]
		address := bridgeRecord.FromAddress
		guid := bridgeRecord.GUID
		points, err := pw.airDropRpcService.SubmitDppLinkPoints("", "0", address.String())
		if err != nil {
			return err
		}
		if points.Code == "200" {
			dbErr := pw.db.BridgeRecord.UpdatePointsStatus(guid)
			if dbErr != nil {
				log.Error("UpdatePointsStatus error", "dbErr", dbErr, "this point is success", guid)
				return err
			}
		}
	}
	return nil
}

func (pw *PointWorker) pointsByStaking() error {
	log.Info(" start pointsByStaking ")
	records := pw.db.StakeRecord.GetNotPointsRecords()
	for i := 0; i < len(records); i++ {
		stakingRecord := records[i]
		address := stakingRecord.UserAddress
		guid := stakingRecord.GUID
		txType := stakingRecord.TxType
		pointsType := pw.getPointsType(txType)
		points, err := pw.airDropRpcService.SubmitDppLinkPoints("", pointsType, address.String())
		if err != nil {
			return err
		}
		if points.Code == "200" {
			dbErr := pw.db.StakeRecord.UpdatePointsStatus(guid)
			if dbErr != nil {
				log.Error("UpdatePointsStatus error", "dbErr", dbErr, "this point is success", guid)
				return err
			}
		}
	}
	return nil
}

func (pw *PointWorker) getPointsType(txType int) string {
	switch txType {
	case global_const.StakingTypeStake:
		return "1"
	case global_const.StakingTypeReward:
		return "3"
	case global_const.StakingTypeWithdraw:
		return "2"
	}
	return ""
}
