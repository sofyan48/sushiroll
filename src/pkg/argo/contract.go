package argo

import "github.com/sofyan48/sushiroll/src/presentations"

type ArgoRolloutLibrary interface {
	GetList() (*presentations.RolloutList, error)
	Detail(service string) (*presentations.RolloutDetail, error)
	Promote(isFull bool, service string) (*presentations.PromoteArgoResponse, error)
	Rollback(revision, service string) ([]byte, error)
	Restart(service string) (*presentations.PromoteArgoResponse, error)
	Retry(service string) (*presentations.PromoteArgoResponse, error)
	Abort(service string) (*presentations.PromoteArgoResponse, error)
}
