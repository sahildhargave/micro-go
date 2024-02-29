### Tools
- [x] Docker
- [x] tablePlus
- [x] redis
- [x] postgres
- [ ] redis
- [ ] jenkins
- [ ] AWS
- [ ] Github 
- [ ] Travis
- [ ] Jenkins
- [ ] CircleCl
- [ ] Actions
- [ ] Many more in future
  

### Database / Sql
😎 Very fast and straight forward
😎 Manual mapping SQL felds to variables
😎 Easy to amke mistakes, not caught until runtime

### GORM
😎  CRUD functions alreday implemented,very short production code

[GORM officical](https://gorm.io/docs/create.html) <a href="https://gorm.io/docs/create.html" style="color: blue;"></a>

- GORM take more time hence SQLX is used for Quite fast & easy to use
- Fields mapping via query text & struct tags
- Failure wont occur until runtime

[SQLC](https://sqlc.dev/) <a href="https://gorm.io/docs/create.html" style="color: blue;"></a>
- very fast and easy to use
- Authomatic code generation
- Catch SQL query errors before genearating codes


### Transation
- example
- [x] Transfer 10 rupess from bank to bank account 2
- [ ]  create a transfer record with amount = 10
- [ ]  Create an account entry for account 1 with amount = -10
- [ ]  Substract 10 from the balance of account 1
- [ ]  Add 10 from the balance of account 2
  
```
- WHY DO WE NEED DB TRANSACTION?
- [x] To provide a reliable and consistent unit of work ,even in case of system failure
- [x] To provide isolation between programs that access the database concurrently 
  
- [] ACID property must follow 
- [x] Atomicity - either all operations complete successfully or the transaction fail and the dbis unchanged.

- [x] Consistency - the db state must be vaild after the transaction. All constraints must be satisfied.

- [x] Isolation - Concurrent transactions must not affect each other.

- [x] Durability - Data written by a successful transaction must be recorded in persistent storage.
  
  How to run sqltx

  BEGIN;
  ...
  COMMIT;

  BEGIN;
  ...
  ROLLLBACK;


#### In Store_test use channel

- channel use 
```
Channel is designed to connect concurretn Go routines,
and allows them to safely share data with each other without explicit locking. In our case , we need 1 channel to receive the errors, and other channel to receive the TransferTxResult.

```

#### DB TXLOCK

debug a deadloak

##### Read Phenomena in ACID

- [x] Dirty Read          - read uncommited   cannot read commited ,cannot repeatable read ,not serializable
- [x] Non-repeatable Read - read uncommited   read commited ,cannot repeatable read ,not serializable
- [x] Phantom Read         - read uncommited  read commited,cannot repeatable read ,not serializable
- [x] Serialization anomaly - read uncommited read commited,
 repeatable read ,not serializable

# Avoiding Deadlocks in Go

## Introduction
- [x] Deadlocks can be a common issue in concurrent programming. This README provides some guidelines to help you avoid deadlocks in Go when working with goroutines and channels.

## Tips to Avoid Deadlocks

### 1. Always Release Locks in the Same Order
   - If you acquire multiple locks, ensure you release them in the same order to prevent potential deadlocks.

### 2. Use `select` with `default` Case
   - When working with channels, prefer using `select` with a `default` case to handle non-blocking operations.
   - This allows you to perform additional checks or take alternative actions when a channel operation would block.

### 3. Avoid Nested Locks
   - Avoid nesting locks whenever possible, as it increases the risk of deadlocks.
   - If you must use nested locks, always release them in the reverse order of acquisition.

### 4. Use `sync.Mutex` and `sync.RWMutex` Carefully
   - When using locks, be cautious with `sync.Mutex` and `sync.RWMutex`.
   - Prefer `sync.RWMutex` when multiple readers are expected, and use `sync.Mutex` when exclusive access is required.

### 5. Be Mindful of Channel Closures
   - Be careful when closing channels to avoid panics or deadlocks.
   - Closing a channel more than once can result in a panic, so ensure proper synchronization.

### 6. Monitor Goroutines
   - Keep an eye on your goroutines using tools like the `goroutine` package or the race detector.
   - Identify any unexpected or unhandled panics in your goroutines that might lead to deadlocks.

### 7. Use `defer` for Unlocking Mutexes
   - When using mutexes, utilize `defer` to ensure that the lock is always released, even if an error occurs or a return statement is encountered.

### 8. Leverage `context` Package for Cancellation
   - Use the `context` package to manage the lifecycle of your goroutines.
   - This allows for clean shutdowns and helps avoid situations where a goroutine is waiting indefinitely.


```

### Workflow 
![alt text](./assets/image.png)
![alt text](./assets/image1.png)

