# File Utility CLI Roadmap

This document outlines the planned features and development milestones for the File Utility CLI tool.

## Current Status

- ✅ Basic CLI structure using Cobra
- ✅ File copy command implementation
- ✅ File move command implementation
- ✅ File delete command implementation
- ✅ File rename command implementation
- ✅ File info command implementation
- ✅ Error handling and logging
- ✅ Version command

## Upcoming Features

### File Operations (Short Term)

- [x] `file move`: Move files from source to destination
- [x] `file delete`: Delete files
- [x] `file rename`: Rename files
- [x] `file info`: Display detailed file information
- [ ] `file search`: Search for files by pattern
- [ ] `file hash`: Generate hash of file contents (MD5, SHA1, SHA256)

### Directory Operations (Medium Term)

- [ ] `dir list`: List directory contents with various options (size, date, etc.)
- [ ] `dir create`: Create directories
- [ ] `dir delete`: Remove directories
- [ ] `dir copy`: Copy directories recursively
- [ ] `dir move`: Move directories
- [ ] `dir size`: Calculate directory size

### File Content Manipulation (Long Term)

- [ ] `content search`: Search for text in files
- [ ] `content replace`: Find and replace text in files
- [ ] `content extract`: Extract specific lines or patterns from files
- [ ] `content encrypt`: Encrypt file contents
- [ ] `content decrypt`: Decrypt file contents
- [ ] `content compress`: Compress files
- [ ] `content decompress`: Decompress files

### Advanced Features (Future)

- [ ] Support for regex patterns in search operations
- [ ] Batch processing of files
- [ ] Configuration file for default settings
- [ ] Plugin system for extending functionality
- [ ] Progress bars for long operations
- [ ] Parallel processing for performance improvement

## Development Milestones

### v0.1.0 (Previous)
- Basic CLI structure
- File copy implementation

### v0.2.0 (Current)
- Complete file operations (move, delete, rename, info)
- Improved error handling and user feedback

### v0.3.0
- Basic directory operations (list, create, delete)
- Configuration file support

### v0.4.0
- Advanced directory operations (copy, move, size)
- File content basic operations (search, replace)

### v0.5.0
- File content advanced operations (extract, encrypt, decrypt)
- Compression and decompression support

### v1.0.0
- Complete feature set with robust error handling
- Comprehensive documentation
- Cross-platform testing and optimization 