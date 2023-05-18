package db

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type CommunityFaq struct {
	ID       uint64 `structs:"id,omitempty" db:"id" json:"id,omitempty"`
	Question string `structs:"q,omitempty" db:"q" json:"q,omitempty"`
	Answer   string `structs:"a,omitempty" db:"a" json:"a,omitempty"`
	Ordering uint64 `structs:"ordering,omitempty" db:"ordering" json:"ordering,omitempty"`
}

func CommunityFaqTableKeys() []string {
	return []string{"id", "q", "a", "ordering"}
}

func GetCommunityFaq[H Handle](h H, id, start, limit uint64) (*[]CommunityFaq, error) {
	const op errors.Op = "GetCommunityFaq"
	query := fmt.Sprintf("select %s from %s.community_faq where id = ? limit ?, ? order by ordering asc", strings.Join(CommunityFaqTableKeys(), ","), NFDDatabase())

	var faq []CommunityFaq
	err := h.Select(&faq, query, id, start, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.E(pkg, op, errors.DatabaseResultNotFound, err, "CommunityFaq Not Found")
		}
		return nil, errors.E(pkg, op, err)
	}

	return &faq, nil
}
