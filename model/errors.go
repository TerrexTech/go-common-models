package model

// InternalError represents an error when something goes wrong, and its our fault.
const InternalError int16 = 1

// DatabaseError is when some operation related to Database,
// such as insert or find, goes wrong and the task cannot proceed.
const DatabaseError int16 = 2

// UserError occurs when there's an error because of user's action.
// An example would be providing invalid input.
const UserError int16 = 3
