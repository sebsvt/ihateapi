package service

import (
	"errors"

	"github.com/sebsvt/ihateapi/model"
)

var (
	ErrToolNotFound = errors.New("tool not found")
)

type WorkflowService interface {
	Start(tool string) (*model.StartWorkFlowResponse, error)
	Upload(file []byte) (*model.UploadWorkFlowResponse, error)
	Process(req model.ProcessWorkFlowRequest) (*model.ProcessWorkFlowResponse, error)
	Download(task string) ([]byte, error)
}
