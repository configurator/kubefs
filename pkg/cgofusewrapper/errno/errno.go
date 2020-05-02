package errno

// cf https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html

const (
	ENOENT  = -2
	EIO     = -5
	EACCES  = -13
	ENOTDIR = -20
	EISDIR  = -21
	ENOSYS  = -38
)
