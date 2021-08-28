package query

import (
	"fmt"

	"github.com/photoprism/photoprism/pkg/txt"

	"github.com/photoprism/photoprism/internal/entity"
)

// Subjects returns subjects from the index.
func Subjects(limit, offset int) (result entity.Subjects, err error) {
	stmt := Db()

	stmt = stmt.Order("subject_name").Limit(limit).Offset(offset)
	err = stmt.Find(&result).Error

	return result, err
}

// SubjectMap returns a map of subjects indexed by UID.
func SubjectMap() (result map[string]entity.Subject, err error) {
	result = make(map[string]entity.Subject)

	var subj entity.Subjects

	stmt := Db().Where("subject_src <> ?", entity.SrcDefault)

	if err = stmt.Find(&subj).Error; err != nil {
		return result, err
	}

	for _, s := range subj {
		result[s.SubjectUID] = s
	}

	return result, err
}

// RemoveDanglingMarkerSubjects permanently deletes dangling marker subjects from the index.
func RemoveDanglingMarkerSubjects() (removed int64, err error) {
	res := UnscopedDb().
		Where("subject_src = ?", entity.SrcMarker).
		Where(fmt.Sprintf("subject_uid NOT IN (SELECT subject_uid FROM %s)", entity.Face{}.TableName())).
		Where(fmt.Sprintf("subject_uid NOT IN (SELECT subject_uid FROM %s)", entity.Marker{}.TableName())).
		Delete(&entity.Subject{})

	return res.RowsAffected, res.Error
}

// AddFaceMarkerSubjects adds and references known marker subjects.
func AddFaceMarkerSubjects() (affected int64, err error) {
	var markers entity.Markers

	if err := Db().
		Where("subject_uid = '' AND marker_name <> ''").
		Where("marker_invalid = 0 AND marker_type = ?", entity.MarkerFace).
		Order("marker_name").
		Find(&markers).Error; err != nil {
		return affected, err
	} else if len(markers) == 0 {
		return affected, nil
	}

	var name string
	var subj *entity.Subject

	for _, m := range markers {
		if name == m.MarkerName && subj != nil {
			// Do nothing.
		} else if subj = entity.NewSubject(m.MarkerName, entity.SubjectPerson, entity.SrcMarker); subj == nil {
			log.Errorf("faces: subject should not be nil - bug?")
			continue
		} else if subj = entity.FirstOrCreateSubject(subj); subj == nil {
			log.Errorf("faces: failed adding subject %s", txt.Quote(m.MarkerName))
			continue
		} else {
			affected++
		}

		name = m.MarkerName

		if err := m.Updates(entity.Values{"SubjectUID": subj.SubjectUID}); err != nil {
			return affected, err
		}

		if m.FaceID == "" {
			continue
		} else if err := Db().Model(&entity.Face{}).Where("id = ? AND subject_uid = ''", m.FaceID).Update("SubjectUID", subj.SubjectUID).Error; err != nil {
			return affected, err
		}
	}

	return affected, err
}
