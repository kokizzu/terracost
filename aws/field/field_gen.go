// Code generated by "enumer -type=Field -output=field_gen.go -linecomment"; DO NOT EDIT.

package field

import (
	"fmt"
	"strings"
)

const _FieldName = "SKUCapacityStatusGroupInstance TypeLocationOperating SystemPre Installed S/WProduct FamilyserviceCodeTenancyusageTypeVolume API NameVolume TypeStorage ClassAccess TypeThroughput ClassCache EngineDatabase EngineDatabase EditionDeployment OptionLicense ModelFile system typeStorage typeThroughput capacityDeployment optionCurrencyPricePerUnitStartingRangeTermTypeUnit"

var _FieldIndex = [...]uint16{0, 3, 17, 22, 35, 43, 59, 76, 90, 101, 108, 117, 132, 143, 156, 167, 183, 195, 210, 226, 243, 256, 272, 284, 303, 320, 328, 340, 353, 361, 365}

const _FieldLowerName = "skucapacitystatusgroupinstance typelocationoperating systempre installed s/wproduct familyservicecodetenancyusagetypevolume api namevolume typestorage classaccess typethroughput classcache enginedatabase enginedatabase editiondeployment optionlicense modelfile system typestorage typethroughput capacitydeployment optioncurrencypriceperunitstartingrangetermtypeunit"

