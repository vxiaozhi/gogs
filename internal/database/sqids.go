package database

import (
	"fmt"
	"time"

	"github.com/sqids/sqids-go"
)

const (
	SQIDS_Alphabet  string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SQIDS_MinLength uint8  = 4

	SQIDS_AppPrefix  uint64 = 1
	SQIDS_UserPrefix uint64 = 1

	SQIDS_TimestampPrefix uint64 = 10
)

var (
	RepoTemplateIdMap = map[string]uint64{
		"webstack": 1,
		"md":       2,
	}
)

func encodeSqids(ids []uint64) string {
	s, _ := sqids.New(sqids.Options{
		Alphabet:  SQIDS_Alphabet,
		MinLength: SQIDS_MinLength,
		Blocklist: []string{"admin"},
	})
	id, _ := s.Encode(ids)
	return id

}

func decodeSqids(id string) []uint64 {
	s, _ := sqids.New(sqids.Options{
		Alphabet:  SQIDS_Alphabet,
		MinLength: SQIDS_MinLength,
		Blocklist: []string{"admin"},
	})
	ids := s.Decode(id)
	return ids
}

func EncodeUserID(userId uint64, userPrefix uint64) string {
	return encodeSqids([]uint64{SQIDS_AppPrefix, userPrefix, userId})
}

func DecodeUserID(id string) (uint64, uint64, uint64) {
	ids := decodeSqids(id)
	return ids[0], ids[1], ids[2]
}

func EncodeRepoID(repoId uint64, repoTemplateId uint64) string {
	return encodeSqids([]uint64{SQIDS_AppPrefix, repoTemplateId, repoId})
}

func DecodeRepoID(id string) (uint64, uint64, uint64) {
	ids := decodeSqids(id)
	return ids[0], ids[1], ids[2]
}

func EncodeRepoIDByTemplateName(repoId uint64, repoTemplateName string) (string, error) {
	repoTemplateId, ok := RepoTemplateIdMap[repoTemplateName]
	if !ok {
		return "", fmt.Errorf("unknown repo template name: %s", repoTemplateName)
	}
	return EncodeRepoID(repoId, repoTemplateId), nil
}

func DecodeRepoIDToTemplateName(id string) (uint64, uint64, string, error) {
	appPrefix, repoTemplateId, repoId := DecodeRepoID(id)
	if appPrefix != SQIDS_AppPrefix {
		return 0, 0, "", fmt.Errorf("invalid app prefix: %d", appPrefix)
	}
	for templateName, templateId := range RepoTemplateIdMap {
		if templateId == repoTemplateId {
			return appPrefix, repoId, templateName, nil
		}
	}
	return 0, 0, "", fmt.Errorf("unknown repo template id: %d", repoTemplateId)
}

func EncodeIDByTimestamp(timestamp uint64) string {
	return encodeSqids([]uint64{SQIDS_TimestampPrefix, timestamp})
}

func EncodeIDByTimestampFromYear2025() string {
	// 计算 2025年到现在的时间戳差值 2025年1月1日的时间戳为 1735689600
	timestampDiff := time.Now().Unix() - 1735689600
	if timestampDiff < 0 {
		timestampDiff = 1000
	}
	return encodeSqids([]uint64{SQIDS_TimestampPrefix, uint64(timestampDiff)})
}
