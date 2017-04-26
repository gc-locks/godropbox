// Package gosqlbuilder is a library for generating sql programmatically.
//
// SQL COMPATIBILITY NOTE: sqlbuilder is designed to generate valid MySQL sql
// statements.  The generated statements may not work for other sql variants.
// For instances, the generated statements does not currently work for
// PostgreSQL since column identifiers are escaped with backquotes.
//
// Known limitations for SELECT queries:
//  - does not support subqueries (since mysql is bad at it)
//  - does not currently support join table alias (and hence self join)
//  - does not support NATURAL joins and join USING
//
// Known limitation for INSERT statements:
//  - does not support "INSERT INTO SELECT"
//
// Known limitation for UPDATE statements:
//  - does not support update without a WHERE clause (since it is dangerous)
//  - does not support multi-table update
//
// Known limitation for DELETE statements:
//  - does not support delete without a WHERE clause (since it is dangerous)
//  - does not support multi-table delete
package gosqlbuilder
