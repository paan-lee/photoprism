package migrate

// Generated code, do not edit.

var DialectSQLite3 = Migrations{
	{
		ID:         "20211121-094727",
		Dialect:    "sqlite3",
		Statements: []string{"DROP INDEX IF EXISTS idx_places_place_label;"},
	},
	{
		ID:         "20211124-120008",
		Dialect:    "sqlite3",
		Statements: []string{"DROP INDEX IF EXISTS uix_places_place_label;", "DROP INDEX IF EXISTS uix_places_label;"},
	},
	{
		ID:         "20220329-040000",
		Dialect:    "sqlite3",
		Statements: []string{"DROP INDEX IF EXISTS idx_albums_album_filter;"},
	},
	{
		ID:         "20220329-050000",
		Dialect:    "sqlite3",
		Statements: []string{"CREATE INDEX IF NOT EXISTS idx_albums_album_filter ON albums (album_filter);"},
	},
	{
		ID:         "20220329-061000",
		Dialect:    "sqlite3",
		Statements: []string{"CREATE INDEX IF NOT EXISTS idx_files_photo_id ON files (photo_id, file_primary);"},
	},
	{
		ID:         "20220329-071000",
		Dialect:    "sqlite3",
		Statements: []string{"UPDATE files SET photo_taken_at = (SELECT taken_at_local FROM photos WHERE photos.id = photo_id) WHERE photo_id IS NOT NULL;"},
	},
	{
		ID:         "20220329-081000",
		Dialect:    "sqlite3",
		Statements: []string{"CREATE UNIQUE INDEX IF NOT EXISTS idx_files_search_media ON files (media_id);"},
	},
	{
		ID:         "20220329-083000",
		Dialect:    "sqlite3",
		Statements: []string{"UPDATE files SET media_id = CASE WHEN photo_id IS NOT NULL AND file_missing = 0 AND deleted_at IS NULL THEN ((10000000000 - photo_id) || '-' || (1 + file_sidecar - file_primary) || '-' || file_uid) END WHERE 1;"},
	},
	{
		ID:         "20220329-091000",
		Dialect:    "sqlite3",
		Statements: []string{"CREATE UNIQUE INDEX IF NOT EXISTS idx_files_search_timeline ON files (time_index);"},
	},
	{
		ID:         "20220329-093000",
		Dialect:    "sqlite3",
		Statements: []string{"UPDATE files SET time_index = CASE WHEN media_id IS NOT NULL AND photo_taken_at IS NOT NULL THEN ((100000000000000 - strftime('%Y%m%d%H%M%S', photo_taken_at)) || '-' || media_id) ELSE NULL END WHERE photo_id IS NOT NULL;"},
	},
	{
		ID:         "20220421-200000",
		Dialect:    "sqlite3",
		Statements: []string{"CREATE INDEX IF NOT EXISTS idx_files_missing_root ON files (file_missing, file_root);"},
	},
	{
		ID:         "20220901-000100",
		Dialect:    "sqlite3",
		Statements: []string{"INSERT OR IGNORE INTO auth_users (id, user_uid, super_admin, user_role, display_name, user_slug, username, email, login_attempts, login_at, created_at, updated_at) SELECT id, user_uid, 1, 'admin', full_name, user_name, user_name, primary_email, login_attempts, login_at, created_at, updated_at FROM users WHERE user_name <> '' AND user_name IS NOT NULL AND user_uid <> '' AND user_uid IS NOT NULL AND role_admin = 1 AND user_disabled = 0 ON CONFLICT DO NOTHING;"},
	},
}
