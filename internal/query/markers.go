package query

import (
	"fmt"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/pkg/txt"
)

// MarkerByID returns a Marker based on the ID.
func MarkerByID(id uint) (marker entity.Marker, err error) {
	if err := UnscopedDb().Where("id = ?", id).
		First(&marker).Error; err != nil {
		return marker, err
	}

	return marker, nil
}

// Markers finds a list of file markers filtered by type, embeddings, and sorted by id.
func Markers(limit, offset int, markerType string, embeddings, subjects bool) (result entity.Markers, err error) {
	db := Db()

	if markerType != "" {
		db = db.Where("marker_type = ?", markerType)
	}

	if embeddings {
		db = db.Where("embeddings_json <> ''")
	}

	if subjects {
		db = db.Where("subject_uid <> ''")
	}

	db = db.Order("id").Limit(limit).Offset(offset)

	err = db.Find(&result).Error

	return result, err
}

// Embeddings returns existing face embeddings.
func Embeddings(single, unclustered bool) (result entity.Embeddings, err error) {
	var col []string

	stmt := Db().
		Model(&entity.Marker{}).
		Where("marker_type = ?", entity.MarkerFace).
		Where("marker_invalid = 0").
		Where("embeddings_json <> ''").
		Order("id")

	if unclustered {
		stmt = stmt.Where("face_id = ''")
	}

	if err := stmt.Pluck("embeddings_json", &col).Error; err != nil {
		return result, err
	}

	for _, embeddingsJson := range col {
		if embeddings := entity.UnmarshalEmbeddings(embeddingsJson); len(embeddings) > 0 {
			if single {
				// Single embedding per face detected.
				result = append(result, embeddings[0])
			} else {
				// Return all embedding otherwise.
				result = append(result, embeddings...)
			}
		}
	}

	return result, nil
}

// AddMarkerSubjects adds and references known marker subjects.
func AddMarkerSubjects() (affected int64, err error) {
	var markers entity.Markers

	if err := Db().
		Where("face_id <> '' AND subject_uid = '' AND subject_src = ?", entity.SrcAuto).
		Where("marker_invalid = 0 AND marker_type = ?", entity.MarkerFace).
		Where("marker_name <> ''").
		Order("marker_name").
		Find(&markers).Error; err != nil {
		return affected, err
	} else if len(markers) == 0 {
		return affected, nil
	}

	for _, m := range markers {
		if faceId := m.FaceID; faceId == "" {
			// Do nothing.
		} else if subj := entity.NewSubject(m.MarkerName, entity.SubjectPerson, entity.SrcMarker); subj == nil {
			log.Errorf("faces: subject should not be nil - bug?")
		} else if subj = entity.FirstOrCreateSubject(subj); subj == nil {
			log.Errorf("faces: failed adding subject %s for marker %d", txt.Quote(m.MarkerName), m.ID)
		} else if err := m.Updates(entity.Values{"SubjectUID": subj.SubjectUID, "SubjectSrc": entity.SrcAuto}); err != nil {
			return affected, err
		} else if err := Db().Model(&entity.Face{}).Where("id = ? AND subject_uid = ''", faceId).Update("SubjectUID", subj.SubjectUID).Error; err != nil {
			return affected, err
		} else {
			affected++
		}
	}

	return affected, err
}

// RemoveInvalidMarkerReferences deletes invalid reference IDs from the markers table.
func RemoveInvalidMarkerReferences() (removed int64, err error) {
	// Remove subject and face relationships for invalid markers.
	if res := Db().
		Model(&entity.Marker{}).
		Where("marker_invalid = 1 AND (subject_uid <> '' OR face_id <> '')").
		UpdateColumns(entity.Values{"subject_uid": "", "face_id": ""}); res.Error != nil {
		return removed, res.Error
	} else {
		removed += res.RowsAffected
	}

	// Remove invalid face IDs.
	if res := Db().
		Model(&entity.Marker{}).
		Where("marker_type = ?", entity.MarkerFace).
		Where(fmt.Sprintf("face_id <> '' AND face_id NOT IN (SELECT id FROM %s)", entity.Face{}.TableName())).
		UpdateColumns(entity.Values{"face_id": ""}); res.Error != nil {
		return removed, res.Error
	} else {
		removed += res.RowsAffected
	}

	// Remove invalid subject UIDs.
	if res := Db().
		Model(&entity.Marker{}).
		Where(fmt.Sprintf("subject_uid <> '' AND subject_uid NOT IN (SELECT subject_uid FROM %s)", entity.Subject{}.TableName())).
		UpdateColumns(entity.Values{"subject_uid": ""}); res.Error != nil {
		return removed, res.Error
	} else {
		removed += res.RowsAffected
	}

	return removed, nil
}

// ResetFaceMarkerMatches removes automatically added subject and face references from the markers table.
func ResetFaceMarkerMatches() (removed int64, err error) {
	res := Db().Model(&entity.Marker{}).
		Where("subject_src <> ? AND marker_type = ?", entity.SrcManual, entity.MarkerFace).
		UpdateColumns(entity.Values{"subject_uid": "", "subject_src": "", "face_id": ""})

	return res.RowsAffected, res.Error
}