func (i Field) String() string {
	if i >= Field(len(_FieldIndex)-1) {
		return fmt.Sprintf("Field(%d)", i)
	}
	return _FieldName[_FieldIndex[i]:_FieldIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FieldNoOp() {
	var x [1]struct{}
	_ = x[SKU-(0)]
	_ = x[CapacityStatus-(1)]
	_ = x[Group-(2)]
	_ = x[InstanceType-(3)]
	_ = x[Location-(4)]
	_ = x[OperatingSystem-(5)]
	_ = x[PreInstalledSW-(6)]
	_ = x[ProductFamily-(7)]
	_ = x[ServiceCode-(8)]
	_ = x[Tenancy-(9)]
	_ = x[UsageType-(10)]
	_ = x[VolumeAPIName-(11)]
	_ = x[VolumeType-(12)]
	_ = x[StorageClass-(13)]
	_ = x[AccessType-(14)]
	_ = x[ThroughputClass-(15)]
	_ = x[CacheEngine-(16)]
	_ = x[DatabaseEngine-(17)]
	_ = x[DatabaseEdition-(18)]
	_ = x[DatabaseDeploymentOption-(19)]
	_ = x[LicenseModel-(20)]
	_ = x[FileSystemType-(21)]
	_ = x[StorageType-(22)]
	_ = x[ThroughputCapacity-(23)]
	_ = x[FileSystemDeploymentOption-(24)]
	_ = x[Currency-(25)]
	_ = x[PricePerUnit-(26)]
	_ = x[StartingRange-(27)]
	_ = x[TermType-(28)]
	_ = x[Unit-(29)]
}

var _FieldValues = []Field{SKU, CapacityStatus, Group, InstanceType, Location, OperatingSystem, PreInstalledSW, ProductFamily, ServiceCode, Tenancy, UsageType, VolumeAPIName, VolumeType, StorageClass, AccessType, ThroughputClass, CacheEngine, DatabaseEngine, DatabaseEdition, DatabaseDeploymentOption, LicenseModel, FileSystemType, StorageType, ThroughputCapacity, FileSystemDeploymentOption, Currency, PricePerUnit, StartingRange, TermType, Unit}

var _FieldNameToValueMap = map[string]Field{
	_FieldName[0:3]:          SKU,
	_FieldLowerName[0:3]:     SKU,
	_FieldName[3:17]:         CapacityStatus,
	_FieldLowerName[3:17]:    CapacityStatus,
	_FieldName[17:22]:        Group,
	_FieldLowerName[17:22]:   Group,
	_FieldName[22:35]:        InstanceType,
	_FieldLowerName[22:35]:   InstanceType,
	_FieldName[35:43]:        Location,
	_FieldLowerName[35:43]:   Location,
	_FieldName[43:59]:        OperatingSystem,
	_FieldLowerName[43:59]:   OperatingSystem,
	_FieldName[59:76]:        PreInstalledSW,
	_FieldLowerName[59:76]:   PreInstalledSW,
	_FieldName[76:90]:        ProductFamily,
	_FieldLowerName[76:90]:   ProductFamily,
	_FieldName[90:101]:       ServiceCode,
	_FieldLowerName[90:101]:  ServiceCode,
	_FieldName[101:108]:      Tenancy,
	_FieldLowerName[101:108]: Tenancy,
	_FieldName[108:117]:      UsageType,
	_FieldLowerName[108:117]: UsageType,
	_FieldName[117:132]:      VolumeAPIName,
	_FieldLowerName[117:132]: VolumeAPIName,
	_FieldName[132:143]:      VolumeType,
	_FieldLowerName[132:143]: VolumeType,
	_FieldName[143:156]:      StorageClass,
	_FieldLowerName[143:156]: StorageClass,
	_FieldName[156:167]:      AccessType,
	_FieldLowerName[156:167]: AccessType,
	_FieldName[167:183]:      ThroughputClass,
	_FieldLowerName[167:183]: ThroughputClass,
	_FieldName[183:195]:      CacheEngine,
	_FieldLowerName[183:195]: CacheEngine,
	_FieldName[195:210]:      DatabaseEngine,
	_FieldLowerName[195:210]: DatabaseEngine,
	_FieldName[210:226]:      DatabaseEdition,
	_FieldLowerName[210:226]: DatabaseEdition,
	_FieldName[226:243]:      DatabaseDeploymentOption,
	_FieldLowerName[226:243]: DatabaseDeploymentOption,
	_FieldName[243:256]:      LicenseModel,
	_FieldLowerName[243:256]: LicenseModel,
	_FieldName[256:272]:      FileSystemType,
	_FieldLowerName[256:272]: FileSystemType,
	_FieldName[272:284]:      StorageType,
	_FieldLowerName[272:284]: StorageType,
	_FieldName[284:303]:      ThroughputCapacity,
	_FieldLowerName[284:303]: ThroughputCapacity,
	_FieldName[303:320]:      FileSystemDeploymentOption,
	_FieldLowerName[303:320]: FileSystemDeploymentOption,
	_FieldName[320:328]:      Currency,
	_FieldLowerName[320:328]: Currency,
	_FieldName[328:340]:      PricePerUnit,
	_FieldLowerName[328:340]: PricePerUnit,
	_FieldName[340:353]:      StartingRange,
	_FieldLowerName[340:353]: StartingRange,
	_FieldName[353:361]:      TermType,
	_FieldLowerName[353:361]: TermType,
	_FieldName[361:365]:      Unit,
	_FieldLowerName[361:365]: Unit,
}

var _FieldNames = []string{
	_FieldName[0:3],
	_FieldName[3:17],
	_FieldName[17:22],
	_FieldName[22:35],
	_FieldName[35:43],
	_FieldName[43:59],
	_FieldName[59:76],
	_FieldName[76:90],
	_FieldName[90:101],
	_FieldName[101:108],
	_FieldName[108:117],
	_FieldName[117:132],
	_FieldName[132:143],
	_FieldName[143:156],
	_FieldName[156:167],
	_FieldName[167:183],
	_FieldName[183:195],
	_FieldName[195:210],
	_FieldName[210:226],
	_FieldName[226:243],
	_FieldName[243:256],
	_FieldName[256:272],
	_FieldName[272:284],
	_FieldName[284:303],
	_FieldName[303:320],
	_FieldName[320:328],
	_FieldName[328:340],
	_FieldName[340:353],
	_FieldName[353:361],
	_FieldName[361:365],
}

// FieldString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FieldString(s string) (Field, error) {
	if val, ok := _FieldNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FieldNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Field values", s)
}

// FieldValues returns all values of the enum
func FieldValues() []Field {
	return _FieldValues
}

// FieldStrings returns a slice of all String values of the enum
func FieldStrings() []string {
	strs := make([]string, len(_FieldNames))
	copy(strs, _FieldNames)
	return strs
}

// IsAField returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Field) IsAField() bool {
	for _, v := range _FieldValues {
		if i == v {
			return true
		}
	}
	return false
}
