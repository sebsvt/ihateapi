package service

import (
	"log"

	"github.com/google/uuid"
	"github.com/sebsvt/ihateapi/model"
	"github.com/sebsvt/ihateapi/repository"
)

type workflowService struct {
	pdfService  PdfService
	fileStorage repository.FileStorageRepository
}

func NewWorkflowService(pdfService PdfService, fileStorage repository.FileStorageRepository) WorkflowService {
	return &workflowService{
		pdfService:  pdfService,
		fileStorage: fileStorage,
	}
}

// Start implements WorkflowService.
func (srv *workflowService) Start(tool string) (*model.StartWorkFlowResponse, error) {
	uuid := uuid.New().String()
	return &model.StartWorkFlowResponse{
		Server:           "https://api.ihateapi.com",
		Task:             uuid,
		RemainingCredits: 10,
	}, nil
}

// Upload implements WorkflowService.
func (srv *workflowService) Upload(file []byte) (*model.UploadWorkFlowResponse, error) {
	uuid := uuid.New().String()
	err := srv.fileStorage.Upload("ihateapi", uuid+".pdf", file)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.UploadWorkFlowResponse{
		ServerFilename: uuid + ".pdf",
	}, nil
}

// Process implements WorkflowService.
func (srv *workflowService) Process(req model.ProcessWorkFlowRequest) (*model.ProcessWorkFlowResponse, error) {
	switch req.Tool {
	// For pdf
	case "merge":
		var files [][]byte
		// read the files from storage
		for _, file := range req.Files {
			fileBytes, err := srv.fileStorage.Download("ihateapi", file.ServerFilename)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			files = append(files, fileBytes)
		}
		mergedFile, err := srv.pdfService.Merge(files)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		// generate a uuid for the merged file
		mergedFileUUID := uuid.New().String()
		// upload the merged file to storage
		err = srv.fileStorage.Upload("ihateapi", mergedFileUUID+".pdf", mergedFile)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &model.ProcessWorkFlowResponse{
			DownloadFilename: mergedFileUUID + ".pdf",
			FileSize:         len(mergedFile),
			OutputFileSize:   10,
			OutputFileNumber: len(mergedFile),
			OutputExtentions: "[\"pdf\"]",
			Timer:            "0.090",
		}, nil
	case "split":
		// do nothing now
	case "compress":
		// do nothing now
	case "pdfocr":
		// do nothing now
	case "pdfjpg":
		// do nothing now
	case "imagepdf":
		// do nothing now
	case "unlock":
		// do nothing now
	case "pagenumber":
		// do nothing now
	case "watermark":
		// do nothing now
	case "officepdf":
		// do nothing now
	case "repair":
		// do nothing now
	case "rotate":
		// do nothing now
	case "protect":
		// do nothing now
	case "pdfa":
		// do nothing now
	case "validatepdfa":
		// do nothing now

	// For image
	case "resizeimage":
		// do nothing now
	case "cropimage":
		// do nothing now
	case "compressimage":
		// do nothing now
	case "upscaleimage":
		// do nothing now
	case "removebackgroundimage":
		// do nothing now
	case "coverimage":
		// do nothing now
	case "watermarkimage":
		// do nothing now
	case "replaceimage":
		// do nothing now
	case "rotateimage":
		// do nothing now
	default:
		return nil, ErrToolNotFound
	}
	return &model.ProcessWorkFlowResponse{
		DownloadFilename: "output.zip",
		FileSize:         10,
		OutputFileSize:   10,
		OutputFileNumber: 10,
		OutputExtentions: "[\"pdf\"]",
		Timer:            "0.090",
		Status:           "TaskSuccess",
	}, nil
}

// Download implements WorkflowService.
func (srv *workflowService) Download(task string) ([]byte, error) {
	return []byte("https://api.ihateapi.com/download"), nil
}
