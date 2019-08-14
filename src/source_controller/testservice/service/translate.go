package service

import (
	"configcenter/src/common"
	"configcenter/src/common/language"
	"configcenter/src/common/metadata"
	"configcenter/src/common/util"
)

var defaultNameLanguagePkg = map[string]map[string][]string{
	common.BKInnerObjIDModule: {
		"1": {"inst_module_idle", common.BKModuleNameField, common.BKModuleIDField},
		"2": {"inst_module_fault", common.BKModuleNameField, common.BKModuleIDField},
	},
	common.BKInnerObjIDApp: {
		"1": {"inst_biz_default", common.BKAppNameField, common.BKAppIDField},
	},
	common.BKInnerObjIDSet: {
		"1": {"inst_set_default", common.BKSetNameField, common.BKSetIDField},
	},
}

func (s *testService) TranslateClassName(defLang language.DefaultCCLanguageIf, att *metadata.Class) string {
	return util.FirstNotEmptyString(defLang.Language("class_"+att.ClassID), att.ClassName, att.ClassID)
}
