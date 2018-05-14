package main

import "fmt"

type (
	// Header represents the header information that's saved along with the ArchiveData
	Header struct {
		Version uint // Indicates the version of the ArchiveData
	}

	// Data represents the data stored in the object store which is used
	// to generate manifests and populate cache tables
	Data struct {
		// The decoder assumes that Version is the first field present
		Version uint // Indicates the version of the ArchiveData
		Value   interface{}
	}

	ArchiveData struct {
		ArchiveBatches  string
		ArchiveSegments string
	}

	PrivateData struct {
		PrivateBatches  string
		PrivateSegments string
		Timelines       string
	}
)

func getData(data *Data) {
	if data.Version == 1 {
		fmt.Println((data.Value).(*PrivateData).PrivateBatches)
	} else {
		fmt.Println((data.Value).(*ArchiveData).ArchiveBatches)
	}
}


func main() {
	privData := PrivateData{
		PrivateBatches: "PrivateBatches",
		PrivateSegments: "PrivateSegments",
		Timelines: "Timelines",
	}
	testData := &Data{
		Version: 1,
		Value: &privData,
	}
	getData(testData)

}
