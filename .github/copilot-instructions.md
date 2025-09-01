# GoSTKB - AI Agent Instructions

This is a Go web application for managing school timetables with a Vue.js frontend. Below are the key patterns and workflows an AI agent needs to know.

## Architecture Overview

### Backend (Go/Gin)
- `main.go` - Entry point, configures Gin router, database connection, and static file serving
- `handlers/` - Request handlers organized by domain entity (giaovien, lophoc, monhoc, phancong)
- `models/` - Data structures matching SQLite schema
- `database.sql` - SQLite schema definition

### Frontend (Vue.js)
- `templates/` - HTML templates using Go template inheritance
  - `layouts/` - Base templates (header.html, footer.html)
  - Domain-specific pages (giaovien.html, lophoc.html, etc.)
- `static/` - Client-side assets
  - `js/lib/` - Third-party libraries (Vue, Axios, Bootstrap)
  - `css/lib/` - Third-party styles
  - `js/vue-config.js` - Vue.js configuration and utilities

## Key Patterns

### Database Operations
```go
// Standard query pattern using sqlite3
rows, err := h.DB.Query("SELECT column FROM table")
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Error message"})
    return
}
defer rows.Close()
```

### API Response Format
```go
// Success response
c.JSON(http.StatusOK, data)

// Error response
c.JSON(http.StatusInternalServerError, gin.H{"error": "Error message"})
```

### Vue.js Component Structure
```javascript
const app = createVueApp({
    data() { return {...} },
    methods: {...},
    mounted() {...}
});
```

## Development Workflows

### Local Development
1. Database must be initialized first:
```bash
sqlite3 tkb.db < database.sql
```

2. Run the server:
```bash
go run main.go
```

3. Access the application at `http://localhost:8080`

### Static Asset Management
- Local copies of third-party libraries are used instead of CDN for offline capability
- Place new libraries in `/static/{js,css}/lib/`
- Update header.html when adding new global dependencies

## Integration Points

### Frontend-Backend Communication
- API endpoints follow RESTful patterns under `/api/` prefix
- Standard CRUD operations per entity
- Excel import/export available via `/api/import/` and `/api/export/` endpoints

### Error Handling
- Backend errors are returned as JSON with an "error" key
- Frontend uses Vue Toastification for error display
- All API calls should be wrapped in try/catch with loading states

## Project-Specific Conventions

### Template Organization
- Use `.tmpl` extension for Go templates
- Base layout in `templates/layouts/`
- Each page should define "content" and optionally "scripts" blocks

### Database Schema
- Tables use snake_case naming
- Foreign keys use `ma_` prefix
- Standard CRUD operations in handlers

Remember to check `database.sql` for the complete schema before making database changes.
