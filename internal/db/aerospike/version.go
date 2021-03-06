package aerospike

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go"
	"github.com/sajeevany/graph-snapper/internal/db/aerospike/record"
	"github.com/sirupsen/logrus"
)

const (
	V1RecordLevel = "1"
)

//GetVersion - returns version as a string. Empty if version not found.
func GetVersion(logger *logrus.Logger, aeroRecord aerospike.BinMap) string {

	if aeroRecord == nil || !hasMetadataBin(logger, aeroRecord) {
		logger.Debugf("Aerospike record is empty or missing metadata bin. Returning empty value access version. Binmap <%v>", aeroRecord)
		return ""
	}

	//Get metadata bin map
	mdBin := aeroRecord[record.MetadataBinName]

	//Get version value
	switch v := mdBin.(type) {
	case map[interface{}]interface{}:
		version := fmt.Sprintf("%s", v[record.VersionAttrName])
		logger.Debugf("Bin map is [interface]interface. Returning version <%v>", version)
		return version
	default:
		logger.Debugf("Bin map is of unsupported type <%T>. Returning empty", v)
		return ""
	}

}

func hasMetadataBin(logger *logrus.Logger, aeroRecord aerospike.BinMap) bool {

	if aeroRecord == nil {
		logger.Debug("Aerospike record is empty. Returning false for hasMetadataBin")
		return false
	}

	_, ok := aeroRecord[record.MetadataBinName]

	return ok
}
