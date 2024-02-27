### Tools
- [x] Docker
- [x] tablePlus
- [x] redis
- [x] postgres
- [ ] redis
- [ ] jenkins
- [ ] AWS
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
  
  