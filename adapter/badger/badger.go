package badger

import (
	"context"
	"database/sql"
	"time"

	"github.com/upper/db/v4"
)

const Adapter = `badger`

type settings struct {
}

// SetPreparedStatementCache enables or disables the prepared statement
// cache.
func (s *settings) SetPreparedStatementCache(bool) {

}

// PreparedStatementCacheEnabled returns true if the prepared statement cache
// is enabled, false otherwise.
func (s *settings) PreparedStatementCacheEnabled() bool {

}

// SetConnMaxLifetime sets the default maximum amount of time a connection
// may be reused.
func (s *settings) SetConnMaxLifetime(time.Duration) {

}

// ConnMaxLifetime returns the default maximum amount of time a connection
// may be reused.
func (s *settings) ConnMaxLifetime() time.Duration {

}

// SetMaxIdleConns sets the default maximum number of connections in the idle
// connection pool.
func (s *settings) SetMaxIdleConns(int) {

}

// MaxIdleConns returns the default maximum number of connections in the idle
// connection pool.
func (s *settings) MaxIdleConns() int {

}

// SetMaxOpenConns sets the default maximum number of open connections to the
// database.
func (s *settings) SetMaxOpenConns(int) {

}

// MaxOpenConns returns the default maximum number of open connections to the
// database.
func (s *settings) MaxOpenConns() int {

}

// SetMaxTransactionRetries sets the number of times a transaction can
// be retried.
func (s *settings) SetMaxTransactionRetries(int) {

}

// MaxTransactionRetries returns the maximum number of times a
// transaction can be retried.
func (s *settings) MaxTransactionRetries() int {

}

type session struct {
	settings
}

func New() (db.Session, error) {
	return nil, nil
}

// ConnectionURL returns the DSN that was used to set up the adapter.
func (s *session) ConnectionURL() db.ConnectionURL {

}

// Name returns the name of the database.
func (s *session) Name() string {

}

// Ping returns an error if the DBMS could not be reached.
func (s *session) Ping() error {

}

// Collection receives a table name and returns a collection reference. The
// information retrieved from a collection is cached.
func (s *session) Collection(name string) db.Collection {

}

// Collections returns a collection reference of all non system tables on the
// database.
func (s *session) Collections() ([]db.Collection, error) {

}

// Save creates or updates a record.
func (s *session) Save(record db.Record) error {

}

// Get retrieves a record that matches the given condition.
func (s *session) Get(record db.Record, cond any) error {

}

// Delete deletes a record.
func (s *session) Delete(record db.Record) error {

}

// Reset resets all the caching mechanisms the adapter is using.
func (s *session) Reset() {

}

// Close terminates the currently active connection to the DBMS and clears
// all caches.
func (s *session) Close() error {

}

// Driver returns the underlying driver of the adapter as an interface.
//
// In order to actually use the driver, the `any` value needs to be
// casted into the appropriate type.
//
// Example:
//  internalSQLDriver := sess.Driver().(*sql.DB)
func (s *session) Driver() any {

}

// SQL returns a special interface for SQL databases.
func (s *session) SQL() db.SQL {

}

// Tx creates a transaction block on the default database context and passes
// it to the function fn. If fn returns no error the transaction is commited,
// else the transaction is rolled back. After being commited or rolled back
// the transaction is closed automatically.
func (s *session) Tx(fn func(sess db.Session) error) error {

}

// TxContext creates a transaction block on the given context and passes it to
// the function fn. If fn returns no error the transaction is commited, else
// the transaction is rolled back. After being commited or rolled back the
// transaction is closed automatically.
func (s *session) TxContext(ctx context.Context, fn func(sess db.Session) error, opts *sql.TxOptions) error {

}

// Context returns the context used as default for queries on this session
// and for new transactions.  If no context has been set, a default
// context.Background() is returned.
func (s *session) Context() context.Context {

}

// WithContext returns a copy of the session that uses the given context as
// default. Copies are safe to use concurrently but they're backed by the
// same Session. You may close a copy at any point but that won't close the
// parent session.
func (s *session) WithContext(ctx context.Context) db.Session {

}

type collection struct {
}

// Name returns the name of the collection.
func (c *collection) Name() string {

}

// Session returns the Session that was used to create the collection
// reference.
func (c *collection) Session() db.Session {

}

// Find defines a new result set.
func (c *collection) Find(...any) db.Result {

}

func (c *collection) Count() (uint64, error) {

}

// Insert inserts a new item into the collection, the type of this item could
// be a map, a struct or pointer to either of them. If the call succeeds and
// if the collection has a primary key, Insert returns the ID of the newly
// added element as an `any`. The underlying type of this ID depends
// on both the database adapter and the column storing the ID.  The ID
// returned by Insert() could be passed directly to Find() to retrieve the
// newly added element.
func (c *collection) Insert(any) (db.InsertResult, error) {

}

// InsertReturning is like Insert() but it takes a pointer to map or struct
// and, if the operation succeeds, updates it with data from the newly
// inserted row. If the database does not support transactions this method
// returns db.ErrUnsupported.
func (c *collection) InsertReturning(any) error {

}

// UpdateReturning takes a pointer to a map or struct and tries to update the
// row the item is refering to. If the element is updated sucessfully,
// UpdateReturning will fetch the row and update the fields of the passed
// item.  If the database does not support transactions this method returns
// db.ErrUnsupported
func (c *collection) UpdateReturning(any) error {

}

// Exists returns true if the collection exists, false otherwise.
func (c *collection) Exists() (bool, error) {

}

// Truncate removes all elements on the collection.
func (c *collection) Truncate() error {

}
