package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sebsvt/ihateapi/model"
	"github.com/sebsvt/ihateapi/repository"
	"github.com/sebsvt/ihateapi/service"
)

func main() {
	pdfService := service.NewPdfService()
	fileStorageRepository := repository.NewFileStorageRepositoryMock()
	workflowService := service.NewWorkflowService(pdfService, fileStorageRepository)

	// read the file
	file, err := os.ReadFile("./assets/file1.pdf")
	if err != nil {
		log.Fatal(err)
	}

	file2, err := os.ReadFile("./assets/file2.pdf")
	if err != nil {
		log.Fatal(err)
	}

	res, err := workflowService.Upload(file)
	if err != nil {
		log.Fatal(err)
	}

	res2, err := workflowService.Upload(file2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server filename: ", res.ServerFilename)

	res3, err := workflowService.Process(model.ProcessWorkFlowRequest{
		Tool: "merge",
		Files: []model.File{
			{
				ServerFilename: res.ServerFilename,
			},
			{
				ServerFilename: res2.ServerFilename,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("download filename: ", res3.DownloadFilename)

	mergedfile, err := fileStorageRepository.Download("ihateapi", res3.DownloadFilename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("merged file size: ", len(mergedfile))
	// write the file
	err = os.WriteFile("./assets/mergedfile.pdf", mergedfile, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// uuid := uuid.New().String()
	// err = fileStorageRepository.Upload("ihateapi", uuid+".pdf", file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// has, err := fileStorageRepository.ObjectExists("ihateapi", uuid+".pdf")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(has)

	// has, err = fileStorageRepository.ObjectExists("ihateapi", uuid+"x.pdf")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(has)

	// // file, err = fileStorageRepository.Download("ihateapi", uuid+".pdf")
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// // fmt.Println(len(file))
	// // // write the file
	// // err = os.WriteFile("./assets/file1_downloaded.pdf", file, 0644)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

}
